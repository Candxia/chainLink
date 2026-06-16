package cmd

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"cl_system/internal/model"
)

// NoticeHandler SSE 推送端点处理器
// 实现服务端推送通知功能，通过 EventSource 协议向客户端推送实时消息
func NoticeHandler(r *ghttp.Request) {
	// 先做 JWT/Token 校验
	token := r.GetHeader("Authorization")
	if token == "" {
		model.SseSys(r.Response, 1, "Unauthorized")
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	user := &model.SseUser{
		Code:   r.GetClientIp() + "-" + time.Now().Format("150405.000"),
		Ctx:    ctx,
		Cancel: cancel,
		Send:   make(chan string, 20),
	}
	defer close(user.Send)

	// 从 token 或上下文获取用户名
	username := r.GetCtxVar("username").String()
	if username == "" {
		// 使用客户端 IP 作为 fallback 标识
		username = "anon-" + r.GetClientIp()
	}

	// 设置 SSE 响应头
	r.Response.Header().Set("Content-Type", "text/event-stream")
	r.Response.Header().Set("Cache-Control", "no-cache")
	r.Response.Header().Set("Connection", "keep-alive")
	r.Response.Header().Set("Access-Control-Allow-Origin", "*")

	// 注册订阅
	model.SseSubscribe.Set(username, user)
	model.SseSys(r.Response, 1, "Connection successful")

	// 心跳保活
	heartbeat := time.NewTicker(15 * time.Second)
	defer heartbeat.Stop()

	eventId := int64(1)
	for {
		select {
		case msg := <-user.Send:
			eventId++
			model.SseDelta(r.Response, eventId, msg)
		case <-heartbeat.C:
			eventId++
			model.SseSys(r.Response, eventId, "ping")
		case <-user.Ctx.Done():
			g.Log().Debug(ctx, "SSE connection done:", username)
			return
		case <-r.Context().Done():
			model.SseSubscribe.Del(username, user.Code)
			g.Log().Debug(ctx, "SSE disconnected:", username)
			return
		}
	}
}
