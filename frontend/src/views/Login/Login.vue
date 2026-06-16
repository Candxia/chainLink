<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 to-indigo-100">
    <el-card class="w-96" shadow="xl">
      <template #header><div class="text-center"><h2 class="text-xl font-bold">ChainLink System</h2></div></template>
      <el-form ref="formRef" :model="form" :rules="rules" label-position="top">
        <el-form-item label="Username" prop="username"><el-input v-model="form.username" placeholder="Username" /></el-form-item>
        <el-form-item label="Password" prop="password"><el-input v-model="form.password" type="password" show-password placeholder="Password" /></el-form-item>
        <el-form-item><el-button type="primary" class="w-full" :loading="loading" @click="handleLogin">Sign In</el-button></el-form-item>
      </el-form>
    </el-card>
  </div>
</template>
<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { loginApi } from '@/api/login'
import { useUserStore } from '@/stores/user'
const router = useRouter()
const userStore = useUserStore()
const formRef = ref(null)
const loading = ref(false)
const form = reactive({ username: '', password: '' })
const rules = { username: [{ required: true, message: 'Required', trigger: 'blur' }], password: [{ required: true, message: 'Required', trigger: 'blur' }] }
async function handleLogin() {
  if (!await formRef.value.validate().catch(() => false)) return
  loading.value = true
  try {
    const res = await loginApi(form)
    if (res.code === 0) {
      userStore.setToken(res.data.token || res.data.access_token)
      userStore.setUserInfo(res.data.userInfo || res.data.user)
      ElMessage.success('Login successful')
      router.push('/')
    } else ElMessage.error(res.msg || 'Login failed')
  } catch (e) { ElMessage.error('Network error') } finally { loading.value = false }
}
</script>

