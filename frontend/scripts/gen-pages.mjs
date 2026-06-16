import { writeFileSync, mkdirSync } from 'fs'
const v = 'D:/CodeByAi/cl_system/frontend/src/views'

function page(path, content) {
  const full = v + '/' + path
  mkdirSync(full.substring(0, full.lastIndexOf('/')), { recursive: true })
  writeFileSync(full, content.trimStart() + '\n', 'utf-8')
}

// Login
page('Login/Login.vue', `
<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 to-indigo-100">
    <el-card class="w-96" shadow="xl">
      <template #header><div class="text-center"><h2 class="text-xl font-bold">ChainLink System</h2></div></template>
      <el-form ref="formRef" :model="form" :rules="rules" label-position="top">
        <el-form-item label="Username" prop="username"><el-input v-model="form.username" placeholder="Username" /></el-form-item>
        <el-form-item label="Password" prop="password"><el-input v-model="form.password" type="password" show-password placeholder="Password" /></el-form-item>
        <el-form-item><el-button type="primary" class="w-full" :loading="loading" @click="handleLogin">Sign In</el-button></el-form-item>
      </el-form>
    </el-card>
  </div>
</template>
<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { loginApi } from '@/api/login'
import { useUserStore } from '@/stores/user'
const router = useRouter()
const userStore = useUserStore()
const formRef = ref(null)
const loading = ref(false)
const form = reactive({ username: '', password: '' })
const rules = { username: [{ required: true, message: 'Required', trigger: 'blur' }], password: [{ required: true, message: 'Required', trigger: 'blur' }] }
async function handleLogin() {
  if (!await formRef.value.validate().catch(() => false)) return
  loading.value = true
  try {
    const res = await loginApi(form)
    if (res.code === 0) {
      userStore.setToken(res.data.token || res.data.access_token)
      userStore.setUserInfo(res.data.userInfo || res.data.user)
      ElMessage.success('Login successful')
      router.push('/')
    } else ElMessage.error(res.msg || 'Login failed')
  } catch (e) { ElMessage.error('Network error') } finally { loading.value = false }
}
</script>
`)

// System pages
const sysPages = [
  { name: 'SysAdmin', title: 'Account', list: 'getAdminListApi', add: 'addAdminApi', edit: 'editAdminApi', del: 'delAdminApi',
    extra: (row) => `          <el-button size="small" link @click="handlePassword(row)">Password</el-button>`,
    extraScript: 'const handlePassword = async (row) => { try { const r = await adminPasswordApi({id:row.id}); if(r.code===0) {ElMessage.success(\'Password reset!\')} } catch(e){} }',
    extraImports: 'adminPasswordApi, adminStatusApi,',
    fields: [{p:'username',l:'Username'},{p:'nickname',l:'Nickname'},{p:'role_name',l:'Role'}],
    formFields: [{p:'username',l:'Username'},{p:'nickname',l:'Nickname'},{p:'role_id',l:'Role ID'}],
    search: [{p:'username',l:'Username'},{p:'status',l:'Status'}],
  },
  { name: 'SysApi', title: 'API', list: 'getApiListApi', add: 'addApiApi', edit: 'editApiApi', del: 'delApiApi',
    fields: [{p:'title',l:'Title'},{p:'path',l:'Path'},{p:'method',l:'Method'}],
    formFields: [{p:'title',l:'Title'},{p:'path',l:'Path'},{p:'method',l:'Method'},{p:'type',l:'Type'}],
    search: [{p:'title',l:'Title'},{p:'path',l:'Path'}],
  },
  { name: 'SysRole', title: 'Role', list: 'getRoleListApi', add: 'addRoleApi', edit: 'editRoleApi', del: 'delRoleApi',
    fields: [{p:'name',l:'Name'},{p:'code',l:'Code'},{p:'status',l:'Status',t:true}],
    formFields: [{p:'name',l:'Name'},{p:'code',l:'Code'}],
    search: [{p:'name',l:'Name'}],
  },
  { name: 'SysDept', title: 'Dept', list: 'getDeptListApi', add: 'addDeptApi', edit: 'editDeptApi', del: 'delDeptApi',
    fields: [{p:'name',l:'Name'},{p:'leader',l:'Leader'},{p:'phone',l:'Phone'},{p:'status',l:'Status',t:true}],
    formFields: [{p:'name',l:'Name'},{p:'leader',l:'Leader'},{p:'phone',l:'Phone'},{p:'status',l:'Status'}],
    search: [{p:'name',l:'Name'}],
  },
  { name: 'SysMenu', title: 'Menu', list: 'getMenuListApi', add: 'addMenuApi', edit: 'editMenuApi', del: 'delMenuApi',
    fields: [{p:'menu_name',l:'Name'},{p:'title',l:'Title'},{p:'path',l:'Path'},{p:'menu_type',l:'Type'},{p:'status',l:'Status',t:true}],
    formFields: [{p:'menu_name',l:'Name'},{p:'title',l:'Title'},{p:'path',l:'Path'},{p:'menu_type',l:'Type'},{p:'parent_id',l:'Parent ID'},{p:'icon',l:'Icon'},{p:'order_num',l:'Order'}],
    search: [{p:'title',l:'Title'}],
  },
]

sysPages.forEach(p => {
  const importNames = [p.list, p.add, p.edit, p.del, p.extraImports || ''].join(', ').replace(/, ,/g, ',').replace(/, $/, '')
  const searchBar = p.search.map(f => `<el-form-item label="${f.l}"><el-input v-model="search.${f.p}" clearable size="small" /></el-form-item>`).join('\n          ')
  const columnDefs = p.fields.map(f => f.t
    ? `\n      <el-table-column prop="${f.p}" label="${f.l}" width="80"><template #default="{row}"><el-tag :type="row.${f.p}===1?'success':'danger'" size="small">{{row.${f.p}===1?'Active':'Inactive'}}</el-tag></template></el-table-column>`
    : `\n      <el-table-column prop="${f.p}" label="${f.l}" />`).join('')
  const formDefs = p.formFields.map(f => `\n        <el-form-item label="${f.l}"><el-input v-model="form.${f.p}" /></el-form-item>`).join('')

  page(`System/${p.name === 'SysAdmin' ? 'admin' : p.name === 'SysApi' ? 'api' : p.name === 'SysRole' ? 'role' : p.name === 'SysDept' ? 'dept' : 'menu'}/${p.name}.vue`, `
<template>
  <div class="p-5">
    <el-form :inline="true" :model="search" class="mb-4">
      ${searchBar}
      <el-form-item><el-button type="primary" @click="handleSearch">Search</el-button><el-button @click="handleReset">Reset</el-button></el-form-item>
    </el-form>
    <div class="flex justify-between mb-4"><h2 class="text-lg font-bold">${p.title} Management</h2><el-button type="primary" @click="handleAdd">Add</el-button></div>
    <el-table :data="list" border stripe v-loading="loading">${columnDefs}
      <el-table-column label="Actions" width="200">
        <template #default="{row}">
          <el-button size="small" link @click="handleEdit(row)">Edit</el-button>
          <el-button size="small" link type="danger" @click="handleDelete(row.id)">Delete</el-button>
          ${p.extra ? p.extra('row') : ''}
        </template>
      </el-table-column>
    </el-table>
    <el-pagination v-model:current-page="page" v-model:page-size="pageSize" :total="total" layout="total,prev,pager,next" class="mt-4" @current-change="getList" />
    <el-dialog v-model="dialog" :title="editId ? 'Edit' : 'Add'" width="500px">
      <el-form :model="form" label-width="80px">${formDefs}
      </el-form>
      <template #footer><el-button @click="dialog=false">Cancel</el-button><el-button type="primary" @click="handleSave" :loading="saving">Save</el-button></template>
    </el-dialog>
  </div>
</template>
<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ${importNames} } from '@/api/system'

const list = ref([]), loading = ref(false), total = ref(0), page = ref(1), pageSize = ref(20)
const dialog = ref(false), editId = ref(null), saving = ref(false)
const search = reactive({})
const form = reactive({})

const getList = async () => {
  loading.value = true
  try { const r = await ${p.list}({ page: page.value, pageSize: pageSize.value, ...search }); list.value = r.data?.list || []; total.value = r.data?.total || 0 }
  catch(e){} finally { loading.value = false }
}
onMounted(() => getList())
const handleSearch = () => { page.value = 1; getList() }
const handleReset = () => { Object.keys(search).forEach(k => search[k] = ''); page.value = 1; getList() }
const handleAdd = () => { editId.value = null; Object.keys(form).forEach(k => delete form[k]); dialog.value = true }
const handleEdit = (row) => { editId.value = row.id; Object.assign(form, row); dialog.value = true }
const handleSave = async () => {
  saving.value = true
  try { const fn = editId.value ? ${p.edit} : ${p.add}; const r = await fn(editId.value ? { id: editId.value, ...form } : form); if (r.code === 0) { ElMessage.success('Saved!'); dialog.value = false; getList() } }
  catch(e){} finally { saving.value = false }
}
const handleDelete = async (id) => {
  try { await ElMessageBox.confirm('Delete?'); const r = await ${p.del}({ id }); if (r.code === 0) { ElMessage.success('Deleted!'); getList() } }
  catch(e){}
}
${p.extraScript || ''}
</script>
`)
})

console.log('All 5+ system pages written')
