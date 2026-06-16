import axios from 'axios'
import { ElMessage } from 'element-plus'
import { TOKEN_KEY, REQUEST_TIMEOUT } from '@/constants'
import { getStorage } from '@/utils/storage'

const service = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api',
  timeout: REQUEST_TIMEOUT,
  headers: { 'Content-Type': 'application/json;charset=UTF-8' },
})

service.interceptors.request.use(
  (config) => {
    const token = getStorage(TOKEN_KEY) || ''
    if (token) config.headers['Authorization'] = 'Bearer ' + token
    return config
  },
  (error) => Promise.reject(error)
)

service.interceptors.response.use(
  (response) => {
    const res = response.data
    if (res.code === 0) return res
    ElMessage.error(res.msg || '请求失败')
    return Promise.reject(res)
  },
  (error) => {
    if (error.response) {
      const { status } = error.response
      if (status === 401) window.location.hash = '#/login'
      else if (status === 403) ElMessage.error('没有权限')
      else if (status === 404) ElMessage.error('请求的资源不存在')
    } else {
      ElMessage.error('网络错误')
    }
    return Promise.reject(error)
  }
)

export default service
