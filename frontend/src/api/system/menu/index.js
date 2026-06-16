import service from '@/api/index'

export const getMenuListApi = (params) => {
  return service({
    url: '/menu/list',
    method: 'get',
    params,
  })
}

export const addMenuApi = (data) => {
  return service({
    url: '/menu/add',
    method: 'post',
    data,
  })
}

export const editMenuApi = (data) => {
  return service({
    url: '/menu/edit',
    method: 'post',
    data,
  })
}

export const delMenuApi = (data) => {
  return service({
    url: '/menu/del',
    method: 'post',
    data,
  })
}

export const getMenuInfoApi = (params) => {
  return service({
    url: '/menu/info',
    method: 'get',
    params,
  })
}

export const getMenuDropdownApi = (params) => {
  return service({
    url: '/menu/dropdown',
    method: 'get',
    params,
  })
}

export const getMenuRoleApi = (params) => {
  return service({
    url: '/menu/role',
    method: 'get',
    params,
  })
}
