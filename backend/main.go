package main

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"cl_system/internal/cmd"
	_ "cl_system/internal/logic"
)

func main() {
	s := g.Server()
	s.BindMiddlewareDefault(MiddlewareCORS)
	s.Group("/api", func(group *ghttp.RouterGroup) {
		cmd.Router(group)
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

func init() {
	fmt.Println("链环系统启动中...")
}
