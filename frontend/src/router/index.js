import { createRouter, createWebHashHistory } from 'vue-router'
import { Layout } from '@/utils/routerHelper'
import { NO_RESET_WHITE_LIST } from '@/constants'

export const constantRouterMap = [
  {
    path: '/',
    component: Layout,
    redirect: '/login',
    name: 'Root',
    meta: { hidden: true }
  },
  {
    path: '/redirect',
    component: Layout,
    name: 'RedirectLayout',
    children: [
      { path: '/redirect/:path(.*)', name: 'Redirect', component: () => import('@/views/Redirect.vue'), meta: {} }
    ],
    meta: { hidden: true, noTagsView: true }
  },
  {
    path: '/login',
    component: () => import('@/views/Login/Login.vue'),
    name: 'Login',
    meta: { hidden: true, title: 'Login', noTagsView: true }
  },
  {
    path: '/404',
    component: () => import('@/views/Error/404.vue'),
    name: 'NoFind',
    meta: { hidden: true, title: '404', noTagsView: true }
  }
]

export const asyncRouterMap = [
  {
    path: '/dashboard',
    component: Layout,
    name: 'Dashboard',
    redirect: '/dashboard/index',
    meta: { title: 'Dashboard', icon: 'DataBoard', noCache: false, hidden: false },
    children: [
      { path: 'index', component: () => import('@/views/Dashboard.vue'), name: 'DashboardIndex', meta: { title: 'Dashboard' } }
    ]
  },
  {
    path: '/system',
    component: Layout,
    name: 'System',
    redirect: '/system/admin',
    permission: [],
    meta: { title: 'System', icon: '', noCache: false, hidden: false },
    children: [
      { path: 'admin', component: () => import('@/views/System/admin/SysAdmin.vue'), name: 'SysAdmin', permission: ['SysAdminAdd', 'SysAdminDel', 'SysAdminEdit', 'SysAdminPassword', 'SysAdminFreeze', 'SysAdminList', 'SysAdminClickOut'], meta: { title: 'Admin', icon: 'user', noCache: false, hidden: false } },
      { path: 'api', component: () => import('@/views/System/api/SysApi.vue'), name: 'SysApi', permission: ['SysApiAdd', 'SysApiDel', 'SysApiEdit', 'SysApiList'], meta: { title: 'API', icon: 'api', noCache: false, hidden: false } },
      { path: 'role', component: () => import('@/views/System/role/SysRole.vue'), name: 'SysRole', permission: ['SysRoleAdd', 'SysRoleDel', 'SysRoleStatus', 'SysRoleEdit', 'SysRoleInfo', 'SysRoleList'], meta: { title: 'Roles', icon: 'peoples', noCache: false, hidden: false } },
      { path: 'dept', component: () => import('@/views/System/dept/SysDept.vue'), name: 'SysDept', permission: ['SysDeptAdd', 'SysDeptDel', 'SysDeptEdit', 'SysDeptInfo', 'SysDeptList'], meta: { title: 'Dept', icon: 'tree', noCache: false, hidden: false } },
      { path: 'menu', component: () => import('@/views/System/menu/SysMenu.vue'), name: 'SysMenu', permission: ['SysMenuAdd', 'SysMenuDel', 'SysMenuEdit', 'SysMenuInfo', 'SysMenuList', 'SysMenuDrop'], meta: { title: 'Menu', icon: 'tree-table', noCache: false, hidden: false } }
    ]
  },
  {
    path: '/personal',
    component: Layout,
    redirect: '/personal/personal-center',
    name: 'Personal',
    meta: { title: 'Personal', hidden: true, canTo: true },
    children: [
      { path: 'personal-center', component: () => import('@/views/Personal/PersonalCenter.vue'), name: 'PersonalCenter', meta: { title: 'Personal Center', hidden: true, canTo: true } }
    ]
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  strict: true,
  routes: constantRouterMap,
  scrollBehavior: () => ({ left: 0, top: 0 })
})

export const resetRouter = () => {
  router.getRoutes().forEach((route) => {
    const { name } = route
    if (name && !NO_RESET_WHITE_LIST.includes(name)) {
      router.hasRoute(name) && router.removeRoute(name)
    }
  })
}

export const setupRouter = (app) => {
  app.use(router)
}

export default router
