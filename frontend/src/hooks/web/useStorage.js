export const useStorage = (type = 'localStorage') => ({
  getStorage: (key) => localStorage.getItem(key),
  setStorage: (key, val) => localStorage.setItem(key, val),
  clear: () => localStorage.clear()
})
