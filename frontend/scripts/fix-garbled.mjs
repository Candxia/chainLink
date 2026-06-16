import { readFileSync, writeFileSync, readdirSync } from 'fs'
import { resolve, dirname, join, relative } from 'path'
import { fileURLToPath } from 'url'

const __dirname = dirname(fileURLToPath(import.meta.url))
const viewsDir = resolve(__dirname, '../src/views/System')

const patches = [
  [/鍒犻[\w]?/g, 'Delete'],
  [/缂栬緫[\w]?/g, 'Edit'],
  [/鏂板[\w]?/g, 'Add New'],
  [/鎺ュ彛[\w]?/g, 'API'],
  [/璐﹀彿[\w]?/g, 'Account'],
  [/绠＄悊[\w]?/g, 'Mgmt'],
  [/绫诲瀷[\w]?/g, 'Type'],
  [/鐘舵[\w]?/g, 'Status'],
  [/鏃堕棿[\w]?/g, 'Time'],
  [/鍒涘缓[\w]?/g, 'Create'],
  [/鏇存柊[\w]?/g, 'Update'],
  [/澶囨敞[\w]?/g, 'Remark'],
  [/璇锋眰[\w]?/g, 'Request'],
  [/鏂瑰紡[\w]?/g, 'Method'],
  [/鎿嶄綔[\w]?/g, 'Action'],
  [/鍏抽敭[\w]?/g, 'Key'],
  [/鍚堣[\w]?/g, 'Allow'],
  [/绂佺[\w]?/g, 'Disable'],
  [/鍐呴[\w]?/g, 'Internal'],
  [/澶栭[\w]?/g, 'External'],
  [/鍚[\w]?/g, 'Enable'],
  [/鍋滅[\w]?/g, 'Disable'],
  [/鎻愮[\w]?/g, 'Tips'],
  [/纭[\w]?/g, 'Confirm'],
  [/涓婚敭[\w]?/g, 'ID'],
  [/鍚嶇[\w]?/g, 'Name'],
  [/鍙风爜[\w]?/g, 'Code'],
  [/璺緞[\w]?/g, 'Path'],
  [/濮撳悕[\w]?/g, 'Name'],
  [/鏄电О[\w]?/g, 'Nickname'],
  [/閭[\w]?/g, 'Email'],
  [/绯荤粺[\w]?/g, 'System'],
  [/鍙傛暟[\w]?/g, 'Params'],
]

function fixFile(fp) {
  let text = readFileSync(fp, 'utf-8')
  const orig = text
  for (const [pattern, replacement] of patches) {
    text = text.replace(pattern, replacement)
  }
  text = text.replace(/(label="[^"]*)\?+/g, '$1')
  text = text.replace(/(placeholder="[^"]*)\?+/g, '$1')
  if (text !== orig) {
    writeFileSync(fp, text, 'utf-8')
    return true
  }
  return false
}

function walk(dir) {
  const entries = readdirSync(dir, { withFileTypes: true })
  for (const e of entries) {
    const p = join(dir, e.name)
    if (e.isDirectory()) walk(p)
    else if (e.name.endsWith('.vue')) {
      if (fixFile(p)) console.log('Fixed: ' + relative(viewsDir, p))
    }
  }
}

walk(viewsDir)
console.log('Done')
