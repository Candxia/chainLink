<template>
  <div class="page-container">
    <div class="page-header"><h2>操作日志</h2></div>
    <el-table :data="logs" border stripe>
      <el-table-column prop="id" label="ID" width="60"/>
      <el-table-column prop="username" label="用户名"/>
      <el-table-column prop="type" label="类型"/>
      <el-table-column prop="action" label="操作内容" min-width="200"/>
      <el-table-column prop="ip" label="IP地址" width="140"/>
      <el-table-column prop="createdAt" label="时间" width="160"/>
      <el-table-column label="操作" width="80">
        <template #default="{row}"><el-button size="small" type="danger" @click="del(row.id)">删除</el-button></template>
      </el-table-column>
    </el-table>
    <el-pagination v-model:current-page="page" :page-size="20" :total="total" layout="total,prev,pager,next" style="margin-top:16px" @current-change="loadData"/>
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import { systemApi } from '@/api'
import { ElMessage } from 'element-plus'
const logs=ref([]),page=ref(1),total=ref(0)
onMounted(()=>loadData())
async function loadData(){try{const r=await systemApi.getLogList({page:page.value,pageSize:20});logs.value=r.data.list;total.value=r.data.total}catch(e){logs.value=[]}}
async function del(id){try{await systemApi.deleteLog(id);ElMessage.success('已删除');loadData()}catch(e){}}
</script>
