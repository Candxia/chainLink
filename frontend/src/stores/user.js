import { defineStore } from 'pinia'
import { store } from './index'
import { useStorage } from '@/hooks/web/useStorage'

export const useUserStore = defineStore('user', {
  state: () => ({ token: '', userInfo: {}, roleRouters: [], rememberMe: false }),
  getters: { getToken: s => s.token, getUserInfo: s => s.userInfo },
  actions: {
    setToken(t) { this.token = t },
    setUserInfo(u) { this.userInfo = u },
    logout() { this.token = ''; this.userInfo = {} }
  },
  persist: true
})

export const useUserStoreWithOut = () => useUserStore(store)
