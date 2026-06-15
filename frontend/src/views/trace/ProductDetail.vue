<template>
  <div class="page-container">
    <div class="page-header"><h2>产品溯源详情</h2></div>
    <el-skeleton :loading="loading" animated>
      <el-descriptions title="产品信息" :column="2" border>
        <el-descriptions-item label="产品名称">{{ product.name }}</el-descriptions-item>
        <el-descriptions-item label="分类">{{ product.category }}</el-descriptions-item>
        <el-descriptions-item label="规格">{{ product.spec }}</el-descriptions-item>
        <el-descriptions-item label="状态"><el-tag size="small" :type="product.status==='active'?'success':'info'">{{product.status}}</el-tag></el-descriptions-item>
      </el-descriptions>
    </el-skeleton>

    <el-divider>溯源链</el-divider>
    <el-steps direction="vertical" :active="traceRecords.length">
      <el-step v-for="(r,i) in traceRecords" :key="i" :title="r.recordType" :description="r.content + ' - ' + r.operator + ' @ ' + r.location">
        <template #status><el-tag size="small" type="success">✓ 已上链</el-tag></template>
      </el-step>
    </el-steps>

    <el-divider>新增溯源记录</el-divider>
    <el-form :model="form" label-width="80px">
      <el-form-item label="记录类型"><el-select v-model="form.recordType"><el-option label="生产" value="生产"/><el-option label="加工" value="加工"/><el-option label="质检" value="质检"/><el-option label="物流" value="物流"/></el-select></el-form-item>
      <el-form-item label="内容"><el-input v-model="form.content" type="textarea" :rows="3"/></el-form-item>
      <el-form-item label="操作人"><el-input v-model="form.operator"/></el-form-item>
      <el-form-item label="地点"><el-input v-model="form.location"/></el-form-item>
      <el-form-item><el-button type="primary" @click="addRecord" :loading="adding">提交上链</el-button></el-form-item>
    </el-form>
  </div>
</template>
<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { traceApi } from '@/api'
import { ElMessage } from 'element-plus'
const route=useRoute(),loading=ref(false),adding=ref(false)
const product=ref({}),traceRecords=ref([])
const form=reactive({productId:'',batchId:0,recordType:'生产',content:'',operator:'',location:''})
onMounted(async()=>{const id=route.params.id;form.productId=id;loading.value=true;try{const r=await traceApi.getProduct(id);product.value=r.data;const r2=await traceApi.getTraceChain(id);traceRecords.value=r2.data||[]}catch(e){}finally{loading.value=false}})
async function addRecord(){adding.value=true;try{const r=await traceApi.createTraceRecord(form);ElMessage.success('溯源记录已上链');traceRecords.value.push(r.data);form.content=''}catch(e){}finally{adding.value=false}}
</script>
