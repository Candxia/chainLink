# -*- coding: utf-8 -*-
"""初始化远程 MySQL 数据库 - 修正版"""
import sys, pymysql
sys.stdout.reconfigure(encoding='utf-8')

DB_HOST = "47.109.157.55"
DB_PORT = 10001
DB_USER = "root"
DB_PASS = "Root@123456"
DB_NAME = "cl_system"

# 1. 先创建数据库
print("🔌 连接远程 MySQL...")
conn = pymysql.connect(host=DB_HOST, port=DB_PORT, user=DB_USER, password=DB_PASS)
cursor = conn.cursor()
cursor.execute(f"CREATE DATABASE IF NOT EXISTS `{DB_NAME}` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci")
print(f"✅ 数据库 `{DB_NAME}` 已就绪")
cursor.close()
conn.close()

# 2. 读 SQL 并分段执行
print("📖 读取 schema.sql...")
with open("D:\\CodeByAi\\cl_system\\backend\\manifest\\sql\\schema.sql", "r", encoding="utf-8") as f:
    raw = f.read()

# 去掉注释行和空行，将整个文件内容按分号拆分
lines = raw.split('\n')
cleaned = []
for line in lines:
    stripped = line.strip()
    if stripped.startswith('--') or stripped.startswith('#') or stripped == '':
        # 保留 USE, DROP, CREATE 等关键字前的注释可以保留
        if any(kw in stripped.upper() for kw in ['DATABASE', 'TABLE', 'INSERT', 'INDEX']):
            continue
        continue
    cleaned.append(line)

sql_text = '\n'.join(cleaned)

# 3. 连接到数据库执行
print("📦 执行 DDL...")
conn2 = pymysql.connect(host=DB_HOST, port=DB_PORT, user=DB_USER, password=DB_PASS, database=DB_NAME)
cursor2 = conn2.cursor()

# 按分号拆分语句
statements = sql_text.split(';')
success = 0
failed = 0

for stmt in statements:
    stmt = stmt.strip()
    if not stmt:
        continue
    try:
        cursor2.execute(stmt)
        conn2.commit()
        first_line = stmt.split('\n')[0].strip()[:70]
        print(f"  ✅ {first_line}")
        success += 1
    except Exception as e:
        err = str(e)
        if "already exists" in err.lower() or "duplicate" in err.lower():
            success += 1
            continue
        print(f"  ⚠️  {err[:80]}")
        failed += 1

cursor2.close()
conn2.close()
print(f"\n✅ DDL 完成: {success} 条成功, {failed} 条失败")

# 4. 验证
print("\n🔍 验证表结构:")
conn3 = pymysql.connect(host=DB_HOST, port=DB_PORT, user=DB_USER, password=DB_PASS, database=DB_NAME)
cursor3 = conn3.cursor()
cursor3.execute("SHOW TABLES")
tables = cursor3.fetchall()
for t in tables:
    cursor3.execute(f"SELECT COUNT(*) FROM `{t[0]}`")
    cnt = cursor3.fetchone()[0]
    print(f"  📋 `{t[0]:20s}` → {cnt} 条记录")
cursor3.close()
conn3.close()
