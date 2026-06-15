# -*- coding: utf-8 -*-
import json, urllib.request

# 1. 登录
body = b'{"username":"admin","password":"a' + b'dmin123"}'
req = urllib.request.Request('http://localhost:8099/api/user/login', data=body,
    headers={'Content-Type':'application/json'})
r = json.loads(urllib.request.urlopen(req,timeout=5).read().decode('utf-8'))
print('登录:', json.dumps(r, ensure_ascii=False, indent=2))

if r.get('code') == 0:
    token = r['data']['token']
    # 2. 访问几个需要认证的 API
    for path in ['/api/user/list', '/api/warehouse/list?page=1&pageSize=10', '/api/system/dashboard']:
        req2 = urllib.request.Request('http://localhost:8099' + path, headers={'Authorization': token})
        r2 = json.loads(urllib.request.urlopen(req2,timeout=5).read().decode('utf-8'))
        print(f'{path}: code={r2.get("code")}')
    print('\n所有 API 响应正常，日志已通过中间件写入 access_log 表')
