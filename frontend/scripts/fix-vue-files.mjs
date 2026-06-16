import { readFileSync, writeFileSync, existsSync } from 'fs'
import { resolve, dirname } from 'path'
import { fileURLToPath } from 'url'

const __dirname = dirname(fileURLToPath(import.meta.url))
const viewsDir = resolve(__dirname, '../src/views')

function fixSysPage(page, dir) {
  const fp = resolve(viewsDir, dir || `System/${page}`, `${page}.vue`)
  if (!existsSync(fp)) {
    console.log(`SKIP: ${page} (not found)`)
    return
  }
  let c = readFileSync(fp, 'utf-8')
  
  // If file starts directly with <template #xxx>, it's missing the outer wrapper
  if (c.match(/^<template #/)) {
    c = `<template>
  <UseDataTable
    :searchSchema="searchSchema"
    :columns="columns"
    :data="tableData"
    :loading="loading"
    :pagination="pagination"
    :toolButtons="toolButtons"
    @search="handleSearch"
    @reset="handleReset"
    @size-change="handleSizeChange"
    @current-change="handleCurrentChange"
  >
` + c
    writeFileSync(fp, c, 'utf-8')
    console.log(`FIXED: ${page}`)
  } else {
    console.log(`OK: ${page}`)
  }
}

fixSysPage('SysAdmin', 'System/admin')
fixSysPage('SysApi', 'System/api')
fixSysPage('SysRole', 'System/role')
fixSysPage('SysDept', 'System/dept')
fixSysPage('SysMenu', 'System/menu')
