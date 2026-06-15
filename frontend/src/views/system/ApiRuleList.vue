<template>
  <div class="api-rule-page">
    <div class="page-header">
      <h3>API 管理</h3>
      <el-tag type="info">共 {{ total }} 条路由规则</el-tag>
    </div>

    <el-table :data="list" border stripe style="width: 100%" v-loading="loading">
      <el-table-column prop="id" label="ID" width="60" />
      <el-table-column prop="title" label="接口名称" min-width="200">
        <template #default="{ row }">
          <el-input v-if="editId === row.id" v-model="editForm.title" size="small" />
          <span v-else>{{ row.title }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="method" label="请求方式" width="100">
        <template #default="{ row }">
          <el-tag :type="methodTag(row.method)" size="small">{{ row.method }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="path" label="请求路径" min-width="280">
        <template #default="{ row }">
          <code>{{ row.path }}</code>
        </template>
      </el-table-column>
      <el-table-column prop="auth_type" label="权限类型" width="130">
        <template #default="{ row }">
          <el-select v-if="editId === row.id" v-model="editForm.auth_type" size="small">
            <el-option label="需登录" value="require" />
            <el-option label="公开" value="public" />
            <el-option label="管理员" value="admin" />
          </el-select>
          <el-tag v-else :type="authTag(row.auth_type)" size="small">
            {{ authMap[row.auth_type] || row.auth_type }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="enable_log" label="是否日志" width="100">
        <template #default="{ row }">
          <el-switch v-if="editId === row.id" v-model="editForm.enable_log" :active-value="1" :inactive-value="0" />
          <el-tag v-else :type="row.enable_log == 1 ? 'success' : 'info'" size="small">
            {{ row.enable_log == 1 ? '是' : '否' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="150" fixed="right">
        <template #default="{ row }">
          <template v-if="editId === row.id">
            <el-button type="primary" size="small" @click="saveEdit(row)">保存</el-button>
            <el-button size="small" @click="cancelEdit">取消</el-button>
          </template>
          <el-button v-else type="primary" size="small" plain @click="startEdit(row)">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>

    <div class="pagination-wrap">
      <el-pagination
        v-model:current-page="page"
        v-model:page-size="pageSize"
        :page-sizes="[10, 20, 50, 100]"
        :total="total"
        layout="total, sizes, prev, pager, next"
        @size-change="fetchList"
        @current-change="fetchList"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const list = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(20)
const loading = ref(false)
const editId = ref(0)
const editForm = ref({ title: '', auth_type: '', enable_log: 1 })

const authMap = { require: '需登录', public: '公开', admin: '管理员' }

const methodTag = (m) => {
  const map = { GET: 'success', POST: 'primary', PUT: 'warning', DELETE: 'danger' }
  return map[m] || 'info'
}
const authTag = (a) => {
  const map = { require: '', public: 'success', admin: 'danger' }
  return map[a] || ''
}

const fetchList = async () => {
  loading.value = true
  try {
    const res = await axios.get('/api/api-rule/list', { params: { page: page.value, pageSize: pageSize.value } })
    if (res.data.code === 0) {
      list.value = res.data.data.list || []
      total.value = res.data.data.total || 0
    }
  } catch (e) {
    console.error(e)
  }
  loading.value = false
}

const startEdit = (row) => {
  editId.value = row.id
  editForm.value = { title: row.title, auth_type: row.auth_type, enable_log: row.enable_log }
}

const cancelEdit = () => {
  editId.value = 0
  editForm.value = { title: '', auth_type: '', enable_log: 1 }
}

const saveEdit = async (row) => {
  const form = new FormData()
  form.append('id', row.id)
  form.append('title', editForm.value.title)
  form.append('auth_type', editForm.value.auth_type)
  form.append('enable_log', editForm.value.enable_log)

  try {
    const res = await axios.put('/api/api-rule/update', form, {
      headers: { 'Content-Type': 'multipart/form-data' }
    })
    if (res.data.code === 0) {
      ElMessage.success('更新成功')
      editId.value = 0
      fetchList()
    } else {
      ElMessage.error(res.data.message)
    }
  } catch (e) {
    ElMessage.error('更新失败')
  }
}

onMounted(fetchList)
</script>

<style scoped>
.api-rule-page { padding: 20px; }
.page-header { display: flex; align-items: center; gap: 12px; margin-bottom: 16px; }
.page-header h3 { margin: 0; }
code { font-size: 12px; background: #f5f7fa; padding: 2px 6px; border-radius: 3px; word-break: break-all; }
.pagination-wrap { margin-top: 16px; display: flex; justify-content: flex-end; }
</style>
