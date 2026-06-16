import service from '../index'

export const inventoryApi = {
  getCarList(params) {
    return service({ url: '/inventory/car/list', method: 'get', params })
  },
  getRawMaterialList(params) {
    return service({ url: '/inventory/raw/list', method: 'get', params })
  },
  getWasteList(params) {
    return service({ url: '/inventory/waste/list', method: 'get', params })
  },
  getPartTraceableList(params) {
    return service({ url: '/inventory/part/traceable/list', method: 'get', params })
  },
  getPartUntraceableList(params) {
    return service({ url: '/inventory/part/untraceable/list', method: 'get', params })
  },
}

