import { readFileSync, writeFileSync, mkdirSync, cpSync, existsSync } from 'fs'
import { resolve, dirname, join } from 'path'
import { fileURLToPath } from 'url'

const __dirname = dirname(fileURLToPath(import.meta.url))

const srcRoot = resolve(__dirname, '../../../sample_framework/frontend/src')
const dstRoot = resolve(__dirname, '../src')

const compPairs = [
  // ContentWrap
  ['components/ContentWrap/src/ContentWrap.vue', 'components/ContentWrap/src/ContentWrap.vue'],
  ['components/ContentWrap/index.ts', 'components/ContentWrap/index.ts'],
  // Table
  ['components/Table/src/Table.vue', 'components/Table/src/Table.vue'],
  ['components/Table/index.ts', 'components/Table/index.ts'],
  // Form
  ['components/Form/src/Form.vue', 'components/Form/src/Form.vue'],
  ['components/Form/index.ts', 'components/Form/index.ts'],
  // Dialog
  ['components/Dialog/src/Dialog.vue', 'components/Dialog/src/Dialog.vue'],
  ['components/Dialog/index.ts', 'components/Dialog/index.ts'],
  // UseDataTable
  ['components/UseDataTable/src/UseDataTable.vue', 'components/UseDataTable/src/UseDataTable.vue'],
  ['components/UseDataTable/index.ts', 'components/UseDataTable/index.ts'],
  ['components/UseDataTable/types/index.ts', 'components/UseDataTable/types/index.ts'],
  // Permission
  ['components/Permission/src/Permission.vue', 'components/Permission/src/Permission.vue'],
  ['components/Permission/index.ts', 'components/Permission/index.ts'],
  // Button
  ['components/Button/src/Button.vue', 'components/Button/src/Button.vue'],
  ['components/Button/index.ts', 'components/Button/index.ts'],
  // Icon
  ['components/Icon/src/Icon.vue', 'components/Icon/src/Icon.vue'],
  ['components/Icon/index.ts', 'components/Icon/index.ts'],
  // Search
  ['components/Search/src/Search.vue', 'components/Search/src/Search.vue'],
  ['components/Search/index.ts', 'components/Search/index.ts'],
  // SearchParams -- actually UseDataTable has its own search
  // Descriptions
  ['components/Descriptions/index.ts', 'components/Descriptions/index.ts'],
  ['components/Descriptions/src/Descriptions.vue', 'components/Descriptions/src/Descriptions.vue'],
  // Highlight
  ['components/Highlight/src/Highlight.vue', 'components/Highlight/src/Highlight.vue'],
  ['components/Highlight/index.ts', 'components/Highlight/index.ts'],
  // components/index.ts
  ['components/index.ts', 'components/index.ts'],
  // hooks
  ['hooks/web/useTable.ts', 'hooks/web/useTable.ts'],
  ['hooks/web/useForm.ts', 'hooks/web/useForm.ts'],
  ['hooks/web/useI18n.ts', 'hooks/web/useI18n.ts'],
  ['hooks/web/useDesign.ts', 'hooks/web/useDesign.ts'],
  ['hooks/web/useValidator.ts', 'hooks/web/useValidator.ts'],
  ['hooks/web/useNProgress.ts', 'hooks/web/useNProgress.ts'],
  ['hooks/web/useTitle.ts', 'hooks/web/useTitle.ts'],
  ['hooks/web/useStorage.ts', 'hooks/web/useStorage.ts'],
  ['hooks/web/useCrudSchemas.ts', 'hooks/web/useCrudSchemas.ts'],
  ['hooks/web/usePageLoading.ts', 'hooks/web/usePageLoading.ts'],
  ['hooks/web/useSearch.ts', 'hooks/web/useSearch.ts'],
  // utils
  ['utils/routerHelper.ts', 'utils/routerHelper.ts'],
  ['utils/is.ts', 'utils/is.ts'],
  ['utils/debounce.ts', 'utils/debounce.ts'],
  ['utils/storage.ts', 'utils/storage.ts'],
  ['utils/dateUtil.ts', 'utils/dateUtil.ts'],
  ['utils/auth.ts', 'utils/auth.ts'],
  ['utils/formatter.ts', 'utils/formatter.ts'],
  ['utils/defaultActSlotBuilder.ts', 'utils/defaultActSlotBuilder.ts'],
  ['utils/roleUtils.ts', 'utils/roleUtils.ts'],
  ['utils/preventDevTools.ts', 'utils/preventDevTools.ts'],
  ['utils/trim.ts', 'utils/trim.ts'],
  ['utils/dateTimeEx.ts', 'utils/dateTimeEx.ts'],
  // directives
  ['directives/permission/index.ts', 'directives/permission/index.ts'],
  ['directives/index.ts', 'directives/index.ts'],
  // constants
  ['constants/index.ts', 'constants/index.ts'],
  // types (minimal set for project)
  ['types/global.d.ts', 'types/global.d.ts'],
]

for (const [relSrc, relDst] of compPairs) {
  const fullSrc = join(srcRoot, relSrc)
  const fullDst = join(dstRoot, relDst)
  
  if (!existsSync(fullSrc)) {
    console.log(`SKIP (not found): ${relSrc}`)
    continue
  }
  
  mkdirSync(dirname(fullDst), { recursive: true })
  
  // For .tsx and .ts files, convert to .js (or .jsx)
  // Actually just copy as-is since Vite can handle TS
  let content = readFileSync(fullSrc, 'utf-8')
  
  // Fix import paths
  content = content.replace(/@\/store\/modules\//g, '@/stores/')
  content = content.replace(/@\/api\/(admin|api|role|dept|menu)\b/g, '@/api/system')
  content = content.replace(/@\/api\/system\/tools/g, '@/api/system/tools')
  
  writeFileSync(fullDst, content, 'utf-8')
  console.log(`Copied: ${relSrc}`)
}

console.log('\nDone copying all components')
