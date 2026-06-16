import service from '@/api/index'

export const getCronListApi = (params) => {
  return service({
    url: '/cron/list',
    method: 'get',
    params,
  })
}

export const addCronApi = (data) => {
  return service({
    url: '/cron/add',
    method: 'post',
    data,
  })
}

export const editCronApi = (data) => {
  return service({
    url: '/cron/edit',
    method: 'post',
    data,
  })
}

export const delCronApi = (data) => {
  return service({
    url: '/cron/del',
    method: 'post',
    data,
  })
}

export const editCronStatusApi = (data) => {
  return service({
    url: '/cron/status',
    method: 'post',
    data,
  })
}
