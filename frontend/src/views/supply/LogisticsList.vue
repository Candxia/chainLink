<template>
  <div class="page-container">
    <div class="page-header"><h2>物流管理</h2><el-button type="primary" @click="showForm=true">新增物流节点</el-button></div>
    <el-table :data="records" border stripe>
      <el-table-column prop="orderId" label="订单ID"/>
      <el-table-column prop="nodeName" label="节点名称"/>
      <el-table-column prop="location" label="位置"/>
      <el-table-column prop="status" label="状态" width="100">
        <template #default="{r}"><el-tag :type="r.status==='completed'?'success':'primary'" size="small">{{r.status}}</el-tag></template>
      </el-table-column>
      <el-table-column prop="operator" label="操作人"/>
      <el-table-column prop="createdAt" label="时间" width="160"/>
    </el-table>
    <el-dialog v-model="showForm" title="新增物流节点" width="400px">
      <el-form :model="form" label-width="80px">
        <el-form-item label="订单ID"><el-input-number v-model="form.orderId" :min="1"/></el-form-item>
        <el-form-item label="节点名称"><el-input v-model="form.nodeName"/></el-form-item>
        <el-form-item label="位置"><el-input v-model="form.location"/></el-form-item>
        <el-form-item label="操作人"><el-input v-model="form.operator"/></el-form-item>
      </el-form>
      <template #footer><el-button @click="showForm=false">取消</el-button><el-button type="primary" @click="save">保存</el-button></template>
    </el-dialog>
  </div>
</template>
<script setup>
import { ref, reactive, onMounted } from 'vue'
import { supplyApi } from '@/api'
import { ElMessage } from 'element-plus'
const records=ref([]),showForm=ref(false)
const form=reactive({orderId:1,nodeName:'',location:'',operator:'',status:'created'})
onMounted(async()=>{try{const r=await supplyApi.getLogisticsList(1);records.value=r.data||[]}catch(e){}})
async function save(){try{await supplyApi.createLogistics(form);ElMessage.success('创建成功');showForm.value=false;const r=await supplyApi.getLogisticsList(form.orderId);records.value=r.data||[]}catch(e){}}
</script>
