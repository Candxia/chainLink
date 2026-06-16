import service from '@/api/index'

export const getTaskListApi = (params) => {
  return service({
    url: '/task/list',
    method: 'get',
    params,
  })
}

export const addTaskApi = (data) => {
  return service({
    url: '/task/add',
    method: 'post',
    data,
  })
}

export const delTaskApi = (data) => {
  return service({
    url: '/task/del',
    method: 'post',
    data,
  })
}
