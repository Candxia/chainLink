package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"cl_system/internal/model"
)

type SysApiAdd struct {
	Title  string `json:"title" v:"required|length:2,32" description:"标题"`
	Path   string `json:"path" v:"required|length:2,64" description:"地址"`
	Type   string `json:"type" v:"required|in:BUS,SYS" description:"接口类型"`
	Logger int    `json:"logger" v:"required|in:1,2" description:"是否记录操作日志(1是,2否)"`
	Action string `json:"action" v:"required|in:GET,POST,PUT,DELETE" description:"请求方式"`
}

func (m *SysApiAdd) Data() g.Map {
	return g.Map{
		"title":  m.Title,
		"path":   m.Path,
		"type":   m.Type,
		"action": m.Action,
		"logger": m.Logger,
	}
}

type SysApiEdit struct {
	ApiId int64 `json:"api_id" v:"required" description:"主键编码"`
	SysApiAdd
}

type SysApiSearch struct {
	model.PageInfo
	Title  string `json:"title" v:"max-length:32" description:"标题"`
	Path   string `json:"path" v:"max-length:64" description:"地址"`
	Action string `json:"action" v:"in:GET,POST,PUT,DELETE" description:"请求方式"`
	Type   string `json:"type" v:"in:BUS,SYS,DEF" description:"接口类型"`
}

func (m *SysApiSearch) Condition() (where g.Map) {
	where = g.Map{}
	if m.Title != "" {
		where["title"] = m.Title
	}
	if m.Type != "" {
		where["type"] = m.Type
	}
	if m.Path != "" {
		where["path"] = m.Path
	}
	if m.Action != "" {
		where["action"] = m.Action
	}
	return where
}
