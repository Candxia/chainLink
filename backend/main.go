package main

import (
	"context"
	"fmt"
	"strings"
	"time"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"cl_system/internal/cmd"
	_ "cl_system/internal/logic"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	s := g.Server()

	// 启动时初始化 access_log 表
	initAccessLogTable()

	// 全局中间件: 顺序执行 CORS → AccessLog → 业务路由
	s.BindMiddlewareDefault(MiddlewareCORS)
	s.BindMiddlewareDefault(MiddlewareAccessLog)
	s.Group("/api", func(group *ghttp.RouterGroup) {
		cmd.Router(group)
	})

	// 同步 API 路由到数据库（在启动前注册）
	initApiRuleTable()
	syncApiRulesToDB(s)

	// 注册 API 规则管理路由（在 /api 组外，避免双前缀）
	s.Group("/api/api-rule", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareAccessLog)
		group.Middleware(cmd.AuthMiddleware)
		group.GET("/list", apiRuleList)
		group.PUT("/update", apiRuleUpdate)
	})

	s.SetServerRoot("resource/public")
	s.SetPort(8099)
	s.Run()
}

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.Header().Set("Access-Control-Allow-Origin", "*")
	r.Response.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	r.Response.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
	if r.Method == "OPTIONS" {
		r.ExitAll()
		return
	}
	r.Middleware.Next()
}

// MiddlewareAccessLog 记录每个 API 请求的完整访问日志
// 包含: 操作人、操作时间、请求信息、接口名称、耗时(ms)、IP地址、客户端密钥、返回详情
func MiddlewareAccessLog(r *ghttp.Request) {
	startTime := time.Now()
	traceId := fmt.Sprintf("%x", time.Now().UnixNano())

	// 先执行业务逻辑
	r.Middleware.Next()

	// 业务完成后，收集日志信息
	durationMs := time.Since(startTime).Milliseconds()

	// 获取操作人信息（从JWT上下文或session）
	operator := "anonymous"
	userId := uint(0)
	usernameCtx := r.GetCtxVar("username")
	userIdCtx := r.GetCtxVar("userId")
	if !usernameCtx.IsEmpty() {
		operator = usernameCtx.String()
	}
	if !userIdCtx.IsEmpty() {
		userId = userIdCtx.Uint()
	}

	// 请求信息
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

	// 接口名称
	apiName := r.Router.Uri
	if apiName == "" {
		apiName = r.URL.Path
	}

	// 客户端信息
	ipAddress := r.GetClientIp()
	userAgent := r.Header.Get("User-Agent")
	if len(userAgent) > 200 {
		userAgent = userAgent[:200]
	}
	authToken := r.Header.Get("Authorization")
	if len(authToken) > 100 {
		authToken = authToken[:100] + "..."
	}

	// 响应信息
	responseCode := r.Response.Status
	responseBody := r.Response.BufferString()
	if len(responseBody) > 1000 {
		responseBody = responseBody[:1000] + "..."
	}

	// 状态判断
	status := 1
	if responseCode >= 400 || strings.Contains(responseBody, `"code":1`) || strings.Contains(responseBody, `"code":401`) {
		status = 0
	}

	// 异步写入数据库（不阻塞请求响应）
	go func() {
		defer func() {
			if err := recover(); err != nil {
				g.Log().Warning(nil, "access log write error:", err)
			}
		}()
		_, err := g.DB().Model("access_log").Insert(g.Map{
			"trace_id":       traceId,
			"operator":       operator,
			"user_id":        userId,
			"action_time":    startTime.Format("2006-01-02 15:04:05"),
			"request_method": r.Method,
			"request_path":   r.URL.Path,
			"request_params": requestParams,
			"api_name":       apiName,
			"duration_ms":    durationMs,
			"ip_address":     ipAddress,
			"user_agent":     userAgent,
			"auth_token":     authToken,
			"response_code":  responseCode,
			"response_body":  responseBody,
			"status":         status,
		})
		if err != nil {
			g.Log().Warning(nil, "access log insert failed:", err)
		}
	}()
}

// initAccessLogTable 启动时自动创建 access_log 表
func initAccessLogTable() {
	defer func() {
		if err := recover(); err != nil {
			g.Log().Warning(nil, "access log table init error:", err)
		}
	}()
	sql := "CREATE TABLE IF NOT EXISTS access_log (" +
		"id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY," +
		"trace_id VARCHAR(64) DEFAULT '' COMMENT '链路追踪ID'," +
		"operator VARCHAR(64) DEFAULT '' COMMENT '操作人'," +
		"user_id BIGINT UNSIGNED DEFAULT 0 COMMENT '操作人ID'," +
		"action_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间'," +
		"request_method VARCHAR(16) DEFAULT '' COMMENT '请求方法'," +
		"request_path VARCHAR(256) DEFAULT '' COMMENT '请求路径'," +
		"request_params TEXT COMMENT '请求参数'," +
		"api_name VARCHAR(256) DEFAULT '' COMMENT '接口名称'," +
		"duration_ms INT DEFAULT 0 COMMENT '耗时(ms)'," +
		"ip_address VARCHAR(64) DEFAULT '' COMMENT 'IP地址'," +
		"user_agent VARCHAR(512) DEFAULT '' COMMENT 'User-Agent'," +
		"auth_token VARCHAR(256) DEFAULT '' COMMENT '客户端Token'," +
		"response_code INT DEFAULT 0 COMMENT '响应状态码'," +
		"response_body TEXT COMMENT '返回详情'," +
		"status TINYINT DEFAULT 1 COMMENT '状态'," +
		"created_at DATETIME DEFAULT CURRENT_TIMESTAMP," +
		"INDEX idx_operator (operator)," +
		"INDEX idx_action_time (action_time)," +
		"INDEX idx_request_path (request_path(64))," +
		"INDEX idx_duration (duration_ms)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='访问日志表';"
	_, err := g.DB().Exec(context.Background(), sql)
	if err != nil {
		g.Log().Warning(nil, "access_log table ensure error:", err)
	} else {
		fmt.Println("access_log 表已就绪")
	}
}

// ==================== API 规则管理 ====================

// initApiRuleTable 创建 api_rule 表
func initApiRuleTable() {
	defer func() {
		if err := recover(); err != nil {
			g.Log().Warning(nil, "api_rule table init error:", err)
		}
	}()
	sql := "CREATE TABLE IF NOT EXISTS api_rule (" +
		"id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY," +
		"title VARCHAR(128) DEFAULT '' COMMENT '标题'," +
		"auth_type VARCHAR(32) DEFAULT 'require' COMMENT '权限类型(require=需登录|public=公开|admin=管理员)'," +
		"path VARCHAR(256) NOT NULL COMMENT '请求路径'," +
		"method VARCHAR(16) NOT NULL COMMENT '请求方式(GET/POST/PUT/DELETE)'," +
		"enable_log TINYINT DEFAULT 1 COMMENT '是否记录日志(1=是 0=否)'," +
		"status TINYINT DEFAULT 1 COMMENT '状态'," +
		"created_at DATETIME DEFAULT CURRENT_TIMESTAMP," +
		"updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		"UNIQUE INDEX idx_path_method (path(128), method)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='API规则表';"
	_, err := g.DB().Exec(context.Background(), sql)
	if err != nil {
		g.Log().Warning(nil, "api_rule table ensure error:", err)
	} else {
		fmt.Println("api_rule 表已就绪")
	}
}

// syncApiRulesToDB 将当前路由自动同步到 api_rule 表
func syncApiRulesToDB(s *ghttp.Server) {
	defer func() {
		if err := recover(); err != nil {
			g.Log().Warning(nil, "sync api rules error:", err)
		}
	}()
	routes := s.GetRoutes()
	count := 0
	for _, r := range routes {
		if r.Route == "/*" || strings.HasPrefix(r.Route, "/api/api") {
			continue
		}
		// 判断是否公开接口
		authType := "require"
		if !strings.HasPrefix(r.Route, "/api/") || r.Route == "/api/user/login" || r.Route == "/api/user/register" || r.Route == "/api/user/logout" {
			authType = "public"
		}
		// 路径拼接
		fullPath := r.Route
		// 提取标题
		title := strings.TrimPrefix(r.Route, "/api/")
		if idx := strings.Index(title, "/"); idx > 0 {
			title = title[:idx]
		}
		title = strings.ReplaceAll(title, "/", " ")

		// 计算唯一键 upsert
		existing, _ := g.DB().Model("api_rule").Where("path", fullPath).Where("method", r.Method).One()
		if existing == nil || existing.IsEmpty() {
			_, err := g.DB().Model("api_rule").Insert(g.Map{
				"title":      title,
				"auth_type":  authType,
				"path":       fullPath,
				"method":     r.Method,
				"enable_log": 1,
			})
			if err == nil {
				count++
			} else {
				g.Log().Debug(nil, "sync route error:", fullPath, err)
			}
		}
	}
	fmt.Printf("api_rule 同步完成: %d 条路由\n", count)
}

// apiRuleList 获取 API 规则列表
func apiRuleList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 20).Int()
	total, _ := g.DB().Model("api_rule").Count()
	list, err := g.DB().Model("api_rule").Page(page, pageSize).Order("id asc").All()
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	// 检查是否有 api-rule 本身的日志
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list.List(), "total": total, "page": page, "pageSize": pageSize}})
}

// apiRuleUpdate 更新 API 规则
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
