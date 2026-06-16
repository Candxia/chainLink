<template>
  <div class="p-5">
    <el-row :gutter="20">
      <el-col :span="6" v-for="stat in stats" :key="stat.label">
        <el-card shadow="hover" class="mb-4">
          <div class="flex items-center">
            <el-icon :size="32" :color="stat.color"><component :is="stat.icon" /></el-icon>
            <div class="ml-4">
              <div class="text-2xl font-bold">{{ stat.value }}</div>
              <div class="text-gray-400 text-sm">{{ stat.label }}</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
    <el-card>
      <template #header><span>Quick Actions</span></template>
      <el-space>
        <el-button @click="$router.push('/system/admin')">Accounts</el-button>
        <el-button @click="$router.push('/system/role')">Roles</el-button>
        <el-button @click="$router.push('/system/menu')">Menus</el-button>
      </el-space>
    </el-card>
  </div>
</template>

<script setup>
import { User, Setting, List, Grid } from '@element-plus/icons-vue'
const stats = [
  { label: 'Users', value: 0, icon: User, color: '#409eff' },
  { label: 'Roles', value: 0, icon: Setting, color: '#67c23a' },
  { label: 'Pages', value: 0, icon: List, color: '#e6a23c' },
  { label: 'Modules', value: 0, icon: Grid, color: '#f56c6c' }
]
</script>
