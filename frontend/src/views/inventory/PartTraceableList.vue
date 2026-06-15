<template>
  <div class="page-container">
    <div class="page-header">
      <h2>配件 - 溯源件</h2>
      <div style="display: flex; gap: 12px">
        <el-input v-model="searchKeyword" placeholder="搜索配件名称/VIN码" clearable style="width: 220px" @clear="loadData" @keyup.enter="loadData" />
        <el-select v-model="typeFilter" placeholder="配件类型" clearable style="width: 140px" @change="loadData">
          <el-option label="发动机" value="engine" />
          <el-option label="变速箱" value="transmission" />
          <el-option label="轮胎" value="tire" />
          <el-option label="轮毂" value="wheel" />
          <el-option label="车身件" value="body" />
          <el-option label="电子件" value="electronic" />
          <el-option label="其他" value="other" />
        </el-select>
        <el-button type="primary" @click="loadData">搜索</el-button>
      </div>
    </div>

    <el-table :data="partList" stripe border v-loading="loading">
      <el-table-column prop="partName" label="配件名称" min-width="140" />
      <el-table-column prop="partType" label="配件类型" width="100">
        <template #default="{ row }">
          <el-tag size="small">{{ typeMap[row.partType] || row.partType }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="vehicleModel" label="车型" width="130" />
      <el-table-column prop="vehicleSeries" label="车系" width="130" />
      <el-table-column prop="vin" label="来源VIN" width="190" />
      <el-table-column prop="quantity" label="库存数量" width="90" />
      <el-table-column prop="warehouseName" label="仓库" width="140" />
      <el-table-column prop="location" label="位置" width="120" />
      <el-table-column label="操作" width="100">
        <template #default="{ row }">
          <el-button size="small" type="warning" @click="handleOutbound(row)">出库</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      v-model:current-page="page"
      :page-size="pageSize"
      :total="total"
      layout="total, prev, pager, next, sizes"
      style="margin-top: 20px; justify-content: flex-end"
      @current-change="loadData"
      @size-change="loadData"
    />

    <el-dialog v-model="showOutbound" title="溯源件出库" width="400px">
      <el-form ref="formRef" :model="outboundForm" label-width="80px">
        <el-form-item label="配件名称">
          <span>{{ outboundForm.partName }}</span>
        </el-form-item>
        <el-form-item label="当前库存">
          <span>{{ outboundForm.currentQty }}</span>
        </el-form-item>
        <el-form-item label="出库数量" prop="quantity">
          <el-input-number v-model="outboundForm.quantity" :min="1" :max="outboundForm.currentQty" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showOutbound = false">取消</el-button>
        <el-button type="warning" @click="handleOutboundConfirm" :loading="saving">确认出库</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { inventoryApi } from '@/api'
import { ElMessage } from 'element-plus'

const loading = ref(false)
const saving = ref(false)
const partList = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const searchKeyword = ref('')
const typeFilter = ref('')
const showOutbound = ref(false)
const formRef = ref(null)

const typeMap = { engine: '发动机', transmission: '变速箱', tire: '轮胎', wheel: '轮毂', body: '车身件', electronic: '电子件', other: '其他' }

const outboundForm = reactive({
  id: null,
  partName: '',
  currentQty: 0,
  quantity: 1,
})

onMounted(() => loadData())

async function loadData() {
  loading.value = true
  try {
    const params = { page: page.value, pageSize: pageSize.value }
    if (searchKeyword.value) params.keyword = searchKeyword.value
    if (typeFilter.value) params.partType = typeFilter.value
    const res = await inventoryApi.getPartTraceableList(params)
    partList.value = res.data.list || []
    total.value = res.data.total || 0
  } catch (e) {
    partList.value = []
  } finally {
    loading.value = false
  }
}

function handleOutbound(row) {
  outboundForm.id = row.id
  outboundForm.partName = row.partName
  outboundForm.currentQty = row.quantity
  outboundForm.quantity = 1
  showOutbound.value = true
}

async function handleOutboundConfirm() {
  saving.value = true
  try {
    await inventoryApi.partTraceableOutbound({ id: outboundForm.id, quantity: outboundForm.quantity })
    ElMessage.success('出库成功')
    showOutbound.value = false
    loadData()
  } catch (e) {
  } finally {
    saving.value = false
  }
}
</script>
