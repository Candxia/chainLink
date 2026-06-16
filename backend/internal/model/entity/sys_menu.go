// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysMenu is the golang structure for table sys_menu.
type SysMenu struct {
	Id        int64   `json:"id"         orm:"id"         description:"show id"`
	MenuId    int64   `json:"menu_id"    orm:"menu_id"    description:"menu_id"`
	ApiId     []int64 `json:"api_id"     orm:"api_id"     description:"关联api_id,json"`
	MenuName  string  `json:"menu_name"  orm:"menu_name"  description:"菜单名称"`
	Title     string  `json:"title"      orm:"title"      description:"显示名称"`
	Icon      string  `json:"icon"       orm:"icon"       description:"菜单图标"`
	ParentId  int64   `json:"parent_id"  orm:"parent_id"  description:"上级菜单id"`
	Component string  `json:"component"  orm:"component"  description:"组件路径"`
	Path      string  `json:"path"       orm:"path"       description:"视图地址"`
	OrderNum  int     `json:"order_num"  orm:"order_num"  description:"排序"`
	MenuType  string  `json:"menu_type"  orm:"menu_type"  description:"菜单类型(M路由,C菜单,T页签,F按钮)"`
	Visible   int     `json:"visible"    orm:"visible"    description:"展示状态(1显示,2隐藏,3超管)"`
	Status    int     `json:"status"     orm:"status"     description:"菜单状态(1正常,2停用)"`
	IsCache   int     `json:"is_cache"   orm:"is_cache"   description:"是否缓存(1是,2否)"`
	UpdatedBy string  `json:"updated_by" orm:"updated_by" description:"更新者uname"`
	CreatedBy string  `json:"created_by" orm:"created_by" description:"创建者uname"`
	CreatedAt int64   `json:"created_at" orm:"created_at" description:"创建时间"`
	UpdatedAt int64   `json:"updated_at" orm:"updated_at" description:"更新时间"`
}
