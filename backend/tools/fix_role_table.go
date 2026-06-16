//go:build ignore

package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Root@123456@tcp(47.109.157.55:10001)/cl_system?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		fmt.Printf("connect fail: %v\n", err)
		return
	}
	defer db.Close()
	db.SetConnMaxLifetime(30 * time.Second)
	db.Ping()

	columns := []string{"menu_all", "is_show_mobile"}
	for _, col := range columns {
		var count int
		db.QueryRow("SELECT COUNT(*) FROM information_schema.COLUMNS WHERE TABLE_SCHEMA='cl_system' AND TABLE_NAME='sys_role' AND COLUMN_NAME=?", col).Scan(&count)
		if count == 0 {
			after := "is_del"
			if col == "menu_all" {
				after = "is_show_mobile"
			}
			sql := fmt.Sprintf("ALTER TABLE sys_role ADD COLUMN %s tinyint(4) DEFAULT 2 COMMENT '' AFTER %s", col, after)
			_, err := db.Exec(sql)
			if err != nil {
				fmt.Printf("fail %s: %v\n", col, err)
			} else {
				fmt.Printf("added %s\n", col)
			}
		} else {
			fmt.Printf("%s exists\n", col)
		}
	}
	fmt.Println("Done")
}
