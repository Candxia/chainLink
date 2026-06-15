import { defineStore } from 'pinia'
import { userApi } from '@/api'

export const useUserStore = defineStore('user', {
  state: () => ({
    token: localStorage.getItem('token') || '',
    userId: 0,
    userName: '',
    role: '',
    avatar: '',
    menus: [],
  }),
  getters: {
    isLoggedIn: (state) => !!state.token,
    isAdmin: (state) => state.role === 'super_admin' || state.role === 'admin',
  },
  actions: {
    async login(credentials) {
      const res = await userApi.login(credentials)
      this.token = res.data.token
      this.userId = res.data.userId
      this.userName = res.data.userName
      this.role = res.data.role
      localStorage.setItem('token', res.data.token)
      return res
    },
    async getUserInfo() {
      try {
        const res = await userApi.getUserInfo()
        this.userId = res.data.id
        this.userName = res.data.username
        this.role = res.data.role
      } catch (e) {
        this.logout()
      }
    },
    logout() {
      this.token = ''
      this.userId = 0
      this.userName = ''
      this.role = ''
      localStorage.removeItem('token')
    },
  },
})
