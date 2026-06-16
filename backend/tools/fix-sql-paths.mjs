import { readFileSync, writeFileSync } from 'fs'

const fp = 'D:/CodeByAi/cl_system/backend/resource/sql/system_mgr.sql'
let sql = readFileSync(fp, 'utf-8')

// 1. Add is_show_mobile column to sys_role CREATE TABLE
sql = sql.replace(
  '`level` int(11) DEFAULT 1 COMMENT',
  '`is_show_mobile` tinyint(4) DEFAULT 1 COMMENT \'显示手机号 1=是 2=否\',\n  `level` int(11) DEFAULT 1 COMMENT'
)

// 2. Update sys_api paths from old format to RESTful
const pathMap = {
  '/api/system/admin/list': '/system/admin',
  '/api/system/admin/info': '/system/admin',
  '/api/system/admin/add': '/system/admin',
  '/api/system/admin/edit': '/system/admin',
  '/api/system/admin/del': '/system/admin',
  '/api/system/admin/password': '/system/admin/password',
  '/api/system/admin/status': '/system/admin/status',
  '/api/system/admin/clickout': '/system/admin/clickout',
  '/api/system/api/list': '/system/api',
  '/api/system/api/add': '/system/api',
  '/api/system/api/edit': '/system/api',
  '/api/system/api/del': '/system/api',
  '/api/system/api/dropdown': '/system/api/dropdown',
  '/api/system/api/path': '/system/api/path',
  '/api/system/role/list': '/system/role',
  '/api/system/role/info': '/system/role',
  '/api/system/role/add': '/system/role',
  '/api/system/role/edit': '/system/role',
  '/api/system/role/del': '/system/role',
  '/api/system/role/status': '/system/role/status',
  '/api/system/role/dropdown': '/system/role/dropdown',
  '/api/system/role/is_mobile': '/system/role/is_mobile',
  '/api/system/menu/list': '/system/menu',
  '/api/system/menu/info': '/system/menu',
  '/api/system/menu/add': '/system/menu',
  '/api/system/menu/edit': '/system/menu',
  '/api/system/menu/del': '/system/menu',
  '/api/system/menu/role': '/system/menu/role',
  '/api/system/menu/dropdown': '/system/menu/dropdown',
  '/api/system/dept/list': '/system/dept',
  '/api/system/dept/info': '/system/dept',
  '/api/system/dept/add': '/system/dept',
  '/api/system/dept/edit': '/system/dept',
  '/api/system/dept/del': '/system/dept',
  '/api/system/dept/dropdown': '/system/dept/dropdown',
}

for (const [oldPath, newPath] of Object.entries(pathMap)) {
  sql = sql.replaceAll(oldPath, newPath)
}

writeFileSync(fp, sql, 'utf-8')
console.log('SQL updated successfully')
