package service

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Response 统一响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ==================== 中间件服务 ====================

type sMiddleware struct{}

func Middleware() *sMiddleware {
	return &sMiddleware{}
}

// Response 统一响应中间件
// 将所有 Controller 返回值统一包装成标准 Response 格式
func (s *sMiddleware) Response(r *ghttp.Request) {
	r.Middleware.Next()

	// 如果已经输出过内容则不处理
	if r.Response.BufferLength() > 0 {
		return
	}

	// 服务端错误统一处理
	if r.Response.Status >= http.StatusInternalServerError {
		r.Response.ClearBuffer()
		r.Response.WriteJson(Response{
			Code:    500,
			Message: "服务器内部错误",
			Data:    nil,
		})
		r.ExitAll()
		return
	}

	var (
		msg  = "success"
		err  = r.GetError()
		res  interface{} = nil
		code              = gerror.Code(err)
	)

	if err != nil {
		if code.Code() > 0 {
			msg = code.Message()
		} else {
			msg = err.Error()
		}
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
		} else {
			if out := r.GetHandlerResponse(); !g.IsNil(out) {
				res = out
			}
			code = gcode.CodeOK
		}
	}

	r.Response.WriteJson(Response{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
}

// ==================== 日志钩子 ====================

type sLogHook struct {
	_logger map[string]ApiHook
	_ctx    context.Context
	_l      sync.RWMutex
	_init   bool
}

type ApiHook struct {
	Title  string
	Action string
}

type LoginLogInfo struct {
	Username string
	Status   int
	Ipaddr   string
	UserAgent string
	Message  string
	Code     int
}

type OperationLogInfo struct {
	AdminId    uint
	Username   string
	Method     string
	Path       string
	Title      string
	Ipaddr     string
	UserAgent  string
	Params     interface{}
	Response   interface{}
	Error      error
	Latency    time.Duration
}

var LogHook = new(sLogHook)

// LoadApi 从 api_rule 表加载需要记录日志的 API 配置
func (s *sLogHook) LoadApi() {
	ctx := context.Background()
	list, err := g.DB().Model("api_rule").
		Where("enable_log", 1).
		Order("id asc").
		All()
	if err != nil {
		g.Log().Warning(ctx, "加载 API 规则失败:", err)
		return
	}
	policy := map[string]ApiHook{}
	for _, item := range list {
		path := item["path"].String()
		method := item["method"].String()
		title := item["title"].String()
		key := method + ":" + path
		policy[key] = ApiHook{
			Title:  title,
			Action: method,
		}
	}
	s._l.Lock()
	defer s._l.Unlock()
	s._logger = policy
	s._ctx = ctx
}

// GetHook 获取 API 的钩子配置
func (s *sLogHook) GetHook(method, path string) (bool, string) {
	s._l.RLock()
	defer s._l.RUnlock()
	key := method + ":" + path
	if v, ok := s._logger[key]; ok {
		return false, v.Title
	}
	return true, ""
}

// Operation 操作日志钩子（通过 BindHookHandler 注册在 HookAfterOutput 时触发）
func (s *sLogHook) Operation(r *ghttp.Request) {
	if !s._init {
		s.LoadApi()
		s._init = true
	}

	// 获取当前用户信息
	userId := r.GetCtxVar("userId").Uint()
	username := r.GetCtxVar("username").String()
	if userId == 0 && username == "" {
		return
	}

	not, title := s.GetHook(r.Method, r.URL.Path)
	if not {
		return
	}

	latency := time.Since(time.Unix(r.EnterTime, 0)).Round(time.Millisecond)
	logInfo := OperationLogInfo{
		AdminId:   userId,
		Username:  username,
		Method:    r.Method,
		Path:      r.URL.Path,
		Title:     title,
		Ipaddr:    r.GetClientIp(),
		UserAgent: r.Request.UserAgent(),
		Params:    r.GetMap(),
		Response:  r.GetHandlerResponse(),
		Error:     r.GetError(),
		Latency:   latency,
	}

	// 异步写入操作日志
	go func() {
		defer func() {
			if err := recover(); err != nil {
				g.Log().Warning(nil, "operation log write error:", err)
			}
		}()
		s.writeOpLog(logInfo)
	}()
}

func (s *sLogHook) writeOpLog(info OperationLogInfo) {
	code := 0
	if info.Error != nil {
		code = gerror.Code(info.Error).Code()
	}

	_, err := g.DB().Model("access_log").Insert(g.Map{
		"operator":      info.Username,
		"user_id":       info.AdminId,
		"action_time":   time.Now().Format("2006-01-02 15:04:05"),
		"request_method": info.Method,
		"request_path":  info.Path,
		"api_name":      info.Title,
		"duration_ms":   info.Latency.Milliseconds(),
		"ip_address":    info.Ipaddr,
		"user_agent":    info.UserAgent,
		"response_code": code,
		"status":        1,
	})
	if err != nil {
		g.Log().Warning(nil, "operation log insert failed:", err)
	}
}

// LoginLog 登录日志钩子
func (s *sLogHook) LoginLog(r *ghttp.Request) {
	username := r.Get("username", "undefined").String()
	err := r.GetError()

	status := 2 // 失败
	if err == nil {
		status = 1 // 成功
	}

	loginLog := LoginLogInfo{
		Username:  username,
		Status:    status,
		Ipaddr:    r.GetClientIp(),
		UserAgent: r.Request.UserAgent(),
	}
	if err != nil {
		loginLog.Message = err.Error()
		loginLog.Code = gerror.Code(err).Code()
	}

	// 异步写入登录日志
	go func() {
		defer func() {
			if err := recover(); err != nil {
				g.Log().Warning(nil, "login log write error:", err)
			}
		}()
		s.writeLoginLog(loginLog)
	}()
}

func (s *sLogHook) writeLoginLog(info LoginLogInfo) {
	_, err := g.DB().Model("login_log").Insert(g.Map{
		"username":   info.Username,
		"status":     info.Status,
		"ipaddr":     info.Ipaddr,
		"user_agent": info.UserAgent,
		"message":    info.Message,
		"code":       info.Code,
		"login_at":   time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		g.Log().Warning(nil, "login log insert failed:", err)
	}
}

// LoginFailed 记录登录失败日志
func (s *sLogHook) LoginFailed(ctx context.Context, err error) {
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return
	}
	username := r.GetRequest("username").String()
	log := LoginLogInfo{
		Username:  username,
		Status:    2,
		Ipaddr:    r.GetClientIp(),
		UserAgent: r.Request.UserAgent(),
		Message:   err.Error(),
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				g.Log().Warning(nil, "login failed log write error:", err)
			}
		}()
		s.writeLoginLog(log)
	}()
}

// ==================== 初始化 ====================

func ensureLogTables() {
	ctx := context.Background()

	// 创建 login_log 表
	loginLogSQL := `CREATE TABLE IF NOT EXISTS login_log (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(64) DEFAULT '' COMMENT '用户名',
		status TINYINT DEFAULT 1 COMMENT '状态(1=成功 2=失败)',
		ipaddr VARCHAR(64) DEFAULT '' COMMENT 'IP地址',
		user_agent VARCHAR(256) DEFAULT '' COMMENT 'User-Agent',
		message VARCHAR(512) DEFAULT '' COMMENT '消息',
		code INT DEFAULT 0 COMMENT '错误码',
		login_at DATETIME DEFAULT NULL COMMENT '登录时间',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		INDEX idx_username (username),
		INDEX idx_login_at (login_at)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='登录日志表'`
	if _, err := g.DB().Exec(ctx, loginLogSQL); err != nil {
		g.Log().Warning(ctx, "login_log table ensure error:", err)
	} else {
		fmt.Println("login_log 表已就绪")
	}
}

// InitHooks 初始化日志钩子（由 main.go 调用）
func InitHooks() {
	ensureLogTables()
}

// EnsureResponseMiddleware 确认 Response 中间件已注入（返回实例供 BindMiddlewareDefault 使用）
func EnsureResponseMiddleware() *sMiddleware {
	return Middleware()
}
