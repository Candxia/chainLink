<template>
  <div class="page-container">
    <div class="page-header"><h2>角色管理</h2><el-button type="primary" @click="showForm=true">新增角色</el-button></div>
    <el-table :data="roles" border stripe>
      <el-table-column prop="id" label="ID" width="60"/>
      <el-table-column prop="name" label="角色名称"/>
      <el-table-column prop="code" label="角色编码"/>
      <el-table-column label="操作" width="160">
        <template #default="{row}">
          <el-button size="small" @click="editRow(row)">编辑</el-button>
          <el-button size="small" type="danger" @click="del(row.id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog v-model="showForm" :title="editId?'编辑角色':'新增角色'" width="400px">
      <el-form :model="form" label-width="80px">
        <el-form-item label="名称"><el-input v-model="form.name"/></el-form-item>
        <el-form-item label="编码"><el-input v-model="form.code"/></el-form-item>
      </el-form>
      <template #footer><el-button @click="showForm=false">取消</el-button><el-button type="primary" @click="saveForm">保存</el-button></template>
    </el-dialog>
  </div>
</template>
<script setup>
import { ref, reactive, onMounted } from 'vue'
import { userApi } from '@/api'
import { ElMessage, ElMessageBox } from 'element-plus'
const roles=ref([]),showForm=ref(false),editId=ref(null)
const form=reactive({name:'',code:''})
onMounted(async()=>{try{const r=await userApi.getRoleList();roles.value=r.data||[]}catch(e){}})
function editRow(row){editId.value=row.id;form.name=row.name;form.code=row.code;showForm.value=true}
async function saveForm(){try{if(editId.value){await userApi.updateRole({id:editId.value,...form})}else{await userApi.createRole(form)}ElMessage.success('保存成功');showForm.value=false;const r=await userApi.getRoleList();roles.value=r.data||[]}catch(e){}}
async function del(id){try{await ElMessageBox.confirm('确定删除？');await userApi.deleteRole(id);const r=await userApi.getRoleList();roles.value=r.data||[];ElMessage.success('已删除')}catch(e){}}
</script>
