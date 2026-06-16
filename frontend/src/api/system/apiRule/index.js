import service from '@/api/index'

export const getApiListApi = (params) => {
  return service({
    url: '/api/list',
    method: 'get',
    params,
  })
}

export const addApiApi = (data) => {
  return service({
    url: '/api/add',
    method: 'post',
    data,
  })
}

export const editApiApi = (data) => {
  return service({
    url: '/api/edit',
    method: 'post',
    data,
  })
}

export const delApiApi = (data) => {
  return service({
    url: '/api/del',
    method: 'post',
    data,
  })
}

export const getApiDropdownApi = (params) => {
  return service({
    url: '/api/dropdown',
    method: 'get',
    params,
  })
}

export const getApiPathApi = (params) => {
  return service({
    url: '/api/path',
    method: 'get',
    params,
  })
}
