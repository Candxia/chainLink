<template>
  <div class="page-container">
    <div class="page-header"><h2>企业管理</h2><el-button type="primary" @click="showForm=true">新增企业</el-button></div>
    <el-table :data="list" border stripe v-loading="loading">
      <el-table-column prop="id" label="ID" width="60"/>
      <el-table-column prop="name" label="企业名称"/>
      <el-table-column prop="code" label="编码"/>
      <el-table-column prop="contact" label="联系人"/>
      <el-table-column prop="phone" label="电话"/>
      <el-table-column prop="status" label="状态" width="80">
        <template #default="{r}"><el-tag :type="r.status===1?'success':'danger'" size="small">{{r.status===1?'启用':'禁用'}}</el-tag></template>
      </el-table-column>
      <el-table-column label="操作" width="120">
        <template #default="{r}"><el-button size="small" @click="editRow(r)">编辑</el-button></template>
      </el-table-column>
    </el-table>
    <el-dialog v-model="showForm" :title="editId?'编辑企业':'新增企业'" width="500px">
      <el-form :model="form" label-width="80px">
        <el-form-item label="名称"><el-input v-model="form.name"/></el-form-item>
        <el-form-item label="编码"><el-input v-model="form.code"/></el-form-item>
        <el-form-item label="联系人"><el-input v-model="form.contact"/></el-form-item>
        <el-form-item label="电话"><el-input v-model="form.phone"/></el-form-item>
      </el-form>
      <template #footer><el-button @click="showForm=false">取消</el-button><el-button type="primary" @click="saveForm">保存</el-button></template>
    </el-dialog>
  </div>
</template>
<script setup>
import { ref, reactive, onMounted } from 'vue'
import { userApi } from '@/api'
import { ElMessage } from 'element-plus'
const loading=ref(false),list=ref([]),showForm=ref(false),editId=ref(null)
const form=reactive({name:'',code:'',contact:'',phone:''})
onMounted(()=>loadData())
async function loadData(){loading.value=true;try{const r=await userApi.getEnterpriseList({page:1,pageSize:100});list.value=r.data.list||[]}catch(e){list.value=[]}finally{loading.value=false}}
function editRow(row){editId.value=row.id;Object.assign(form,{name:row.name,code:row.code,contact:row.contact,phone:row.phone});showForm.value=true}
async function saveForm(){try{if(editId.value){await userApi.updateEnterprise({id:editId.value,...form})}else{await userApi.createEnterprise(form)}ElMessage.success('保存成功');showForm.value=false;loadData()}catch(e){}}
</script>
