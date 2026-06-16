package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	systemV1 "cl_system/api/system/v1"
	"cl_system/internal/model"
	mdlSys "cl_system/internal/model/system"
	"cl_system/internal/service"
)

func (c *Controller) AdminList(r *ghttp.Request) {
	var req systemV1.AdminListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	search := mdlSys.SysAdminSearch{
		PageInfo: model.PageInfo{Page: req.Page},
		Username: req.Username,
		RoleId:   req.RoleId,
		DeptId:   req.DeptId,
		Status:   req.Status,
	}
	list, total, err := service.Admin().List(r.Context(), search)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total}})
}

func (c *Controller) AdminInfo(r *ghttp.Request) {
	var req systemV1.AdminInfoReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	info, err := service.Admin().Info(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": info})
}

func (c *Controller) AdminAdd(r *ghttp.Request) {
	var req systemV1.AdminAddReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Admin().Add(r.Context(), mdlSys.SysAdminAdd{
		Username:   req.Username,
		Nickname:   req.Nickname,
		Password:   req.Password,
		RePassword: req.RePassword,
		RoleId:     req.RoleId,
		DeptId:     req.DeptId,
	})
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "添加成功"})
}

func (c *Controller) AdminEdit(r *ghttp.Request) {
	var req systemV1.AdminEditReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Admin().Edit(r.Context(), mdlSys.SysAdminEdit{
		Id:       req.Id,
		Username: req.Username,
		Nickname: req.Nickname,
		RoleId:   req.RoleId,
		DeptId:   req.DeptId,
		Status:   req.Status,
	})
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) AdminDel(r *ghttp.Request) {
	var req systemV1.AdminDelReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	if err := service.Admin().Del(r.Context(), req.Id); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除成功"})
}

func (c *Controller) AdminPassword(r *ghttp.Request) {
	var req systemV1.AdminPasswordReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	if err := service.Admin().Password(r.Context(), req.Id, req.Password); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "密码修改成功"})
}

func (c *Controller) AdminStatus(r *ghttp.Request) {
	var req systemV1.AdminStatusReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	if err := service.Admin().Status(r.Context(), "", req.Status); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "状态修改成功"})
}

func (c *Controller) AdminClickOut(r *ghttp.Request) {
	var req systemV1.AdminClickOutReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	if err := service.Admin().ClickOut(r.Context(), req.Username); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "已踢下线"})
}
