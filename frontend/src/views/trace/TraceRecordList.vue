<template>
  <div class="page-container">
    <div class="page-header"><h2>溯源记录</h2></div>
    <el-table :data="records" border stripe v-loading="loading">
      <el-table-column prop="id" label="ID" width="60"/>
      <el-table-column prop="productId" label="产品ID"/>
      <el-table-column prop="recordType" label="类型"><template #default="{r}"><el-tag size="small">{{r.recordType}}</el-tag></template></el-table-column>
      <el-table-column prop="content" label="内容" min-width="200"/>
      <el-table-column prop="operator" label="操作人"/>
      <el-table-column prop="location" label="地点"/>
      <el-table-column prop="txHash" label="交易哈希" min-width="200"><template #default="{r}"><span style="font-size:12px;font-family:monospace">{{r.txHash?.substring(0,20)}}...</span></template></el-table-column>
      <el-table-column prop="blockNumber" label="区块高度" width="100"/>
    </el-table>
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import { traceApi } from '@/api'
const loading=ref(false),records=ref([])
onMounted(async()=>{loading.value=true;try{const r=await traceApi.getProductList({page:1,pageSize:1});if(r.data.list?.length){const r2=await traceApi.getTraceRecordList(r.data.list[0].id);records.value=r2.data||[]}}catch(e){}finally{loading.value=false}})
</script>
