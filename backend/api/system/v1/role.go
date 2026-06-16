package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 角色管理 ====================

type RoleListReq struct {
	g.Meta   `path:"/role" method:"get" tags:"系统管理-角色管理" summary:"角色列表"`
	Page     int    `json:"page" d:"1" dc:"页码"`
	PageSize int    `json:"pageSize" d:"10" dc:"每页条数"`
	Name     string `json:"name" dc:"角色名称"`
	Status   int    `json:"status" dc:"状态"`
}
type RoleListRes struct {
	g.Meta `mime:"application/json"`
	Total  int           `json:"total" dc:"总数"`
	List   []interface{} `json:"list" dc:"列表"`
}

type RoleInfoReq struct {
	g.Meta `path:"/role/:id" method:"get" tags:"系统管理-角色管理" summary:"角色详情"`
	Id     int64 `json:"id" v:"required" dc:"角色ID"`
}
type RoleInfoRes struct {
	g.Meta `mime:"application/json"`
	Data   interface{} `json:"data" dc:"详情"`
}

type RoleAddReq struct {
	g.Meta       `path:"/role" method:"post" tags:"系统管理-角色管理" summary:"添加角色"`
	Name         string  `json:"name" v:"required|max-length:32" dc:"角色名称"`
	Status       int     `json:"status" d:"1" v:"in:1,2" dc:"状态 1=启用 2=禁用"`
	Remark       string  `json:"remark" dc:"备注"`
	MenuAll      int     `json:"menu_all" d:"2" v:"in:1,2" dc:"菜单全选 1=是 2=否"`
	MenuIds      []int64 `json:"menu_ids" dc:"菜单IDS"`
	IsShowMobile int     `json:"is_show_mobile" d:"2" v:"in:1,2" dc:"显示手机号 1=是 2=否"`
}
type RoleAddRes struct {
	g.Meta `mime:"application/json"`
}

type RoleEditReq struct {
	g.Meta       `path:"/role/:id" method:"put" tags:"系统管理-角色管理" summary:"编辑角色"`
	Id           int64   `json:"id" v:"required|min:2" dc:"角色ID"`
	Name         string  `json:"name" v:"required|max-length:32" dc:"角色名称"`
	Status       int     `json:"status" v:"in:1,2" dc:"状态 1=启用 2=禁用"`
	Remark       string  `json:"remark" dc:"备注"`
	MenuAll      int     `json:"menu_all" v:"in:1,2" dc:"菜单全选 1=是 2=否"`
	MenuIds      []int64 `json:"menu_ids" dc:"菜单IDS"`
	IsShowMobile int     `json:"is_show_mobile" v:"in:1,2" dc:"显示手机号 1=是 2=否"`
}
type RoleEditRes struct {
	g.Meta `mime:"application/json"`
}

type RoleDelReq struct {
	g.Meta `path:"/role/:id" method:"delete" tags:"系统管理-角色管理" summary:"删除角色"`
	Id     int64 `json:"id" v:"required|min:1" dc:"角色ID"`
}
type RoleDelRes struct {
	g.Meta `mime:"application/json"`
}

type RoleStatusReq struct {
	g.Meta `path:"/role/:id/status" method:"put" tags:"系统管理-角色管理" summary:"修改角色状态"`
	Id     int64 `json:"id" v:"required|min:2" dc:"角色ID"`
	Status int   `json:"status" v:"required|in:1,2" dc:"状态 1=启用 2=禁用"`
}
type RoleStatusRes struct {
	g.Meta `mime:"application/json"`
}

type RoleDropdownReq struct {
	g.Meta `path:"/role/dropdown" method:"get" tags:"系统管理-角色管理" summary:"角色下拉"`
}
type RoleDropdownRes struct {
	g.Meta `mime:"application/json"`
	List   []map[string]interface{} `json:"list" dc:"下拉列表"`
}

type RoleIsMobileReq struct {
	g.Meta       `path:"/role/:id/is_mobile" method:"put" tags:"系统管理-角色管理" summary:"是否显示手机号"`
	Id           int64 `json:"id" v:"required|min:2" dc:"角色ID"`
	IsShowMobile int   `json:"is_show_mobile" v:"required|in:1,2" dc:"显示手机号 1=是 2=否"`
}
type RoleIsMobileRes struct {
	g.Meta `mime:"application/json"`
}
