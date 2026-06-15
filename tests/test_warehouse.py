# -*- coding: utf-8 -*-
"""仓库管理模块完整 API 验证"""
import sys, json, urllib.request
sys.stdout.reconfigure(encoding='utf-8')

API = "http://localhost:8099"
TOKEN = None

def api(method, path, data=None):
    url = API + path
    body = json.dumps(data).encode('utf-8') if data else None
    req = urllib.request.Request(url, data=body, method=method)
    req.add_header("Content-Type", "application/json")
    if TOKEN:
        req.add_header("Authorization", TOKEN)
    try:
        resp = urllib.request.urlopen(req, timeout=5)
        return json.loads(resp.read().decode('utf-8'))
    except Exception as e:
        return {"error": str(e)}

# 登录
r = api("POST", "/api/user/login", {"username": "admin", "password": "admin123"})
TOKEN = r.get("data", {}).get("token", "")
print(f"登录: code={r.get('code')}, message={r.get('message')}")

ok, fail = 0, 0
def test(name, result):
    global ok, fail
    if result.get("code") in [0, None] and "error" not in result:
        ok += 1
        print(f"  ✅ {name}")
    else:
        fail += 1
        err = result.get("message", result.get("error", str(result)))[:50]
        print(f"  ❌ {name} → {err}")

# === 仓库信息 ===
print("\n【仓库信息】")
test("仓库列表", api("GET", "/api/warehouse/list?page=1&pageSize=10"))
test("添加仓库", api("POST", "/api/warehouse", {"code":"WH-TEST","name":"测试仓库","address":"测试地址","company":"测试公司"}))
test("区域列表", api("GET", "/api/warehouse/area/list"))
test("新增区域", api("POST", "/api/warehouse/area", {"warehouseId":1,"name":"E区","description":"测试区域"}))
test("货架列表", api("GET", "/api/warehouse/shelf/list"))
test("新增货架", api("POST", "/api/warehouse/shelf", {"areaId":1,"name":"A-4","description":"测试货架"}))
test("盘存列表", api("GET", "/api/warehouse/stocktake/list"))
test("新建盘存", api("POST", "/api/warehouse/stocktake", {"warehouseId":1,"type":1,"operator":"测试员"}))

# === 库存管理 ===
print("\n【库存管理】")
test("整车库存", api("GET", "/api/inventory/car/list?page=1&pageSize=5"))
test("原材料库存", api("GET", "/api/inventory/raw-material/list"))
test("原材料入库", api("POST", "/api/inventory/raw-material/in", {"warehouseId":1,"materialType":"steel","quantity":500}))
test("原材料出库", api("POST", "/api/inventory/raw-material/out", {"warehouseId":1,"materialType":"steel","quantity":200}))
test("危固废库存", api("GET", "/api/inventory/waste/list"))
test("危废入库", api("POST", "/api/inventory/waste/in", {"warehouseId":1,"wasteType":"hazardous","quantity":100}))
test("溯源件列表", api("GET", "/api/inventory/part/traceable/list"))
test("非溯源件列表", api("GET", "/api/inventory/part/untraceable/list"))
test("非溯源件入库", api("POST", "/api/inventory/part/untraceable/in", {"warehouseId":1,"partName":"测试零件","partType":"测试","quantity":50}))

# === 操作记录 ===
print("\n【操作记录】")
test("入库记录", api("GET", "/api/records/inbound"))
test("出库记录", api("GET", "/api/records/outbound"))
test("调拨记录", api("GET", "/api/records/transfer"))

print(f"\n{'='*40}")
print(f"结果: ✅ {ok} 通过 | ❌ {fail} 失败 | 总计 {ok+fail}")
