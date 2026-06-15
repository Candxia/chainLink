<template>
  <div class="page-container">
    <div class="page-header">
      <h2>出库记录</h2>
      <div style="display: flex; gap: 12px">
        <el-select v-model="typeFilter" placeholder="出库类型" clearable style="width: 140px" @change="loadData">
          <el-option label="原材料" value="raw_material" />
          <el-option label="整车" value="car" />
          <el-option label="配件" value="part" />
          <el-option label="危固废" value="waste" />
        </el-select>
        <el-select v-model="methodFilter" placeholder="出库方式" clearable style="width: 150px" @change="loadData">
          <el-option label="销售出库" value="sale" />
          <el-option label="调拨出库" value="transfer" />
          <el-option label="报废出库" value="scrap" />
          <el-option label="其他" value="other" />
        </el-select>
        <el-select v-model="statusFilter" placeholder="审核状态" clearable style="width: 140px" @change="loadData">
          <el-option label="审核中" value="pending" />
          <el-option label="已通过" value="approved" />
          <el-option label="未通过" value="rejected" />
        </el-select>
        <el-button type="primary" @click="loadData">搜索</el-button>
      </div>
    </div>

    <el-table :data="outboundList" stripe border v-loading="loading">
      <el-table-column prop="outboundNo" label="出库编号" width="170" />
      <el-table-column prop="outboundType" label="出库类型" width="100">
        <template #default="{ row }">
          <el-tag size="small">{{ typeMap[row.outboundType] || row.outboundType }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="outboundMethod" label="出库方式" width="100">
        <template #default="{ row }">
          {{ methodMap[row.outboundMethod] || row.outboundMethod }}
        </template>
      </el-table-column>
      <el-table-column prop="outboundTime" label="出库时间" width="170" />
      <el-table-column prop="warehouseName" label="出库仓库" width="140" />
      <el-table-column prop="operator" label="出库人" width="100" />
      <el-table-column prop="totalQuantity" label="出库数量" width="90" />
      <el-table-column prop="status" label="审核状态" width="100">
        <template #default="{ row }">
          <el-tag :type="row.status === 'approved' ? 'success' : row.status === 'rejected' ? 'danger' : 'warning'" size="small">
            {{ statusMap[row.status] || row.status }}
          </el-tag>
        </template>
      </el-table-column>
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

    <el-dialog v-model="showDetail" title="出库详情" width="800px">
      <el-descriptions :column="3" border style="margin-bottom: 16px">
        <el-descriptions-item label="出库编号">{{ detail.outboundNo }}</el-descriptions-item>
        <el-descriptions-item label="出库类型">{{ typeMap[detail.outboundType] }}</el-descriptions-item>
        <el-descriptions-item label="出库方式">{{ methodMap[detail.outboundMethod] }}</el-descriptions-item>
        <el-descriptions-item label="出库时间">{{ detail.outboundTime }}</el-descriptions-item>
        <el-descriptions-item label="出库仓库">{{ detail.warehouseName }}</el-descriptions-item>
        <el-descriptions-item label="出库人">{{ detail.operator }}</el-descriptions-item>
        <el-descriptions-item label="总数量">{{ detail.totalQuantity }}</el-descriptions-item>
        <el-descriptions-item label="审核状态">
          <el-tag :type="detail.status === 'approved' ? 'success' : detail.status === 'rejected' ? 'danger' : 'warning'" size="small">
            {{ statusMap[detail.status] }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="备注" :span="3">{{ detail.remark || '-' }}</el-descriptions-item>
      </el-descriptions>
      <h4 style="margin-bottom: 12px">出库明细</h4>
      <el-table :data="detail.items || []" border size="small">
        <el-table-column prop="materialName" label="物料名称" min-width="140" />
        <el-table-column prop="materialType" label="类型" width="100" />
        <el-table-column prop="quantity" label="数量" width="80" />
        <el-table-column prop="unit" label="单位" width="60" />
      </el-table>
      <template #footer>
        <el-button @click="showDetail = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { recordApi } from '@/api'

const loading = ref(false)
const outboundList = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const typeFilter = ref('')
const methodFilter = ref('')
const statusFilter = ref('')
const showDetail = ref(false)
const detail = ref({ items: [] })

const typeMap = { raw_material: '原材料', car: '整车', part: '配件', waste: '危固废' }
const methodMap = { sale: '销售出库', transfer: '调拨出库', scrap: '报废出库', other: '其他' }
const statusMap = { pending: '审核中', approved: '已通过', rejected: '未通过' }

onMounted(() => loadData())

async function loadData() {
  loading.value = true
  try {
    const params = { page: page.value, pageSize: pageSize.value }
    if (typeFilter.value) params.outboundType = typeFilter.value
    if (methodFilter.value) params.outboundMethod = methodFilter.value
    if (statusFilter.value) params.status = statusFilter.value
    const res = await recordApi.getOutboundList(params)
    outboundList.value = res.data.list || []
    total.value = res.data.total || 0
  } catch (e) {
    outboundList.value = []
  } finally {
    loading.value = false
  }
}

async function handleView(row) {
  try {
    const res = await recordApi.getOutboundDetail(row.id)
    detail.value = res.data || row
  } catch (e) {
    detail.value = row
  }
  showDetail.value = true
}
</script>
