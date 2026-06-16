
import router from './router'

router.beforeEach((to, from, next) => {
  // token 检查在 router/index.js 中
  next()
})
