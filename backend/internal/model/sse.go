package model

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
)

// ==================== SSE 消息体 ====================

// SseMsgRes 推送给客户端的消息结构
type SseMsgRes struct {
	Code    int           `json:"code"`
	Type    string        `json:"type"`
	Content SseMsgContent `json:"content"`
}

// SseMsg 推送队列中的消息结构
type SseMsg struct {
	Code     int           `json:"code"`
	Type     string        `json:"type"`
	Content  SseMsgContent `json:"content"`
	Topic    string        `json:"topic"`
	Username string        `json:"username"`
}

// SseMsgContent 消息内容
type SseMsgContent struct {
	Msg    string   `json:"msg"`
	Action string   `json:"action"`
	Title  string   `json:"title"`
	Param  []string `json:"param"`
}

// SseUser SSE 连接用户
type SseUser struct {
	Code   string
	Cancel context.CancelFunc
	Ctx    context.Context
	Send   chan string
}

// SseUserMap 线程安全的 SSE 用户映射表
type SseUserMap struct {
	mu   sync.Mutex
	List map[string]*SseUser
}

func (s *SseUserMap) Get(key string) *SseUser {
	s.mu.Lock()
	defer s.mu.Unlock()
	if val, ok := s.List[key]; ok {
		return val
	}
	return nil
}

func (s *SseUserMap) Set(key string, val *SseUser) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if old, ok := s.List[key]; ok {
		old.Cancel()
	}
	s.List[key] = val
}

func (s *SseUserMap) Del(key, code string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if val, ok := s.List[key]; ok && val.Code == code {
		delete(s.List, key)
	}
}

func (s *SseUserMap) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.List)
}

// SseSubscribe 全局 SSE 订阅管理
var SseSubscribe = SseUserMap{
	List: map[string]*SseUser{},
}

// SseQueue 全局 SSE 消息队列
var SseQueue = gqueue.New()

// SsePush 向 SSE 队列推送消息
func SsePush(info SseMsg) {
	SseQueue.Push(info)
}

// SseDelta 向客户端发送 delta 事件
func SseDelta(r *ghttp.Response, id int64, data string) {
	r.Writeln("id: " + strconv.FormatInt(id, 10))
	r.Writeln("event: delta")
	r.Writeln("data: " + data)
	r.Writeln("")
	r.Flush()
}

// SseSys 向客户端发送系统事件
func SseSys(r *ghttp.Response, id int64, data string) {
	r.Writeln("id: " + strconv.FormatInt(id, 10))
	r.Writeln("event: system")
	r.Writeln("data: " + data)
	r.Writeln("")
	r.Flush()
}

// SseNotice SSE 消息消费循环
// 从 SseQueue 取出消息并分发给对应订阅用户
func SseNotice() {
	defer func() {
		if rc := recover(); rc != nil {
			g.Log().Errorf(context.TODO(), "Panic recovered in SseNotice: %+v\n%s", rc, debug.Stack())
		}
	}()
	for {
		info := SseQueue.Pop()
		if info == nil {
			break
		}
		msg := info.(SseMsg)
		if msg.Username != "" {
			users := strings.Split(msg.Username, ",")
			for _, user := range users {
				if u := SseSubscribe.Get(user); u != nil {
					dataBytes, _ := json.Marshal(SseMsgRes{msg.Code, msg.Type, msg.Content})
					select {
					case u.Send <- string(dataBytes):
					default:
						// 通道已满，跳过
					}
				}
			}
		} else {
			SseSubscribe.mu.Lock()
			for _, u := range SseSubscribe.List {
				dataBytes, _ := json.Marshal(SseMsgRes{msg.Code, msg.Type, msg.Content})
				select {
				case u.Send <- string(dataBytes):
				default:
				}
			}
			SseSubscribe.mu.Unlock()
		}
	}
}
