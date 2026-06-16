package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 账号管理 ====================

type AdminListReq struct {
	g.Meta   `path:"/admin" method:"get" tags:"系统管理-账号管理" summary:"管理员列表"`
	Page     int    `json:"page" d:"1" dc:"页码"`
	PageSize int    `json:"pageSize" d:"10" dc:"每页条数"`
	Username string `json:"username" dc:"用户名"`
	RoleId   int64  `json:"role_id" dc:"角色ID"`
	DeptId   int64  `json:"dept_id" dc:"部门ID"`
	Status   int    `json:"status" dc:"状态 1=正常 2=冻结"`
}
type AdminListRes struct {
	g.Meta `mime:"application/json"`
	Total  int           `json:"total" dc:"总数"`
	List   []interface{} `json:"list" dc:"列表"`
}

type AdminInfoReq struct {
	g.Meta `path:"/admin/:id" method:"get" tags:"系统管理-账号管理" summary:"管理员详情"`
	Id     int64 `json:"id" v:"required" dc:"编号"`
}
type AdminInfoRes struct {
	g.Meta `mime:"application/json"`
	Id        int64  `json:"id" dc:"编号"`
	Username  string `json:"username" dc:"用户名"`
	Nickname  string `json:"nickname" dc:"昵称"`
	RoleId    int64  `json:"role_id" dc:"角色ID"`
	RoleName  string `json:"role_name" dc:"角色名称"`
	DeptId    int64  `json:"dept_id" dc:"部门ID"`
	DeptName  string `json:"dept_name" dc:"部门名称"`
	Status    int    `json:"status" dc:"状态"`
	Online    int    `json:"online" dc:"在线状态"`
	LoginAt   int64  `json:"login_at" dc:"最后登录"`
	CreatedBy string `json:"created_by" dc:"创建者"`
	CreatedAt int64  `json:"created_at" dc:"创建时间"`
}

type AdminAddReq struct {
	g.Meta     `path:"/admin" method:"post" tags:"系统管理-账号管理" summary:"添加管理员"`
	Username   string `json:"username" v:"required|length:3,32" dc:"用户名"`
	Nickname   string `json:"nickname" v:"required|length:1,20" dc:"昵称"`
	Password   string `json:"password" v:"required|length:6,32" dc:"密码"`
	RePassword string `json:"re_password" v:"required|same:Password" dc:"确认密码"`
	RoleId     int64  `json:"role_id" v:"required" dc:"角色"`
	DeptId     int64  `json:"dept_id" v:"required" dc:"部门"`
}
type AdminAddRes struct {
	g.Meta `mime:"application/json"`
}

type AdminEditReq struct {
	g.Meta   `path:"/admin/:id" method:"put" tags:"系统管理-账号管理" summary:"编辑管理员"`
	Id       int64  `json:"id" v:"required" dc:"编号"`
	Username string `json:"username" v:"required|length:3,32" dc:"用户名"`
	Nickname string `json:"nickname" v:"required|length:1,20" dc:"昵称"`
	RoleId   int64  `json:"role_id" v:"required" dc:"角色"`
	DeptId   int64  `json:"dept_id" v:"required" dc:"部门"`
	Status   int    `json:"status" v:"in:1,2" dc:"状态 1=正常 2=冻结"`
}
type AdminEditRes struct {
	g.Meta `mime:"application/json"`
}

type AdminDelReq struct {
	g.Meta `path:"/admin/:id" method:"delete" tags:"系统管理-账号管理" summary:"删除管理员"`
	Id     int64 `json:"id" v:"required" dc:"编号"`
}
type AdminDelRes struct {
	g.Meta `mime:"application/json"`
}

type AdminPasswordReq struct {
	g.Meta   `path:"/admin/:id/password" method:"put" tags:"系统管理-账号管理" summary:"修改密码"`
	Id       int64  `json:"id" v:"required" dc:"编号"`
	Password string `json:"password" v:"required|length:6,32" dc:"新密码"`
}
type AdminPasswordRes struct {
	g.Meta `mime:"application/json"`
}

type AdminStatusReq struct {
	g.Meta   `path:"/admin/:id/status" method:"put" tags:"系统管理-账号管理" summary:"修改状态"`
	Id       int64  `json:"id" v:"required" dc:"编号"`
	Status   int    `json:"status" v:"in:1,2,3" dc:"状态 1=正常 2=冻结 3=注销"`
}
type AdminStatusRes struct {
	g.Meta `mime:"application/json"`
}

type AdminClickOutReq struct {
	g.Meta   `path:"/admin/:id/clickout" method:"put" tags:"系统管理-账号管理" summary:"踢下线"`
	Username string `json:"username" v:"required" dc:"用户名"`
}
type AdminClickOutRes struct {
	g.Meta `mime:"application/json"`
}
