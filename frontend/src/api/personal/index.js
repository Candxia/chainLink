import service from '@/api/index'

export const getPersonalInfoApi = () => {
  return service({
    url: '/personal/info',
    method: 'get',
  })
}

export const optResetApi = (data) => {
  return service({
    url: '/personal/optReset',
    method: 'post',
    data,
  })
}

export const getPersonalRouterApi = () => {
  return service({
    url: '/personal/router',
    method: 'get',
  })
}
