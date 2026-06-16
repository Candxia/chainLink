import service from '@/api/index'

export const getDeptListApi = (params) => {
  return service({
    url: '/dept/list',
    method: 'get',
    params,
  })
}

export const addDeptApi = (data) => {
  return service({
    url: '/dept/add',
    method: 'post',
    data,
  })
}

export const editDeptApi = (data) => {
  return service({
    url: '/dept/edit',
    method: 'post',
    data,
  })
}

export const delDeptApi = (data) => {
  return service({
    url: '/dept/del',
    method: 'post',
    data,
  })
}

export const getDeptInfoApi = (params) => {
  return service({
    url: '/dept/info',
    method: 'get',
    params,
  })
}

export const getDeptDropdownApi = (params) => {
  return service({
    url: '/dept/dropdown',
    method: 'get',
    params,
  })
}
