import service from '@/api/index'
export default { getList(params) { return service({ url: '/pkg/list', method: 'get', params }).catch(() => ({ data: [] })) } }
