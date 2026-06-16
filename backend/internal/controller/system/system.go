package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	systemV1 "cl_system/api/system/v1"
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
	var req systemV1.UpdateConfigReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.System().UpdateConfig(r.Context(), req.Data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) GetLogList(r *ghttp.Request) {
	var req systemV1.GetLogListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.System().GetLogList(r.Context(), req.Page, req.PageSize, req.Type)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total}})
}

func (c *Controller) DeleteLog(r *ghttp.Request) {
	var req systemV1.DeleteLogReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.System().DeleteLog(r.Context(), req.Id)
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
	var req systemV1.MarkNotificationReadReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.System().MarkNotificationRead(r.Context(), req.Id)
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
