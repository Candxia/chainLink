<template>
  <canvas
    ref="canvasRef"
    :width="width"
    :height="height"
    class="captcha-canvas cursor-pointer rounded"
    @click="handleClick"
  />
</template>

<script setup>
import { ref, onMounted, onUnmounted, nextTick, defineExpose } from 'vue'

const props = defineProps({
  width: { type: Number, default: 140 },
  height: { type: Number, default: 40 }
})

const emit = defineEmits(['refresh'])

const canvasRef = ref(null)
const value = ref('')
const isExpired = ref(false)
const remaining = ref(120)
const EXPIRY_MS = 120000 // 120 seconds
let countInterval = null

const chars = 'ABCDEFGHJKMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz23456789'

function randomNum(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min
}

function randomColor(opacity = 1) {
  const r = randomNum(0, 200)
  const g = randomNum(0, 200)
  const b = randomNum(0, 200)
  return opacity < 1 ? `rgba(${r},${g},${b},${opacity})` : `rgb(${r},${g},${b})`
}

function drawExpired() {
  const canvas = canvasRef.value
  if (!canvas) return
  const ctx = canvas.getContext('2d')
  if (!ctx) return

  clearRect(ctx, 0, 0, props.width, props.height)
  ctx.fillStyle = '#fce4e4'
  ctx.fillRect(0, 0, props.width, props.height)

  ctx.fillStyle = '#e74c3c'
  ctx.font = 'bold 18px "Arial", sans-serif'
  ctx.textAlign = 'center'
  ctx.textBaseline = 'middle'
  ctx.fillText('已失效 · 点击刷新', props.width / 2, props.height / 2)
}

function clearRect(ctx, x, y, w, h) {
  // Save context, clear, restore to avoid transform stack issues
  ctx.save()
  ctx.setTransform(1, 0, 0, 1, 0, 0)
  ctx.clearRect(x, y, w, h)
  ctx.restore()
}

function drawCountdown(ctx) {
  const s = Math.floor(remaining.value)
  if (s <= 0) return
  const text = `${s}s`
  // Clear bottom-right area before redrawing time
  clearRect(ctx, props.width - 42, props.height - 18, 42, 18)
  // Semi-transparent pill background
  ctx.fillStyle = 'rgba(0,0,0,0.4)'
  ctx.font = '11px "Arial", sans-serif'
  ctx.textAlign = 'right'
  ctx.textBaseline = 'bottom'
  ctx.fillText(text, props.width - 4, props.height - 3)
}

function startCountdown() {
  clearCountInterval()
  remaining.value = Math.floor(EXPIRY_MS / 1000)

  countInterval = setInterval(() => {
    remaining.value--
    if (remaining.value <= 0) {
      clearCountInterval()
      isExpired.value = true
      drawExpired()
    } else {
      const canvas = canvasRef.value
      if (!canvas) return
      const ctx = canvas.getContext('2d')
      if (!ctx) return
      drawCountdown(ctx)
    }
  }, 1000)
}

function clearCountInterval() {
  if (countInterval) {
    clearInterval(countInterval)
    countInterval = null
  }
}

function draw() {
  isExpired.value = false

  const canvas = canvasRef.value
  if (!canvas) return
  const ctx = canvas.getContext('2d')
  if (!ctx) return

  // Generate 4-char code
  let code = ''
  for (let i = 0; i < 4; i++) {
    code += chars[randomNum(0, chars.length - 1)]
  }
  value.value = code

  // Clear canvas
  clearRect(ctx, 0, 0, props.width, props.height)

  // Background
  ctx.fillStyle = '#f0f2f5'
  ctx.fillRect(0, 0, props.width, props.height)

  // Interference lines
  for (let i = 0; i < 3; i++) {
    ctx.beginPath()
    ctx.moveTo(randomNum(0, props.width), randomNum(0, props.height))
    ctx.lineTo(randomNum(0, props.width), randomNum(0, props.height))
    ctx.strokeStyle = randomColor(0.4)
    ctx.lineWidth = randomNum(1, 2)
    ctx.stroke()
  }

  // Noise dots
  for (let i = 0; i < 40; i++) {
    ctx.beginPath()
    ctx.arc(randomNum(0, props.width), randomNum(0, props.height), randomNum(1, 2), 0, 2 * Math.PI)
    ctx.fillStyle = randomColor(0.5)
    ctx.fill()
  }

  // Characters
  const charSpacing = props.width / 5
  for (let i = 0; i < code.length; i++) {
    const x = charSpacing * (i + 0.5) + randomNum(-4, 4)
    const y = props.height / 2 + randomNum(-6, 6)
    const angle = (randomNum(-25, 25) * Math.PI) / 180
    const fontSize = randomNum(18, 24)

    ctx.save()
    ctx.translate(x, y)
    ctx.rotate(angle)
    ctx.font = `bold ${fontSize}px "Arial", sans-serif`
    ctx.fillStyle = randomColor()
    ctx.textAlign = 'center'
    ctx.textBaseline = 'middle'
    ctx.fillText(code[i], 0, 0)
    ctx.restore()
  }

  // Initial countdown draw + start interval
  drawCountdown(ctx)
  startCountdown()
}

function handleClick() {
  if (isExpired.value) {
    refresh()
  } else {
    refresh()
  }
}

function refresh() {
  clearCountInterval()
  isExpired.value = false
  nextTick(() => draw())
  emit('refresh')
}

defineExpose({ value, refresh, isExpired })

onMounted(() => {
  draw()
})

onUnmounted(() => {
  clearCountInterval()
})
</script>

<style scoped>
.captcha-canvas {
  border: 1px solid var(--el-border-color);
  display: block;
  border-radius: 4px;
  cursor: pointer;
}
.captcha-canvas:hover {
  border-color: var(--el-color-primary);
}
</style>
