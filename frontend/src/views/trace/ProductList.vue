<template>
  <div class="page-container">
    <div class="page-header"><h2>产品管理</h2><el-button type="primary" @click="showForm=true">新增产品</el-button></div>
    <el-form :inline="true" class="search-bar">
      <el-form-item><el-input v-model="searchName" placeholder="产品名称" clearable @input="loadData" /></el-form-item>
      <el-form-item><el-button @click="loadData">查询</el-button></el-form-item>
    </el-form>
    <el-table :data="list" border stripe v-loading="loading">
      <el-table-column prop="id" label="ID" width="60"/>
      <el-table-column prop="name" label="产品名称"/>
      <el-table-column prop="category" label="分类"/>
      <el-table-column prop="spec" label="规格"/>
      <el-table-column prop="unit" label="单位" width="60"/>
      <el-table-column prop="status" label="状态" width="80">
        <template #default="{row}"><el-tag :type="row.status==='active'?'success':'info'" size="small">{{row.status==='active'?'启用':'停用'}}</el-tag></template>
      </el-table-column>
      <el-table-column label="操作" width="220">
        <template #default="{row}">
          <el-button size="small" @click="$router.push(`/product/detail/${row.id}`)">详情</el-button>
          <el-button size="small" @click="editRow(row)">编辑</el-button>
          <el-button size="small" type="danger" @click="delRow(row.id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination v-model:current-page="page" :page-size="pageSize" :total="total" layout="total,prev,pager,next" style="margin-top:16px;justify-content:flex-end" @current-change="loadData"/>

    <el-dialog v-model="showForm" :title="editId?'编辑产品':'新增产品'" width="600px">
      <el-form :model="form" label-width="80px">
        <el-form-item label="名称"><el-input v-model="form.name"/></el-form-item>
        <el-form-item label="分类"><el-input v-model="form.category"/></el-form-item>
        <el-form-item label="规格"><el-input v-model="form.spec"/></el-form-item>
        <el-form-item label="单位"><el-select v-model="form.unit"><el-option label="个" value="个"/><el-option label="箱" value="箱"/><el-option label="吨" value="吨"/><el-option label="千克" value="千克"/></el-select></el-form-item>
        <el-form-item label="描述"><el-input v-model="form.description" type="textarea" :rows="3"/></el-form-item>
      </el-form>
      <template #footer><el-button @click="showForm=false">取消</el-button><el-button type="primary" @click="saveForm" :loading="saving">保存</el-button></template>
    </el-dialog>
  </div>
</template>
<script setup>
import { ref, reactive, onMounted } from 'vue'
import { traceApi } from '@/api'
import { ElMessage, ElMessageBox } from 'element-plus'
const loading=ref(false),saving=ref(false),list=ref([]),page=ref(1),pageSize=ref(10),total=ref(0),showForm=ref(false),editId=ref(null),searchName=ref('')
const form=reactive({name:'',category:'',spec:'',unit:'个',description:''})
onMounted(()=>loadData())
async function loadData(){loading.value=true;try{const r=await traceApi.getProductList({page:page.value,pageSize:pageSize.value,name:searchName.value});list.value=r.data.list;total.value=r.data.total}catch(e){list.value=[]}finally{loading.value=false}}
function editRow(row){editId.value=row.id;Object.assign(form,{name:row.name,category:row.category,spec:row.spec,unit:row.unit,description:row.description});showForm.value=true}
async function delRow(id){try{await ElMessageBox.confirm('确定删除？');await traceApi.deleteProduct(id);ElMessage.success('已删除');loadData()}catch(e){}}
async function saveForm(){saving.value=true;try{if(editId.value){await traceApi.updateProduct({id:editId.value,...form})}else{await traceApi.createProduct(form)}ElMessage.success('保存成功');showForm.value=false;loadData()}catch(e){}finally{saving.value=false}}
</script>
