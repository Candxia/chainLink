package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 部门管理 ====================

type DeptListReq struct {
	g.Meta `path:"/system/dept/list" method:"get" tags:"系统管理-部门管理" summary:"部门列表"`
}
type DeptListRes struct {
	g.Meta `mime:"application/json"`
	List   []interface{} `json:"list" dc:"部门树"`
}

type DeptInfoReq struct {
	g.Meta `path:"/system/dept/info" method:"get" tags:"系统管理-部门管理" summary:"部门详情"`
	Id     int64 `json:"id" v:"required" dc:"编号"`
}
type DeptInfoRes struct {
	g.Meta `mime:"application/json"`
	Data   interface{} `json:"data" dc:"详情"`
}

type DeptAddReq struct {
	g.Meta   `path:"/system/dept/add" method:"post" tags:"系统管理-部门管理" summary:"添加部门"`
	ParentId int64  `json:"parent_id" dc:"上级部门ID"`
	Name     string `json:"name" v:"required|max-length:64" dc:"部门名称"`
	OrderNum int    `json:"order_num" d:"0" dc:"排序"`
	Leader   string `json:"leader" dc:"负责人"`
	Phone    string `json:"phone" dc:"联系电话"`
	Email    string `json:"email" dc:"邮箱"`
	Status   int    `json:"status" d:"1" v:"in:1,2" dc:"状态 1=正常 2=停用"`
}
type DeptAddRes struct {
	g.Meta `mime:"application/json"`
}

type DeptEditReq struct {
	g.Meta   `path:"/system/dept/edit" method:"put" tags:"系统管理-部门管理" summary:"编辑部门"`
	Id       int64  `json:"id" v:"required" dc:"编号"`
	ParentId int64  `json:"parent_id" dc:"上级部门ID"`
	Name     string `json:"name" v:"required|max-length:64" dc:"部门名称"`
	OrderNum int    `json:"order_num" d:"0" dc:"排序"`
	Leader   string `json:"leader" dc:"负责人"`
	Phone    string `json:"phone" dc:"联系电话"`
	Email    string `json:"email" dc:"邮箱"`
	Status   int    `json:"status" v:"in:1,2" dc:"状态 1=正常 2=停用"`
}
type DeptEditRes struct {
	g.Meta `mime:"application/json"`
}

type DeptDelReq struct {
	g.Meta `path:"/system/dept/del" method:"delete" tags:"系统管理-部门管理" summary:"删除部门"`
	Id     int64 `json:"id" v:"required" dc:"编号"`
}
type DeptDelRes struct {
	g.Meta `mime:"application/json"`
}

type DeptDropdownReq struct {
	g.Meta `path:"/system/dept/dropdown" method:"get" tags:"系统管理-部门管理" summary:"部门下拉"`
}
type DeptDropdownRes struct {
	g.Meta `mime:"application/json"`
	List   []interface{} `json:"list" dc:"下拉列表"`
}
