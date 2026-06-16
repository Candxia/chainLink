// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysDept is the golang structure for table sys_dept.
type SysDept struct {
	Id        int64  `json:"id"         orm:"id"         description:"编号"`
	ParentId  int64  `json:"parent_id"  orm:"parent_id"  description:"上级部门id"`
	Name      string `json:"name"       orm:"name"       description:"部门名称"`
	OrderNum  int    `json:"order_num"  orm:"order_num"  description:"排序"`
	Leader    string `json:"leader"     orm:"leader"     description:"负责人"`
	Phone     string `json:"phone"      orm:"phone"      description:"联系电话"`
	Email     string `json:"email"      orm:"email"      description:"邮箱"`
	Status    int    `json:"status"     orm:"status"     description:"状态 1=正常 2=停用"`
	IsDel     int    `json:"is_del"     orm:"is_del"     description:"删除 1=是 2=否"`
	CreatedAt int64  `json:"created_at" orm:"created_at" description:"创建时间"`
	UpdatedAt int64  `json:"updated_at" orm:"updated_at" description:"更新时间"`
}
