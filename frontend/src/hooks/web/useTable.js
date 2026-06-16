import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
export const useTable = (fetchApi) => {
  const loading = ref(false)
  const tableData = ref([])
  const total = ref(0)
  const page = ref(1)
  const pageSize = ref(20)
  const searchParams = ref({})
  const getList = async () => {
    loading.value = true
    try {
      const res = await fetchApi({ page: page.value, pageSize: pageSize.value, ...searchParams.value })
      tableData.value = res.data?.list || res.data || []
      total.value = res.data?.total || 0
    } catch (e) { tableData.value = [] } finally { loading.value = false }
  }
  const handleSearch = () => { page.value = 1; getList() }
  const handleReset = () => { searchParams.value = {}; page.value = 1; getList() }
  return { loading, tableData, total, page, pageSize, searchParams, getList, handleSearch, handleReset }
}
