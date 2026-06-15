package user

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"cl_system/internal/model"
	"cl_system/internal/service"
)

type Controller struct{}
var Ctl = &Controller{}

func (c *Controller) Login(r *ghttp.Request) {
	var req model.UserLoginInput
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	result, err := service.User().Login(r.Context(), req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "登录成功", "data": result})
}

func (c *Controller) Register(r *ghttp.Request) {
	var req model.UserRegisterInput
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.User().Register(r.Context(), req)
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
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	users, total, err := service.User().GetUserList(r.Context(), page, pageSize)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": users, "total": total, "page": page, "pageSize": pageSize}})
}

func (c *Controller) UpdateUser(r *ghttp.Request) {
	userId := r.GetCtxVar("userId").Uint()
	var data map[string]interface{}
	if err := r.Parse(&data); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.User().UpdateUser(r.Context(), userId, data)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) DeleteUser(r *ghttp.Request) {
	id := r.Get("id").Uint()
	err := service.User().DeleteUser(r.Context(), id)
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
	var role model.RoleInfo
	if err := r.Parse(&role); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.User().CreateRole(r.Context(), role)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "创建成功"})
}

func (c *Controller) UpdateRole(r *ghttp.Request) {
	var role model.RoleInfo
	if err := r.Parse(&role); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.User().UpdateRole(r.Context(), role)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}

func (c *Controller) DeleteRole(r *ghttp.Request) {
	id := r.Get("id").Uint()
	err := service.User().DeleteRole(r.Context(), id)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "删除成功"})
}

func (c *Controller) GetEnterpriseList(r *ghttp.Request) {
	page := r.GetQuery("page", 1).Int()
	pageSize := r.GetQuery("pageSize", 10).Int()
	list, total, err := service.User().GetEnterpriseList(r.Context(), page, pageSize)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": g.Map{"list": list, "total": total}})
}

func (c *Controller) CreateEnterprise(r *ghttp.Request) {
	var ent model.EnterpriseInfo
	if err := r.Parse(&ent); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.User().CreateEnterprise(r.Context(), ent)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "创建成功"})
}

func (c *Controller) UpdateEnterprise(r *ghttp.Request) {
	var ent model.EnterpriseInfo
	if err := r.Parse(&ent); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	err := service.User().UpdateEnterprise(r.Context(), ent)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "message": "更新成功"})
}
