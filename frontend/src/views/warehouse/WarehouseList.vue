<template>
  <div class="page-container">
    <div class="page-header">
      <h2>仓库信息</h2>
      <div style="display: flex; gap: 12px">
        <el-input v-model="searchKeyword" placeholder="搜索仓库名称/编码" clearable style="width: 240px" @clear="loadData" @keyup.enter="loadData" />
        <el-button type="primary" @click="loadData">搜索</el-button>
        <el-button type="primary" @click="handleAdd">新增仓库</el-button>
      </div>
    </div>

    <el-table :data="warehouseList" stripe border v-loading="loading">
      <el-table-column prop="warehouseCode" label="仓库编码" width="140" />
      <el-table-column prop="warehouseName" label="仓库名称" min-width="160" />
      <el-table-column prop="address" label="地址" min-width="220" show-overflow-tooltip />
      <el-table-column prop="companyName" label="所属公司" width="160" />
      <el-table-column prop="status" label="状态" width="80">
        <template #default="{ row }">
          <el-tag :type="row.status === 1 ? 'success' : 'danger'" size="small">
            {{ row.status === 1 ? '启用' : '停用' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="createdAt" label="创建时间" width="170" />
      <el-table-column label="操作" width="260" fixed="right">
        <template #default="{ row }">
          <el-button size="small" @click="handleEdit(row)">编辑</el-button>
          <el-button size="small" :type="row.status === 1 ? 'warning' : 'success'" @click="handleToggleStatus(row)">
            {{ row.status === 1 ? '停用' : '启用' }}
          </el-button>
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

    <el-dialog v-model="showDialog" :title="editId ? '编辑仓库' : '新增仓库'" width="600px">
      <el-form ref="formRef" :model="form" label-width="100px">
        <el-form-item label="仓库编码" prop="warehouseCode">
          <el-input v-model="form.warehouseCode" :disabled="!!editId" />
        </el-form-item>
        <el-form-item label="仓库名称" prop="warehouseName">
          <el-input v-model="form.warehouseName" />
        </el-form-item>
        <el-form-item label="仓库地址" prop="address">
          <el-input v-model="form.address" type="textarea" :rows="2" />
        </el-form-item>
        <el-form-item label="所属公司" prop="companyName">
          <el-input v-model="form.companyName" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio :value="1">启用</el-radio>
            <el-radio :value="0">停用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showDialog = false">取消</el-button>
        <el-button type="primary" @click="handleSave" :loading="saving">保存</el-button>
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
const warehouseList = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const showDialog = ref(false)
const editId = ref(null)
const formRef = ref(null)
const searchKeyword = ref('')

const form = reactive({
  warehouseCode: '',
  warehouseName: '',
  address: '',
  companyName: '',
  status: 1,
})

onMounted(() => loadData())

async function loadData() {
  loading.value = true
  try {
    const params = { page: page.value, pageSize: pageSize.value }
    if (searchKeyword.value) params.keyword = searchKeyword.value
    const res = await warehouseApi.getWarehouseList(params)
    warehouseList.value = res.data.list || []
    total.value = res.data.total || 0
  } catch (e) {
    warehouseList.value = []
  } finally {
    loading.value = false
  }
}

function handleAdd() {
  editId.value = null
  Object.assign(form, { warehouseCode: '', warehouseName: '', address: '', companyName: '', status: 1 })
  showDialog.value = true
}

function handleEdit(row) {
  editId.value = row.id
  Object.assign(form, {
    warehouseCode: row.warehouseCode,
    warehouseName: row.warehouseName,
    address: row.address,
    companyName: row.companyName,
    status: row.status,
  })
  showDialog.value = true
}

async function handleDelete(id) {
  try {
    await ElMessageBox.confirm('确定删除该仓库？删除前请确认无关联库存。')
    await warehouseApi.deleteWarehouse(id)
    ElMessage.success('删除成功')
    loadData()
  } catch (e) {}
}

async function handleToggleStatus(row) {
  const newStatus = row.status === 1 ? 0 : 1
  const label = newStatus === 1 ? '启用' : '停用'
  try {
    await ElMessageBox.confirm(`确定${label}该仓库？`)
    await warehouseApi.updateWarehouse({ id: row.id, status: newStatus })
    ElMessage.success(`${label}成功`)
    loadData()
  } catch (e) {}
}

async function handleSave() {
  saving.value = true
  try {
    if (editId.value) {
      await warehouseApi.updateWarehouse({ id: editId.value, ...form })
    } else {
      await warehouseApi.createWarehouse(form)
    }
    ElMessage.success('保存成功')
    showDialog.value = false
    loadData()
  } catch (e) {
  } finally {
    saving.value = false
  }
}
</script>
