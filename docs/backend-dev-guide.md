# cl_system 后端开发文档

> 项目路径: `D:\CodeByAi\cl_system\backend`
> 工程: cl_system | Framework: GoFrame v2.5.6 | Database: MySQL

---

## 一、项目结构

```
backend/
├── api/                            ← 请求/响应定义 (GoFrame 标准)
│   ├── system/v1/                  ← 系统管理 API
│   │   ├── admin.go                ← 账号管理 Req/Res
│   │   ├── role.go                 ← 角色管理 Req/Res
│   │   ├── menu.go                 ← 菜单管理 Req/Res
│   │   ├── dept.go                 ← 部门管理 Req/Res
│   │   ├── api.go                  ← 接口管理 Req/Res
│   │   └── system.go               ← 仪表盘/配置
│   ├── user/v1/                    ← 用户认证 API
│   ├── trace/v1/                   ← 产品溯源 API
│   ├── supply/v1/                  ← 供应链 API
│   ├── warehouse/v1/               ← 仓库管理 API
│   └── blockchain/v1/              ← 区块链 API
├── hack/
│   └── config.yaml                 ← gf CLI 配置 (gf gen dao)
├── internal/
│   ├── cmd/
│   │   ├── http.go                 ← 启动命令 (gcmd.Command)
│   │   ├── middleware.go           ← CORS / Response / AccessLog / Auth
│   │   ├── router.go               ← 路由注册表
│   │   ├── rules.go                ← 自定义校验规则
│   │   └── sse.go                  ← SSE 推送端点
│   ├── consts/consts.go            ← 系统常量
│   ├── controller/                 ← 控制器 (5 个模块)
│   │   └── system/                 ← 系统管理: admin/api/role/dept/menu
│   ├── dao/                        ← DAO (gf gen dao 输出)
│   ├── logic/                      ← 业务逻辑实现
│   │   └── system/                 ← 系统管理: admin/api/role/dept/menu
│   ├── model/
│   │   ├── model.go                ← 公共模型 (PageInfo)
│   │   ├── sse.go                  ← SSE 推送模型
│   │   ├── system/                 ← 系统管理业务对象 (BO)
│   │   ├── entity/                 ← 数据库实体 (gf gen dao)
│   │   └── do/                     ← Domain Object (gf gen dao)
│   ├── router/                     ← (占位)
│   └── service/
│       ├── service.go              ← 服务接口 + 注册函数
│       ├── middleware.go           ← LogHook + Response 结构体
│       └── export/                 ← CSV 导出框架
├── main.go                         ← 入口 (调用 cmd.Main.Run)
├── manifest/config/config.yaml     ← 数据库/服务配置
├── resource/sql/                   ← 建表 SQL
├── tools/                          ← 工具脚本
├── utility/                        ← 工具函数
│   ├── decimal.go                  ← 高精度运算 (shopspring/decimal)
│   └── password.go                 ← 加盐 HMAC-SHA256 密码
├── go.mod
└── go.sum
```

---

## 二、API 规范 (GoFrame 标准)

### 2.1 请求/响应结构体

所有 API Req/Res 定义在 `api/{module}/v1/` 中，**不嵌入 model 类型**：

```go
// api/system/v1/admin.go — 标准写法
type AdminAddReq struct {
    g.Meta   `path:"/system/admin/add" method:"post" tags:"系统管理-账号管理" summary:"添加管理员"`
    Username string `json:"username" v:"required|length:3,32" dc:"用户名"`
    Nickname string `json:"nickname" v:"required|length:1,20" dc:"昵称"`
    Password string `json:"password" v:"required|length:6,32" dc:"密码"`
    RoleId   int64  `json:"role_id" v:"required" dc:"角色"`
}
type AdminAddRes struct {
    g.Meta `mime:"application/json"`
}
```

| 标签 | 用途 |
|------|------|
| `g.Meta` | 路由绑定 + OpenAPI 文档 |
| `path` | 接口路径 (通过路由组 `/api` 前缀) |
| `method` | HTTP 方法 |
| `tags` | 分组标签 (OpenAPI) |
| `summary` | 接口说明 |
| `v:"required"` | 校验规则 (GoFrame validator) |
| `dc:"说明"` | 字段文档说明 |
| `d:"1"` | 默认值 |
| `mime` | 响应 MIME 类型 |

### 2.2 分层流转

```
客户端请求
    ↓
api/system/v1/XxxReq        ← Request 解析 + 校验
    ↓
controller/system/xxx.go    ← 构造 model 类型 → 调用 service
    ↓
service interface           ← 服务接口 IAdmin / IApi / IRole / IDept / IMenu
    ↓
logic/system/xxx.go         ← 业务逻辑 + DB 操作
    ↓
MySQL 表
```

### 2.3 系统管理 API 路由表

| 模块 | API | Method | 路径 |
|------|-----|--------|------|
| 账号管理 | 列表 | GET | `/api/system/admin/list` |
| | 详情 | GET | `/api/system/admin/info` |
| | 添加 | POST | `/api/system/admin/add` |
| | 编辑 | PUT | `/api/system/admin/edit` |
| | 删除 | DELETE | `/api/system/admin/del` |
| | 密码 | PUT | `/api/system/admin/password` |
| | 状态 | PUT | `/api/system/admin/status` |
| | 踢下线 | PUT | `/api/system/admin/clickout` |
| 角色管理 | 列表 | GET | `/api/system/role/list` |
| | 详情 | GET | `/api/system/role/info` |
| | 添加 | POST | `/api/system/role/add` |
| | 编辑 | PUT | `/api/system/role/edit` |
| | 删除 | DELETE | `/api/system/role/del` |
| | 状态 | PUT | `/api/system/role/status` |
| | 下拉 | GET | `/api/system/role/dropdown` |
| | 手机号 | PUT | `/api/system/role/is_mobile` |
| 菜单管理 | 列表 | GET | `/api/system/menu/list` |
| | 详情 | GET | `/api/system/menu/info` |
| | 添加 | POST | `/api/system/menu/add` |
| | 编辑 | PUT | `/api/system/menu/edit` |
| | 删除 | DELETE | `/api/system/menu/del` |
| | 下拉 | GET | `/api/system/menu/dropdown` |
| | 角色菜单 | GET | `/api/system/menu/role` |
| 部门管理 | 列表 | GET | `/api/system/dept/list` |
| | 详情 | GET | `/api/system/dept/info` |
| | 添加 | POST | `/api/system/dept/add` |
| | 编辑 | PUT | `/api/system/dept/edit` |
| | 删除 | DELETE | `/api/system/dept/del` |
| | 下拉 | GET | `/api/system/dept/dropdown` |
| 接口管理 | 列表 | GET | `/api/system/api/list` |
| | 添加 | POST | `/api/system/api/add` |
| | 编辑 | PUT | `/api/system/api/edit` |
| | 删除 | DELETE | `/api/system/api/del` |
| | 下拉 | GET | `/api/system/api/dropdown` |
| | 路径 | GET | `/api/system/api/path` |

---

## 三、数据库表

| 表 | 说明 | 位置 |
|---|------|------|
| `sys_admin` | 系统管理员 | `resource/sql/system_mgr.sql` |
| `sys_role` | 角色表 | 同上 |
| `sys_menu` | 菜单表 | 同上 |
| `sys_role_menu` | 角色-菜单关联 | 同上 |
| `sys_dept` | 部门表 | 同上 |
| `sys_api` | 接口表 | 同上 |
| `login_log` | 登录日志 | 自动建表 |
| `access_log` | 访问日志 | 自动建表 |
| `api_rule` | API 规则 | 自动建表 |

---

## 四、Controller → Model 构造模式

Controller 从 API Req 解析后，构造 model 类型传入 Service：

```go
func (c *Controller) AdminAdd(r *ghttp.Request) {
    var req systemV1.AdminAddReq
    if err := r.Parse(&req); err != nil {
        r.Response.WriteJson(g.Map{"code": 1, "message": err.Error()})
        return
    }
    // 构造 model 类型 → 调用 service
    err := service.Admin().Add(r.Context(), mdlSys.SysAdminAdd{
        Username: req.Username,
        Nickname: req.Nickname,
        Password: req.Password,
        RoleId:   req.RoleId,
        DeptId:   req.DeptId,
    })
    // ...处理结果
}
```

---

## 五、Service 接口

定义在 `internal/service/service.go`，5 个系统管理接口：

| 接口 | 方法 | 说明 |
|------|------|------|
| `IAdmin` | Add/Edit/Del/Info/List/Password/Status/ClickOut/SetOnline | 账号管理 |
| `IApi` | Add/Edit/Del/List/Drop/Path | 接口管理 |
| `IRole` | Add/Edit/Del/Info/List/Drop/Status/IsMobile | 角色管理 |
| `IDept` | Add/Edit/Del/Info/List/Drop/Exist | 部门管理 |
| `IMenu` | Add/Edit/Del/Info/List/Drop/Role | 菜单管理 |

---

## 六、自定义验证规则

注册在 `internal/cmd/rules.go`：

| 规则 | 说明 |
|------|------|
| `cPwd` | 密码 6-18 位，含字母+数字 |
| `cName` | 管理员用户名 (字母开头，6-32位) |
| `mustDate` | 必填日期，不能早于当天 |
| `limitDate` | 日期范围 (Y-m-d) |
| `limitDateTime` | 日期时间范围 (Y-m-d H:i:s) |
| `lenRid` | 数组长度限制 |

---

## 七、SSE 推送

- 端点: `/api/system/notice`
- 订阅管理: `internal/model/sse.go` (SseUserMap)
- 消息队列: `internal/model/sse.go` (SseQueue)
- 使用 `go model.SseNotice()` 在启动时消费队列
- 通过 `model.SsePush(info)` 推送消息

---

## 八、构建与运行

```bash
# 数据库建表
cd backend
go run tools/init_system_mgr/main.go

# 启动服务
go run main.go
# 默认端口: 8099

# gf CLI 生成 DAO
gf gen dao

# 构建
go build -o cl_system.exe .
```
