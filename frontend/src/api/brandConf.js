import service from '@/api/index'
export default { getList(params) { return service({ url: '/brandConf/list', method: 'get', params }).catch(() => ({ data: [] })) } }
