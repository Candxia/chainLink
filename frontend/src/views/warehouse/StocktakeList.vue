<template>
  <div class="page-container">
    <div class="page-header">
      <h2>仓库盘存</h2>
      <div style="display: flex; gap: 12px">
        <el-select v-model="statusFilter" placeholder="盘点状态" clearable style="width: 140px" @change="loadData">
          <el-option label="待盘点" value="pending" />
          <el-option label="有差异" value="difference" />
          <el-option label="无差异" value="normal" />
        </el-select>
        <el-select v-model="typeFilter" placeholder="盘点类型" clearable style="width: 140px" @change="loadData">
          <el-option label="定期盘点" value="regular" />
          <el-option label="不定期盘点" value="irregular" />
        </el-select>
        <el-button type="primary" @click="handleAdd">新增盘存</el-button>
      </div>
    </div>

    <el-table :data="stocktakeList" stripe border v-loading="loading">
      <el-table-column prop="stocktakeNo" label="盘存编号" width="170" />
      <el-table-column prop="type" label="盘点类型" width="120">
        <template #default="{ row }">
          <el-tag :type="row.type === 'regular' ? 'primary' : 'warning'" size="small">
            {{ row.type === 'regular' ? '定期盘点' : '不定期盘点' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="status" label="盘点状态" width="100">
        <template #default="{ row }">
          <el-tag :type="row.status === 'pending' ? 'info' : row.status === 'difference' ? 'danger' : 'success'" size="small">
            {{ row.status === 'pending' ? '待盘点' : row.status === 'difference' ? '有差异' : '无差异' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="operator" label="盘点人员" width="120" />
      <el-table-column prop="stocktakeTime" label="盘点时间" width="170" />
      <el-table-column label="操作" width="160">
        <template #default="{ row }">
          <el-button size="small" type="primary" @click="handleView(row)">查看详情</el-button>
          <el-button size="small" type="danger" @click="handleDelete(row.id)">删除</el-button>
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

    <!-- 新增盘存 -->
    <el-dialog v-model="showDialog" title="新增盘存" width="500px">
      <el-form ref="formRef" :model="form" label-width="100px">
        <el-form-item label="盘点类型" prop="type">
          <el-select v-model="form.type">
            <el-option label="定期盘点" value="regular" />
            <el-option label="不定期盘点" value="irregular" />
          </el-select>
        </el-form-item>
        <el-form-item label="盘点人员" prop="operator">
          <el-input v-model="form.operator" />
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="form.remark" type="textarea" :rows="3" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showDialog = false">取消</el-button>
        <el-button type="primary" @click="handleSave" :loading="saving">确认创建</el-button>
      </template>
    </el-dialog>

    <!-- 详情弹窗 -->
    <el-dialog v-model="showDetail" title="盘存详情" width="700px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="盘存编号">{{ detail.stocktakeNo }}</el-descriptions-item>
        <el-descriptions-item label="盘点类型">{{ detail.type === 'regular' ? '定期盘点' : '不定期盘点' }}</el-descriptions-item>
        <el-descriptions-item label="盘点状态">
          <el-tag :type="detail.status === 'pending' ? 'info' : detail.status === 'difference' ? 'danger' : 'success'" size="small">
            {{ detail.status === 'pending' ? '待盘点' : detail.status === 'difference' ? '有差异' : '无差异' }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="盘点人员">{{ detail.operator }}</el-descriptions-item>
        <el-descriptions-item label="盘点时间">{{ detail.stocktakeTime }}</el-descriptions-item>
        <el-descriptions-item label="备注">{{ detail.remark || '-' }}</el-descriptions-item>
      </el-descriptions>
      <el-divider />
      <h4 style="margin-bottom: 12px">盘点明细</h4>
      <el-table :data="detail.items || []" border size="small">
        <el-table-column prop="materialName" label="物料名称" />
        <el-table-column prop="expectedQty" label="账面数量" />
        <el-table-column prop="actualQty" label="实际数量" />
        <el-table-column prop="diffQty" label="差异" />
        <el-table-column prop="remark" label="说明" />
      </el-table>
      <template #footer>
        <el-button @click="showDetail = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { warehouseApi } from '@/api'
import { ElMessage, ElMessageBox } from 'element-plus'

const loading = ref(false)
const saving = ref(false)
const stocktakeList = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const statusFilter = ref('')
const typeFilter = ref('')
const showDialog = ref(false)
const showDetail = ref(false)
const formRef = ref(null)
const detail = ref({ items: [] })

const form = reactive({
  type: 'regular',
  operator: '',
  remark: '',
})

onMounted(() => loadData())

async function loadData() {
  loading.value = true
  try {
    const params = { page: page.value, pageSize: pageSize.value }
    if (statusFilter.value) params.status = statusFilter.value
    if (typeFilter.value) params.type = typeFilter.value
    const res = await warehouseApi.getStocktakeList(params)
    stocktakeList.value = res.data.list || []
    total.value = res.data.total || 0
  } catch (e) {
    stocktakeList.value = []
  } finally {
    loading.value = false
  }
}

function handleAdd() {
  Object.assign(form, { type: 'regular', operator: '', remark: '' })
  showDialog.value = true
}

async function handleSave() {
  saving.value = true
  try {
    await warehouseApi.createStocktake(form)
    ElMessage.success('创建成功')
    showDialog.value = false
    loadData()
  } catch (e) {
  } finally {
    saving.value = false
  }
}

async function handleDelete(id) {
  try {
    await ElMessageBox.confirm('确定删除该盘存记录？')
    await warehouseApi.deleteStocktake(id)
    ElMessage.success('删除成功')
    loadData()
  } catch (e) {}
}

async function handleView(row) {
  try {
    const res = await warehouseApi.getStocktakeList({ id: row.id })
    detail.value = res.data || row
  } catch (e) {
    detail.value = row
  }
  showDetail.value = true
}
</script>
