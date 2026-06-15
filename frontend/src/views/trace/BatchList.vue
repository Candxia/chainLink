<template>
  <div class="page-container">
    <div class="page-header"><h2>批次管理</h2><el-button type="primary" @click="showForm=true">新增批次</el-button></div>
    <el-table :data="list" border stripe v-loading="loading">
      <el-table-column prop="batchNo" label="批次号" width="160"/>
      <el-table-column prop="productId" label="产品ID"/>
      <el-table-column prop="quantity" label="数量"/>
      <el-table-column prop="produceDate" label="生产日期"/>
      <el-table-column prop="expireDate" label="到期日期"/>
      <el-table-column prop="status" label="状态" width="80"><template #default="{row}"><el-tag :type="row.status==='active'?'success':'info'" size="small">{{row.status}}</el-tag></template></el-table-column>
    </el-table>
    <el-pagination v-model:current-page="page" :page-size="pageSize" :total="total" layout="total,prev,pager,next" style="margin-top:16px" @current-change="loadData"/>
    <el-dialog v-model="showForm" title="新增批次" width="500px">
      <el-form :model="form" label-width="80px">
        <el-form-item label="批次号"><el-input v-model="form.batchNo"/></el-form-item>
        <el-form-item label="产品ID"><el-input-number v-model="form.productId" :min="1"/></el-form-item>
        <el-form-item label="数量"><el-input-number v-model="form.quantity" :min="1"/></el-form-item>
        <el-form-item label="生产日期"><el-date-picker v-model="form.produceDate" type="date" value-format="YYYY-MM-DD"/></el-form-item>
        <el-form-item label="到期日期"><el-date-picker v-model="form.expireDate" type="date" value-format="YYYY-MM-DD"/></el-form-item>
      </el-form>
      <template #footer><el-button @click="showForm=false">取消</el-button><el-button type="primary" @click="saveForm">保存</el-button></template>
    </el-dialog>
  </div>
</template>
<script setup>
import { ref, reactive, onMounted } from 'vue'
import { traceApi } from '@/api'
import { ElMessage } from 'element-plus'
const loading=ref(false),list=ref([]),page=ref(1),pageSize=ref(10),total=ref(0),showForm=ref(false)
const form=reactive({batchNo:'',productId:1,quantity:100,produceDate:'',expireDate:''})
onMounted(()=>loadData())
async function loadData(){loading.value=true;try{const r=await traceApi.getBatchList({page:page.value,pageSize:pageSize.value});list.value=r.data.list;total.value=r.data.total}catch(e){list.value=[]}finally{loading.value=false}}
async function saveForm(){try{await traceApi.createBatch(form);ElMessage.success('创建成功');showForm.value=false;loadData()}catch(e){}}
</script>
