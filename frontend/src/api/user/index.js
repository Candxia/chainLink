import service from '../index'

export const userApi = {
  getUserList(params) {
    return service({ url: '/user/list', method: 'get', params })
  },
  register(data) {
    return service({ url: '/user/register', method: 'post', data })
  },
  updateUser(data) {
    return service({ url: '/user/update', method: 'post', data })
  },
  deleteUser(id) {
    return service({ url: '/user/delete', method: 'post', data: { id } })
  },
}

