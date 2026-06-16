<template>
  <div class="p-5">
    <el-form :inline="true" :model="search" class="mb-4">
      <el-form-item label="Title"><el-input v-model="search.title" clearable size="small" /></el-form-item>
          <el-form-item label="Path"><el-input v-model="search.path" clearable size="small" /></el-form-item>
      <el-form-item><el-button type="primary" @click="handleSearch">Search</el-button><el-button @click="handleReset">Reset</el-button></el-form-item>
    </el-form>
    <div class="flex justify-between mb-4"><h2 class="text-lg font-bold">API Management</h2><el-button type="primary" @click="handleAdd">Add</el-button></div>
    <el-table :data="list" border stripe v-loading="loading">
      <el-table-column prop="title" label="Title" />
      <el-table-column prop="path" label="Path" />
      <el-table-column prop="method" label="Method" />
      <el-table-column label="Actions" width="200">
        <template #default="{row}">
          <el-button size="small" link @click="handleEdit(row)">Edit</el-button>
          <el-button size="small" link type="danger" @click="handleDelete(row.id)">Delete</el-button>
          
        </template>
      </el-table-column>
    </el-table>
    <el-pagination v-model:current-page="page" v-model:page-size="pageSize" :total="total" layout="total,prev,pager,next" class="mt-4" @current-change="getList" />
    <el-dialog v-model="dialog" :title="editId ? 'Edit' : 'Add'" width="500px">
      <el-form :model="form" label-width="80px">
        <el-form-item label="Title"><el-input v-model="form.title" /></el-form-item>
        <el-form-item label="Path"><el-input v-model="form.path" /></el-form-item>
        <el-form-item label="Method"><el-input v-model="form.method" /></el-form-item>
        <el-form-item label="Type"><el-input v-model="form.type" /></el-form-item>
      </el-form>
      <template #footer><el-button @click="dialog=false">Cancel</el-button><el-button type="primary" @click="handleSave" :loading="saving">Save</el-button></template>
    </el-dialog>
  </div>
</template>
<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getApiListApi, addApiApi, editApiApi, delApiApi } from '@/api/system'

const list = ref([]), loading = ref(false), total = ref(0), page = ref(1), pageSize = ref(20)
const dialog = ref(false), editId = ref(null), saving = ref(false)
const search = reactive({})
const form = reactive({})

const getList = async () => {
  loading.value = true
  try { const r = await getApiListApi({ page: page.value, pageSize: pageSize.value, ...search }); list.value = r.data?.list || []; total.value = r.data?.total || 0 }
  catch(e){} finally { loading.value = false }
}
onMounted(() => getList())
const handleSearch = () => { page.value = 1; getList() }
const handleReset = () => { Object.keys(search).forEach(k => search[k] = ''); page.value = 1; getList() }
const handleAdd = () => { editId.value = null; Object.keys(form).forEach(k => delete form[k]); dialog.value = true }
const handleEdit = (row) => { editId.value = row.id; Object.assign(form, row); dialog.value = true }
const handleSave = async () => {
  saving.value = true
  try { const fn = editId.value ? editApiApi : addApiApi; const r = await fn(editId.value ? { id: editId.value, ...form } : form); if (r.code === 0) { ElMessage.success('Saved!'); dialog.value = false; getList() } }
  catch(e){} finally { saving.value = false }
}
const handleDelete = async (id) => {
  try { await ElMessageBox.confirm('Delete?'); const r = await delApiApi({ id }); if (r.code === 0) { ElMessage.success('Deleted!'); getList() } }
  catch(e){}
}

</script>

