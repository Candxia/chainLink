import service from '@/api/index'

export const getAdminListApi = (params) => {
  return service({
    url: '/admin/list',
    method: 'get',
    params,
  })
}

export const addAdminApi = (data) => {
  return service({
    url: '/admin/add',
    method: 'post',
    data,
  })
}

export const editAdminApi = (data) => {
  return service({
    url: '/admin/edit',
    method: 'post',
    data,
  })
}

export const delAdminApi = (data) => {
  return service({
    url: '/admin/del',
    method: 'post',
    data,
  })
}

export const adminPasswordApi = (data) => {
  return service({
    url: '/admin/password',
    method: 'post',
    data,
  })
}

export const adminStatusApi = (data) => {
  return service({
    url: '/admin/status',
    method: 'post',
    data,
  })
}

export const adminClickOutApi = (data) => {
  return service({
    url: '/admin/clickout',
    method: 'post',
    data,
  })
}
