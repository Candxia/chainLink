package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	systemV1 "cl_system/api/system/v1"
	mdlSys "cl_system/internal/model/system"
	"cl_system/internal/service"
)

func (c *Controller) DeptAdd(r *ghttp.Request) {
	var req systemV1.DeptAddReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Dept().Add(r.Context(), mdlSys.SysDeptAdd{
		ParentId: req.ParentId,
		Name:     req.Name,
		OrderNum: req.OrderNum,
		Leader:   req.Leader,
		Phone:    req.Phone,
		Email:    req.Email,
		Status:   req.Status,
	})
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "添加成功"})
}

func (c *Controller) DeptDel(r *ghttp.Request) {
	var req systemV1.DeptDelReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	if err := service.Dept().Del(r.Context(), req.Id); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除成功"})
}

func (c *Controller) DeptEdit(r *ghttp.Request) {
	var req systemV1.DeptEditReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.Dept().Edit(r.Context(), mdlSys.SysDeptEdit{
		Id: req.Id,
		SysDeptAdd: mdlSys.SysDeptAdd{
			ParentId: req.ParentId,
			Name:     req.Name,
			OrderNum: req.OrderNum,
			Leader:   req.Leader,
			Phone:    req.Phone,
			Email:    req.Email,
			Status:   req.Status,
		},
	})
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) DeptInfo(r *ghttp.Request) {
	var req systemV1.DeptInfoReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	info, err := service.Dept().Info(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": info})
}

func (c *Controller) DeptList(r *ghttp.Request) {
	list, err := service.Dept().List(r.Context(), mdlSys.SysDeptSearch{})
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list}})
}

func (c *Controller) DeptDrop(r *ghttp.Request) {
	list, err := service.Dept().Drop(r.Context())
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list}})
}
