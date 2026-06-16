<script setup lang="tsx">
import { computed, h, onBeforeMount, onMounted, reactive, ref, unref, watch } from 'vue'
import { Form, FormSchema } from '@/components/Form'
import { useI18n } from '@/hooks/web/useI18n'
import { ElButton, ElDialog, ElLoading, ElMessage, ElOption, ElSelect } from 'element-plus'
import { useForm } from '@/hooks/web/useForm'
import { getImgCodeAPi, getShowImgCodeAPi, loginApi } from '@/api/login'
import { useAppStore } from '@/stores/app'
import type { RouteLocationNormalizedLoaded } from 'vue-router'
import { UserType } from '@/api/login/types'
import { useValidator } from '@/hooks/web/useValidator'
import { useUserStore } from '@/stores/user'
import { BaseButton } from '@/components/Button'
import { Icon } from '@iconify/vue'

import { getOptBindApi, getOptCheckApi } from '@/api/personal'
import OptVar from './optVar.vue'
import { useLocaleStore } from '@/stores/locale'
import { useLocale } from '@/hooks/web/useLocale'
import { useRouter } from 'vue-router'
import { getRole } from '@/utils/roleUtils'

const { required } = useValidator()

const emit = defineEmits(['to-register'])

const appStore = useAppStore()

const userStore = useUserStore()

const { currentRoute, push } = useRouter()

const { t } = useI18n()


const btn_code = ref('get')
const btn_code_label = ref('')
const clickThumb = ref()

const localeStore = useLocaleStore()
const langMap = computed(() => localeStore.getLocaleMap)
const currentLang = computed(() => localeStore.getCurrentLocale)
const setLang = (lang: LocaleType) => {
  if (lang === unref(currentLang).lang) return
  
  window.location.reload()
  localeStore.setCurrentLocale({
    lang
  })
  const { changeLocale } = useLocale()
  changeLocale(lang)
}

async function getCodeImgShow() {
  const { code, data } = await getShowImgCodeAPi()
  console.log(data)
  if (code == 0 && data.captcha) {
    await setSchema([{ field: 'tool', path: 'remove', value: false }])
  }
}

const rules = {
  username: [
    required(),
    {
      pattern: /^[a-zA-Z0-9]{6,32}$/, 
      message: t('system.admin.usernameMessage'),
      trigger: 'blur'
    }
  ],
  password: [
    required(),
    {
      pattern: /^(?=[A-Za-z])(?=.*[a-z])(?=.*[A-Z]).{8,32}$/,
      message: t('system.admin.passwordMessage'),
      trigger: 'blur'
    }
  ]
}

const schema = reactive<FormSchema[]>([
  {
    field: 'title',
    colProps: {
      span: 24
    },
    formItemProps: {
      slots: {
        default: () => {
          return <h2 class="text-2xl font-bold text-center w-[100%]">{t('login.login')}</h2>
        }
      }
    }
  },
  {
    field: 'username',
    label: t('login.username'),
    // value: 'admin',
    component: 'Input',
    colProps: {
      span: 24
    },
    componentProps: {
      placeholder: '输入账号'
    }
  },
  {
    field: 'password',
    label: t('login.password'),
    // value: 'admin',
    component: 'InputPassword',
    colProps: {
      span: 24
    },
    componentProps: {
      style: {
        width: '100%'
      },
      placeholder: '输入密码'
    }
  },
  {
    field: 'language',
    label: t('setting.chooseLanguage'),
    colProps: {
      span: 24
    },
    formItemProps: {
      slots: {
        default: () =>
          h(
            ElSelect,
            {
              modelValue: currentLang.value.lang,
              onChange: (value: any) => {
                setLang(value)
              }
            },
            () =>
              langMap.value?.map((item: any) =>
                h(ElOption, { label: item.name, value: item.lang }, () => item.name)
              )
          )
      }
    }
  },
  {
    field: 'tool',
    colProps: {
      span: 24
    },
    remove: true,
    formItemProps: {
      slots: {
        default: () => {
          if (unref(btn_code.value == 'get')) {
            return (
              <>
                <div class="flex justify-between items-center w-[100%]">
                  <ElButton type="primary" plain class="w-[100%]" onClick={getCaptcha}>
                    <Icon icon="bi:shield-fill" />
                    {t('login.getCode')}
                  </ElButton>
                </div>
              </>
            )
          } else if (btn_code.value == 'loading') {
            return (
              <>
                <div class="flex justify-between items-center w-[100%]">
                  <ElButton type="warning" plain class="w-[100%]" onClick={getCaptcha} loading>
                    {t('login.setCode')}
                  </ElButton>
                </div>
              </>
            )
          } else if (btn_code.value == 'error') {
            return (
              <>
                <div class="flex justify-between items-center w-[100%]">
                  <ElButton type="danger" plain class="w-[100%]" onClick={getCaptcha}>
                    <Icon icon="iconamoon:shield-no-fill" />
                    {btn_code_label.value}
                  </ElButton>
                </div>
              </>
            )
          } else {
            return (
              <>
                <div class="flex justify-between items-center w-[100%]">
                  <ElButton type="success" plain class="w-[100%]" onClick={getCaptcha}>
                    {/* <ElIcon><SuccessFilled /></ElIcon> */}
                    <Icon icon="iconamoon:shield-yes-fill" />
                    {t('login.codeSucess')}
                  </ElButton>
                </div>
              </>
            )
          }
        }
      }
    }
  },
  {
    field: 'login',
    colProps: {
      span: 24
    },
    formItemProps: {
      slots: {
        default: () => {
          return (
            <>
              <div class="w-[100%]">
                <BaseButton
                  loading={loading.value}
                  type="primary"
                  class="w-[100%]"
                  onClick={signIn}
                >
                  {t('login.login')}
                </BaseButton>
              </div>
            </>
          )
        }
      }
    }
  }
])

const remember = ref(userStore.getRememberMe)

const initLoginInfo = () => {
  const loginInfo = userStore.getLoginInfo
  if (loginInfo) {
    const { username, password } = loginInfo
    setValues({ username, password })
  }
}

onMounted(() => {
  initLoginInfo()
})

onBeforeMount(() => {
  getCodeImgShow()
})

const { formRegister, formMethods } = useForm()
const { getFormData, getElFormExpose, setValues, setSchema } = formMethods

const loading = ref(false)

const redirect = ref<string>('')

watch(
  () => currentRoute.value,
  (route: RouteLocationNormalizedLoaded) => {
    redirect.value = route?.query?.redirect as string
  },
  {
    immediate: true
  }
)

const signIn = async () => {
  const formRef = await getElFormExpose()
  await formRef?.validate(async (isValid) => {
    if (isValid) {
      loading.value = true
      const formData = await getFormData<UserType>()
      try {
        formData.captcha = captcha.value
        formData.captcha_id = captcha_id.value
        const res: any = await loginApi(formData)
        if (res.code == 0) {
          userStore.setToken(res.data.token)
          userStore.setRememberMe(unref(remember.value))
          const userInfo = {
            ...res.data,
            username: formData.username
          }
          userStore.setUserInfo(userInfo)
          
          const optStats = res.data.redirect
          if (optStats == 2) {
            
            dialogOpt.value = true
          } else if (optStats == 3) {
            
            dialogVaropt.value = true
          } else {
            loading.value = false
            setRouters()
          }
        } else {
          if (res.code == '15132') {
            btn_code.value = 'error'
            btn_code_label.value = res.message
          } else {
            btn_code.value = 'get'
          }
          loading.value = false
        }
      } finally {
        captcha.value = []
        formData.captcha = []
      }
    }
  })
}

const dialogVisible = ref(false)

const getCaptcha = () => {
  btn_code.value = 'loading'
  dialogVisible.value = true
  getCode()
}

const clickImage = ref()
const captcha_id = ref()
const captcha: any = ref([])
const getCode = () => {
  captcha.value = []
  getImgCodeAPi().then((res: any) => {
    if (res.code == 0) {
      clickImage.value = res.data.image
      clickThumb.value = res.data.thumb
      captcha_id.value = res.data.captcha_id
    }
  })
}

const clickData = reactive({
  image: clickImage,
  thumb: clickThumb
})


const clickEvents = {
  click(x: number, y: number): void {
    console.log('click >>>>>>>', x, y)
  },
  confirm(dots: any, clear: Function): void {
    if (dots.length <= 0) {
      ElMessage.warning(t('请进行人机验证再操作'))
      return
    }
    dots.filter((dots) => {
      captcha.value.push(dots.x)
      captcha.value.push(dots.y)
    })
    dialogVisible.value = false
    btn_code.value = 'over'
    setTimeout(() => {
      clear()
    }, 100)
  },
  refresh(): void {
    captcha.value = []
    getCode()
  },
  close(): void {
    captcha.value = []
    btn_code.value = 'get'
    dialogVisible.value = false
  }
}

const dialogOpt = ref(false)
const optData = ref()
const getOptData = (data) => {
  optData.value = data
}
const optMsg: any = ref<string>('')
const optChagne = async () => {
  if (optData.value?.code) {
    const res: any = await getOptBindApi(optData.value)
    if (!unref(res.code)) {
      ElMessage.success('设置成功')
      await setRouters()
      dialogOpt.value = false
    } else {
      const public_Key: any = localStorage.getItem('public_Key')
      const refresh_cdn: any = localStorage.getItem('refresh_cdn')
      localStorage.clear()
      localStorage.setItem('public_Key', public_Key)
      localStorage.setItem('refresh_cdn', refresh_cdn)
      optcode.value = undefined
      optMsg.value = res.message
    }
  } else {
    ElMessage.error(t('login.codePlaceholder'))
    optMsg.value = t('login.codePlaceholder')
  }
}

const dialogVaropt = ref(false)
const optcode = ref()
const getOptCodeData = (data) => {
  optcode.value = data
}
const isSlow = ref(false)
const sure = async () => {
  if (optcode.value) {
    const res: any = await getOptCheckApi({ code: optcode.value })
    if (!unref(res.code)) {
      optMsg.value = null
      dialogVaropt.value = false
      isSlow.value = true
      await setRouters()
      ElLoading.service({
        lock: true,
        text: 'Loading',
        background: 'rgba(250, 246, 246, 0.7)'
      })
    } else {
      const public_Key: any = localStorage.getItem('public_Key')
      const refresh_cdn: any = localStorage.getItem('refresh_cdn')
      localStorage.clear()
      localStorage.setItem('public_Key', public_Key)
      localStorage.setItem('refresh_cdn', refresh_cdn)
      optcode.value = undefined
      optMsg.value = res.message
    }
  } else {
    optMsg.value = t('login.codePlaceholder')
  }
}

const setRouters = async () => {
  
  if (appStore.getDynamicRouter) {
    await getRole()
    if (redirect.value && redirect.value === '/user/uinfo') {
      await push({ path: '/' })
    } else {
      await push({ path: redirect.value || '/' })
    }
  } else {
    await getRole('static')
    await push({ path: redirect.value || '/' })
  }
}
const optClose = () => {
  optMsg.value = null
  optData.value = undefined
  if (!isSlow.value) {
    loading.value = false
    btn_code.value = 'get'
  }
}
</script>

<template>
  <Form
    :schema="schema"
    :rules="rules"
    label-position="top"
    hide-required-asterisk
    size="large"
    class="dark:(border-1 border-[var(--el-border-color)] border-solid)"
    @register="formRegister"
  />
  <ElDialog
    v-model="dialogVisible"
    :show-close="false"
    :close-on-click-modal="false"
    width="360"
    :title="t('login.verification')"
    center
  >
    <gocaptcha-click :data="clickData" :events="clickEvents" />
  </ElDialog>
  <ElDialog
    v-model="dialogOpt"
    title=""
    width="800px"
    :close-on-click-modal="false"
    destroy-on-close
  >
    <Opt ref="optRef" @back-data="getOptData" :optMsg="optMsg" />
    <template #footer>
      <ElButton type="primary" @click="optChagne"> {{ t('common.ok') }}</ElButton>
    </template>
  </ElDialog>
  <ElDialog
    v-model="dialogVaropt"
    title=""
    width="800px"
    :close-on-click-modal="false"
    destroy-on-close
    @closed="optClose"
  >
    <OptVar @back-data="getOptCodeData" :optcode="optcode" :optMsg="optMsg" />
    <template #footer>
      <ElButton type="primary" @click="sure"> {{ t('common.ok') }}</ElButton>
    </template>
  </ElDialog>
</template>

