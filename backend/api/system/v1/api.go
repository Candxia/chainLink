package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 接口管理 ====================

type ApiListReq struct {
	g.Meta   `path:"/api" method:"get" tags:"系统管理-接口管理" summary:"接口列表"`
	Page     int    `json:"page" d:"1" dc:"页码"`
	PageSize int    `json:"pageSize" d:"10" dc:"每页条数"`
	Title    string `json:"title" dc:"标题"`
	Path     string `json:"path" dc:"路径"`
	Method   string `json:"method" dc:"请求方式"`
	Type     string `json:"type" dc:"接口类型"`
}
type ApiListRes struct {
	g.Meta `mime:"application/json"`
	Total  int           `json:"total" dc:"总数"`
	List   []interface{} `json:"list" dc:"列表"`
}

type ApiAddReq struct {
	g.Meta `path:"/api" method:"post" tags:"系统管理-接口管理" summary:"添加接口"`
	Title  string `json:"title" v:"required|max-length:128" dc:"标题"`
	Path   string `json:"path" v:"required|max-length:128" dc:"路径"`
	Type   string `json:"type" d:"BUS" v:"in:BUS,SYS,DEF" dc:"接口类型 BUS=业务 SYS=系统 DEF=自定义"`
	Action string `json:"action" d:"GET" v:"in:GET,POST,PUT,DELETE" dc:"请求方式"`
	Logger int    `json:"logger" d:"1" v:"in:1,2" dc:"记录日志 1=是 2=否"`
}
type ApiAddRes struct {
	g.Meta `mime:"application/json"`
}

type ApiEditReq struct {
	g.Meta `path:"/api/:id" method:"put" tags:"系统管理-接口管理" summary:"编辑接口"`
	Id     int64  `json:"id" v:"required" dc:"主键"`
	Title  string `json:"title" v:"required|max-length:128" dc:"标题"`
	Path   string `json:"path" v:"required|max-length:128" dc:"路径"`
	Type   string `json:"type" v:"in:BUS,SYS,DEF" dc:"接口类型"`
	Action string `json:"action" v:"in:GET,POST,PUT,DELETE" dc:"请求方式"`
	Logger int    `json:"logger" v:"in:1,2" dc:"记录日志"`
}
type ApiEditRes struct {
	g.Meta `mime:"application/json"`
}

type ApiDelReq struct {
	g.Meta `path:"/api/:id" method:"delete" tags:"系统管理-接口管理" summary:"删除接口"`
	Id     int64 `json:"id" v:"required" dc:"主键"`
}
type ApiDelRes struct {
	g.Meta `mime:"application/json"`
}

type ApiDropdownReq struct {
	g.Meta `path:"/api/dropdown" method:"get" tags:"系统管理-接口管理" summary:"接口下拉"`
}
type ApiDropdownRes struct {
	g.Meta `mime:"application/json"`
	List   []map[string]interface{} `json:"list" dc:"下拉列表"`
}

type ApiPathReq struct {
	g.Meta `path:"/api/path" method:"get" tags:"系统管理-接口管理" summary:"接口路径下拉"`
	Path   string `json:"path" dc:"路径关键字"`
}
type ApiPathRes struct {
	g.Meta `mime:"application/json"`
	List   []interface{} `json:"list" dc:"路径列表"`
}
