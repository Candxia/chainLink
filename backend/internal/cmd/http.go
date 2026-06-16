package cmd

import (
	"context"
	"os"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gproc"
	"cl_system/internal/model"
	_ "cl_system/internal/logic"
)

var (
	Main = &gcmd.Command{
		Name:        "main",
		Brief:       "start cl_system server",
		Description: "启动超链系统服务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 初始化日志钩子
			initHooks()

			// 启动 SSE 消息消费
			go model.SseNotice()

			// 初始化 access_log 表
			initAccessLogTable()

			// 注册全局中间件
			s.BindMiddlewareDefault(middlewareCORS)
			s.BindMiddlewareDefault(middlewareResponse)
			s.BindMiddlewareDefault(middlewareAccessLog)

			// 注册 Hook 钩子
			s.BindHookHandler("/api/personal/login", "HookAfterOutput", loginLogHook)
			s.BindHookHandler("/api/*", "HookAfterOutput", operationLogHook)

			// 注册 SSE 端点
			s.BindHandler("/api/system/notice", NoticeHandler)

			// 注册业务路由
			s.Group("/api", func(group *ghttp.RouterGroup) {
				Router(group)
			})

			// 同步 API 规则
			initApiRuleTable()
			syncApiRulesToDB(s)
			registerApiRuleRoutes(s)

			s.SetServerRoot("resource/public")
			s.SetPort(8099)
			err = s.Start()
			if err != nil {
				panic(any("服务启动失败: " + err.Error()))
			}
			gproc.AddSigHandlerShutdown(func(sig os.Signal) {
				g.Log().Debug(ctx, "收到关闭信号:", sig)
			})
			gproc.Listen()
			return
		},
	}
)
