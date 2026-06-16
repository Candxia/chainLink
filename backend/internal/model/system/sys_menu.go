package system

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SysMenuTree struct {
	MenuId    int64         `json:"menu_id" description:"menu_id"`
	Title     string        `json:"title" description:"显示名称"`
	Icon      string        `json:"icon" description:"菜单图标"`
	OrderNum  int           `json:"order_num" description:"排序"`
	Component string        `json:"component" description:"组件路径"`
	Visible   int           `json:"visible" description:"展示状态(1显示,2隐藏)"`
	MenuType  string        `json:"menu_type" description:"菜单类型(M路由,C菜单,T页签,F按钮)"`
	Path      string        `json:"path" description:"路径"`
	CreatedAt int64         `json:"created_at" description:"创建时间"`
	Children  []SysMenuTree `json:"children,omitempty" description:"下级菜单"`
}

type SysMenuMeta struct {
	Title   string `json:"title" description:"显示名称"`
	Icon    string `json:"icon" description:"菜单图标"`
	NoCache bool   `json:"noCache" description:"是否缓存"`
	Hidden  bool   `json:"hidden" description:"是否隐藏"`
}

type SysMenuRole struct {
	Path       string        `json:"path" description:"路径"`
	Component  string        `json:"component" description:"组件路径"`
	Name       string        `json:"name" description:"菜单名称"`
	Permission []string      `json:"permission" description:"包含按钮permission"`
	Tags       []string      `json:"tags" description:"包含页签"`
	Meta       SysMenuMeta   `json:"meta" description:"基础信息"`
	Children   []SysMenuRole `json:"children" description:"下级菜单"`
}

type SysMenuAdd struct {
	MenuName   string  `json:"menu_name" v:"required" description:"菜单名称"`
	Title      string  `json:"title" v:"required" description:"显示名称"`
	Path       string  `json:"path" description:"路径"`
	Icon       string  `json:"icon" description:"菜单图标"`
	ParentId   int64   `json:"parent_id" v:"required" description:"上级菜单id"`
	Component  string  `json:"component" description:"组件路径"`
	Permission string  `json:"permission" description:"权限"`
	OrderNum   int     `json:"order_num" v:"required" description:"排序"`
	MenuType   string  `json:"menu_type" v:"in:M,C,F,T" description:"菜单类型(M路由,C菜单,T页签,F按钮)"`
	Status     int     `json:"status" description:"菜单状态(1正常,2停用)"`
	Visible    int     `json:"visible" v:"in:1,2" description:"展示状态(1显示,2隐藏)"`
	IsFrame    int     `json:"is_frame" v:"in:1,2" description:"是否为外链(1是,2否)"`
	IsCache    int     `json:"is_cache" v:"in:1,2" description:"是否缓存(1是,2否)"`
	ApiId      []int64 `json:"api_id" description:"绑定api"`
}

type SysMenuEdit struct {
	MenuId int64 `json:"menu_id" v:"required" description:"menu_id"`
	SysMenuAdd
}

type SysMenuSearch struct {
	MenuName string `json:"menu_name" v:"max-length:64" description:"菜单名称"`
	Visible  int    `json:"visible" v:"in:1,2" description:"展示状态(1显示,2隐藏)"`
}

func (m *SysMenuSearch) Condition() (where g.Map) {
	where = g.Map{}
	if m.Visible > 0 {
		where["visible"] = m.Visible
	}
	if m.MenuName != "" {
		where["title"] = m.MenuName
	}
	return where
}
