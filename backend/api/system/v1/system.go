package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DashboardReq struct {
	g.Meta `path:"/system/dashboard" method:"get" tags:"系统管理" summary:"仪表盘数据"`
}
type DashboardRes struct {
	g.Meta `mime:"application/json"`
	Data   interface{} `json:"data"`
}

type GetConfigReq struct {
	g.Meta `path:"/system/config" method:"get" tags:"系统管理" summary:"配置列表"`
}
type GetConfigRes struct {
	g.Meta `mime:"application/json"`
	Data   interface{} `json:"data"`
}

type UpdateConfigReq struct {
	g.Meta `path:"/system/config" method:"put" tags:"系统管理" summary:"更新配置"`
	Data   map[string]interface{} `json:"data"`
}
type UpdateConfigRes struct {
	g.Meta `mime:"application/json"`
}

type GetLogListReq struct {
	g.Meta   `path:"/system/logs" method:"get" tags:"系统管理" summary:"日志列表"`
	Page     int    `json:"page" d:"1"`
	PageSize int    `json:"pageSize" d:"20"`
	Type     string `json:"type"`
}
type GetLogListRes struct {
	g.Meta   `mime:"application/json"`
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
}

type DeleteLogReq struct {
	g.Meta `path:"/system/logs/:id" method:"delete" tags:"系统管理" summary:"删除日志"`
	Id     uint `json:"id"`
}
type DeleteLogRes struct {
	g.Meta `mime:"application/json"`
}

type GetNotificationListReq struct {
	g.Meta `path:"/system/notifications" method:"get" tags:"系统管理" summary:"通知列表"`
}
type GetNotificationListRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"list"`
}

type MarkNotificationReadReq struct {
	g.Meta `path:"/system/notification/read/:id" method:"post" tags:"系统管理" summary:"标记通知已读"`
	Id     uint `json:"id"`
}
type MarkNotificationReadRes struct {
	g.Meta `mime:"application/json"`
}
