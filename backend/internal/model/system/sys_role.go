package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"cl_system/internal/consts"
)

type SysRoleSearch struct {
	Name   string `json:"name" description:"角色名称"`
	Status int    `json:"status" v:"in:1,2" description:"状态1启用2禁用"`
}

func (m *SysRoleSearch) Condition() g.Map {
	var where = g.Map{
		"is_del": consts.EnumNot,
	}
	if m.Status > 0 {
		where["status"] = m.Status
	}
	if m.Name != "" {
		where["name"] = m.Name
	}
	return where
}

type SysRoleAdd struct {
	Name         string  `json:"name" v:"required|max-length:32" description:"角色名称"`
	Status       int     `json:"status" d:"1" v:"in:1,2" description:"状态1启用2禁用"`
	Remark       string  `json:"remark" description:"备注"`
	MenuAll      int     `json:"menu_all" d:"2" v:"in:1,2" description:"菜单全选:1是2否"`
	MenuIds      []int64 `json:"menu_ids" description:"菜单ids"`
	IsShowMobile int     `json:"is_show_mobile" d:"2" v:"in:1,2" description:"是否显示手机号 1是2否"`
}

func (m *SysRoleAdd) Valid() bool {
	if m.MenuAll == consts.EnumIs {
		m.MenuIds = make([]int64, 0)
	} else if len(m.MenuIds) == 0 {
		return false
	}
	return true
}

type SysRoleEdit struct {
	SysRoleAdd
	Id        int64  `json:"id" v:"required|min:2" description:"角色id"`
	UpdatedBy string `json:"updated_by"`
}

type SysRoleInfo struct {
	Id           int64   `json:"id" description:"Id"`
	ParentId     int64   `json:"parent_id" description:"父级id"`
	Link         string  `json:"link" description:"链路"`
	Code         string  `json:"code" description:"角色代码"`
	Name         string  `json:"name" description:"角色名称"`
	Status       int     `json:"status" description:"状态1启用2关闭"`
	IsDel        int     `json:"is_del" description:"删除1是2否"`
	IsShowMobile int     `json:"is_show_mobile" description:"是否显示手机号 1是2否"`
	Level        int     `json:"level" description:"层级：默认1级"`
	Remark       string  `json:"remark" description:"备注"`
	MenuAll      int     `json:"menu_all" description:"菜单全选:1是2否"`
	CreatedAt    int64   `json:"created_at" description:"创建时间"`
	CreatedBy    string  `json:"created_by" description:"创建者uname"`
	UpdatedAt    int64   `json:"updated_at" description:"更新时间"`
	UpdatedBy    string  `json:"updated_by" description:"更新者uname"`
	MenuIds      []int64 `json:"menu_ids" description:"菜单信息"`
}

type SysRoleStatus struct {
	Id     int64 `json:"id" v:"required|min:2" description:"角色id"`
	Status int   `json:"status" v:"required|in:1,2" description:"状态1启用2禁用"`
}

type SysRoleIsMobile struct {
	Id           int64 `json:"id" v:"required|min:2" description:"角色id"`
	IsShowMobile int   `json:"is_show_mobile" v:"required|in:1,2" description:"显示手机号:1启用2禁用"`
}

type SysRoleList struct {
	Id           int64  `json:"id" description:"Id"`
	ParentId     int64  `json:"parent_id" description:"父级id"`
	Link         string `json:"link" description:"链路"`
	Code         string `json:"code" description:"角色代码"`
	Name         string `json:"name" description:"角色名称"`
	Status       int    `json:"status" description:"状态1启用2关闭"`
	IsDel        int    `json:"is_del" description:"删除1是2否"`
	IsShowMobile int    `json:"is_show_mobile" description:"是否显示手机号 1是2否"`
	Level        int    `json:"level" description:"层级：默认1级"`
	Remark       string `json:"remark" description:"备注"`
	MenuAll      int    `json:"menu_all" description:"菜单全选:1是2否"`
	CreatedAt    int64  `json:"created_at" description:"创建时间"`
	CreatedBy    string `json:"created_by" description:"创建者uname"`
	UpdatedAt    int64  `json:"updated_at" description:"更新时间"`
	UpdatedBy    string `json:"updated_by" description:"更新者uname"`
}

type SysRoleTree struct {
	SysRoleList
	Children []SysRoleTree `json:"children" description:"下级角色"`
}
