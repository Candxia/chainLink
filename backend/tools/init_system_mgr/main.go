//go:build ignore

package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:Root@123456@tcp(47.109.157.55:10001)/cl_system?charset=utf8mb4&parseTime=true&loc=Local&multiStatements=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("数据库连接失败: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	db.SetConnMaxLifetime(30 * time.Second)
	db.SetMaxOpenConns(1)

	if err := db.Ping(); err != nil {
		fmt.Printf("数据库 ping 失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ 数据库连接成功")

	sqlBytes, err := os.ReadFile("resource/sql/system_mgr.sql")
	if err != nil {
		fmt.Printf("读取SQL文件失败: %v\n", err)
		os.Exit(1)
	}

	// 先清空旧数据（忽略表不存在的错误）
	tables := []string{"sys_dept", "sys_role_menu", "sys_api", "sys_menu", "sys_role", "sys_admin"}
	for _, t := range tables {
		db.Exec("DROP TABLE IF EXISTS " + t)
	}

	// 一次性执行整个文件（multiStatements=true 支持多条语句）
	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		fmt.Printf("SQL执行失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ SQL 全部执行成功")

	// 验证
	fmt.Println("\n📊 验证数据表:")
	for _, t := range tables {
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM " + t).Scan(&count)
		if err != nil {
			fmt.Printf("   ❌ %s: %v\n", t, err)
		} else {
			fmt.Printf("   ✅ %s: %d 条记录\n", t, count)
		}
	}
}
