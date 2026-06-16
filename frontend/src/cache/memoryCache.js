class MemoryCache {
  constructor() {
    this.cache = new Map()
    this.timeouts = new Map()
  }

  set(key, value, expire = 0) {
    this.cache.set(key, value)
    if (expire > 0) {
      const timeout = setTimeout(() => {
        this.delete(key)
      }, expire)
      this.timeouts.set(key, timeout)
    }
  }

  get(key) {
    return this.cache.get(key)
  }

  delete(key) {
    this.cache.delete(key)
    if (this.timeouts.has(key)) {
      clearTimeout(this.timeouts.get(key))
      this.timeouts.delete(key)
    }
  }

  has(key) {
    return this.cache.has(key)
  }

  clear() {
    this.cache.clear()
    this.timeouts.forEach((timeout) => clearTimeout(timeout))
    this.timeouts.clear()
  }
}

export const memoryCache = new MemoryCache()
