import { hasPermi } from './permission/hasPermi'

export const setupPermission = (app) => {
  app.directive('permi', hasPermi)
}
