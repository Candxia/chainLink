<template>
  <div class="page-container">
    <div class="page-header">
      <h2>用户管理</h2>
      <el-button type="primary" @click="showDialog = true">新增用户</el-button>
    </div>

    <el-table :data="userList" stripe border v-loading="loading">
      <el-table-column prop="id" label="ID" width="60" />
      <el-table-column prop="username" label="用户名" />
      <el-table-column prop="email" label="邮箱" />
      <el-table-column prop="phone" label="手机号" />
      <el-table-column prop="role" label="角色">
        <template #default="{ row }">
          <el-tag :type="row.role === 'super_admin' ? 'danger' : row.role === 'admin' ? 'warning' : 'info'" size="small">
            {{ roleMap[row.role] || row.role }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="status" label="状态" width="80">
        <template #default="{ row }">
          <el-tag :type="row.status === 1 ? 'success' : 'danger'" size="small">
            {{ row.status === 1 ? '启用' : '禁用' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180">
        <template #default="{ row }">
          <el-button size="small" @click="handleEdit(row)">编辑</el-button>
          <el-button size="small" type="danger" @click="handleDelete(row.id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      v-model:current-page="page"
      :page-size="pageSize"
      :total="total"
      layout="total, prev, pager, next"
      style="margin-top: 20px; justify-content: flex-end"
      @current-change="loadData"
    />

    <!-- 编辑弹窗 -->
    <el-dialog v-model="showDialog" :title="editId ? '编辑用户' : '新增用户'" width="500px">
      <el-form ref="formRef" :model="form" label-width="80px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="form.username" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="form.email" />
        </el-form-item>
        <el-form-item label="手机号">
          <el-input v-model="form.phone" />
        </el-form-item>
        <el-form-item label="角色">
          <el-select v-model="form.role">
            <el-option label="普通用户" value="user" />
            <el-option label="企业管理员" value="admin" />
            <el-option label="质检员" value="inspector" />
          </el-select>
        </el-form-item>
        <el-form-item v-if="!editId" label="密码" prop="password">
          <el-input v-model="form.password" type="password" />
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
import { userApi } from '@/api'
import { ElMessage, ElMessageBox } from 'element-plus'

const loading = ref(false)
const saving = ref(false)
const userList = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const showDialog = ref(false)
const editId = ref(null)
const formRef = ref(null)

const roleMap = { super_admin: '超级管理员', admin: '企业管理员', inspector: '质检员', user: '普通用户' }

const form = reactive({
  username: '',
  email: '',
  phone: '',
  role: 'user',
  password: '',
})

onMounted(() => loadData())

async function loadData() {
  loading.value = true
  try {
    const res = await userApi.getUserList({ page: page.value, pageSize: pageSize.value })
    userList.value = res.data.list
    total.value = res.data.total
  } catch (e) {
    userList.value = []
  } finally {
    loading.value = false
  }
}

function handleEdit(row) {
  editId.value = row.id
  Object.assign(form, { username: row.username, email: row.email, phone: row.phone, role: row.role })
  showDialog.value = true
}

async function handleDelete(id) {
  try {
    await ElMessageBox.confirm('确定删除该用户？')
    await userApi.deleteUser(id)
    ElMessage.success('删除成功')
    loadData()
  } catch (e) {}
}

async function handleSave() {
  saving.value = true
  try {
    if (editId.value) {
      await userApi.updateUser({ id: editId.value, ...form })
    } else {
      await userApi.register(form)
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
