import service from '@/api/index'

export const getRoleListApi = (params) => {
  return service({
    url: '/role/list',
    method: 'get',
    params,
  })
}

export const addRoleApi = (data) => {
  return service({
    url: '/role/add',
    method: 'post',
    data,
  })
}

export const editRoleApi = (data) => {
  return service({
    url: '/role/edit',
    method: 'post',
    data,
  })
}

export const delRoleApi = (data) => {
  return service({
    url: '/role/del',
    method: 'post',
    data,
  })
}

export const getRoleInfoApi = (params) => {
  return service({
    url: '/role/info',
    method: 'get',
    params,
  })
}

export const roleStatusApi = (data) => {
  return service({
    url: '/role/status',
    method: 'post',
    data,
  })
}

export const getRoleDropdownApi = (params) => {
  return service({
    url: '/role/dropdown',
    method: 'get',
    params,
  })
}

export const roleIsMobileApi = (data) => {
  return service({
    url: '/role/is_mobile',
    method: 'post',
    data,
  })
}
