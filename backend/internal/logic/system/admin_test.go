package system

import (
	"testing"

	"cl_system/internal/consts"
	mdlSys "cl_system/internal/model/system"
	"cl_system/utility"
)

// TestPasswordEncrypt 测试密码加密与验证逻辑
func TestPasswordEncrypt(t *testing.T) {
	pwd := "test123456"
	encrypted, err := utility.Password.Encrypt(pwd)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}
	if encrypted == "" {
		t.Fatal("encrypted password should not be empty")
	}
	// 验证格式: salt$hash (32 hex chars + $ + 64 hex chars = 97 chars)
	if len(encrypted) != 97 {
		t.Fatalf("unexpected encrypted length: got %d, want 97", len(encrypted))
	}
	// 验证正确密码比对
	if !utility.Password.Equal(pwd, encrypted) {
		t.Fatal("Equal should return true for correct password")
	}
	// 验证错误密码比对
	if utility.Password.Equal("wrong_password", encrypted) {
		t.Fatal("Equal should return false for wrong password")
	}
	// 验证非法格式
	if utility.Password.Equal(pwd, "invalid_format_no_dollar") {
		t.Fatal("Equal should return false for invalid format")
	}
}

// TestAdminAdd 验证管理员添加所需的结构体定义
func TestAdminAdd(t *testing.T) {
	admin := mdlSys.SysAdminAdd{
		Username:   "testuser",
		Nickname:   "测试用户",
		Password:   "password123",
		RePassword: "password123",
		RoleId:     2,
		DeptId:     2,
	}
	if admin.Username != "testuser" {
		t.Fatal("SysAdminAdd.Username field mismatch")
	}
	if admin.Password != admin.RePassword {
		t.Fatal("Password and RePassword should match")
	}
	if admin.RoleId <= 0 {
		t.Fatal("RoleId must be positive")
	}
	if admin.DeptId <= 0 {
		t.Fatal("DeptId must be positive")
	}
	t.Log("AdminAdd struct fields verified successfully")
}

// TestAdminPassword 测试密码修改逻辑中的加密验证（不依赖 DB）
func TestAdminPassword(t *testing.T) {
	// 验证空密码加密不报错
	emptyEncrypted, err := utility.Password.Encrypt("")
	if err != nil {
		t.Fatalf("Encrypt empty password failed: %v", err)
	}
	if len(emptyEncrypted) != 97 {
		t.Fatalf("empty password encrypted length: got %d, want 97", len(emptyEncrypted))
	}
	// 验证短密码加密
	shortEncrypted, _ := utility.Password.Encrypt("ab")
	if !utility.Password.Equal("ab", shortEncrypted) {
		t.Fatal("Equal should match short password")
	}
	// 验证长密码加密
	longPwd := "thisIsAVeryLongPassword123!!!@#$%^&*()_+"
	longEncrypted, _ := utility.Password.Encrypt(longPwd)
	if !utility.Password.Equal(longPwd, longEncrypted) {
		t.Fatal("Equal should match long password with special chars")
	}
	// 验证同一密码每次加密结果不同（盐值不同）
	enc1, _ := utility.Password.Encrypt("s3cur3P@ss")
	enc2, _ := utility.Password.Encrypt("s3cur3P@ss")
	if enc1 == enc2 {
		t.Fatal("same password should produce different encrypted results due to random salt")
	}
}

// TestAdminList 测试列表查询条件构造
func TestAdminList(t *testing.T) {
	// 1. 默认条件应排除超管
	search := mdlSys.SysAdminSearch{}
	cond := search.Condition()
	if cond == nil {
		t.Fatal("Condition() should not return nil")
	}
	if cond["id >"] != consts.SuperAdminId {
		t.Fatalf("default condition should exclude super admin (id > %d), got %v", consts.SuperAdminId, cond["id >"])
	}

	// 2. 按用户名过滤
	search2 := mdlSys.SysAdminSearch{Username: "admin"}
	cond2 := search2.Condition()
	if cond2["username"] != "admin" {
		t.Fatal("condition should include username filter when set")
	}

	// 3. 按部门过滤
	search3 := mdlSys.SysAdminSearch{DeptId: 3}
	cond3 := search3.Condition()
	if cond3["dept_id"] != int64(3) {
		t.Fatal("condition should include dept_id filter when > 0")
	}

	// 4. 按状态过滤
	search4 := mdlSys.SysAdminSearch{Status: 2}
	cond4 := search4.Condition()
	if cond4["status"] != 2 {
		t.Fatal("condition should include status filter when > 0")
	}

	// 5. 按在线状态过滤
	search5 := mdlSys.SysAdminSearch{Online: 1}
	cond5 := search5.Condition()
	if cond5["online"] != 1 {
		t.Fatal("condition should include online filter when > 0")
	}

	// 6. 组合条件过滤
	search6 := mdlSys.SysAdminSearch{
		Username: "zhangsan",
		Status:   1,
		DeptId:   2,
	}
	cond6 := search6.Condition()
	if cond6["username"] != "zhangsan" {
		t.Fatal("combined condition should include username")
	}
	if cond6["status"] != 1 {
		t.Fatal("combined condition should include status")
	}
	if cond6["dept_id"] != int64(2) {
		t.Fatal("combined condition should include dept_id")
	}
	if cond6["id >"] != consts.SuperAdminId {
		t.Fatal("combined condition should still exclude super admin")
	}
}
