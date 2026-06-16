import { readFileSync, writeFileSync } from 'fs'

const files = ['src/views/system/LogList.vue', 'src/views/system/ConfigList.vue', 'src/views/system/ApiRuleList.vue']

for (const fp of files) {
  let c = readFileSync(fp, 'utf-8')
  if (c.startsWith('<template>')) {
    console.log('OK: ' + fp)
    continue
  }
  // Remove prefix garbage, ensure outer template
  const templateIdx = c.indexOf('<template')
  if (templateIdx > 0) c = c.substring(templateIdx)
  if (!c.startsWith('<template>')) {
    c = '<template>\n' + c
  }
  // Ensure closing </template>
  const scriptMatch = c.match(/\n<script/s)
  if (scriptMatch && c.includes('</div>') && !c.includes('</template>')) {
    const lastDivIdx = c.lastIndexOf('</div>')
    c = c.slice(0, lastDivIdx + 7) + '\n</template>' + c.slice(lastDivIdx + 7)
  }
  writeFileSync(fp, c, 'utf-8')
  console.log('Fixed: ' + fp)
}
console.log('Done')
