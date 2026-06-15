<template>
  <div class="page-container">
    <div class="page-header">
      <h2>区域货架管理</h2>
      <el-select v-model="warehouseId" placeholder="选择仓库" clearable style="width: 240px" @change="loadData">
        <el-option v-for="w in warehouseOptions" :key="w.id" :label="w.warehouseName" :value="w.id" />
      </el-select>
      <el-button type="primary" :disabled="!warehouseId" @click="handleAddArea">新增区域</el-button>
    </div>

    <div v-loading="loading">
      <el-card v-for="area in areaList" :key="area.id" shadow="hover" style="margin-bottom: 16px">
        <template #header>
          <div style="display: flex; justify-content: space-between; align-items: center">
            <span style="font-weight: bold; font-size: 15px">{{ area.areaName }}</span>
            <div>
              <el-button size="small" type="primary" @click="handleAddShelf(area)">新增货架</el-button>
              <el-button size="small" @click="handleEditArea(area)">编辑区域</el-button>
              <el-button size="small" type="danger" @click="handleDeleteArea(area.id)">删除区域</el-button>
            </div>
          </div>
        </template>
        <div v-if="area.shelves && area.shelves.length">
          <el-table :data="area.shelves" stripe border size="small">
            <el-table-column prop="shelfName" label="货架名称" width="160" />
            <el-table-column prop="description" label="货架说明" min-width="240" />
            <el-table-column label="操作" width="180">
              <template #default="{ row }">
                <el-button size="small" @click="handleEditShelf(row)">编辑</el-button>
                <el-button size="small" type="danger" @click="handleDeleteShelf(row.id)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>
        <el-empty v-else description="暂无货架，请新增" :image-size="60" />
      </el-card>
      <el-empty v-if="!areaList.length && warehouseId" description="暂无区域，请新增" />
      <el-empty v-if="!warehouseId" description="请先选择仓库" />
    </div>

    <!-- 区域弹窗 -->
    <el-dialog v-model="showAreaDialog" :title="areaEditId ? '编辑区域' : '新增区域'" width="500px">
      <el-form ref="areaFormRef" :model="areaForm" label-width="80px">
        <el-form-item label="区域名称" prop="areaName">
          <el-input v-model="areaForm.areaName" placeholder="如 A区、B区、AK区" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showAreaDialog = false">取消</el-button>
        <el-button type="primary" @click="handleSaveArea" :loading="saving">保存</el-button>
      </template>
    </el-dialog>

    <!-- 货架弹窗 -->
    <el-dialog v-model="showShelfDialog" :title="shelfEditId ? '编辑货架' : '新增货架'" width="500px">
      <el-form ref="shelfFormRef" :model="shelfForm" label-width="80px">
        <el-form-item label="所属区域">
          <span style="font-weight: bold">{{ currentAreaName }}</span>
        </el-form-item>
        <el-form-item label="货架名称" prop="shelfName">
          <el-input v-model="shelfForm.shelfName" placeholder="如 AE-1、AE-2" />
        </el-form-item>
        <el-form-item label="货架说明" prop="description">
          <el-input v-model="shelfForm.description" type="textarea" :rows="2" placeholder="存放内容说明" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showShelfDialog = false">取消</el-button>
        <el-button type="primary" @click="handleSaveShelf" :loading="saving">保存</el-button>
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
const warehouseId = ref(null)
const warehouseOptions = ref([])
const areaList = ref([])

// Area dialog
const showAreaDialog = ref(false)
const areaEditId = ref(null)
const areaFormRef = ref(null)
const areaForm = reactive({ areaName: '' })

// Shelf dialog
const showShelfDialog = ref(false)
const shelfEditId = ref(null)
const shelfFormRef = ref(null)
const shelfForm = reactive({ shelfName: '', description: '' })
const currentAreaId = ref(null)
const currentAreaName = ref('')

onMounted(async () => {
  try {
    const res = await warehouseApi.getWarehouseList({ page: 1, pageSize: 999 })
    warehouseOptions.value = res.data.list || []
  } catch (e) {
    warehouseOptions.value = []
  }
})

async function loadData() {
  if (!warehouseId.value) {
    areaList.value = []
    return
  }
  loading.value = true
  try {
    const res = await warehouseApi.getAreaList({ warehouseId: warehouseId.value })
    areaList.value = res.data.list || []
  } catch (e) {
    areaList.value = []
  } finally {
    loading.value = false
  }
}

function handleAddArea() {
  areaEditId.value = null
  areaForm.areaName = ''
  showAreaDialog.value = true
}

function handleEditArea(row) {
  areaEditId.value = row.id
  areaForm.areaName = row.areaName
  showAreaDialog.value = true
}

async function handleDeleteArea(id) {
  try {
    await ElMessageBox.confirm('确定删除该区域及其所有货架？')
    await warehouseApi.deleteArea(id)
    ElMessage.success('删除成功')
    loadData()
  } catch (e) {}
}

async function handleSaveArea() {
  saving.value = true
  try {
    const data = { ...areaForm, warehouseId: warehouseId.value }
    if (areaEditId.value) {
      await warehouseApi.updateArea({ id: areaEditId.value, ...data })
    } else {
      await warehouseApi.createArea(data)
    }
    ElMessage.success('保存成功')
    showAreaDialog.value = false
    loadData()
  } catch (e) {
  } finally {
    saving.value = false
  }
}

function handleAddShelf(area) {
  shelfEditId.value = null
  currentAreaId.value = area.id
  currentAreaName.value = area.areaName
  shelfForm.shelfName = ''
  shelfForm.description = ''
  showShelfDialog.value = true
}

function handleEditShelf(row) {
  shelfEditId.value = row.id
  currentAreaId.value = row.areaId
  currentAreaName.value = row.areaName || ''
  shelfForm.shelfName = row.shelfName
  shelfForm.description = row.description || ''
  showShelfDialog.value = true
}

async function handleDeleteShelf(id) {
  try {
    await ElMessageBox.confirm('确定删除该货架？')
    await warehouseApi.deleteShelf(id)
    ElMessage.success('删除成功')
    loadData()
  } catch (e) {}
}

async function handleSaveShelf() {
  saving.value = true
  try {
    const data = { ...shelfForm, areaId: currentAreaId.value }
    if (shelfEditId.value) {
      await warehouseApi.updateShelf({ id: shelfEditId.value, ...data })
    } else {
      await warehouseApi.createShelf(data)
    }
    ElMessage.success('保存成功')
    showShelfDialog.value = false
    loadData()
  } catch (e) {
  } finally {
    saving.value = false
  }
}
</script>
