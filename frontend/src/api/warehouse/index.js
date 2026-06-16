import service from '../index'

export const warehouseApi = {
  getWarehouseList(params) {
    return service({ url: '/warehouse/list', method: 'get', params })
  },
  getWarehouseArea(params) {
    return service({ url: '/warehouse/area/list', method: 'get', params })
  },
  getStocktakeList(params) {
    return service({ url: '/warehouse/stocktake/list', method: 'get', params })
  },
}

