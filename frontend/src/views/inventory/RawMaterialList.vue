<template>
  <div class="page-container">
    <div class="page-header">
      <h2>原材料库存</h2>
      <div style="display: flex; gap: 12px">
        <el-select v-model="typeFilter" placeholder="原材料类型" clearable style="width: 150px" @change="loadData">
          <el-option label="钢铁" value="steel" />
          <el-option label="铜材" value="copper" />
          <el-option label="铝材" value="aluminum" />
          <el-option label="塑料" value="plastic" />
          <el-option label="橡胶" value="rubber" />
          <el-option label="其他" value="other" />
        </el-select>
        <el-select v-model="warehouseFilter" placeholder="选择仓库" clearable style="width: 180px" @change="loadData">
          <el-option v-for="w in warehouseOptions" :key="w.id" :label="w.warehouseName" :value="w.id" />
        </el-select>
        <el-button type="primary" @click="loadData">搜索</el-button>
      </div>
    </div>

    <el-table :data="materialList" stripe border v-loading="loading">
      <el-table-column prop="materialNo" label="编号" width="150" />
      <el-table-column prop="materialType" label="原材料类型" width="120">
        <template #default="{ row }">
          <el-tag size="small">{{ typeMap[row.materialType] || row.materialType }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="quantity" label="数量(kg)" width="120" />
      <el-table-column prop="warehouseName" label="仓库" width="160" />
      <el-table-column prop="createdAt" label="入库时间" width="170" />
      <el-table-column prop="updatedAt" label="更新时间" width="170" />
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
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { inventoryApi, warehouseApi } from '@/api'

const loading = ref(false)
const materialList = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const typeFilter = ref('')
const warehouseFilter = ref(null)
const warehouseOptions = ref([])

const typeMap = { steel: '钢铁', copper: '铜材', aluminum: '铝材', plastic: '塑料', rubber: '橡胶', other: '其他' }

onMounted(async () => {
  await loadWarehouses()
  await loadData()
})

async function loadWarehouses() {
  try {
    const res = await warehouseApi.getWarehouseList({ page: 1, pageSize: 999 })
    warehouseOptions.value = res.data.list || []
  } catch (e) {
    warehouseOptions.value = []
  }
}

async function loadData() {
  loading.value = true
  try {
    const params = { page: page.value, pageSize: pageSize.value }
    if (typeFilter.value) params.materialType = typeFilter.value
    if (warehouseFilter.value) params.warehouseId = warehouseFilter.value
    const res = await inventoryApi.getRawMaterialList(params)
    materialList.value = res.data.list || []
    total.value = res.data.total || 0
  } catch (e) {
    materialList.value = []
  } finally {
    loading.value = false
  }
}
</script>
