import router from './index'
import { usePermissionStore } from '@/stores/permission'
import { useUserStore } from '@/stores/user'
import { NO_REDIRECT_WHITE_LIST } from '@/constants'

const whiteList = ['/login', '/404', '/redirect']

router.beforeEach(async (to, from, next) => {
  const userStore = useUserStore()
  const permissionStore = usePermissionStore()

  if (userStore.getUserInfo && userStore.getUserInfo.username) {
    // Already logged in
    if (to.path === '/login') {
      next({ path: '/' })
    } else {
      if (permissionStore.getIsAddRouters) {
        // Routes already loaded
        next()
        return
      }
      // First login after page load: generate routes and add them
      await permissionStore.generateRoutes('frontEnd')
      permissionStore.getAddRouters.forEach((route) => {
        router.addRoute(route)
      })
      permissionStore.setIsAddRouters(true)
      next({ ...to, replace: true })
    }
  } else {
    // Not logged in
    if (whiteList.indexOf(to.path) !== -1) {
      next()
    } else {
      next(`/login?redirect=${to.path}`)
    }
  }
})

router.afterEach(() => {
  // Future: could set page title here
})
