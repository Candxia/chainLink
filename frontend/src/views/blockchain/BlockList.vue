<template>
  <div class="page-container">
    <div class="page-header"><h2>区块浏览器</h2></div>
    <el-table :data="blocks" border stripe v-loading="loading">
      <el-table-column prop="number" label="区块高度" width="120"/>
      <el-table-column prop="hash" label="区块哈希" min-width="250"><template #default="{row}"><span style="font-family:monospace;font-size:12px">{{row.hash?.substring(0,30)}}...</span></template></el-table-column>
      <el-table-column prop="parentHash" label="父哈希" min-width="200"><template #default="{row}"><span style="font-family:monospace;font-size:12px">{{row.parentHash?.substring(0,20)}}...</span></template></el-table-column>
      <el-table-column prop="txCount" label="交易数" width="80"/>
      <el-table-column prop="timestamp" label="时间" width="120"/>
    </el-table>
    <el-pagination v-model:current-page="page" :page-size="10" :total="total" layout="total,prev,pager,next" style="margin-top:16px" @current-change="loadData"/>
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import { blockchainApi } from '@/api'
const loading=ref(false),blocks=ref([]),page=ref(1),total=ref(0)
onMounted(()=>loadData())
async function loadData(){loading.value=true;try{const r=await blockchainApi.getBlockList({page:page.value,pageSize:10});blocks.value=r.data.list;total.value=r.data.total}catch(e){blocks.value=[]}finally{loading.value=false}}
</script>
