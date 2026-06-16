import service from '@/api/index'

export const getBrandTreeApi = (params) => {
  return service({
    url: '/common/brandTree',
    method: 'get',
    params,
  })
}

export const getCommonChannelApi = (params) => {
  return service({
    url: '/common/channel',
    method: 'get',
    params,
  })
}

export const getCommonPkgApi = (params) => {
  return service({
    url: '/common/pkg',
    method: 'get',
    params,
  })
}
