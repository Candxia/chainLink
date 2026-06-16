package v1

import (
	"cl_system/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 登录 ====================

type LoginReq struct {
	g.Meta `path:"/user/login" method:"post" tags:"用户认证" summary:"登录"`
	model.UserLoginInput
}
type LoginRes struct {
	g.Meta `mime:"application/json"`
	*model.UserLoginOutput
}

type RegisterReq struct {
	g.Meta `path:"/user/register" method:"post" tags:"用户认证" summary:"注册"`
	model.UserRegisterInput
}
type RegisterRes struct {
	g.Meta `mime:"application/json"`
}

type LogoutReq struct {
	g.Meta `path:"/user/logout" method:"post" tags:"用户认证" summary:"退出"`
}
type LogoutRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 用户管理 ====================

type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"用户管理" summary:"用户详情"`
}
type UserInfoRes struct {
	g.Meta `mime:"application/json"`
	*model.UserInfo
}

type UserListReq struct {
	g.Meta   `path:"/user/list" method:"get" tags:"用户管理" summary:"用户列表"`
	Page     int `json:"page" d:"1"`
	PageSize int `json:"pageSize" d:"10"`
}
type UserListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.UserInfo `json:"list"`
	Total  int               `json:"total"`
	Page   int               `json:"page"`
	PageSize int             `json:"pageSize"`
}

type UserUpdateReq struct {
	g.Meta `path:"/user/update" method:"put" tags:"用户管理" summary:"更新用户"`
	Id     uint                   `json:"id" v:"required"`
	Data   map[string]interface{} `json:"data"`
}
type UserUpdateRes struct {
	g.Meta `mime:"application/json"`
}

type UserDeleteReq struct {
	g.Meta `path:"/user/delete/:id" method:"delete" tags:"用户管理" summary:"删除用户"`
	Id     uint `json:"id"`
}
type UserDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 角色管理 ====================

type RoleListReq struct {
	g.Meta `path:"/user/roles" method:"get" tags:"角色管理" summary:"角色列表"`
}
type RoleListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.RoleInfo `json:"list"`
}

type RoleCreateReq struct {
	g.Meta `path:"/user/role" method:"post" tags:"角色管理" summary:"创建角色"`
	model.RoleInfo
}
type RoleCreateRes struct {
	g.Meta `mime:"application/json"`
}

type RoleUpdateReq struct {
	g.Meta `path:"/user/role" method:"put" tags:"角色管理" summary:"更新角色"`
	model.RoleInfo
}
type RoleUpdateRes struct {
	g.Meta `mime:"application/json"`
}

type RoleDeleteReq struct {
	g.Meta `path:"/user/role/:id" method:"delete" tags:"角色管理" summary:"删除角色"`
	Id     uint `json:"id"`
}
type RoleDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 企业管理 ====================

type EnterpriseListReq struct {
	g.Meta   `path:"/user/enterprise" method:"get" tags:"企业管理" summary:"企业列表"`
	Page     int `json:"page" d:"1"`
	PageSize int `json:"pageSize" d:"10"`
}
type EnterpriseListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.EnterpriseInfo `json:"list"`
	Total  int                     `json:"total"`
}

type EnterpriseCreateReq struct {
	g.Meta `path:"/user/enterprise" method:"post" tags:"企业管理" summary:"创建企业"`
	model.EnterpriseInfo
}
type EnterpriseCreateRes struct {
	g.Meta `mime:"application/json"`
}

type EnterpriseUpdateReq struct {
	g.Meta `path:"/user/enterprise" method:"put" tags:"企业管理" summary:"更新企业"`
	model.EnterpriseInfo
}
type EnterpriseUpdateRes struct {
	g.Meta `mime:"application/json"`
}
