<template>
  <el-container style="height: 100vh">
    <!-- 侧边栏 -->
    <el-aside :width="isCollapse ? '64px' : '220px'" style="background: #304156; transition: width 0.3s">
      <div class="logo" :style="{ width: isCollapse ? '64px' : '220px' }">
        <span v-if="!isCollapse" class="logo-text">🔗 链环系统</span>
        <span v-else class="logo-text-sm">🔗</span>
      </div>
      <el-menu
        :default-active="activeMenu"
        :collapse="isCollapse"
        :collapse-transition="false"
        background-color="#304156"
        text-color="#bfcbd9"
        active-text-color="#409eff"
        router
      >
        <el-menu-item index="/dashboard">
          <el-icon><Odometer /></el-icon>
          <template #title>仪表盘</template>
        </el-menu-item>

        <el-sub-menu index="user">
          <template #title>
            <el-icon><User /></el-icon>
            <span>用户与权限</span>
          </template>
          <el-menu-item index="/user">用户管理</el-menu-item>
          <el-menu-item index="/user/role">角色管理</el-menu-item>
          <el-menu-item index="/enterprise">企业管理</el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="trace">
          <template #title>
            <el-icon><Search /></el-icon>
            <span>产品溯源</span>
          </template>
          <el-menu-item index="/product">产品管理</el-menu-item>
          <el-menu-item index="/batch">批次管理</el-menu-item>
          <el-menu-item index="/trace-record">溯源记录</el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="supply">
          <template #title>
            <el-icon><Truck /></el-icon>
            <span>供应链管理</span>
          </template>
          <el-menu-item index="/supplier">供应商管理</el-menu-item>
          <el-menu-item index="/order">订单管理</el-menu-item>
          <el-menu-item index="/warehouse">仓储管理</el-menu-item>
          <el-menu-item index="/logistics">物流管理</el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="blockchain">
          <template #title>
            <el-icon><Link /></el-icon>
            <span>区块链</span>
          </template>
          <el-menu-item index="/blockchain">区块浏览器</el-menu-item>
          <el-menu-item index="/blockchain/transaction">交易记录</el-menu-item>
          <el-menu-item index="/blockchain/contract">合约管理</el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="warehouse">
          <template #title>
            <el-icon><HomeFilled /></el-icon>
            <span>仓库管理</span>
          </template>
          <el-menu-item index="/warehouse/list">仓库信息</el-menu-item>
          <el-menu-item index="/warehouse-area">区域货架管理</el-menu-item>
          <el-menu-item index="/stocktake">仓库盘存</el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="inventory">
          <template #title>
            <el-icon><Box /></el-icon>
            <span>库存管理</span>
          </template>
          <el-menu-item index="/car-inventory">整车库存</el-menu-item>
          <el-menu-item index="/raw-material">原材料库存</el-menu-item>
          <el-menu-item index="/waste">危固废管理</el-menu-item>
          <el-menu-item index="/part-traceable">配件-溯源件</el-menu-item>
          <el-menu-item index="/part-untraceable">配件-非溯源件</el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="records">
          <template #title>
            <el-icon><Document /></el-icon>
            <span>操作记录</span>
          </template>
          <el-menu-item index="/inbound">入库记录</el-menu-item>
          <el-menu-item index="/outbound">出库记录</el-menu-item>
          <el-menu-item index="/transfer">调拨记录</el-menu-item>
        </el-sub-menu>

        <el-sub-menu index="system">
          <template #title>
            <el-icon><Setting /></el-icon>
            <span>系统管理</span>
          </template>
          <el-menu-item index="/config">系统配置</el-menu-item>
          <el-menu-item index="/logs">操作日志</el-menu-item>
        </el-sub-menu>
      </el-menu>
    </el-aside>

    <!-- 主区域 -->
    <el-container>
      <!-- 顶部导航 -->
      <el-header style="height: 60px; background: #fff; border-bottom: 1px solid #e6e6e6; display: flex; align-items: center; justify-content: space-between; padding: 0 20px">
        <div style="display: flex; align-items: center">
          <el-button @click="isCollapse = !isCollapse" text>
            <el-icon><Fold v-if="!isCollapse" /><Expand v-else /></el-icon>
          </el-button>
          <el-breadcrumb separator="/" style="margin-left: 16px">
            <el-breadcrumb-item :to="{ path: '/dashboard' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item v-if="currentTitle">{{ currentTitle }}</el-breadcrumb-item>
          </el-breadcrumb>
        </div>
        <div style="display: flex; align-items: center; gap: 16px">
          <el-badge :value="unreadCount" :hidden="unreadCount === 0" class="item">
            <el-button @click="showNotifications" :icon="Bell" circle text />
          </el-badge>
          <el-dropdown @command="handleCommand">
            <span style="cursor: pointer; display: flex; align-items: center; gap: 6px">
              <el-avatar :size="32" icon="UserFilled" />
              <span>{{ userName }}</span>
              <el-icon><ArrowDown /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="profile">个人信息</el-dropdown-item>
                <el-dropdown-item command="logout" divided>退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>

      <!-- 内容区域 -->
      <el-main style="background: #f0f2f5; padding: 20px">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { systemApi } from '@/api'
import { Bell, ArrowDown, Fold, Expand } from '@element-plus/icons-vue'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

const isCollapse = ref(false)
const unreadCount = ref(0)
const userName = ref(userStore.userName || 'Admin')

const activeMenu = computed(() => route.path)
const currentTitle = computed(() => route.meta?.title)

onMounted(async () => {
  if (userStore.token) {
    await userStore.getUserInfo()
    userName.value = userStore.userName
  }
})

async function showNotifications() {
  try {
    const res = await systemApi.getNotificationList()
    // 简化：弹出通知列表
    console.log('Notifications:', res.data)
  } catch (e) {}
}

function handleCommand(command) {
  if (command === 'logout') {
    userStore.logout()
    router.push('/login')
  }
}
</script>

<style scoped>
.logo {
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #2b3a4a;
  border-bottom: 1px solid #3a4a5a;
}
.logo-text {
  color: #fff;
  font-size: 18px;
  font-weight: bold;
  white-space: nowrap;
}
.logo-text-sm {
  color: #fff;
  font-size: 24px;
}
.el-menu {
  border-right: none;
}
</style>
