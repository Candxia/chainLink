import 'uno.css'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPersist from 'pinia-plugin-persistedstate'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import DisableDevtool from 'disable-devtool'
import { openDevTools, formatError, setupKeyInterceptor, crashesReportLog } from '@/utils/preventDevTools'
import App from './App.vue'
import router from './router'
import '@/router/permission'
import './utils/storage'

const pinia = createPinia()
pinia.use(piniaPersist)

const app = createApp(App)

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

// 键盘拦截（F12 / Ctrl+Shift+I / 右键 全部屏蔽）
setupKeyInterceptor()

// 禁用开发者工具（仅生产环境）- 必须在 app.mount() 之前初始化
if (!import.meta.env.DEV) {
  DisableDevtool({
    md5: 'e10adc3949ba59abbe56e057f20f883e',
    url: '',
    detectors: 'all',
    interval: 200,
    disableMenu: true,
    disableSelect: false,
    disableCopy: false,
    disableCut: false,
    disablePaste: false,
    clearLog: false,
    ignore: ['?Omniscience=true'],
    ondevtoolopen: function (type) {
      openDevTools()
    },
    ondevtoolclose: function () {
    }
  })

  // 全局未捕获异常上报
  window.onerror = function (message) {
    crashesReportLog && crashesReportLog(typeof message === 'string' ? message : '')
  }
}

app.use(pinia)
app.use(router)
app.use(ElementPlus, { size: 'default' })
app.mount('#app')
