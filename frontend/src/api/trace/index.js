import service from '../index'

export const traceApi = {
  getProductList(params) {
    return service({ url: '/trace/product/list', method: 'get', params })
  },
  getProductDetail(id) {
    return service({ url: '/trace/product/detail', method: 'get', params: { id } })
  },
  getBatchList(params) {
    return service({ url: '/trace/batch/list', method: 'get', params })
  },
  getTraceRecordList(params) {
    return service({ url: '/trace/record/list', method: 'get', params })
  },
  getTraceChain(productId) {
    return service({ url: '/trace/chain', method: 'get', params: { productId } })
  },
}

