package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 菜单管理 ====================

type MenuListReq struct {
	g.Meta `path:"/system/menu/list" method:"get" tags:"系统管理-菜单管理" summary:"菜单列表"`
	Title  string `json:"title" dc:"菜单名称"`
	Status int    `json:"status" dc:"状态"`
}
type MenuListRes struct {
	g.Meta `mime:"application/json"`
	List   []interface{} `json:"list" dc:"菜单树"`
}

type MenuInfoReq struct {
	g.Meta `path:"/system/menu/info" method:"get" tags:"系统管理-菜单管理" summary:"菜单详情"`
	Id     int64 `json:"id" v:"required" dc:"菜单ID"`
}
type MenuInfoRes struct {
	g.Meta `mime:"application/json"`
	Data   interface{} `json:"data" dc:"详情"`
}

type MenuAddReq struct {
	g.Meta     `path:"/system/menu/add" method:"post" tags:"系统管理-菜单管理" summary:"添加菜单"`
	MenuName   string `json:"menu_name" v:"required|max-length:64" dc:"菜单名称"`
	Title      string `json:"title" v:"required|max-length:64" dc:"显示名称"`
	ParentId   int64  `json:"parent_id" dc:"上级菜单ID"`
	Path       string `json:"path" dc:"视图地址"`
	Icon       string `json:"icon" dc:"图标"`
	Component  string `json:"component" dc:"组件路径"`
	Permission string `json:"permission" dc:"权限标识"`
	OrderNum   int    `json:"order_num" d:"0" dc:"排序"`
	MenuType   string `json:"menu_type" d:"M" v:"in:M,C,F" dc:"菜单类型 M=路由 C=菜单 F=按钮"`
	Visible    int    `json:"visible" d:"1" v:"in:1,2" dc:"展示状态 1=显示 2=隐藏"`
	Status     int    `json:"status" d:"1" v:"in:1,2" dc:"菜单状态 1=正常 2=停用"`
	IsFrame    int    `json:"is_frame" d:"2" v:"in:1,2" dc:"是否外链 1=是 2=否"`
	IsCache    int    `json:"is_cache" d:"1" v:"in:1,2" dc:"是否缓存 1=是 2=否"`
}
type MenuAddRes struct {
	g.Meta `mime:"application/json"`
}

type MenuEditReq struct {
	g.Meta     `path:"/system/menu/edit" method:"put" tags:"系统管理-菜单管理" summary:"编辑菜单"`
	MenuId     int64  `json:"menu_id" v:"required" dc:"菜单编号"`
	MenuName   string `json:"menu_name" v:"required|max-length:64" dc:"菜单名称"`
	Title      string `json:"title" v:"required|max-length:64" dc:"显示名称"`
	ParentId   int64  `json:"parent_id" dc:"上级菜单ID"`
	Path       string `json:"path" dc:"视图地址"`
	Icon       string `json:"icon" dc:"图标"`
	Component  string `json:"component" dc:"组件路径"`
	Permission string `json:"permission" dc:"权限标识"`
	OrderNum   int    `json:"order_num" d:"0" dc:"排序"`
	MenuType   string `json:"menu_type" v:"in:M,C,F" dc:"菜单类型"`
	Visible    int    `json:"visible" v:"in:1,2" dc:"展示状态"`
	Status     int    `json:"status" v:"in:1,2" dc:"菜单状态"`
	IsFrame    int    `json:"is_frame" v:"in:1,2" dc:"是否外链"`
	IsCache    int    `json:"is_cache" v:"in:1,2" dc:"是否缓存"`
}
type MenuEditRes struct {
	g.Meta `mime:"application/json"`
}

type MenuDelReq struct {
	g.Meta `path:"/system/menu/del" method:"delete" tags:"系统管理-菜单管理" summary:"删除菜单"`
	Id     int64 `json:"id" v:"required" dc:"菜单ID"`
}
type MenuDelRes struct {
	g.Meta `mime:"application/json"`
}

type MenuDropdownReq struct {
	g.Meta `path:"/system/menu/dropdown" method:"get" tags:"系统管理-菜单管理" summary:"菜单下拉"`
}
type MenuDropdownRes struct {
	g.Meta `mime:"application/json"`
	List   []interface{} `json:"list" dc:"下拉列表"`
}

type MenuRoleReq struct {
	g.Meta `path:"/system/menu/role" method:"get" tags:"系统管理-菜单管理" summary:"角色菜单权限"`
	RoleId int64 `json:"role_id" v:"required" dc:"角色ID"`
}
type MenuRoleRes struct {
	g.Meta `mime:"application/json"`
	List   []interface{} `json:"list" dc:"菜单列表"`
}
