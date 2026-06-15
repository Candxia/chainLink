package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"cl_system/internal/service"
)

type Controller struct{}
var Ctl = &Controller{}

func (c *Controller) GetDashboard(r *ghttp.Request) {
	data, err := service.System().GetDashboard(r.Context())
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": data})
}

func (c *Controller) GetConfig(r *ghttp.Request) {
	configs, err := service.System().GetConfigList(r.Context())
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": configs})
}

func (c *Controller) UpdateConfig(r *ghttp.Request) {
	var data map[string]interface{}
	if err := r.Parse(&data); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.System().UpdateConfig(r.Context(), data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) GetLogList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 20).Int()
	logType := r.GetQuery("type").String()
	list, total, err := service.System().GetLogList(r.Context(), page, pageSize, logType)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total}})
}

func (c *Controller) DeleteLog(r *ghttp.Request) {
	id := r.Get("id").Uint()
	err := service.System().DeleteLog(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除成功"})
}

func (c *Controller) GetNotificationList(r *ghttp.Request) {
	userId := r.GetCtxVar("userId").Uint()
	list, err := service.System().GetNotificationList(r.Context(), userId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": list})
}

func (c *Controller) MarkNotificationRead(r *ghttp.Request) {
	id := r.Get("id").Uint()
	err := service.System().MarkNotificationRead(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "已标记已读"})
}

func (c *Controller) UploadFile(r *ghttp.Request) {
	file := r.GetUploadFile("file")
	if file == nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": "请选择文件"})
		return
	}
	url, err := service.System().UploadFile(r.Context(), file)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "上传成功", "data": g.Map{"url": url}})
}
