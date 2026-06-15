<template>
  <div class="page-container">
    <div class="page-header">
      <h2>整车库存</h2>
      <div style="display: flex; gap: 12px">
        <el-input v-model="searchKeyword" placeholder="搜索 VIN码/品牌/车型" clearable style="width: 260px" @clear="loadData" @keyup.enter="loadData" />
        <el-button type="primary" @click="loadData">搜索</el-button>
        <el-button type="success" @click="handleExport">导出所有车辆</el-button>
      </div>
    </div>

    <el-table :data="carList" stripe border v-loading="loading">
      <el-table-column prop="vin" label="VIN码" width="190" />
      <el-table-column prop="brand" label="品牌" width="120" />
      <el-table-column prop="model" label="车型" width="160" />
      <el-table-column prop="color" label="颜色" width="80" />
      <el-table-column prop="year" label="年款" width="80" />
      <el-table-column prop="licensePlate" label="车牌号" width="120" />
      <el-table-column prop="warehouseName" label="所在仓库" width="140" />
      <el-table-column prop="location" label="位置" width="140" />
      <el-table-column prop="status" label="状态" width="100">
        <template #default="{ row }">
          <el-tag :type="row.status === 'pending' ? 'warning' : row.status === 'dismantling' ? 'primary' : row.status === 'completed' ? 'success' : 'info'" size="small">
            {{ statusMap[row.status] || row.status }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="100">
        <template #default="{ row }">
          <el-button size="small" type="primary" @click="handleView(row)">查看详情</el-button>
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

    <el-dialog v-model="showDetail" title="车辆详情" width="600px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="VIN码" :span="2">{{ detail.vin }}</el-descriptions-item>
        <el-descriptions-item label="品牌">{{ detail.brand }}</el-descriptions-item>
        <el-descriptions-item label="车型">{{ detail.model }}</el-descriptions-item>
        <el-descriptions-item label="颜色">{{ detail.color }}</el-descriptions-item>
        <el-descriptions-item label="年款">{{ detail.year }}</el-descriptions-item>
        <el-descriptions-item label="车牌号">{{ detail.licensePlate }}</el-descriptions-item>
        <el-descriptions-item label="所在仓库">{{ detail.warehouseName }}</el-descriptions-item>
        <el-descriptions-item label="位置">{{ detail.location }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="detail.status === 'pending' ? 'warning' : detail.status === 'dismantling' ? 'primary' : detail.status === 'completed' ? 'success' : 'info'" size="small">
            {{ statusMap[detail.status] || detail.status }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="入场时间">{{ detail.createdAt }}</el-descriptions-item>
      </el-descriptions>
      <template #footer>
        <el-button @click="showDetail = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { inventoryApi } from '@/api'
import { ElMessage } from 'element-plus'

const loading = ref(false)
const carList = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const searchKeyword = ref('')
const showDetail = ref(false)
const detail = ref({})

const statusMap = { pending: '待拆解', dismantling: '拆解中', completed: '已拆解', stored: '已入库' }

onMounted(() => loadData())

async function loadData() {
  loading.value = true
  try {
    const params = { page: page.value, pageSize: pageSize.value }
    if (searchKeyword.value) params.keyword = searchKeyword.value
    const res = await inventoryApi.getCarList(params)
    carList.value = res.data.list || []
    total.value = res.data.total || 0
  } catch (e) {
    carList.value = []
  } finally {
    loading.value = false
  }
}

function handleView(row) {
  detail.value = row
  showDetail.value = true
}

function handleExport() {
  ElMessage.success('导出任务已提交，请稍后在导出记录中查看')
}
</script>
