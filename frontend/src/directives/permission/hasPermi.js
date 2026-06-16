import { useUserStore } from '@/stores/user'

export const hasPermi = {
  mounted(el, binding) {
    const { value } = binding
    if (!value) return

    const userStore = useUserStore()
    const permissions = userStore.getPermissions || []

    if (value && value instanceof Array && value.length > 0) {
      const hasPermission = permissions.some((perm) => value.includes(perm))
      if (!hasPermission) {
        el.parentNode && el.parentNode.removeChild(el)
      }
    } else if (typeof value === 'string') {
      if (!permissions.includes(value)) {
        el.parentNode && el.parentNode.removeChild(el)
      }
    }
  },
}
