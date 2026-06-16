export const isUrl = (path) => {
  if (!path) return false
  return /^(https?:|mailto:|tel:)/.test(path)
}
export const isExternal = (path) => /^(https?:|mailto:|tel:)/.test(path)
