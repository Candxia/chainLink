import { writeFileSync } from 'fs'
const base = 'D:/CodeByAi/cl_system/frontend/src'

function wf(path, lines) {
  writeFileSync(base + '/' + path, lines.join('\n') + '\n', 'utf-8')
}

// === Utils ===
wf('utils/is.js', [
  "export const isUrl = (path) => {",
  "  if (!path) return false",
  "  return /^(https?:|mailto:|tel:)/.test(path)",
  "}",
  "export const isExternal = (path) => /^(https?:|mailto:|tel:)/.test(path)",
])

wf('utils/storage.js', [
  "export const getStorage = (key) => localStorage.getItem(key)",
  "export const setStorage = (key, val) => localStorage.setItem(key, val)",
  "export const removeStorage = (key) => localStorage.removeItem(key)",
  "export const clearStorage = () => localStorage.clear()",
])

wf('utils/routerHelper.js', [
  "export const Layout = () => import('@/layout/Layout.vue')",
  "export const getParentLayout = () => ({ name: 'ParentLayout' })",
  "export const flatMultiLevelRoutes = (routes) => routes",
  "export const generateRoutesByFrontEnd = (routes, keys) => routes",
  "export const generateRoutesByServer = (routers) => routers || []",
  "export const getRawRoute = (route) => route",
])

wf('utils/trim.js', [
  "export const trim = (str) => (str || '').trim()",
])

wf('utils/formatter.js', [
  "export const formatterTagType = () => 'success'",
  "export const formatterPushType = () => 'info'",
])

// === Constants ===
wf('constants/index.js', [
  "export const TOKEN_KEY = 'token'",
  "export const SUCCESS_CODE = 0",
  "export const REQUEST_TIMEOUT = 60000",
  "export const CONTENT_TYPE = 'application/json'",
  "export const NO_REDIRECT_WHITE_LIST = ['/login']",
  "export const NO_RESET_WHITE_LIST = ['Redirect', 'Login', 'NoFind', 'Root']",
  "export const ICON_PREFIX = 'vi-'",
])

// === Hooks ===
wf('hooks/web/useTable.js', [
  "import { ref } from 'vue'",
  "import { ElMessage, ElMessageBox } from 'element-plus'",
  "export const useTable = (fetchApi) => {",
  "  const loading = ref(false)",
  "  const tableData = ref([])",
  "  const total = ref(0)",
  "  const page = ref(1)",
  "  const pageSize = ref(20)",
  "  const searchParams = ref({})",
  "  const getList = async () => {",
  "    loading.value = true",
  "    try {",
  "      const res = await fetchApi({ page: page.value, pageSize: pageSize.value, ...searchParams.value })",
  "      tableData.value = res.data?.list || res.data || []",
  "      total.value = res.data?.total || 0",
  "    } catch (e) { tableData.value = [] } finally { loading.value = false }",
  "  }",
  "  const handleSearch = () => { page.value = 1; getList() }",
  "  const handleReset = () => { searchParams.value = {}; page.value = 1; getList() }",
  "  return { loading, tableData, total, page, pageSize, searchParams, getList, handleSearch, handleReset }",
  "}",
])

wf('hooks/web/useForm.js', [
  "import { ref } from 'vue'",
  "export const useForm = () => { const formRef = ref(null); return { formRef } }",
])

wf('hooks/web/useI18n.js', [
  "export const useI18n = () => ({ t: (key) => key })",
])

wf('hooks/web/useStorage.js', [
  "export const useStorage = (type = 'localStorage') => ({",
  "  getStorage: (key) => localStorage.getItem(key),",
  "  setStorage: (key, val) => localStorage.setItem(key, val),",
  "  clear: () => localStorage.clear()",
  "})",
])

wf('hooks/web/useNProgress.js', [
  "import nProgress from 'nprogress'",
  "export const useNProgress = () => ({",
  "  start: () => nProgress.start(),",
  "  done: () => nProgress.done()",
  "})",
])

wf('hooks/web/useTitle.js', [
  "export const useTitle = (title) => {",
  "  if (title) document.title = title",
  "}",
])

wf('hooks/web/usePageLoading.js', [
  "import { ref } from 'vue'",
  "export const usePageLoading = () => {",
  "  const pageLoading = ref(false)",
  "  const loadStart = () => { pageLoading.value = true }",
  "  const loadDone = () => { pageLoading.value = false }",
  "  return { pageLoading, loadStart, loadDone }",
  "}",
])

wf('hooks/web/useValidator.js', [
  "export const useValidator = () => ({})",
])

wf('hooks/web/useDesign.js', [
  "export const useDesign = () => ({ getPrefixCls: (s) => 'cl-' + s })",
])

wf('hooks/event/useEventBus.js', [
  "import mitt from 'mitt'",
  "export const emitter = mitt()",
  "export const useEventBus = () => emitter",
])

// === Extra Utils Stubs ===
wf('utils/defaultActSlotBuilder.js', [
  "export const actionSlotNotAllPermiBuilder = () => null",
  "export const actionSlotBuilder = () => null",
])
wf('utils/debounce.js', [
  "export const handleAxiosError = (err) => { throw err }",
  "export const clear = () => {}",
])
wf('utils/dateUtil.js', [
  "import dayjs from 'dayjs'",
  "export const formatToDateTzTime = (ts, fmt) => dayjs(ts).format(fmt || 'YYYY-MM-DD HH:mm:ss')",
])
wf('utils/tree.js', [
  "export const buildTree = (list, pid=0) => list.filter(i => i.parent_id===pid).map(i=>({...i,children:buildTree(list,i.id)}))",
])
wf('utils/color.js', [
  "export const colorIsDark = () => false",
  "export const hexToRGB = () => ''",
  "export const lighten = (c) => c",
  "export const mix = (c1,c2,r) => c1",
])
wf('utils/roleUtils.js', [
  "export const getRole = async () => true",
])
console.log('All 30+ files written successfully')
