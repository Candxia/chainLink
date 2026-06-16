package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	systemV1 "cl_system/api/system/v1"
	mdlSys "cl_system/internal/model/system"
	"cl_system/internal/service"
)

func (c *Controller) MenuAdd(r *ghttp.Request) {
	var req systemV1.MenuAddReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Menu().Add(r.Context(), mdlSys.SysMenuAdd{
		MenuName:   req.MenuName,
		Title:      req.Title,
		ParentId:   req.ParentId,
		Path:       req.Path,
		Icon:       req.Icon,
		Component:  req.Component,
		Permission: req.Permission,
		OrderNum:   req.OrderNum,
		MenuType:   req.MenuType,
		Visible:    req.Visible,
		Status:     req.Status,
		IsFrame:    req.IsFrame,
		IsCache:    req.IsCache,
	})
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "添加成功"})
}

func (c *Controller) MenuDel(r *ghttp.Request) {
	var req systemV1.MenuDelReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	if err := service.Menu().Del(r.Context(), req.Id); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除成功"})
}

func (c *Controller) MenuEdit(r *ghttp.Request) {
	var req systemV1.MenuEditReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Menu().Edit(r.Context(), mdlSys.SysMenuEdit{
		MenuId: req.MenuId,
		SysMenuAdd: mdlSys.SysMenuAdd{
			MenuName:   req.MenuName,
			Title:      req.Title,
			ParentId:   req.ParentId,
			Path:       req.Path,
			Icon:       req.Icon,
			Component:  req.Component,
			Permission: req.Permission,
			OrderNum:   req.OrderNum,
			MenuType:   req.MenuType,
			Visible:    req.Visible,
			Status:     req.Status,
			IsFrame:    req.IsFrame,
			IsCache:    req.IsCache,
		},
	})
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) MenuInfo(r *ghttp.Request) {
	var req systemV1.MenuInfoReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	info, err := service.Menu().Info(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": info})
}

func (c *Controller) MenuList(r *ghttp.Request) {
	list, err := service.Menu().List(r.Context())
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": list})
}

func (c *Controller) MenuRole(r *ghttp.Request) {
	var req systemV1.MenuRoleReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, err := service.Menu().Role(r.Context())
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": list})
}

func (c *Controller) MenuDrop(r *ghttp.Request) {
	list, total, err := service.Menu().Drop(r.Context(), nil)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total}})
}

func (c *Controller) MenuDropdown(r *ghttp.Request) {
	list, _, err := service.Menu().Drop(r.Context(), nil)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list}})
}
