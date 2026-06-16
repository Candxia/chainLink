<template>
  <div class="p-5">
    <el-row :gutter="20">
      <el-col :span="8">
        <el-card>
          <template #header><span>Profile</span></template>
          <div class="text-center">
            <el-avatar :size="80" :icon="UserFilled" />
            <h3 class="mt-4 text-lg">{{ userInfo.nickname || 'Admin' }}</h3>
            <p class="text-gray-400 mt-2">{{ userInfo.email || '' }}</p>
          </div>
        </el-card>
      </el-col>
      <el-col :span="16">
        <el-card>
          <template #header><span>Basic Info</span></template>
          <el-descriptions :column="1" border>
            <el-descriptions-item label="Username">{{ userInfo.username || '-' }}</el-descriptions-item>
            <el-descriptions-item label="Nickname">{{ userInfo.nickname || '-' }}</el-descriptions-item>
            <el-descriptions-item label="Email">{{ userInfo.email || '-' }}</el-descriptions-item>
            <el-descriptions-item label="Phone">{{ userInfo.phone || '-' }}</el-descriptions-item>
            <el-descriptions-item label="Created">{{ userInfo.createTime || '-' }}</el-descriptions-item>
          </el-descriptions>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { UserFilled } from '@element-plus/icons-vue'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const userInfo = computed(() => userStore.getUserInfo)
</script>
