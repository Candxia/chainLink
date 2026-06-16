// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysAdmin is the golang structure for table sys_admin.
type SysAdmin struct {
	Id        int64  `json:"id"         orm:"id"         description:"编号"`
	Username  string `json:"username"   orm:"username"   description:"用户名"`
	Nickname  string `json:"nickname"   orm:"nickname"   description:"昵称"`
	Password  string `json:"password"   orm:"password"   description:"密码"`
	RoleId    int64  `json:"role_id"    orm:"role_id"    description:"角色"`
	DeptId    int64  `json:"dept_id"    orm:"dept_id"    description:"部门"`
	Avatar    string `json:"avatar"     orm:"avatar"     description:"头像"`
	Remark    string `json:"remark"     orm:"remark"     description:"备注"`
	Status    int    `json:"status"     orm:"status"     description:"账号状态 1=正常,2=冻结,3=注销"`
	Online    int    `json:"online"     orm:"online"     description:"在线状态 2=离线,1=在线"`
	LoginAt   int64  `json:"login_at"   orm:"login_at"   description:"最后登录时间"`
	CreatedBy string `json:"created_by" orm:"created_by" description:"创建者"`
	CreatedAt int64  `json:"created_at" orm:"created_at" description:"创建时间"`
	UpdatedBy string `json:"updated_by" orm:"updated_by" description:"更新者"`
	UpdatedAt int64  `json:"updated_at" orm:"updated_at" description:"更新时间"`
}
