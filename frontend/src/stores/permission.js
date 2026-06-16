import { defineStore } from 'pinia'
import { store } from '@/stores'
import { asyncRouterMap } from '@/router'

export const usePermissionStore = defineStore('permission', {
  state: () => ({
    routers: [],
    addRouters: [],
    isAddRouters: false
  }),
  getters: {
    getAddRouters: (state) => state.addRouters,
    getIsAddRouters: (state) => state.isAddRouters,
    getRouters: (state) => state.routers
  },
  actions: {
    async generateRoutes(type = 'frontEnd') {
      if (type === 'static') {
        this.addRouters = asyncRouterMap
      } else {
        // frontEnd: use all async routes
        this.addRouters = asyncRouterMap
      }
      this.routers = this.addRouters
    },
    setIsAddRouters(v) {
      this.isAddRouters = v
    }
  },
  persist: true
})

export const usePermissionStoreWithOut = () => usePermissionStore(store)
