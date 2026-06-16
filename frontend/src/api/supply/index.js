import service from '../index'

export const supplyApi = {
  getSupplierList(params) {
    return service({ url: '/supply/supplier/list', method: 'get', params })
  },
  getOrderList(params) {
    return service({ url: '/supply/order/list', method: 'get', params })
  },
  getLogisticsList(params) {
    return service({ url: '/supply/logistics/list', method: 'get', params })
  },
  getStockList(params) {
    return service({ url: '/supply/stock/list', method: 'get', params })
  },
  warehouseIn(data) {
    return service({ url: '/supply/warehouse/in', method: 'post', data })
  },
  warehouseOut(data) {
    return service({ url: '/supply/warehouse/out', method: 'post', data })
  },
}

