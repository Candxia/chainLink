<script lang="ts" setup>
import { ref, watch } from 'vue'
import { ElImage, ElInput, ElText } from 'element-plus'

import defaultAvatar from '@/assets/imgs/google.webp'

const props = defineProps({
  optMsg: {
    type: String,
    default: () => undefined
  },
  optcode: {
    default: () => undefined
  }
})
const emits = defineEmits(['backData'])
const vercode = ref()

const exportData = () => {
  emits('backData', vercode.value)
}
watch(
  () => vercode.value,
  (val) => {
    exportData()
  }
)
watch(
  () => props.optcode,
  (val) => {
    vercode.value = val
  }
)
</script>

<template>
  <div class="w-full h-full flex flex-col">
    <div class="flex justify-center">
      <h2>校验OTP</h2>
    </div>
    <div class="flex-1 flex my-2 justify-center items-center">
      <div>
        <div class="avatar w-[230px] h-[230px] relative cursor-pointer">
          <ElImage class="w-[220px] h-[220px]" :src="defaultAvatar" fit="fill" />
        </div>
      </div>
      <div class="ml-3">
        <div class="w-full">您已经绑定并开启二次校验！</div>
        <div class="my-2">请在绑定设备查看OTP验证码</div>
        <div class="flex inline-grid flex-col">如有疑问请联系管理员!</div>
        <div class="my-2">
          <div class="my-2">4.请输入6位OTP验证码</div>
          <ElInput v-model.number:="vercode" maxlength="6" minlength="6" show-word-limit />
          <div v-if="!!props.optMsg">
            <ElText type="danger"> {{ props.optMsg }}</ElText>
          </div>
        </div>
      </div>
    </div>
    <div class="flex justify-center">
      <el-text type="danger">温馨提示: 更换设备或不慎删除OTP,请联系管理员!</el-text>
    </div>
  </div>
</template>
