
import service from '@/api/index'

// ==================== 浠〃鐩?/ 閰嶇疆 / 鏃ュ織 ====================

export function getDashboard(params) {
  return service({ url: '/system/dashboard', method: 'get', params })
}
export function getConfigList(params) {
  return service({ url: '/system/config/list', method: 'get', params })
}
export function updateConfig(data) {
  return service({ url: '/system/config/update', method: 'post', data })
}
export function getLogList(params) {
  return service({ url: '/system/logs', method: 'get', params })
}
export function deleteLog(id) {
  return service({ url: '/system/logs/' + id, method: 'delete' })
}

// ==================== 鑱氬悎瀵煎嚭 (鍏煎鏃х増) ====================

export const systemApi = {
  getConfigList,
  getDashboard,
  updateConfig,
  getLogList,
  deleteLog,
  getAdminListApi, addAdminApi, editAdminApi, delAdminApi, adminPasswordApi, adminStatusApi, adminClickOutApi,
  getApiListApi, addApiApi, editApiApi, delApiApi, getApiDropdownApi, getApiPathApi,
  getRoleListApi, addRoleApi, editRoleApi, delRoleApi, getRoleInfoApi, roleStatusApi, getRoleDropdownApi, roleIsMobileApi,
  getDeptListApi, addDeptApi, editDeptApi, delDeptApi, getDeptInfoApi, getDeptDropdownApi,
  getMenuListApi, addMenuApi, editMenuApi, delMenuApi, getMenuInfoApi, getMenuDropdownApi, getMenuRoleApi,
  getDictListApi, addDictApi, editDictApi, delDictApi,
  getCronListApi, editCronApi,
  getTaskListApi, addTaskApi, delTaskApi
}

// ==================== 璐﹀彿绠＄悊 ====================

export function getAdminListApi(params) {
  return service({ url: '/system/admin', method: 'get', params })
}
export function addAdminApi(data) {
  return service({ url: '/system/admin', method: 'post', data })
}
export function editAdminApi(data) {
  return service({ url: '/system/admin', method: 'put', data })
}
export function delAdminApi(data) {
  return service({ url: `/system/admin/${data.id}`, method: 'delete' })
}
export function adminPasswordApi(data) {
  return service({ url: `/system/admin/${data.id}/password`, method: 'post', data })
}
export function adminStatusApi(data) {
  return service({ url: `/system/admin/${data.id}/status`, method: 'post', data })
}
export function adminClickOutApi(data) {
  return service({ url: `/system/admin/${data.id}/clickout`, method: 'post', data })
}

// ==================== 鎺ュ彛绠＄悊 ====================

export function getApiListApi(params) {
  return service({ url: '/system/api', method: 'get', params })
}
export function addApiApi(data) {
  return service({ url: '/system/api', method: 'post', data })
}
export function editApiApi(data) {
  return service({ url: '/system/api', method: 'post', data })
}
export function delApiApi(data) {
  return service({ url: `/system/api/${data.id}`, method: 'delete' })
}
export function getApiDropdownApi(params) {
  return service({ url: '/system/api/dropdown', method: 'get', params })
}
export function getApiPathApi(params) {
  return service({ url: '/system/api/path', method: 'get', params })
}

// ==================== 瑙掕壊绠＄悊 ====================

export function getRoleListApi(params) {
  return service({ url: '/system/role', method: 'get', params })
}
export function addRoleApi(data) {
  return service({ url: '/system/role', method: 'post', data })
}
export function editRoleApi(data) {
  return service({ url: `/system/role/${data.id}`, method: 'put', data })
}
export function delRoleApi(data) {
  return service({ url: `/system/role/${data.id}`, method: 'delete' })
}
export function getRoleInfoApi(params) {
  return service({ url: `/system/role/${params.id}`, method: 'get', params })
}
export function roleStatusApi(data) {
  return service({ url: `/system/role/${data.id}/status`, method: 'post', data })
}
export function getRoleDropdownApi(params) {
  return service({ url: '/system/role/dropdown', method: 'get', params })
}
export function roleIsMobileApi(data) {
  return service({ url: `/system/role/${data.id}/is_mobile`, method: 'post', data })
}

// ==================== 閮ㄩ棬绠＄悊 ====================

export function getDeptListApi(params) {
  return service({ url: '/system/dept', method: 'get', params })
}
export function addDeptApi(data) {
  return service({ url: '/system/dept', method: 'post', data })
}
export function editDeptApi(data) {
  return service({ url: `/system/dept/${data.id}`, method: 'put', data })
}
export function delDeptApi(data) {
  return service({ url: `/system/dept/${data.id}`, method: 'delete' })
}
export function getDeptInfoApi(params) {
  return service({ url: `/system/dept/${params.id}`, method: 'get', params })
}
export function getDeptDropdownApi(params) {
  return service({ url: '/system/dept/dropdown', method: 'get', params })
}

// ==================== 鑿滃崟绠＄悊 ====================

export function getMenuListApi(params) {
  return service({ url: '/system/menu', method: 'get', params })
}
export function addMenuApi(data) {
  return service({ url: '/system/menu', method: 'post', data })
}
export function editMenuApi(data) {
  return service({ url: `/system/menu/${data.id}`, method: 'put', data })
}
export function delMenuApi(data) {
  return service({ url: `/system/menu/${data.id}`, method: 'delete' })
}
export function getMenuInfoApi(params) {
  return service({ url: `/system/menu/${params.id}`, method: 'get', params })
}
export function getMenuDropdownApi(params) {
  return service({ url: '/system/menu/dropdown', method: 'get', params })
}
export function getMenuRoleApi(params) {
  return service({ url: '/system/menu/role', method: 'get', params })
}

// ==================== 寮€鍙戝伐鍏?====================

export function getDictListApi(params) {
  return service({ url: '/system/tools/dict/list', method: 'get', params })
}
export function addDictApi(data) {
  return service({ url: '/system/tools/dict/add', method: 'post', data })
}
export function editDictApi(data) {
  return service({ url: '/system/tools/dict/edit', method: 'post', data })
}
export function delDictApi(data) {
  return service({ url: '/system/tools/dict/del', method: 'post', data })
}

export function getCronListApi(params) {
  return service({ url: '/system/tools/cron/list', method: 'get', params })
}
export function editCronApi(data) {
  return service({ url: '/system/tools/cron/edit', method: 'post', data })
}

export function getTaskListApi(params) {
  return service({ url: '/system/tools/task/list', method: 'get', params })
}
export function addTaskApi(data) {
  return service({ url: '/system/tools/task/add', method: 'post', data })
}
export function delTaskApi(data) {
  return service({ url: '/system/tools/task/del', method: 'post', data })
}
