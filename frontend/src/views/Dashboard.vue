<template>
  <div class="page-container">
    <div class="page-header">
      <h2>仪表盘</h2>
    </div>

    <!-- 统计卡片 -->
    <el-row :gutter="20" style="margin-bottom: 20px">
      <el-col :span="6" v-for="item in stats" :key="item.label">
        <el-card shadow="hover">
          <div class="stat-card">
            <div class="stat-icon" :style="{ background: item.color }">
              <el-icon :size="24"><component :is="item.icon" /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ item.value }}</div>
              <div class="stat-label">{{ item.label }}</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 快捷操作 -->
    <el-row :gutter="20">
      <el-col :span="12">
        <el-card>
          <template #header>
            <span>快捷操作</span>
          </template>
          <div style="display: flex; flex-wrap: wrap; gap: 12px">
            <el-button type="primary" @click="$router.push('/product')">产品管理</el-button>
            <el-button type="success" @click="$router.push('/trace-record')">新增溯源</el-button>
            <el-button type="warning" @click="$router.push('/order')">订单管理</el-button>
            <el-button type="info" @click="$router.push('/blockchain')">区块浏览器</el-button>
          </div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card>
          <template #header>
            <span>系统信息</span>
          </template>
          <el-descriptions :column="2" border size="small">
            <el-descriptions-item label="系统名称">链环系统 v1.0.0</el-descriptions-item>
            <el-descriptions-item label="运行状态">
              <el-tag type="success">运行中</el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="区块链节点">{{ blockchainStatus }}</el-descriptions-item>
            <el-descriptions-item label="数据库状态">
              <el-tag type="success">已连接</el-tag>
            </el-descriptions-item>
          </el-descriptions>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { systemApi } from '@/api'
import { User, Goods, ShoppingCart, Connection } from '@element-plus/icons-vue'

const blockchainStatus = ref('模拟节点')

const stats = reactive([
  { label: '用户总数', value: 0, icon: User, color: '#409eff' },
  { label: '产品总数', value: 0, icon: Goods, color: '#67c23a' },
  { label: '订单总数', value: 0, icon: ShoppingCart, color: '#e6a23c' },
  { label: '区块高度', value: 0, icon: Connection, color: '#f56c6c' },
])

onMounted(async () => {
  try {
    const res = await systemApi.getDashboard()
    if (res.data) {
      stats[0].value = res.data.totalUsers
      stats[1].value = res.data.totalProducts
      stats[2].value = res.data.totalOrders
      stats[3].value = res.data.totalBlockCount
    }
  } catch (e) {
    // 设置默认值
    stats[0].value = '--'
    stats[1].value = '--'
    stats[2].value = '--'
    stats[3].value = '--'
  }
})
</script>

<style scoped>
.stat-card {
  display: flex;
  align-items: center;
  gap: 16px;
}
.stat-icon {
  width: 56px;
  height: 56px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
}
.stat-value {
  font-size: 28px;
  font-weight: bold;
  color: #303133;
}
.stat-label {
  font-size: 13px;
  color: #909399;
  margin-top: 4px;
}
</style>
