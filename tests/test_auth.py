# -*- coding: utf-8 -*-
import sys
sys.stdout.reconfigure(encoding='utf-8')

from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from webdriver_manager.chrome import ChromeDriverManager
import json, urllib.request, time

BASE = "http://localhost:8080"
API = "http://localhost:8099"

options = webdriver.ChromeOptions()
options.add_argument('--headless')
options.add_argument('--no-sandbox')
options.add_argument('--disable-dev-shm-usage')
options.add_argument('--window-size=1280,800')

driver = webdriver.Chrome(
    service=Service(ChromeDriverManager().install()),
    options=options
)

try:
    # 1. 打开登录页
    driver.get(BASE + "/")
    time.sleep(1.5)
    print("✅ 页面加载完成")

    # 2. 填登录表单
    inputs = driver.find_elements(By.TAG_NAME, "input")
    btns = driver.find_elements(By.TAG_NAME, "button")

    print("输入框:", len(inputs))
    print("按钮:", len(btns))

    # 尝试各种选择器
    username_input = None
    pwd_input = None
    login_btn = None

    for inp in inputs:
        ph = (inp.get_attribute("placeholder") or "").lower()
        if "用户" in ph or "账号" in ph or "admin" in ph:
            username_input = inp
        if "密码" in ph:
            pwd_input = inp

    for btn in btns:
        txt = btn.text.strip()
        if "登录" in txt or "登 录" in txt:
            login_btn = btn

    if username_input and pwd_input and login_btn:
        username_input.clear()
        username_input.send_keys("admin")
        pwd_input.clear()
        pwd_input.send_keys("admin123")
        login_btn.click()
        time.sleep(2)
        print("✅ 已尝试登录")

        # 尝试从 localStorage 获取 token
        token = driver.execute_script("return localStorage.getItem('token')")
        if token:
            print("✅ 获取到 token:", token[:40] + "...")
        else:
            print("⚠️ localStorage 无 token，尝试检查页面变化")
    else:
        print("⚠️ 未能定位登录元素, 尝试直接注入 token")

    # 3. 无论是否成功登录，看看当前页面
    html = driver.find_element(By.TAG_NAME, "body").text[:300]
    print("\n页面内容:", html)

    # 截图
    driver.save_screenshot("D:\\CodeByAi\\cl_system\\tests\\login_result.png")
    print("✅ 截图已保存")

finally:
    driver.quit()

# 4. 测试所有需要认证的 API（无 token）
print("\n" + "=" * 60)
print("  [带认证] API 测试")
print("=" * 60)

api_list = [
    ("用户列表", "GET", "/api/user/list"),
    ("用户信息", "GET", "/api/user/info"),
    ("角色列表", "GET", "/api/user/roles"),
    ("企业列表", "GET", "/api/user/enterprise"),
    ("产品列表", "GET", "/api/trace/product/list"),
    ("批次列表", "GET", "/api/trace/batch/list"),
    ("溯源记录", "GET", "/api/trace/record/list/1"),
    ("追溯链", "GET", "/api/trace/chain/1"),
    ("供应商列表", "GET", "/api/supply/supplier/list"),
    ("订单列表", "GET", "/api/supply/order/list"),
    ("库存查询", "GET", "/api/supply/warehouse/stock"),
    ("区块列表", "GET", "/api/blockchain/blocks"),
    ("交易列表", "GET", "/api/blockchain/transactions"),
    ("合约列表", "GET", "/api/blockchain/contract/list"),
    ("仪表盘", "GET", "/api/system/dashboard"),
    ("系统配置", "GET", "/api/system/config"),
    ("操作日志", "GET", "/api/system/logs"),
    ("通知列表", "GET", "/api/system/notifications"),
]

for name, method, path in api_list:
    req = urllib.request.Request(API + path, method=method)
    try:
        resp = urllib.request.urlopen(req, timeout=5)
        body = resp.read().decode('utf-8')
        data = json.loads(body)
        code = data.get('code', '?')
        msg = data.get('message', '')[:40]
        status = data.get('status', data.get('data', None))
        print(f"  {'✅' if code in [0, 200, None] else '⚠️'} {name:10s} ({path:35s}) → code={code}, msg={msg}")
    except urllib.error.HTTPError as e:
        print(f"  ❌ {name:10s} ({path:35s}) → HTTP {e.code}: {str(e)[:60]}")
    except Exception as e:
        print(f"  ❌ {name:10s} ({path:35s}) → {str(e)[:60]}")

print("\n✅ 测试完成")
