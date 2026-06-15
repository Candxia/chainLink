# -*- coding: utf-8 -*-
import sys, json, urllib.request, time
sys.stdout.reconfigure(encoding='utf-8')

API = "http://localhost:8099"

def api_call(method, path, data=None, token=None):
    url = API + path
    body = None
    if data:
        body = json.dumps(data).encode('utf-8')
    req = urllib.request.Request(url, data=body, method=method)
    req.add_header("Content-Type", "application/json")
    if token:
        req.add_header("Authorization", token)
    try:
        resp = urllib.request.urlopen(req, timeout=5)
        return resp.status, json.loads(resp.read().decode('utf-8'))
    except urllib.error.HTTPError as e:
        body = e.read().decode('utf-8') if e.fp else ""
        return e.code, body[:200]
    except Exception as e:
        return 0, str(e)

# 1. 尝试各种方式登录
print("=" * 60)
print("  1. 尝试登录获取 Token")
print("=" * 60)

login_attempts = [
    {"username": "admin", "password": "admin123"},
    {"username": "admin", "password": "***"},
    {"username": "admin", "password": "Root@123456"},
    {"phone": "admin", "password": "admin123"},
    {"mobile": "admin", "password": "admin123"},
    {"email": "admin@test.com", "password": "admin123"},
    {"username": "admin", "password": "admin123", "captcha": ""},
]

for i, creds in enumerate(login_attempts):
    status, resp = api_call("POST", "/api/user/login", creds)
    if isinstance(resp, dict):
        code = resp.get("code", "?")
        msg = resp.get("message", "")[:40]
        token = resp.get("data", {}).get("token", resp.get("token", ""))
        if token:
            print(f"  ✅ [尝试{i+1}] {creds} → 登录成功!")
            print(f"     Token: {token[:50]}...")
            TOKEN = token
            break
        else:
            print(f"  ⚠️  [尝试{i+1}] {json.dumps(creds):40s} → code={code}, msg={msg}")
    else:
        print(f"  ❌ [尝试{i+1}] {json.dumps(creds):40s} → {str(resp)[:60]}")
else:
    print("\n  ❌ 所有登录方式均失败，后端当前为模拟数据层，无真实用户数据")
    print("  下一步建议: 查看后端 login controller 逻辑")
    TOKEN = None

# 2. 用 Token 测试所有接口
if TOKEN:
    print("\n" + "=" * 60)
    print("  2. 带 Token 测试认证接口")
    print("=" * 60)

    api_list = [
        ("用户列表", "GET", "/api/user/list"),
        ("用户信息", "GET", "/api/user/info"),
        ("角色列表", "GET", "/api/user/roles"),
        ("企业列表", "GET", "/api/user/enterprise"),
        ("产品列表", "GET", "/api/trace/product/list"),
        ("批次列表", "GET", "/api/trace/batch/list"),
        ("供应商列表", "GET", "/api/supply/supplier/list"),
        ("订单列表", "GET", "/api/supply/order/list"),
        ("库存查询", "GET", "/api/supply/warehouse/stock"),
        ("区块列表", "GET", "/api/blockchain/blocks"),
        ("交易列表", "GET", "/api/blockchain/transactions"),
        ("仪表盘", "GET", "/api/system/dashboard"),
        ("系统配置", "GET", "/api/system/config"),
        ("操作日志", "GET", "/api/system/logs"),
    ]

    pass_count = 0
    fail_count = 0
    for name, method, path in api_list:
        status, resp = api_call(method, path, token=TOKEN)
        if isinstance(resp, dict) and resp.get("code") in [0, 200, None]:
            print(f"  ✅ {name:10s} → HTTP {status}, code={resp.get('code')}")
            pass_count += 1
        elif status == 200:
            print(f"  ⚠️ {name:10s} → HTTP 200, 响应: {str(resp)[:80]}")
            pass_count += 1
        else:
            print(f"  ❌ {name:10s} → {str(resp)[:80]}")
            fail_count += 1

    print(f"\n  结果: ✅ {pass_count} 通过, ❌ {fail_count} 失败")
else:
    # 3. 直接看 login 控制器逻辑
    print("\n" + "=" * 60)
    print("  2. 检查登录接口逻辑")
    print("=" * 60)

    # 检查各种可能的登录端点
    extra_endpoints = [
        ("POST", "/api/user/login", {"username": "admin"}),
        ("POST", "/api/user/login", {}),
        ("POST", "/api/user/login", {"phone": "13800138000", "password": "***"}),
        ("GET", "/api/user/login", None),
    ]
    for method, path, data in extra_endpoints:
        status, resp = api_call(method, path, data)
        detail = str(resp)[:80] if not isinstance(resp, dict) else f"code={resp.get('code')}, msg={resp.get('message','')[:30]}"
        print(f"  {method} {path} {data or ''} → {detail}")

print("\n✅ 测试完成")
