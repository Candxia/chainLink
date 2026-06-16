<template>
  <el-container class="layout-container">
    <el-aside :width="collapse ? '64px' : '220px'" class="layout-aside">
      <div class="logo">
        <h3 v-show="!collapse" class="logo-title">ChainLink System</h3>
        <h3 v-show="collapse" class="logo-title-mini">CL</h3>
      </div>
      <el-menu
        :default-active="route.path"
        :collapse="collapse"
        :router="true"
        class="layout-menu"
        background-color="#001529"
        text-color="#bfcbd9"
        active-text-color="#fff"
      >
        <template v-for="item in menuList" :key="item.path">
          <el-menu-item v-if="!item.children?.length" :index="item.path">
            <el-icon><component :is="item.meta?.icon || 'Menu'" /></el-icon>
            <span>{{ item.meta?.title || item.name }}</span>
          </el-menu-item>
          <el-sub-menu v-else :index="item.path">
            <template #title>
              <el-icon><component :is="item.meta?.icon || 'Menu'" /></el-icon>
              <span>{{ item.meta?.title || item.name }}</span>
            </template>
            <el-menu-item v-for="child in item.children" :key="child.path" :index="item.path + '/' + child.path">
              <span>{{ child.meta?.title || child.name }}</span>
            </el-menu-item>
          </el-sub-menu>
        </template>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header class="layout-header">
        <div class="header-left">
          <el-icon class="collapse-btn" @click="collapse = !collapse" size="20">
            <Fold v-if="!collapse" /><Expand v-else />
          </el-icon>
          <el-breadcrumb separator="/" class="ml-4">
            <el-breadcrumb-item v-for="item in breadcrumbs" :key="item.path">{{ item.title }}</el-breadcrumb-item>
          </el-breadcrumb>
        </div>
        <div class="header-right">
          <el-dropdown @command="handleUserCommand">
            <span class="user-info">
              <el-icon><User /></el-icon>
              <span class="ml-1">{{ userStore.getUserInfo?.username || 'Admin' }}</span>
              <el-icon class="el-icon--right"><ArrowDown /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="personal">Profile</el-dropdown-item>
                <el-dropdown-item command="logout" divided>Logout</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>
      <el-main class="layout-main"><router-view /></el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/stores/user'
import { Fold, Expand, Menu, User, ArrowDown } from '@element-plus/icons-vue'
import { asyncRouterMap } from '@/router'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()
const collapse = ref(false)

const menuList = computed(() => asyncRouterMap)

const breadcrumbs = computed(() => {
  const parts = route.path.split('/').filter(Boolean)
  return parts.map(p => ({ path: p, title: p.charAt(0).toUpperCase() + p.slice(1) }))
})

function handleUserCommand(cmd) {
  if (cmd === 'logout') {
    userStore.logout()
    router.push('/login')
  } else if (cmd === 'personal') {
    router.push('/personal/personal-center')
  }
}
</script>

<style scoped>
.layout-container { height: 100vh; }
.layout-aside { background: #001529; transition: width 0.3s; overflow: hidden; }
.logo { height: 60px; display: flex; align-items: center; justify-content: center; border-bottom: 1px solid rgba(255,255,255,0.1); }
.logo-title { color: #fff; font-size: 18px; white-space: nowrap; }
.logo-title-mini { color: #fff; font-size: 16px; }
.layout-menu { border-right: none; }
.layout-header { display: flex; align-items: center; justify-content: space-between; padding: 0 20px; background: #fff; border-bottom: 1px solid #eee; height: 50px; }
.header-left, .header-right { display: flex; align-items: center; }
.collapse-btn { cursor: pointer; }
.user-info { cursor: pointer; display: flex; align-items: center; }
.layout-main { background: #f5f7fa; padding: 0; }
</style>
