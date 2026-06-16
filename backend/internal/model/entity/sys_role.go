// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	Id           int64  `json:"id"             orm:"id"             description:"Id"`
	ParentId     int64  `json:"parent_id"      orm:"parent_id"      description:"父级id"`
	Link         string `json:"link"           orm:"link"           description:"链路"`
	Code         string `json:"code"           orm:"code"           description:"角色代码"`
	Name         string `json:"name"           orm:"name"           description:"角色名称"`
	Status       int    `json:"status"         orm:"status"         description:"状态1启用2关闭"`
	IsDel        int    `json:"is_del"         orm:"is_del"         description:"删除1是2否"`
	IsShowMobile int    `json:"is_show_mobile" orm:"is_show_mobile" description:"是否显示手机号 1是2否"`
	Level        int    `json:"level"          orm:"level"          description:"层级：默认1级"`
	Remark       string `json:"remark"         orm:"remark"         description:"备注"`
	MenuAll      int    `json:"menu_all"       orm:"menu_all"       description:"菜单全选:1是2否"`
	CreatedAt    int64  `json:"created_at"     orm:"created_at"     description:"创建时间"`
	CreatedBy    string `json:"created_by"     orm:"created_by"     description:"创建者uname"`
	UpdatedAt    int64  `json:"updated_at"     orm:"updated_at"     description:"更新时间"`
	UpdatedBy    string `json:"updated_by"     orm:"updated_by"     description:"更新者uname"`
}
