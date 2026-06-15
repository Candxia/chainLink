import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue'),
    meta: { title: '登录' },
  },
  {
    path: '/',
    component: () => import('@/layouts/MainLayout.vue'),
    redirect: '/dashboard',
    children: [
      { path: 'dashboard', name: 'Dashboard', component: () => import('@/views/Dashboard.vue'), meta: { title: '仪表盘' } },
      // 用户管理
      { path: 'user', name: 'UserList', component: () => import('@/views/user/UserList.vue'), meta: { title: '用户管理' } },
      { path: 'user/role', name: 'RoleList', component: () => import('@/views/user/RoleList.vue'), meta: { title: '角色管理' } },
      { path: 'enterprise', name: 'EnterpriseList', component: () => import('@/views/user/EnterpriseList.vue'), meta: { title: '企业管理' } },
      // 产品溯源
      { path: 'product', name: 'ProductList', component: () => import('@/views/trace/ProductList.vue'), meta: { title: '产品管理' } },
      { path: 'product/detail/:id', name: 'ProductDetail', component: () => import('@/views/trace/ProductDetail.vue'), meta: { title: '产品详情' } },
      { path: 'batch', name: 'BatchList', component: () => import('@/views/trace/BatchList.vue'), meta: { title: '批次管理' } },
      { path: 'trace-record', name: 'TraceRecordList', component: () => import('@/views/trace/TraceRecordList.vue'), meta: { title: '溯源记录' } },
      { path: 'trace/chain/:productId', name: 'TraceChain', component: () => import('@/views/trace/TraceChain.vue'), meta: { title: '追溯链' } },
      // 供应链管理
      { path: 'supplier', name: 'SupplierList', component: () => import('@/views/supply/SupplierList.vue'), meta: { title: '供应商管理' } },
      { path: 'order', name: 'OrderList', component: () => import('@/views/supply/OrderList.vue'), meta: { title: '订单管理' } },
      { path: 'warehouse', name: 'WarehouseList', component: () => import('@/views/supply/WarehouseList.vue'), meta: { title: '仓储管理' } },
      { path: 'logistics', name: 'LogisticsList', component: () => import('@/views/supply/LogisticsList.vue'), meta: { title: '物流管理' } },
      // 区块链
      { path: 'blockchain', name: 'BlockList', component: () => import('@/views/blockchain/BlockList.vue'), meta: { title: '区块浏览器' } },
      { path: 'blockchain/transaction', name: 'TransactionList', component: () => import('@/views/blockchain/TransactionList.vue'), meta: { title: '交易记录' } },
      { path: 'blockchain/contract', name: 'ContractList', component: () => import('@/views/blockchain/ContractList.vue'), meta: { title: '合约管理' } },
      // 系统管理
      { path: 'config', name: 'ConfigList', component: () => import('@/views/system/ConfigList.vue'), meta: { title: '系统配置' } },
      { path: 'logs', name: 'LogList', component: () => import('@/views/system/LogList.vue'), meta: { title: '操作日志' } },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  if (to.name !== 'Login' && !token) {
    next({ name: 'Login' })
  } else {
    next()
  }
})

export default router
