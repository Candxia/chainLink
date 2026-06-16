package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	systemV1 "cl_system/api/system/v1"
	mdlSys "cl_system/internal/model/system"
	"cl_system/internal/service"
)

func (c *Controller) ApiAdd(r *ghttp.Request) {
	var req systemV1.ApiAddReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Api().Add(r.Context(), mdlSys.SysApiAdd{
		Title:  req.Title,
		Path:   req.Path,
		Type:   req.Type,
		Action: req.Action,
		Logger: req.Logger,
	})
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "添加成功"})
}

func (c *Controller) ApiDel(r *ghttp.Request) {
	var req systemV1.ApiDelReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	if err := service.Api().Del(r.Context(), req.Id); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除成功"})
}

func (c *Controller) ApiEdit(r *ghttp.Request) {
	var req systemV1.ApiEditReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Api().Edit(r.Context(), mdlSys.SysApiEdit{
		ApiId: req.Id,
		SysApiAdd: mdlSys.SysApiAdd{
			Title:  req.Title,
			Path:   req.Path,
			Type:   req.Type,
			Action: req.Action,
			Logger: req.Logger,
		},
	})
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) ApiList(r *ghttp.Request) {
	var req systemV1.ApiListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.Api().List(r.Context(), mdlSys.SysApiSearch{
		Title:  req.Title,
		Path:   req.Path,
		Action: req.Method,
		Type:   req.Type,
	})
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total}})
}

func (c *Controller) ApiDrop(r *ghttp.Request) {
	list, err := service.Api().Drop(r.Context())
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": list})
}

func (c *Controller) ApiPath(r *ghttp.Request) {
	list, err := service.Api().Path(r.Context())
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": list})
}
