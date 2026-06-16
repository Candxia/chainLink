import service from '@/api/index'

export const getDictListApi = (params) => {
  return service({
    url: '/dict/list',
    method: 'get',
    params,
  })
}

export const addDictApi = (data) => {
  return service({
    url: '/dict/add',
    method: 'post',
    data,
  })
}

export const editDictApi = (data) => {
  return service({
    url: '/dict/edit',
    method: 'post',
    data,
  })
}

export const delDictApi = (data) => {
  return service({
    url: '/dict/del',
    method: 'post',
    data,
  })
}
