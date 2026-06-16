// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysApi is the golang structure for table sys_api.
type SysApi struct {
	ApiId  int64  `json:"api_id" orm:"api_id" description:"主键编码"`
	Title  string `json:"title"  orm:"title"  description:"标题"`
	Path   string `json:"path"   orm:"path"   description:"地址"`
	Type   string `json:"type"   orm:"type"   description:"接口类型(BUS业务接口,SYS系统接口,DEF下拉/选项)"`
	Action string `json:"action" orm:"action" description:"请求方式"`
	Logger int    `json:"logger" orm:"logger" description:"是否记录操作日志(1是,2否)"`
}
