import service from '@/api/index'

export const loginApi = (data) => {
  return service({
    url: '/login',
    method: 'post',
    data,
  })
}

export const logoutApi = () => {
  return service({
    url: '/logout',
    method: 'post',
  })
}

export const refreshApi = () => {
  return service({
    url: '/refresh',
    method: 'post',
  })
}

export const updateTokenApi = () => {
  return service({
    url: '/updateToken',
    method: 'post',
  })
}
