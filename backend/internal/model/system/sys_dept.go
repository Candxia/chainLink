package system

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SysDeptAdd struct {
	ParentId int64  `json:"parent_id" v:"required"         description:"上级部门"`
	Name     string `json:"name"      v:"required|length:2,32" description:"部门名称"`
	OrderNum int    `json:"order_num" v:"required"         description:"排序"`
	Leader   string `json:"leader"    v:"required|length:2,32" description:"负责人"`
	Phone    string `json:"phone"     description:"联系电话"`
	Email    string `json:"email"     description:"邮箱"`
	Status   int    `json:"status"    v:"required|in:1,2"  description:"状态(1正常,2禁用)"`
}

func (m *SysDeptAdd) Data() g.Map {
	where := g.Map{
		"order_num": m.OrderNum,
		"parent_id": m.ParentId,
		"status":    m.Status,
		"name":      m.Name,
		"leader":    m.Leader,
		"phone":     m.Phone,
		"email":     m.Email,
	}
	return where
}

type SysDeptEdit struct {
	Id int64 `json:"id" v:"required" description:"主键编码"`
	SysDeptAdd
}

func (m *SysDeptEdit) Update() g.Map {
	data := g.Map{}
	if m.ParentId > 0 {
		data["parent_id"] = m.ParentId
	}
	if m.Name != "" {
		data["name"] = m.Name
	}
	if m.OrderNum > 0 {
		data["order_num"] = m.OrderNum
	}
	if m.Leader != "" {
		data["leader"] = m.Leader
	}
	if m.Phone != "" {
		data["phone"] = m.Phone
	}
	if m.Email != "" {
		data["email"] = m.Email
	}
	if m.Status > 0 {
		data["status"] = m.Status
	}
	return data
}

type SysDeptSearch struct {
	Name   string `json:"name"   v:"max-length:32" description:"部门名称"`
	Leader string `json:"leader" v:"max-length:32" description:"负责人"`
	Status int    `json:"status" v:"in:1,2"        description:"状态(1正常,2禁用)"`
}

func (m *SysDeptSearch) Condition() (where g.Map) {
	where = g.Map{}
	if m.Status > 0 {
		where["status"] = m.Status
	}
	if m.Name != "" {
		where["name"] = m.Name
	}
	if m.Leader != "" {
		where["leader"] = m.Leader
	}
	return where
}

type SysDeptInfo struct {
	Id        int64         `json:"id"         description:"Id"`
	ParentId  int64         `json:"parent_id"  description:"上级部门"`
	OrderNum  int           `json:"order_num"  description:"排序"`
	Name      string        `json:"name"       description:"部门名称"`
	Leader    string        `json:"leader"     description:"负责人"`
	Phone     string        `json:"phone"      description:"联系电话"`
	Email     string        `json:"email"      description:"邮箱"`
	Status    int           `json:"status"     description:"状态(1正常,2禁用)"`
	CreatedAt int64         `json:"created_at" description:"创建时间"`
	Children  []SysDeptInfo `json:"children"   description:"下级部门"`
}

type SysDeptExist struct {
	Id       int64  `json:"id"        description:"编号"`
	ParentId int64  `json:"parent_id" description:"上级部门"`
	Name     string `json:"name"      description:"部门名称"`
}

func (m *SysDeptExist) Condition() g.Map {
	where := g.Map{}
	if m.Id > 0 {
		where["id"] = m.Id
	} else {
		where["parent_id"] = m.ParentId
	}
	if m.Name != "" {
		where["name"] = m.Name
	}
	return where
}
