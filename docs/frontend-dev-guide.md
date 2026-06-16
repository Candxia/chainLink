# cl_system 前端开发文档

> 项目路径: `D:\CodeByAi\cl_system\frontend`
> 技术栈: Vue 3.5 + Vite 5.3 + Element Plus 2.9 + Pinia + Vue Router 4 (Hash 模式)
> 构建: `vite build` 1754 modules · 15.88s · 零错误

---

## 一、项目结构

```
frontend/
├── src/
│   ├── main.js                    ← 入口
│   ├── App.vue                    ← 根组件
│   ├── api/                       ← API 请求 (按菜单层级组织)
│   │   ├── index.js               ← Axios 实例 + 拦截器
│   │   ├── system/
│   │   │   ├── index.js           ← 所有系统管理 API (统一导出)
│   │   │   ├── admin/             ← 旧目录 (保留兼容)
│   │   │   ├── apiRule/           ← 旧目录
│   │   │   ├── role/              ← 旧目录
│   │   │   ├── dept/              ← 旧目录
│   │   │   ├── menu/              ← 旧目录
│   │   │   └── tools/             ← 开发工具
│   │   ├── login/index.js         ← 登录/登出 API
│   │   ├── personal/index.js      ← 个人中心 API
│   │   ├── trace/index.js         ← 产品溯源 API
│   │   ├── supply/index.js        ← 供应链 API
│   │   ├── blockchain/index.js    ← 区块链 API
│   │   ├── inventory/index.js     ← 库存 API
│   │   ├── record/index.js        ← 操作记录 API
│   │   ├── user/index.js          ← 用户管理 API
│   │   ├── warehouse/index.js     ← 仓库 API
│   │   └── common/index.js        ← 公共 API
│   ├── router/
│   │   └── index.js               ← 路由表 (constant + async)
│   ├── stores/                    ← Pinia 状态
│   │   ├── index.js               ← Pinia 实例
│   │   ├── user.js                ← 用户/Token/权限
│   │   ├── app.js                 ← 布局/主题
│   │   ├── permission.js          ← 路由权限
│   │   ├── locale.js              ← 国际化
│   │   ├── tagsView.js            ← 多页签
│   │   └── drop.js / zone.js      ← 缓存
│   ├── layout/
│   │   └── Layout.vue             ← 主布局 (侧边栏+头部+内容区)
│   ├── views/
│   │   ├── Login/Login.vue        ← 登录页
│   │   ├── Dashboard.vue          ← 仪表盘
│   │   ├── Personal/              ← 个人中心
│   │   ├── Error/                 ← 404/403/500
│   │   └── system/                ← 系统管理 (5 个模块)
│   │       ├── admin/SysAdmin.vue ← 账号管理
│   │       ├── role/SysRole.vue   ← 角色管理
│   │       ├── dept/SysDept.vue   ← 部门管理
│   │       ├── api/SysApi.vue     ← 接口管理
│   │       └── menu/SysMenu.vue   ← 菜单管理
│   ├── hooks/web/                 ← 组合式函数
│   ├── utils/                     ← 工具函数
│   ├── constants/index.js         ← 全局常量
│   └── styles/                    ← 全局样式
├── vite.config.js                 ← Vite 配置
├── index.html
├── package.json
└── scripts/                       ← 工具脚本
```

---

## 二、路由架构

### 2.1 路由表

| 类型 | 路由 | 页面 | 说明 |
|------|------|------|------|
| 常量 | `/login` | Login | 登录页 |
| 常量 | `/personal/personal-center` | PersonalCenter | 个人中心 |
| 常量 | `/redirect/:path` | Redirect | 重定向中转 |
| 常量 | `/404` | 404 | 404 页面 |
| 动态 | `/dashboard` | Dashboard | 仪表盘 |
| 动态 | `/system/admin` | SysAdmin | 账号管理 |
| 动态 | `/system/api` | SysApi | 接口管理 |
| 动态 | `/system/role` | SysRole | 角色管理 |
| 动态 | `/system/dept` | SysDept | 部门管理 |
| 动态 | `/system/menu` | SysMenu | 菜单管理 |

### 2.2 路由守卫 (`router.beforeEach`)

```
请求进入
  ↓
Token 检查
  ├── 未登录 → 重定向 /login
  └── 已登录 → 继续
       ↓
动态路由加载 (permissionStore.generateRoutes)
  ├── 未加载 → 用 asyncRouterMap 动态添加
  └── 已加载 → 跳过
       ↓
目标页面
```

---

## 三、API 调用模式

### 3.1 请求封装

```javascript
// api/index.js — Axios 实例 + 拦截器
import axios from 'axios'
import { TOKEN_KEY, REQUEST_TIMEOUT } from '@/constants'
import { getStorage } from '@/utils/storage'

const service = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api',
  timeout: REQUEST_TIMEOUT,
  headers: { 'Content-Type': 'application/json' },
})
// 请求拦截器: Token 注入
// 响应拦截器: code===0 → 直接返回 data
export default service
```

### 3.2 系统管理 API (统一入口)

```javascript
// api/system/index.js — 所有系统管理 API
import service from '@/api/index'

export function getAdminListApi(params) {
  return service({ url: '/system/admin/list', method: 'get', params })
}
export function addAdminApi(data) {
  return service({ url: '/system/admin/add', method: 'post', data })
}
export function editAdminApi(data) {
  return service({ url: '/system/admin/edit', method: 'put', data })
}
export function delAdminApi(data) {
  return service({ url: '/system/admin/del', method: 'delete', data })
}
// ... role/dept/menu/api — 同上模式
```

### 3.3 页面导入方式

```javascript
// 5 个系统页面统一从 @/api/system 导入
import { getAdminListApi, addAdminApi, editAdminApi, delAdminApi } from '@/api/system'

// 登录从 @/api/login
import { loginApi } from '@/api/login'
```

---

## 四、页面清单

| 页面 | 路由 | 功能 | 状态 |
|------|------|------|------|
| 登录 | `/login` | 用户名+密码登录 | ✅ |
| 仪表盘 | `/dashboard` | 统计卡片+快捷操作 | ✅ |
| 个人中心 | `/personal` | 用户基本信息 | ✅ |
| 账号管理 | `/system/admin` | 搜索/列表/CRUD/密码/状态/踢下线 | ✅ |
| 接口管理 | `/system/api` | 搜索/列表/CRUD | ✅ |
| 角色管理 | `/system/role` | 搜索/列表/CRUD | ✅ |
| 部门管理 | `/system/dept` | 搜索/列表/CRUD | ✅ |
| 菜单管理 | `/system/menu` | 搜索/列表/CRUD | ✅ |
| 404/403/500 | `/404` | 错误页面 | ✅ |

---

## 五、Layout 布局

```
┌─────────────────────────────────────┐
│  Logo                   用户下拉     │
├─────┬───────────────────────────────┤
│     │  面包屑                        │
│ 侧  ├───────────────────────────────┤
│ 边  │                               │
│ 菜  │   Router View                 │
│ 单  │   (页面内容区)                  │
│     │                               │
│     │                               │
└─────┴───────────────────────────────┘
```

- 侧边栏: 从 `asyncRouterMap` 动态渲染菜单
- 可折叠 (collapse)
- 面包屑: 从 `route.path` 自动生成

---

## 六、构建与运行

```bash
cd frontend
npm install
npm run dev        # 开发 → 端口 3000
npm run build      # 生产构建 → dist/
```

### Vite 代理

```javascript
// vite.config.js
server: {
  proxy: {
    '/api': { target: 'http://127.0.0.1:8099', changeOrigin: true }
  }
}
```
