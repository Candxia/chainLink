package v1

import (
	"context"
	"strings"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
)

// TestAdminAddReqValidation 测试 AdminAddReq 必填字段校验规则
func TestAdminAddReqValidation(t *testing.T) {
	ctx := context.TODO()
	// 空请求 - 所有必填字段缺失
	req := AdminAddReq{}
	err := g.Validator().Data(req).Run(ctx)
	if err == nil {
		t.Fatal("empty AdminAddReq should trigger validation errors")
	}
	errStr := err.Error()
	// 验证关键必填字段错误（GoFrame 使用字段名如 Username, Nickname 等）
	requiredFields := []string{"Username", "Nickname", "Password", "RePassword"}
	for _, field := range requiredFields {
		if !strings.Contains(errStr, field) {
			t.Fatalf("validation error should mention missing field: %s, got: %s", field, errStr)
		}
	}
}

// TestAdminAddReqValidationSuccess 测试 AdminAddReq 合法数据通过校验
func TestAdminAddReqValidationSuccess(t *testing.T) {
	ctx := context.TODO()
	req := AdminAddReq{
		Username:   "testuser",
		Nickname:   "测试用户",
		Password:   "pass123456",
		RePassword: "pass123456",
		RoleId:     2,
		DeptId:     3,
	}
	err := g.Validator().Data(req).Run(ctx)
	if err != nil {
		t.Fatalf("valid AdminAddReq should pass validation: %v", err)
	}
}

// TestAdminAddReqRePasswordMismatch 测试确认密码不匹配
func TestAdminAddReqRePasswordMismatch(t *testing.T) {
	ctx := context.TODO()
	req := AdminAddReq{
		Username:   "testuser",
		Nickname:   "测试用户",
		Password:   "pass123456",
		RePassword: "different_pass",
		RoleId:     2,
		DeptId:     3,
	}
	err := g.Validator().Data(req).Run(ctx)
	if err == nil {
		t.Fatal("re_password mismatch should trigger validation error")
	}
	if !strings.Contains(err.Error(), "RePassword") {
		t.Fatalf("validation should mention RePassword mismatch, got: %s", err.Error())
	}
}

// TestAdminAddReqUsernameLength 测试用户名长度校验
func TestAdminAddReqUsernameLength(t *testing.T) {
	ctx := context.TODO()
	// 用户名过短 (<3)
	req := AdminAddReq{
		Username:   "ab",
		Nickname:   "测试用户",
		Password:   "pass123456",
		RePassword: "pass123456",
		RoleId:     2,
		DeptId:     3,
	}
	err := g.Validator().Data(req).Run(ctx)
	if err == nil {
		t.Fatal("short username (length 2) should trigger validation error")
	}
	if !strings.Contains(err.Error(), "length") {
		t.Fatalf("validation should mention length constraint, got: %s", err.Error())
	}
}

// TestAdminAddReqPasswordLength 测试密码长度校验
func TestAdminAddReqPasswordLength(t *testing.T) {
	ctx := context.TODO()
	// 密码过短 (<6)
	req := AdminAddReq{
		Username:   "testuser",
		Nickname:   "测试用户",
		Password:   "12345",
		RePassword: "12345",
		RoleId:     2,
		DeptId:     3,
	}
	err := g.Validator().Data(req).Run(ctx)
	if err == nil {
		t.Fatal("short password (length 5) should trigger validation error")
	}
	if !strings.Contains(err.Error(), "length") {
		t.Fatalf("validation should mention length constraint, got: %s", err.Error())
	}
}

// TestAdminStatusReqValidation 测试 AdminStatusReq status in 规则
func TestAdminStatusReqValidation(t *testing.T) {
	ctx := context.TODO()
	// 非法 status 值
	req := AdminStatusReq{
		Id:     1,
		Status: 99,
	}
	err := g.Validator().Data(req).Run(ctx)
	if err == nil {
		t.Fatal("invalid status value 99 should trigger validation error")
	}

	// 合法 status 值
	req2 := AdminStatusReq{
		Id:     1,
		Status: 2,
	}
	err2 := g.Validator().Data(req2).Run(ctx)
	if err2 != nil {
		t.Fatalf("valid status value 2 should pass validation: %v", err2)
	}
}


