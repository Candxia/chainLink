<template>
  <div class="page-container">
    <div class="page-header"><h2>订单管理</h2><el-button type="primary" @click="showForm=true">新增订单</el-button></div>
    <el-table :data="list" border stripe v-loading="loading">
      <el-table-column prop="orderNo" label="订单号" width="160"/>
      <el-table-column prop="productId" label="产品ID"/>
      <el-table-column prop="supplierId" label="供应商ID"/>
      <el-table-column prop="quantity" label="数量"/>
      <el-table-column prop="totalAmount" label="金额(元)"/>
      <el-table-column prop="status" label="状态" width="100">
        <template #default="{row}"><el-tag :type="statusMap[row.status]||'info'" size="small">{{row.status}}</el-tag></template>
      </el-table-column>
      <el-table-column label="操作" width="160">
        <template #default="{row}">
          <el-button size="small" @click="editRow(row)">编辑</el-button>
          <el-button size="small" type="danger" @click="del(row.id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination v-model:current-page="page" :page-size="10" :total="total" layout="total,prev,pager,next" style="margin-top:16px" @current-change="loadData"/>
    <el-dialog v-model="showForm" :title="editId?'编辑订单':'新增订单'" width="500px">
      <el-form :model="form" label-width="80px">
        <el-form-item label="订单号"><el-input v-model="form.orderNo"/></el-form-item>
        <el-form-item label="产品ID"><el-input-number v-model="form.productId" :min="1"/></el-form-item>
        <el-form-item label="供应商ID"><el-input-number v-model="form.supplierId" :min="1"/></el-form-item>
        <el-form-item label="数量"><el-input-number v-model="form.quantity" :min="1"/></el-form-item>
        <el-form-item label="金额"><el-input-number v-model="form.totalAmount" :min="0" :precision="2"/></el-form-item>
      </el-form>
      <template #footer><el-button @click="showForm=false">取消</el-button><el-button type="primary" @click="saveForm">保存</el-button></template>
    </el-dialog>
  </div>
</template>
<script setup>
import { ref, reactive, onMounted } from 'vue'
import { supplyApi } from '@/api'
import { ElMessage, ElMessageBox } from 'element-plus'
const statusMap={pending:'warning',confirmed:'primary',shipped:'info',delivered:'success',cancelled:'danger'}
const loading=ref(false),list=ref([]),page=ref(1),total=ref(0),showForm=ref(false),editId=ref(null)
const form=reactive({orderNo:'',productId:1,supplierId:1,quantity:100,totalAmount:0})
onMounted(()=>loadData())
async function loadData(){loading.value=true;try{const r=await supplyApi.getOrderList({page:page.value,pageSize:10});list.value=r.data.list;total.value=r.data.total}catch(e){list.value=[]}finally{loading.value=false}}
function editRow(row){editId.value=row.id;Object.assign(form,{orderNo:row.orderNo,productId:row.productId,supplierId:row.supplierId,quantity:row.quantity,totalAmount:row.totalAmount});showForm.value=true}
async function saveForm(){try{if(editId.value){await supplyApi.updateOrder({id:editId.value,...form})}else{await supplyApi.createOrder(form)}ElMessage.success('保存成功');showForm.value=false;loadData()}catch(e){}}
async function del(id){try{await ElMessageBox.confirm('确定删除？');await supplyApi.deleteOrder(id);ElMessage.success('已删除');loadData()}catch(e){}}
</script>
