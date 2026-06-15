import axios from 'axios'
import { ElMessage } from 'element-plus'
import router from '@/router'

const request = axios.create({
  baseURL: '/api',
  timeout: 30000,
})

request.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = token
    }
    return config
  },
  (error) => Promise.reject(error)
)

request.interceptors.response.use(
  (response) => {
    const res = response.data
    if (res.code !== 0) {
      ElMessage.error(res.message || '请求失败')
      if (res.code === 401) {
        localStorage.removeItem('token')
        router.push('/login')
      }
      return Promise.reject(new Error(res.message))
    }
    return res
  },
  (error) => {
    ElMessage.error(error.message || '网络错误')
    return Promise.reject(error)
  }
)

export default request

// ==================== 用户与权限 ====================
export const userApi = {
  login: (data) => request.post('/user/login', data),
  register: (data) => request.post('/user/register', data),
  getUserInfo: () => request.get('/user/info'),
  getUserList: (params) => request.get('/user/list', { params }),
  updateUser: (data) => request.put('/user/update', data),
  deleteUser: (id) => request.delete(`/user/delete/${id}`),
  getRoleList: () => request.get('/user/roles'),
  createRole: (data) => request.post('/user/role', data),
  updateRole: (data) => request.put('/user/role', data),
  deleteRole: (id) => request.delete(`/user/role/${id}`),
  getEnterpriseList: (params) => request.get('/user/enterprise', { params }),
  createEnterprise: (data) => request.post('/user/enterprise', data),
  updateEnterprise: (data) => request.put('/user/enterprise', data),
}

// ==================== 产品溯源 ====================
export const traceApi = {
  createProduct: (data) => request.post('/trace/product', data),
  updateProduct: (data) => request.put('/trace/product', data),
  deleteProduct: (id) => request.delete(`/trace/product/${id}`),
  getProduct: (id) => request.get(`/trace/product/${id}`),
  getProductList: (params) => request.get('/trace/product/list', { params }),
  createBatch: (data) => request.post('/trace/batch', data),
  getBatch: (id) => request.get(`/trace/batch/${id}`),
  getBatchList: (params) => request.get('/trace/batch/list', { params }),
  createTraceRecord: (data) => request.post('/trace/record', data),
  getTraceRecord: (id) => request.get(`/trace/record/${id}`),
  getTraceRecordList: (productId) => request.get(`/trace/record/list/${productId}`),
  getTraceChain: (productId) => request.get(`/trace/chain/${productId}`),
  generateQRCode: (productId) => request.get(`/trace/qrcode/${productId}`),
  publicQuery: (params) => request.get('/trace/public/query', { params }),
}

// ==================== 供应链管理 ====================
export const supplyApi = {
  createSupplier: (data) => request.post('/supply/supplier', data),
  updateSupplier: (data) => request.put('/supply/supplier', data),
  getSupplier: (id) => request.get(`/supply/supplier/${id}`),
  getSupplierList: (params) => request.get('/supply/supplier/list', { params }),
  deleteSupplier: (id) => request.delete(`/supply/supplier/${id}`),
  createOrder: (data) => request.post('/supply/order', data),
  updateOrder: (data) => request.put('/supply/order', data),
  getOrder: (id) => request.get(`/supply/order/${id}`),
  getOrderList: (params) => request.get('/supply/order/list', { params }),
  deleteOrder: (id) => request.delete(`/supply/order/${id}`),
  warehouseIn: (data) => request.post('/supply/warehouse/in', data),
  warehouseOut: (data) => request.post('/supply/warehouse/out', data),
  getStockList: (params) => request.get('/supply/warehouse/stock', { params }),
  createLogistics: (data) => request.post('/supply/logistics', data),
  getLogisticsList: (orderId) => request.get(`/supply/logistics/${orderId}`),
}

// ==================== 区块链管理 ====================
export const blockchainApi = {
  getBlockList: (params) => request.get('/blockchain/blocks', { params }),
  getBlockByHash: (hash) => request.get(`/blockchain/block/${hash}`),
  getTransactionList: (params) => request.get('/blockchain/transactions', { params }),
  getTransactionDetail: (txId) => request.get(`/blockchain/transaction/${txId}`),
  deployContract: (data) => request.post('/blockchain/contract', data),
  getContractInfo: (address) => request.get(`/blockchain/contract/${address}`),
  getContractList: () => request.get('/blockchain/contract/list'),
  uploadToChain: (data) => request.post('/blockchain/data/upload', data),
  verifyData: (params) => request.get('/blockchain/data/verify', { params }),
}

// ==================== 仓库管理 ====================
export const warehouseApi = {
  // 仓库信息
  getWarehouseList: (params) => request.get('/v1/warehouse/list', { params }),
  createWarehouse: (data) => request.post('/v1/warehouse', data),
  updateWarehouse: (data) => request.put('/v1/warehouse', data),
  deleteWarehouse: (id) => request.delete(`/v1/warehouse/${id}`),
  // 区域货架
  getAreaList: (params) => request.get('/v1/warehouse/area/list', { params }),
  createArea: (data) => request.post('/v1/warehouse/area', data),
  updateArea: (data) => request.put('/v1/warehouse/area', data),
  deleteArea: (id) => request.delete(`/v1/warehouse/area/${id}`),
  createShelf: (data) => request.post('/v1/warehouse/shelf', data),
  updateShelf: (data) => request.put('/v1/warehouse/shelf', data),
  deleteShelf: (id) => request.delete(`/v1/warehouse/shelf/${id}`),
  // 盘存
  getStocktakeList: (params) => request.get('/v1/warehouse/stocktakes', { params }),
  createStocktake: (data) => request.post('/v1/warehouse/stocktakes', data),
  deleteStocktake: (id) => request.delete(`/v1/warehouse/stocktakes/${id}`),
}

// ==================== 库存管理 ====================
export const inventoryApi = {
  // 整车
  getCarList: (params) => request.get('/v1/inventory/cars', { params }),
  // 原材料
  getRawMaterialList: (params) => request.get('/v1/inventory/raw-materials', { params }),
  // 危固废
  getWasteList: (params) => request.get('/v1/inventory/waste', { params }),
  // 配件-溯源件
  getPartTraceableList: (params) => request.get('/v1/inventory/parts/traceable', { params }),
  partTraceableOutbound: (data) => request.post('/v1/inventory/parts/traceable/outbound', data),
  // 配件-非溯源件
  getPartUntraceableList: (params) => request.get('/v1/inventory/parts/non-traceable', { params }),
  partUntraceableInbound: (data) => request.post('/v1/inventory/parts/non-traceable/inbound', data),
  partUntraceableOutbound: (data) => request.post('/v1/inventory/parts/non-traceable/outbound', data),
}

// ==================== 操作记录 ====================
export const recordApi = {
  getInboundList: (params) => request.get('/v1/records/inbound', { params }),
  getInboundDetail: (id) => request.get(`/v1/records/inbound/${id}`),
  getOutboundList: (params) => request.get('/v1/records/outbound', { params }),
  getOutboundDetail: (id) => request.get(`/v1/records/outbound/${id}`),
  getTransferList: (params) => request.get('/v1/records/transfer', { params }),
  getTransferDetail: (id) => request.get(`/v1/records/transfer/${id}`),
}

// ==================== 系统管理 ====================
export const systemApi = {
  getDashboard: () => request.get('/system/dashboard'),
  getConfigList: () => request.get('/system/config'),
  updateConfig: (data) => request.put('/system/config', data),
  getLogList: (params) => request.get('/system/logs', { params }),
  deleteLog: (id) => request.delete(`/system/logs/${id}`),
  getNotificationList: () => request.get('/system/notifications'),
  markNotificationRead: (id) => request.post(`/system/notification/read/${id}`),
  uploadFile: (file) => {
    const formData = new FormData()
    formData.append('file', file)
    return request.post('/upload', formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
    })
  },
}
