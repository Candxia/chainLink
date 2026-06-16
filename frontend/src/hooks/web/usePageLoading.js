import { ref } from 'vue'
export const usePageLoading = () => {
  const pageLoading = ref(false)
  const loadStart = () => { pageLoading.value = true }
  const loadDone = () => { pageLoading.value = false }
  return { pageLoading, loadStart, loadDone }
}
