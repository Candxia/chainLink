<template>
  <div class="login-container">
    <!-- Left: gradient background + welcome -->
    <div class="login-left">
      <div class="left-content">
        <div class="welcome-title">ChainLink System</div>
        <div class="welcome-subtitle">Welcome Back</div>
        <div class="welcome-desc">Please sign in to your account to continue</div>
      </div>
    </div>

    <!-- Right: login form card -->
    <div class="login-right">
      <div class="login-card-wrapper">
        <el-card class="login-card" shadow="xl">
          <template #header>
            <div class="card-header">
              <h2 class="card-title">Sign In</h2>
            </div>
          </template>

          <el-form
            ref="formRef"
            :model="form"
            :rules="rules"
            label-position="top"
            size="large"
            @keyup.enter="handleLogin"
          >
            <el-form-item label="Username" prop="username">
              <el-input
                v-model="form.username"
                placeholder="Username"
                :prefix-icon="User"
              />
            </el-form-item>

            <el-form-item label="Password" prop="password">
              <el-input
                v-model="form.password"
                type="password"
                show-password
                placeholder="Password"
                :prefix-icon="Lock"
              />
            </el-form-item>

            <el-form-item label="验证码" prop="captchaInput">
              <div class="captcha-row">
                <el-input
                  v-model="form.captchaInput"
                  placeholder="验证码"
                  class="captcha-input"
                  maxlength="4"
                  @keyup.enter="handleLogin"
                />
                <Captcha ref="captchaRef" class="captcha-display" />
              </div>
            </el-form-item>

            <el-form-item>
              <el-button
                type="primary"
                class="w-full"
                :loading="loading"
                @click="handleLogin"
              >
                登 录
              </el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { User, Lock } from '@element-plus/icons-vue'
import { loginApi } from '@/api/login'
import { useUserStore } from '@/stores/user'
import Captcha from './components/Captcha.vue'

const router = useRouter()
const userStore = useUserStore()

const formRef = ref(null)
const captchaRef = ref(null)
const loading = ref(false)

const form = reactive({
  username: '',
  password: '',
  captchaInput: '',
})

const rules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' }
  ],
  captchaInput: [
    { required: true, message: '请输入验证码', trigger: 'blur' },
    { min: 4, max: 4, message: '验证码为4位字符', trigger: 'blur' }
  ]
}

async function handleLogin() {
  if (!await formRef.value.validate().catch(() => false)) return

  // Check if captcha expired
  if (captchaRef.value?.isExpired) {
    ElMessage.warning('验证码已失效，请点击刷新')
    form.captchaInput = ''
    nextTick(() => captchaRef.value?.refresh())
    return
  }

  // Validate captcha locally (case-insensitive)
  const captchaValue = captchaRef.value?.value
  if (!captchaValue || form.captchaInput.toLowerCase() !== captchaValue.toLowerCase()) {
    ElMessage.error('验证码不正确')
    form.captchaInput = ''
    nextTick(() => captchaRef.value?.refresh())
    return
  }

  loading.value = true
  try {
    const res = await loginApi({
      username: form.username,
      password: form.password
    })
    if (res.code === 0) {
      userStore.setToken(res.data.token || res.data.access_token)
      userStore.setUserInfo(res.data.userInfo || res.data.user)
      ElMessage.success('登录成功')
      router.push('/')
    } else {
      ElMessage.error(res.msg || '登录失败')
      // Refresh captcha on failure
      nextTick(() => captchaRef.value?.refresh())
    }
  } catch (e) {
    ElMessage.error('网络错误，请重试')
    nextTick(() => captchaRef.value?.refresh())
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  width: 100%;
  height: 100vh;
  overflow: hidden;
}

/* Left side ~50% */
.login-left {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #1a73e8 0%, #0d47a1 50%, #001f3f 100%);
  position: relative;
  overflow: hidden;
}

.login-left::before {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle, rgba(255,255,255,0.03) 0%, transparent 70%);
  animation: pulse 8s ease-in-out infinite;
}

@keyframes pulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.1); }
}

.left-content {
  position: relative;
  z-index: 1;
  text-align: center;
  color: #fff;
  padding: 40px;
}

.welcome-title {
  font-size: 36px;
  font-weight: 700;
  margin-bottom: 16px;
  letter-spacing: 2px;
}

.welcome-subtitle {
  font-size: 24px;
  font-weight: 500;
  margin-bottom: 12px;
  opacity: 0.9;
}

.welcome-desc {
  font-size: 14px;
  opacity: 0.7;
  line-height: 1.6;
}

/* Right side ~50% */
.login-right {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--el-bg-color-page, #f5f7fa);
  padding: 24px;
}

.login-card-wrapper {
  width: 100%;
  max-width: 420px;
}

.login-card {
  border-radius: 12px;
}

.card-header {
  display: flex;
  justify-content: center;
}

.card-title {
  font-size: 22px;
  font-weight: 600;
  margin: 0;
  color: var(--el-text-color-primary);
}

.captcha-row {
  display: flex;
  gap: 10px;
  align-items: center;
  width: 100%;
}

.captcha-input {
  width: 110px;
  flex-shrink: 0;
}

.captcha-display {
  flex-shrink: 0;
}

/* Responsive: stack on small screens */
@media (max-width: 768px) {
  .login-container {
    flex-direction: column;
  }
  .login-left {
    height: 30vh;
  }
  .login-right {
    height: 70vh;
  }
}
</style>
