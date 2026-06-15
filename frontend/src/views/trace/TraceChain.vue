<template>
  <div class="page-container">
    <div class="page-header"><h2>追溯链</h2></div>
    <el-steps direction="vertical" :active="records.length" v-if="records.length">
      <el-step v-for="(r,i) in records" :key="i" :title="r.recordType" :description="r.content">
        <template #status>
          <el-tag size="small" type="success">✓ 已上链</el-tag>
        </template>
      </el-step>
    </el-steps>
    <el-empty v-else description="暂无溯源记录" />
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { traceApi } from '@/api'
const route = useRoute()
const records = ref([])
onMounted(async () => {
  try {
    const r = await traceApi.getTraceChain(route.params.productId)
    records.value = r.data || []
  } catch (e) { records.value = [] }
})
</script>
