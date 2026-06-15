<template>
  <div class="page-container">
    <div class="page-header"><h2>供应商管理</h2><el-button type="primary" @click="showForm=true">新增供应商</el-button></div>
    <el-table :data="list" border stripe v-loading="loading">
      <el-table-column prop="name" label="名称"/><el-table-column prop="code" label="编码"/>
      <el-table-column prop="contact" label="联系人"/><el-table-column prop="phone" label="电话"/>
      <el-table-column prop="qualLevel" label="资质" width="80">
        <template #default="{row}"><el-tag :type="row.qualLevel==='A'?'success':row.qualLevel==='B'?'warning':'info'" size="small">{{row.qualLevel}}</el-tag></template>
      </el-table-column>
      <el-table-column prop="status" label="状态" width="80">
        <template #default="{row}"><el-tag :type="row.status==='active'?'success':'danger'" size="small">{{row.status==='active'?'启用':'停用'}}</el-tag></template>
      </el-table-column>
      <el-table-column label="操作" width="120"><template #default="{row}"><el-button size="small" type="danger" @click="del(row.id)">删除</el-button></template></el-table-column>
    </el-table>
    <el-pagination v-model:current-page="page" :page-size="10" :total="total" layout="total,prev,pager,next" style="margin-top:16px" @current-change="loadData"/>
    <el-dialog v-model="showForm" title="新增供应商" width="500px">
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
import { supplyApi } from '@/api'
import { ElMessage, ElMessageBox } from 'element-plus'
const loading=ref(false),list=ref([]),page=ref(1),total=ref(0),showForm=ref(false)
const form=reactive({name:'',code:'',contact:'',phone:''})
onMounted(()=>loadData())
async function loadData(){loading.value=true;try{const r=await supplyApi.getSupplierList({page:page.value,pageSize:10});list.value=r.data.list;total.value=r.data.total}catch(e){list.value=[]}finally{loading.value=false}}
async function saveForm(){try{await supplyApi.createSupplier(form);ElMessage.success('创建成功');showForm.value=false;loadData()}catch(e){}}
async function del(id){try{await ElMessageBox.confirm('确定删除？');await supplyApi.deleteSupplier(id);ElMessage.success('已删除');loadData()}catch(e){}}
</script>
