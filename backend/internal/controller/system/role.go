package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	systemV1 "cl_system/api/system/v1"
	mdlSys "cl_system/internal/model/system"
	"cl_system/internal/service"
)

func (c *Controller) RoleAdd(r *ghttp.Request) {
	var req systemV1.RoleAddReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Role().Add(r.Context(), mdlSys.SysRoleAdd{
		Name:         req.Name,
		Status:       req.Status,
		Remark:       req.Remark,
		MenuAll:      req.MenuAll,
		MenuIds:      req.MenuIds,
		IsShowMobile: req.IsShowMobile,
	})
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "添加成功"})
}

func (c *Controller) RoleDel(r *ghttp.Request) {
	var req systemV1.RoleDelReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	if err := service.Role().Del(r.Context(), req.Id); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除成功"})
}

func (c *Controller) RoleEdit(r *ghttp.Request) {
	var req systemV1.RoleEditReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Role().Edit(r.Context(), mdlSys.SysRoleEdit{
		Id: req.Id,
		SysRoleAdd: mdlSys.SysRoleAdd{
			Name:         req.Name,
			Status:       req.Status,
			Remark:       req.Remark,
			MenuAll:      req.MenuAll,
			MenuIds:      req.MenuIds,
			IsShowMobile: req.IsShowMobile,
		},
	})
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) RoleStatus(r *ghttp.Request) {
	var req systemV1.RoleStatusReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Role().Status(r.Context(), mdlSys.SysRoleStatus{
		Id:     req.Id,
		Status: req.Status,
	})
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "状态更新成功"})
}

func (c *Controller) RoleInfo(r *ghttp.Request) {
	var req systemV1.RoleInfoReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	info, err := service.Role().Info(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": info})
}

func (c *Controller) RoleList(r *ghttp.Request) {
	var req systemV1.RoleListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.Role().List(r.Context(), mdlSys.SysRoleSearch{
		Name:   req.Name,
		Status: req.Status,
	})
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total}})
}

func (c *Controller) RoleDrop(r *ghttp.Request) {
	list, err := service.Role().Drop(r.Context())
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": list})
}

func (c *Controller) RoleIsMobile(r *ghttp.Request) {
	var req systemV1.RoleIsMobileReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Role().IsMobile(r.Context(), mdlSys.SysRoleIsMobile{
		Id:           req.Id,
		IsShowMobile: req.IsShowMobile,
	})
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}
