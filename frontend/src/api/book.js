import service from '@/api/index'

export const getOutlineMenusApi = (params) => {
  return service({ url: '/book/outline/menus', method: 'get', params })
}
