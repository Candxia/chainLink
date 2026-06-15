<template>
  <div class="page-container">
    <div class="page-header">
      <h2>调拨记录</h2>
      <div style="display: flex; gap: 12px">
        <el-select v-model="typeFilter" placeholder="调拨类型" clearable style="width: 140px" @change="loadData">
          <el-option label="原材料" value="raw_material" />
          <el-option label="配件" value="part" />
          <el-option label="整车" value="car" />
          <el-option label="危固废" value="waste" />
        </el-select>
        <el-select v-model="statusFilter" placeholder="审核状态" clearable style="width: 140px" @change="loadData">
          <el-option label="审核中" value="pending" />
          <el-option label="已通过" value="approved" />
          <el-option label="未通过" value="rejected" />
        </el-select>
        <el-button type="primary" @click="loadData">搜索</el-button>
      </div>
    </div>

    <el-table :data="transferList" stripe border v-loading="loading">
      <el-table-column prop="transferNo" label="调拨编号" width="170" />
      <el-table-column prop="transferType" label="调拨类型" width="100">
        <template #default="{ row }">
          <el-tag size="small">{{ typeMap[row.transferType] || row.transferType }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="transferTime" label="调拨时间" width="170" />
      <el-table-column prop="fromWarehouse" label="调出仓库" width="150" />
      <el-table-column prop="toWarehouse" label="调入仓库" width="150" />
      <el-table-column prop="goodsType" label="商品类型" min-width="120" />
      <el-table-column prop="totalCategories" label="总品类" width="80" />
      <el-table-column prop="totalQuantity" label="总数量" width="80" />
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

    <el-dialog v-model="showDetail" title="调拨详情" width="800px">
      <el-descriptions :column="3" border style="margin-bottom: 16px">
        <el-descriptions-item label="调拨编号">{{ detail.transferNo }}</el-descriptions-item>
        <el-descriptions-item label="调拨类型">{{ typeMap[detail.transferType] }}</el-descriptions-item>
        <el-descriptions-item label="调拨时间">{{ detail.transferTime }}</el-descriptions-item>
        <el-descriptions-item label="调出仓库">{{ detail.fromWarehouse }}</el-descriptions-item>
        <el-descriptions-item label="调入仓库">{{ detail.toWarehouse }}</el-descriptions-item>
        <el-descriptions-item label="商品类型">{{ detail.goodsType }}</el-descriptions-item>
        <el-descriptions-item label="总品类">{{ detail.totalCategories }}</el-descriptions-item>
        <el-descriptions-item label="总数量">{{ detail.totalQuantity }}</el-descriptions-item>
        <el-descriptions-item label="审核状态">
          <el-tag :type="detail.status === 'approved' ? 'success' : detail.status === 'rejected' ? 'danger' : 'warning'" size="small">
            {{ statusMap[detail.status] }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="备注" :span="3">{{ detail.remark || '-' }}</el-descriptions-item>
      </el-descriptions>
      <h4 style="margin-bottom: 12px">调拨明细</h4>
      <el-table :data="detail.items || []" border size="small">
        <el-table-column prop="materialName" label="物料名称" min-width="140" />
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
const transferList = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const typeFilter = ref('')
const statusFilter = ref('')
const showDetail = ref(false)
const detail = ref({ items: [] })

const typeMap = { raw_material: '原材料', part: '配件', car: '整车', waste: '危固废' }
const statusMap = { pending: '审核中', approved: '已通过', rejected: '未通过' }

onMounted(() => loadData())

async function loadData() {
  loading.value = true
  try {
    const params = { page: page.value, pageSize: pageSize.value }
    if (typeFilter.value) params.transferType = typeFilter.value
    if (statusFilter.value) params.status = statusFilter.value
    const res = await recordApi.getTransferList(params)
    transferList.value = res.data.list || []
    total.value = res.data.total || 0
  } catch (e) {
    transferList.value = []
  } finally {
    loading.value = false
  }
}

async function handleView(row) {
  try {
    const res = await recordApi.getTransferDetail(row.id)
    detail.value = res.data || row
  } catch (e) {
    detail.value = row
  }
  showDetail.value = true
}
</script>
