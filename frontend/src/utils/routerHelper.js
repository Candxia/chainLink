export const Layout = () => import('@/layout/Layout.vue')
export const getParentLayout = () => ({ name: 'ParentLayout' })
export const flatMultiLevelRoutes = (routes) => routes
export const generateRoutesByFrontEnd = (routes, keys) => routes
export const generateRoutesByServer = (routers) => routers || []
export const getRawRoute = (route) => route
