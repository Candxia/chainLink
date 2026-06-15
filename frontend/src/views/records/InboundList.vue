<template>
  <div class="page-container">
    <div class="page-header">
      <h2>入库记录</h2>
      <div style="display: flex; gap: 12px">
        <el-select v-model="typeFilter" placeholder="入库类型" clearable style="width: 140px" @change="loadData">
          <el-option label="原材料" value="raw_material" />
          <el-option label="整车" value="car" />
          <el-option label="配件" value="part" />
          <el-option label="危固废" value="waste" />
        </el-select>
        <el-select v-model="methodFilter" placeholder="入库方式" clearable style="width: 150px" @change="loadData">
          <el-option label="拆解入库" value="dismantle" />
          <el-option label="调拨入库" value="transfer" />
          <el-option label="采购入库" value="purchase" />
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

    <el-table :data="inboundList" stripe border v-loading="loading">
      <el-table-column prop="inboundNo" label="入库编号" width="170" />
      <el-table-column prop="inboundType" label="入库类型" width="100">
        <template #default="{ row }">
          <el-tag size="small">{{ typeMap[row.inboundType] || row.inboundType }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="inboundMethod" label="入库方式" width="100">
        <template #default="{ row }">
          {{ methodMap[row.inboundMethod] || row.inboundMethod }}
        </template>
      </el-table-column>
      <el-table-column prop="inboundTime" label="入库时间" width="170" />
      <el-table-column prop="warehouseName" label="目标仓库" width="140" />
      <el-table-column prop="operator" label="入库人" width="100" />
      <el-table-column prop="totalQuantity" label="入库数目" width="90" />
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

    <el-dialog v-model="showDetail" title="入库详情" width="800px">
      <el-descriptions :column="3" border style="margin-bottom: 16px">
        <el-descriptions-item label="入库编号">{{ detail.inboundNo }}</el-descriptions-item>
        <el-descriptions-item label="入库类型">{{ typeMap[detail.inboundType] }}</el-descriptions-item>
        <el-descriptions-item label="入库方式">{{ methodMap[detail.inboundMethod] }}</el-descriptions-item>
        <el-descriptions-item label="入库时间">{{ detail.inboundTime }}</el-descriptions-item>
        <el-descriptions-item label="目标仓库">{{ detail.warehouseName }}</el-descriptions-item>
        <el-descriptions-item label="入库人">{{ detail.operator }}</el-descriptions-item>
        <el-descriptions-item label="总数目">{{ detail.totalQuantity }}</el-descriptions-item>
        <el-descriptions-item label="审核状态">
          <el-tag :type="detail.status === 'approved' ? 'success' : detail.status === 'rejected' ? 'danger' : 'warning'" size="small">
            {{ statusMap[detail.status] }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="备注" :span="3">{{ detail.remark || '-' }}</el-descriptions-item>
      </el-descriptions>
      <h4 style="margin-bottom: 12px">入库明细</h4>
      <el-table :data="detail.items || []" border size="small">
        <el-table-column prop="materialName" label="物料名称" min-width="140" />
        <el-table-column prop="materialType" label="类型" width="100" />
        <el-table-column prop="quantity" label="数量" width="80" />
        <el-table-column prop="unit" label="单位" width="60" />
        <el-table-column prop="location" label="存放位置" width="140" />
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
const inboundList = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const typeFilter = ref('')
const methodFilter = ref('')
const statusFilter = ref('')
const showDetail = ref(false)
const detail = ref({ items: [] })

const typeMap = { raw_material: '原材料', car: '整车', part: '配件', waste: '危固废' }
const methodMap = { dismantle: '拆解入库', transfer: '调拨入库', purchase: '采购入库', other: '其他' }
const statusMap = { pending: '审核中', approved: '已通过', rejected: '未通过' }

onMounted(() => loadData())

async function loadData() {
  loading.value = true
  try {
    const params = { page: page.value, pageSize: pageSize.value }
    if (typeFilter.value) params.inboundType = typeFilter.value
    if (methodFilter.value) params.inboundMethod = methodFilter.value
    if (statusFilter.value) params.status = statusFilter.value
    const res = await recordApi.getInboundList(params)
    inboundList.value = res.data.list || []
    total.value = res.data.total || 0
  } catch (e) {
    inboundList.value = []
  } finally {
    loading.value = false
  }
}

async function handleView(row) {
  try {
    const res = await recordApi.getInboundDetail(row.id)
    detail.value = res.data || row
  } catch (e) {
    detail.value = row
  }
  showDetail.value = true
}
</script>
