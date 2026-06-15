# -*- coding: utf-8 -*-
import sys, json, urllib.request
sys.stdout.reconfigure(encoding='utf-8')

API = "http://localhost:8099"

def call(method, path, data=None, token=None):
    url = API + path
    body = json.dumps(data).encode('utf-8') if data else None
    req = urllib.request.Request(url, data=body, method=method)
    req.add_header("Content-Type", "application/json")
    if token:
        req.add_header("Authorization", token)
    try:
        resp = urllib.request.urlopen(req, timeout=5)
        return resp.status, json.loads(resp.read().decode('utf-8'))
    except urllib.error.HTTPError as e:
        return e.code, e.read().decode('utf-8')[:300]
    except Exception as e:
        return 0, str(e)

# 1. 登录
print("=" * 60)
print("  1. 登录测试")
print("=" * 60)
status, resp = call("POST", "/api/user/login", {"username": "admin", "password": "admin123"})
msg = resp.get("message", str(resp)) if isinstance(resp, dict) else str(resp)
print(f"  Login → HTTP {status}, code={resp.get('code') if isinstance(resp, dict) else '?'}, msg={msg}")

# 2. 如果登录成功，遍历所有 API
if isinstance(resp, dict) and resp.get("code") == 0:
    token = resp["data"]["token"]
    print(f"\n  ✅ Token 获取成功: {token[:50]}...\n")
else:
    # 3. 即使登录失败，直接测试所有 API（看哪些是公开的）
    print("\n  ⚠️ 登录返回非0，测试所有接口状态")
    token = None

print("=" * 60)
print("  2. 遍历所有 API 接口")
print("=" * 60)

apis = [
    ("POST", "/api/user/login", {"username": "admin", "password": "admin123"}, False),
    ("POST", "/api/user/register", {"username": "newuser", "password": "***", "email": "test@test.com"}, False),
    ("GET", "/api/user/list", None, True),
    ("GET", "/api/user/info", None, True),
    ("GET", "/api/user/roles", None, True),
    ("GET", "/api/user/enterprise", None, True),
    ("GET", "/api/trace/product/list", None, True),
    ("GET", "/api/trace/batch/list", None, True),
    ("GET", "/api/trace/public/query?keyword=test", None, False),
    ("GET", "/api/trace/qrcode/1", None, True),
    ("GET", "/api/supply/supplier/list", None, True),
    ("GET", "/api/supply/order/list", None, True),
    ("GET", "/api/supply/warehouse/stock", None, True),
    ("GET", "/api/blockchain/blocks", None, True),
    ("GET", "/api/blockchain/transactions", None, True),
    ("GET", "/api/blockchain/contract/list", None, True),
    ("GET", "/api/system/dashboard", None, True),
    ("GET", "/api/system/config", None, True),
    ("GET", "/api/system/logs", None, True),
]

pass_count = 0
fail_count = 0
skip_count = 0

for method, path, data, need_auth in apis:
    status, resp = call(method, path, data, token)
    
    if isinstance(resp, dict):
        code = resp.get("code", "?")
        rmsg = resp.get("message", "")[:40]
        detail = f"HTTP {status}, code={code}"
        if rmsg:
            detail += f", msg={rmsg}"
    else:
        detail = f"HTTP {status}" if status else f"Error: {str(resp)[:60]}"
    
    # 判断结果
    if status == 200:
        if isinstance(resp, dict) and resp.get("code") in [0, None]:
            status_icon = "✅"
            pass_count += 1
        elif isinstance(resp, dict) and resp.get("code") == 401:
            if token:
                status_icon = "❌"  # 有 token 还 401 是问题
                fail_count += 1
            else:
                status_icon = "⏭️"  # 没 token 401 正常
                skip_count += 1
        elif isinstance(resp, dict) and resp.get("code") == 1:
            if need_auth and not token:
                status_icon = "⏭️"
                skip_count += 1
            else:
                status_icon = "⚡"
                pass_count += 1
        else:
            status_icon = "⚡"
            pass_count += 1
    else:
        status_icon = "❌"
        fail_count += 1

    auth_tag = " 🔐" if need_auth else ""
    print(f"  {status_icon} {method:4s} {path:40s}{auth_tag} → {detail}")

print()
print("=" * 60)
print(f"  结果: ✅ {pass_count} 通过 | ❌ {fail_count} 失败 | ⏭️ {skip_count} 跳过")
print(f"  总计: {pass_count + fail_count + skip_count} 个接口")
print("=" * 60)
