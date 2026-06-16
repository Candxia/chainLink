package cmd

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"cl_system/internal/service"
)

// middlewareCORS CORS 跨域中间件
func middlewareCORS(r *ghttp.Request) {
	r.Response.Header().Set("Access-Control-Allow-Origin", "*")
	r.Response.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	r.Response.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
	if r.Method == "OPTIONS" {
		r.ExitAll()
		return
	}
	r.Middleware.Next()
}

// middlewareResponse 统一响应包装中间件
func middlewareResponse(r *ghttp.Request) {
	r.Middleware.Next()
	if r.Response.BufferLength() > 0 {
		return
	}
	if r.Response.Status >= 500 {
		r.Response.ClearBuffer()
		r.Response.WriteJson(service.Response{
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
		if r.Response.Status > 0 && r.Response.Status != 200 {
			msg = http.StatusText(r.Response.Status)
		} else {
			if out := r.GetHandlerResponse(); !g.IsNil(out) {
				res = out
			}
			code = gcode.CodeOK
		}
	}
	r.Response.WriteJson(service.Response{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
}

// middlewareAccessLog 记录完整访问日志
func middlewareAccessLog(r *ghttp.Request) {
	startTime := time.Now()
	traceId := fmt.Sprintf("%x", time.Now().UnixNano())
	r.Middleware.Next()

	durationMs := time.Since(startTime).Milliseconds()
	operator := "anonymous"
	userId := uint(0)
	if v := r.GetCtxVar("username"); !v.IsEmpty() {
		operator = v.String()
	}
	if v := r.GetCtxVar("userId"); !v.IsEmpty() {
		userId = v.Uint()
	}

	// 请求参数
	queryStr := r.URL.RawQuery
	bodyBytes := r.GetBody()
	requestParams := queryStr
	if len(bodyBytes) > 0 && len(string(bodyBytes)) < 500 {
		if requestParams != "" {
			requestParams += " | body: " + string(bodyBytes)
		} else {
			requestParams = string(bodyBytes)
		}
	}
	if len(requestParams) > 500 {
		requestParams = requestParams[:500] + "..."
	}

	apiName := r.Router.Uri
	if apiName == "" {
		apiName = r.URL.Path
	}

	userAgent := r.Header.Get("User-Agent")
	if len(userAgent) > 200 {
		userAgent = userAgent[:200]
	}
	authToken := r.Header.Get("Authorization")
	if len(authToken) > 100 {
		authToken = authToken[:100] + "..."
	}

	responseCode := r.Response.Status
	responseBody := r.Response.BufferString()
	if len(responseBody) > 1000 {
		responseBody = responseBody[:1000] + "..."
	}
	status := 1
	if responseCode >= 400 || strings.Contains(responseBody, `"code":1`) || strings.Contains(responseBody, `"code":401`) {
		status = 0
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				g.Log().Warning(nil, "access log write error:", err)
			}
		}()
		g.DB().Model("access_log").Insert(g.Map{
			"trace_id":       traceId,
			"operator":       operator,
			"user_id":        userId,
			"action_time":    startTime.Format("2006-01-02 15:04:05"),
			"request_method": r.Method,
			"request_path":   r.URL.Path,
			"request_params": requestParams,
			"api_name":       apiName,
			"duration_ms":    durationMs,
			"ip_address":     r.GetClientIp(),
			"user_agent":     userAgent,
			"auth_token":     authToken,
			"response_code":  responseCode,
			"response_body":  responseBody,
			"status":         status,
		})
	}()
}

// AuthMiddleware JWT 认证中间件
func AuthMiddleware(r *ghttp.Request) {
	token := r.GetHeader("Authorization")
	if token == "" {
		r.Response.WriteJson(g.Map{"code": 401, "message": "未登录"})
		r.Exit()
		return
	}
	r.Middleware.Next()
}

// ==================== 日志钩子 ====================

// loginLogHook 登录日志钩子
func loginLogHook(r *ghttp.Request) {
	service.LogHook.LoginLog(r)
}

// operationLogHook 操作日志钩子
func operationLogHook(r *ghttp.Request) {
	service.LogHook.Operation(r)
}

// ==================== 初始化 ====================

// initHooks 初始化日志表
func initHooks() {
	service.InitHooks()
}

// initAccessLogTable 启动时自动创建 access_log 表
func initAccessLogTable() {
	defer func() {
		if err := recover(); err != nil {
			g.Log().Warning(nil, "access log table init error:", err)
		}
	}()
	sql := `CREATE TABLE IF NOT EXISTS access_log (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		trace_id VARCHAR(64) DEFAULT '' COMMENT '链路追踪ID',
		operator VARCHAR(64) DEFAULT '' COMMENT '操作人',
		user_id BIGINT UNSIGNED DEFAULT 0 COMMENT '操作人ID',
		action_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
		request_method VARCHAR(16) DEFAULT '' COMMENT '请求方法',
		request_path VARCHAR(256) DEFAULT '' COMMENT '请求路径',
		request_params TEXT COMMENT '请求参数',
		api_name VARCHAR(256) DEFAULT '' COMMENT '接口名称',
		duration_ms INT DEFAULT 0 COMMENT '耗时(ms)',
		ip_address VARCHAR(64) DEFAULT '' COMMENT 'IP地址',
		user_agent VARCHAR(512) DEFAULT '' COMMENT 'User-Agent',
		auth_token VARCHAR(256) DEFAULT '' COMMENT '客户端Token',
		response_code INT DEFAULT 0 COMMENT '响应状态码',
		response_body TEXT COMMENT '返回详情',
		status TINYINT DEFAULT 1 COMMENT '状态',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		INDEX idx_operator (operator),
		INDEX idx_action_time (action_time),
		INDEX idx_request_path (request_path(64)),
		INDEX idx_duration (duration_ms)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='访问日志表'`
	if _, err := g.DB().Exec(context.Background(), sql); err != nil {
		g.Log().Warning(nil, "access_log table ensure error:", err)
	} else {
		fmt.Println("access_log 表已就绪")
	}
}

// initApiRuleTable 创建 api_rule 表
func initApiRuleTable() {
	defer func() {
		if err := recover(); err != nil {
			g.Log().Warning(nil, "api_rule table init error:", err)
		}
	}()
	sql := `CREATE TABLE IF NOT EXISTS api_rule (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(128) DEFAULT '' COMMENT '标题',
		auth_type VARCHAR(32) DEFAULT 'require' COMMENT '权限类型(require=需登录|public=公开|admin=管理员)',
		path VARCHAR(256) NOT NULL COMMENT '请求路径',
		method VARCHAR(16) NOT NULL COMMENT '请求方式(GET/POST/PUT/DELETE)',
		enable_log TINYINT DEFAULT 1 COMMENT '是否记录日志(1=是 0=否)',
		status TINYINT DEFAULT 1 COMMENT '状态',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		UNIQUE INDEX idx_path_method (path(128), method)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='API规则表'`
	if _, err := g.DB().Exec(context.Background(), sql); err != nil {
		g.Log().Warning(nil, "api_rule table ensure error:", err)
	} else {
		fmt.Println("api_rule 表已就绪")
	}
}

// syncApiRulesToDB 自动同步路由到 api_rule 表
func syncApiRulesToDB(s *ghttp.Server) {
	defer func() {
		if err := recover(); err != nil {
			g.Log().Warning(nil, "sync api rules error:", err)
		}
	}()
	routes := s.GetRoutes()
	count := 0
	for _, r := range routes {
		if r.Route == "/*" || strings.HasPrefix(r.Route, "/api/api") || strings.HasPrefix(r.Route, "/api/system/notice") {
			continue
		}
		authType := "require"
		if !strings.HasPrefix(r.Route, "/api/") || isPublicRoute(r.Route) {
			authType = "public"
		}
		fullPath := r.Route
		title := extractTitle(fullPath)

		existing, _ := g.DB().Model("api_rule").Where("path", fullPath).Where("method", r.Method).One()
		if existing == nil || existing.IsEmpty() {
			g.DB().Model("api_rule").Insert(g.Map{
				"title":      title,
				"auth_type":  authType,
				"path":       fullPath,
				"method":     r.Method,
				"enable_log": 1,
			})
			count++
		}
	}
	fmt.Printf("api_rule 同步完成: %d 条路由\n", count)
}

func isPublicRoute(route string) bool {
	pubPaths := []string{"/api/user/login", "/api/user/register", "/api/user/logout", "/api/trace/public/query"}
	for _, p := range pubPaths {
		if route == p {
			return true
		}
	}
	return false
}

func extractTitle(fullPath string) string {
	title := strings.TrimPrefix(fullPath, "/api/")
	if idx := strings.Index(title, "/"); idx > 0 {
		title = title[:idx]
	}
	return strings.ReplaceAll(title, "/", " ")
}

// registerApiRuleRoutes 注册 API 规则管理路由
func registerApiRuleRoutes(s *ghttp.Server) {
	s.Group("/api/api-rule", func(group *ghttp.RouterGroup) {
		group.Middleware(middlewareAccessLog)
		group.Middleware(AuthMiddleware)
		group.GET("/list", apiRuleList)
		group.PUT("/update", apiRuleUpdate)
	})
}

func apiRuleList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 20).Int()
	total, _ := g.DB().Model("api_rule").Count()
	list, err := g.DB().Model("api_rule").Page(page, pageSize).Order("id asc").All()
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list.List(), "total": total, "page": page, "pageSize": pageSize}})
}

func apiRuleUpdate(r *ghttp.Request) {
	id := r.GetForm("id").Uint()
	data := g.Map{}
	if title := r.GetForm("title"); !title.IsEmpty() {
		data["title"] = title.String()
	}
	if authType := r.GetForm("auth_type"); !authType.IsEmpty() {
		data["auth_type"] = authType.String()
	}
	if enableLog := r.GetForm("enable_log"); !enableLog.IsEmpty() {
		data["enable_log"] = enableLog.Int()
	}
	if len(data) == 0 || id == 0 {
		r.Response.WriteJson(g.Map{"code": 1, "message": "参数错误"})
		return
	}
	_, err := g.DB().Model("api_rule").Where("id", id).Update(data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func init() {
	fmt.Println("链环系统启动中...")
}
