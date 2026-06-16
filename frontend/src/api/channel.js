import service from '@/api/index'
export default { getList(params) { return service({ url: '/channel/list', method: 'get', params }).catch(() => ({ data: [] })) } }
