package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"cl_system/internal/consts"
	"cl_system/internal/model"
)

type SysAdminAdd struct {
	Username   string `json:"username" v:"required|length:3,32" description:"用户名"`
	Nickname   string `json:"nickname" v:"required|length:1,20" description:"昵称"`
	Password   string `json:"password" v:"required|length:6,32" description:"密码"`
	RePassword string `json:"re_password" v:"required|same:Password" description:"确认密码"`
	RoleId     int64  `json:"role_id" v:"required" description:"角色"`
	DeptId     int64  `json:"dept_id" v:"required" description:"部门"`
}

type SysAdminEdit struct {
	Id       int64  `json:"id" v:"required" description:"编号"`
	Username string `json:"username" v:"required|length:3,32" description:"用户名"`
	Nickname string `json:"nickname" v:"required|length:1,20" description:"昵称"`
	RoleId   int64  `json:"role_id" v:"required" description:"角色"`
	DeptId   int64  `json:"dept_id" v:"required" description:"部门"`
	Status   int    `json:"status" v:"required|in:1,2" description:"账号状态 1=正常,2=冻结"`
}

func (m *SysAdminEdit) Update() (data g.Map) {
	data = g.Map{}
	if m.RoleId > 0 {
		data["role_id"] = m.RoleId
	}
	if m.DeptId > 0 {
		data["dept_id"] = m.DeptId
	}
	if m.Status > 0 {
		data["status"] = m.Status
		if m.Status == consts.OnDisabled {
			data["online"] = consts.IsOffline
		}
	}
	if m.Nickname != "" {
		data["nickname"] = m.Nickname
	}
	return data
}

type SysAdminInfo struct {
	Id        int64  `json:"id" description:"编号"`
	Username  string `json:"username" description:"用户名"`
	Nickname  string `json:"nickname" description:"昵称"`
	RoleId    int64  `json:"role_id" description:"角色id"`
	DeptId    int64  `json:"dept_id" description:"部门id"`
	RoleName  string `json:"role_name" description:"角色名称"`
	DeptName  string `json:"dept_name" description:"部门名称"`
	Remark    string `json:"remark" description:"auth_key"`
	Status    int    `json:"status" description:"账号状态 1=正常,2=冻结,3=注销"`
	Online    int    `json:"online" description:"在线状态 2=离线,1=在线"`
	CreatedBy string `json:"created_by" description:"创建者uname"`
	CreatedAt int64  `json:"created_at" description:"创建时间"`
	LoginAt   int64  `json:"login_at" description:"上次登录时间"`
}

type SysAdminSearch struct {
	model.PageInfo
	Username string `json:"username" v:"max-length:32" description:"用户名"`
	RoleId   int64  `json:"role_id" description:"角色"`
	DeptId   int64  `json:"dept_id" description:"部门"`
	Status   int    `json:"status" v:"in:1,2,3" description:"账号状态 1=正常,2=冻结,3=注销"`
	Online   int    `json:"online" v:"in:1,2" description:"在线状态 2=离线,1=在线"`
}

func (m *SysAdminSearch) Condition() (where g.Map) {
	where = g.Map{
		"id >": consts.SuperAdminId,
	}
	if m.Username != "" {
		where["username"] = m.Username
	}
	if m.DeptId > 0 {
		where["dept_id"] = m.DeptId
	}
	if m.Status > 0 {
		where["status"] = m.Status
	}
	if m.Online > 0 {
		where["online"] = m.Online
	}
	return where
}

type SysAdminOnline struct {
	Id       int64  `json:"id" v:"required" description:"编号"`
	Username string `json:"username" description:"用户名"`
	AuthKey  string `json:"auth_key"` // 随机码
	LoginAt  int64  `json:"login_at" orm:"login_at" description:"登录时间戳"`
}

func (m *SysAdminOnline) Condition() (where g.Map) {
	where = g.Map{}
	if m.Username != "" {
		where["username"] = m.Username
	}
	if m.Id > 0 {
		where["id"] = m.Id
	}
	return where
}
