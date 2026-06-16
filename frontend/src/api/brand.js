import service from '@/api/index'


export const getBrandConfTreeDropApi = (params) => {
  return service({ url: '/system/brand/tree', method: 'get', params })
}
