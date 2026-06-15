<template>
  <div class="page-container">
    <div class="page-header"><h2>系统配置</h2></div>
    <el-table :data="configs" border stripe>
      <el-table-column prop="key" label="配置键" min-width="180"/>
      <el-table-column prop="value" label="配置值" min-width="250">
        <template #default="{row,$index}"><el-input v-model="configs[$index].value" size="small"/></template>
      </el-table-column>
      <el-table-column prop="desc" label="说明"/>
      <el-table-column label="操作" width="100">
        <template #default="{row}"><el-button size="small" type="primary" @click="updateConfig(row)">保存</el-button></template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import { systemApi } from '@/api'
import { ElMessage } from 'element-plus'
const configs=ref([])
onMounted(async()=>{try{const r=await systemApi.getConfigList();configs.value=r.data||[]}catch(e){configs.value=[]}})
async function updateConfig(row){try{await systemApi.updateConfig({id:row.id,value:row.value});ElMessage.success('已更新')}catch(e){}}
</script>
