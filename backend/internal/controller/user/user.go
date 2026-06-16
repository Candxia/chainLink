package user

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	userV1 "cl_system/api/user/v1"
	"cl_system/internal/service"
)

type Controller struct{}
var Ctl = &Controller{}

func (c *Controller) Login(r *ghttp.Request) {
	var req userV1.LoginReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	result, err := service.User().Login(r.Context(), req.UserLoginInput)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "登录成功", "data": result})
}

func (c *Controller) Register(r *ghttp.Request) {
	var req userV1.RegisterReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.User().Register(r.Context(), req.UserRegisterInput)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "注册成功"})
}

func (c *Controller) Logout(r *ghttp.Request) {
	r.Response.WriteJson(g.Map{"code": 0, "message": "已退出"})
}

func (c *Controller) GetUserInfo(r *ghttp.Request) {
	userId := r.GetCtxVar("userId").Uint()
	userInfo, err := service.User().GetUserInfo(r.Context(), userId)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": userInfo})
}

func (c *Controller) GetUserList(r *ghttp.Request) {
	var req userV1.UserListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	users, total, err := service.User().GetUserList(r.Context(), req.Page, req.PageSize)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": users, "total": total, "page": req.Page, "pageSize": req.PageSize}})
}

func (c *Controller) UpdateUser(r *ghttp.Request) {
	var req userV1.UserUpdateReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.User().UpdateUser(r.Context(), req.Id, req.Data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) DeleteUser(r *ghttp.Request) {
	var req userV1.UserDeleteReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.User().DeleteUser(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除成功"})
}

func (c *Controller) GetRoleList(r *ghttp.Request) {
	roles, err := service.User().GetRoleList(r.Context())
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": roles})
}

func (c *Controller) CreateRole(r *ghttp.Request) {
	var req userV1.RoleCreateReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.User().CreateRole(r.Context(), req.RoleInfo)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "创建成功"})
}

func (c *Controller) UpdateRole(r *ghttp.Request) {
	var req userV1.RoleUpdateReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.User().UpdateRole(r.Context(), req.RoleInfo)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) DeleteRole(r *ghttp.Request) {
	var req userV1.RoleDeleteReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.User().DeleteRole(r.Context(), req.Id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除成功"})
}

func (c *Controller) GetEnterpriseList(r *ghttp.Request) {
	var req userV1.EnterpriseListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	list, total, err := service.User().GetEnterpriseList(r.Context(), req.Page, req.PageSize)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total}})
}

func (c *Controller) CreateEnterprise(r *ghttp.Request) {
	var req userV1.EnterpriseCreateReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.User().CreateEnterprise(r.Context(), req.EnterpriseInfo)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "创建成功"})
}

func (c *Controller) UpdateEnterprise(r *ghttp.Request) {
	var req userV1.EnterpriseUpdateReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.User().UpdateEnterprise(r.Context(), req.EnterpriseInfo)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}
