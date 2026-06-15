<template>
  <div class="page-container">
    <div class="page-header">
      <h2>危固废管理</h2>
      <div style="display: flex; gap: 12px">
        <el-select v-model="typeFilter" placeholder="废料类型" clearable style="width: 150px" @change="loadData">
          <el-option label="危险废物" value="hazardous" />
          <el-option label="固体废物" value="solid" />
        </el-select>
        <el-select v-model="warehouseFilter" placeholder="选择仓库" clearable style="width: 180px" @change="loadData">
          <el-option v-for="w in warehouseOptions" :key="w.id" :label="w.warehouseName" :value="w.id" />
        </el-select>
        <el-button type="primary" @click="loadData">搜索</el-button>
      </div>
    </div>

    <el-table :data="wasteList" stripe border v-loading="loading">
      <el-table-column prop="wasteNo" label="废料编号" width="150" />
      <el-table-column prop="wasteType" label="废料类型" width="120">
        <template #default="{ row }">
          <el-tag :type="row.wasteType === 'hazardous' ? 'danger' : 'warning'" size="small">
            {{ row.wasteType === 'hazardous' ? '危险废物' : '固体废物' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="quantity" label="数量(kg)" width="120" />
      <el-table-column prop="warehouseName" label="仓库" width="160" />
      <el-table-column prop="createdAt" label="入库时间" width="170" />
      <el-table-column label="操作" width="100">
        <template #default="{ row }">
          <el-button size="small" type="primary" @click="handleView(row)">详情</el-button>
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

    <el-dialog v-model="showDetail" title="废料详情" width="500px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="废料编号">{{ detail.wasteNo }}</el-descriptions-item>
        <el-descriptions-item label="废料类型">
          <el-tag :type="detail.wasteType === 'hazardous' ? 'danger' : 'warning'" size="small">
            {{ detail.wasteType === 'hazardous' ? '危险废物' : '固体废物' }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="数量(kg)">{{ detail.quantity }}</el-descriptions-item>
        <el-descriptions-item label="仓库">{{ detail.warehouseName }}</el-descriptions-item>
        <el-descriptions-item label="入库时间">{{ detail.createdAt }}</el-descriptions-item>
        <el-descriptions-item label="更新时间">{{ detail.updatedAt }}</el-descriptions-item>
      </el-descriptions>
      <template #footer>
        <el-button @click="showDetail = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { inventoryApi, warehouseApi } from '@/api'

const loading = ref(false)
const wasteList = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const typeFilter = ref('')
const warehouseFilter = ref(null)
const warehouseOptions = ref([])
const showDetail = ref(false)
const detail = ref({})

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
    if (typeFilter.value) params.wasteType = typeFilter.value
    if (warehouseFilter.value) params.warehouseId = warehouseFilter.value
    const res = await inventoryApi.getWasteList(params)
    wasteList.value = res.data.list || []
    total.value = res.data.total || 0
  } catch (e) {
    wasteList.value = []
  } finally {
    loading.value = false
  }
}

function handleView(row) {
  detail.value = row
  showDetail.value = true
}
</script>
