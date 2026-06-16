import { readFileSync, writeFileSync, mkdirSync, cpSync } from 'fs'
import { resolve, dirname, join } from 'path'
import { fileURLToPath } from 'url'

const __dirname = dirname(fileURLToPath(import.meta.url))

const src = resolve(__dirname, '../../../sample_framework/frontend/src')
const dst = resolve(__dirname, '../src')

const pairs = [
  // Login
  ['views/Login/Login.vue', 'views/Login/Login.vue'],
  ['views/Login/components/LoginForm.vue', 'views/Login/components/LoginForm.vue'],
  ['views/Login/components/optVar.vue', 'views/Login/components/optVar.vue'],
  ['views/Login/components/index.ts', 'views/Login/components/index.ts'],
  // System admin
  ['views/System/admin/SysAdmin.vue', 'views/System/admin/SysAdmin.vue'],
  ['views/System/admin/components/adminWrite.tsx', 'views/System/admin/components/adminWrite.tsx'],
  ['views/System/admin/components/UploadAvatar.vue', 'views/System/admin/components/UploadAvatar.vue'],
  // System api
  ['views/System/api/SysApi.vue', 'views/System/api/SysApi.vue'],
  ['views/System/api/components/apiWrite.tsx', 'views/System/api/components/apiWrite.tsx'],
  // System role
  ['views/System/role/SysRole.vue', 'views/System/role/SysRole.vue'],
  ['views/System/role/components/roleWrite.tsx', 'views/System/role/components/roleWrite.tsx'],
  // System dept
  ['views/System/dept/SysDept.vue', 'views/System/dept/SysDept.vue'],
  ['views/System/dept/components/deptWrite.tsx', 'views/System/dept/components/deptWrite.tsx'],
  // System menu
  ['views/System/menu/SysMenu.vue', 'views/System/menu/SysMenu.vue'],
  ['views/System/menu/components/menuWrite.tsx', 'views/System/menu/components/menuWrite.tsx'],
]

for (const [relSrc, relDst] of pairs) {
  const fullSrc = join(src, relSrc)
  const fullDst = join(dst, relDst)
  
  // Create destination directory
  mkdirSync(dirname(fullDst), { recursive: true })
  
  // Read source content and fix import paths
  let content = readFileSync(fullSrc, 'utf-8')
  
  // Fix API import paths: from '@/api/admin' → '@/api/system'
  content = content.replace(/from '@\/api\/admin'/g, "from '@/api/system'")
  content = content.replace(/from '@\/api\/api'/g, "from '@/api/system'")
  content = content.replace(/from '@\/api\/role'/g, "from '@/api/system'")
  content = content.replace(/from '@\/api\/dept'/g, "from '@/api/system'")
  content = content.replace(/from '@\/api\/menu'/g, "from '@/api/system'")
  content = content.replace(/from '@\/api\/personal'/g, "from '@/api/personal'")
  content = content.replace(/from '@\/api\/system\/tools'/g, "from '@/api/system/tools'")
  
  // Fix store import paths
  content = content.replace(/from '@\/store\/modules\//g, "from '@/stores/")
  content = content.replace(/'\.\.\/\.\.\/\.\.\/store\/modules\//g, "'../../stores/")
  
  // Fix Layout import
  content = content.replace(/from '@\/utils\/routerHelper'/g, "from '@/utils/routerHelper'")
  
  // Fix hooks imports
  content = content.replace(/from '@\/hooks\/web\/useI18n'/g, "from '@/hooks/web/useI18n'")
  content = content.replace(/from '@\/hooks\/web\/useValidator'/g, "from '@/hooks/web/useValidator'")
  content = content.replace(/from '@\/hooks\/web\/useDesign'/g, "from '@/hooks/web/useDesign'")
  content = content.replace(/from '@\/hooks\/web\/useTable'/g, "from '@/hooks/web/useTable'")
  
  // Fix type imports (remove TypeScript type-only imports)
  // Convert TSX to JSX by removing TypeScript-specific syntax
  if (relSrc.endsWith('.tsx')) {
    content = content.replace(/import type \{/g, 'import {')
  }
  
  writeFileSync(fullDst, content, 'utf-8')
  console.log(`Copied: ${relSrc} → ${relDst}`)
}

console.log('\nDone copying all system pages')
