import service from '../index'

export const recordApi = {
  getInboundList(params) {
    return service({ url: '/record/inbound/list', method: 'get', params })
  },
  getOutboundList(params) {
    return service({ url: '/record/outbound/list', method: 'get', params })
  },
  getTransferList(params) {
    return service({ url: '/record/transfer/list', method: 'get', params })
  },
}

