package user

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
	"github.com/gogf/gf/v2/frame/g"
	"cl_system/internal/model"
	"cl_system/internal/service"
)

type sUser struct{}
func init() {
	service.RegisterUser(&sUser{})
}

func (s *sUser) Login(ctx context.Context, req model.UserLoginInput) (*model.UserLoginOutput, error) {
	user, err := g.DB().Model("sys_user").Where("username", req.Username).One()
	if err != nil {
		return nil, fmt.Errorf("用户不存在")
	}
	password := encryptPassword(req.Password)
	if user["password"].String() != password {
		return nil, fmt.Errorf("密码错误")
	}
	if user["status"].Int() != 1 {
		return nil, fmt.Errorf("账号已被禁用")
	}
	token, err := generateToken(user["id"].Uint(), user["username"].String(), user["role"].String())
	if err != nil {
		return nil, fmt.Errorf("生成令牌失败")
	}
	return &model.UserLoginOutput{
		Token:    token,
		UserId:   user["id"].Uint(),
		UserName: user["username"].String(),
		Role:     user["role"].String(),
	}, nil
}

func (s *sUser) Register(ctx context.Context, req model.UserRegisterInput) error {
	count, err := g.DB().Model("sys_user").Where("username", req.Username).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("用户名已存在")
	}
	_, err = g.DB().Model("sys_user").Insert(g.Map{
		"username": req.Username,
		"password": encryptPassword(req.Password),
		"email":    req.Email,
		"phone":    req.Phone,
		"role":     "user",
		"status":   1,
	})
	return err
}

func (s *sUser) GetUserInfo(ctx context.Context, userId uint) (*model.UserInfo, error) {
	user, err := g.DB().Model("sys_user").Where("id", userId).One()
	if err != nil {
		return nil, err
	}
	info := &model.UserInfo{
		Id:       user["id"].Uint(),
		Username: user["username"].String(),
		Email:    user["email"].String(),
		Phone:    user["phone"].String(),
		Role:     user["role"].String(),
		Status:   user["status"].Int(),
	}
	return info, nil
}

func (s *sUser) GetUserList(ctx context.Context, page, pageSize int) ([]*model.UserInfo, int, error) {
	total, _ := g.DB().Model("sys_user").Count()
	users, err := g.DB().Model("sys_user").Page(page, pageSize).Order("id desc").All()
	if err != nil {
		return nil, 0, err
	}
	var list []*model.UserInfo
	for _, u := range users {
		list = append(list, &model.UserInfo{
			Id:       u["id"].Uint(),
			Username: u["username"].String(),
			Email:    u["email"].String(),
			Phone:    u["phone"].String(),
			Role:     u["role"].String(),
			Status:   u["status"].Int(),
		})
	}
	return list, total, nil
}

func (s *sUser) UpdateUser(ctx context.Context, userId uint, data map[string]interface{}) error {
	if pw, ok := data["password"]; ok {
		data["password"] = encryptPassword(pw.(string))
	}
	_, err := g.DB().Model("sys_user").Where("id", userId).Update(data)
	return err
}

func (s *sUser) DeleteUser(ctx context.Context, id uint) error {
	_, err := g.DB().Model("sys_user").Where("id", id).Delete()
	return err
}

func (s *sUser) GetRoleList(ctx context.Context) ([]*model.RoleInfo, error) {
	roles, err := g.DB().Model("sys_role").All()
	if err != nil {
		return nil, err
	}
	var list []*model.RoleInfo
	for _, r := range roles {
		list = append(list, &model.RoleInfo{
			Id:          r["id"].Uint(),
			Name:        r["name"].String(),
			Code:        r["code"].String(),
			Permissions: r["permissions"].String(),
		})
	}
	return list, nil
}

func (s *sUser) CreateRole(ctx context.Context, role model.RoleInfo) error {
	_, err := g.DB().Model("sys_role").Insert(g.Map{
		"name":        role.Name,
		"code":        role.Code,
		"permissions": role.Permissions,
	})
	return err
}

func (s *sUser) UpdateRole(ctx context.Context, role model.RoleInfo) error {
	_, err := g.DB().Model("sys_role").Where("id", role.Id).Update(g.Map{
		"name":        role.Name,
		"code":        role.Code,
		"permissions": role.Permissions,
	})
	return err
}

func (s *sUser) DeleteRole(ctx context.Context, id uint) error {
	_, err := g.DB().Model("sys_role").Where("id", id).Delete()
	return err
}

func (s *sUser) GetEnterpriseList(ctx context.Context, page, pageSize int) ([]*model.EnterpriseInfo, int, error) {
	total, _ := g.DB().Model("enterprise").Count()
	ents, err := g.DB().Model("enterprise").Page(page, pageSize).Order("id desc").All()
	if err != nil {
		return nil, 0, err
	}
	var list []*model.EnterpriseInfo
	for _, e := range ents {
		list = append(list, &model.EnterpriseInfo{
			Id:      e["id"].Uint(),
			Name:    e["name"].String(),
			Code:    e["code"].String(),
			Contact: e["contact"].String(),
			Phone:   e["phone"].String(),
			Address: e["address"].String(),
			Status:  e["status"].Int(),
			UserId:  e["user_id"].Uint(),
		})
	}
	return list, total, nil
}

func (s *sUser) CreateEnterprise(ctx context.Context, ent model.EnterpriseInfo) error {
	_, err := g.DB().Model("enterprise").Insert(g.Map{
		"name":    ent.Name,
		"code":    ent.Code,
		"contact": ent.Contact,
		"phone":   ent.Phone,
		"address": ent.Address,
		"status":  1,
	})
	return err
}

func (s *sUser) UpdateEnterprise(ctx context.Context, ent model.EnterpriseInfo) error {
	_, err := g.DB().Model("enterprise").Where("id", ent.Id).Update(g.Map{
		"name":    ent.Name,
		"code":    ent.Code,
		"contact": ent.Contact,
		"phone":   ent.Phone,
		"address": ent.Address,
		"status":  ent.Status,
	})
	return err
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(password + "cl_system_salt"))
	return hex.EncodeToString(h.Sum(nil))
}

func generateToken(userId uint, username, role string) (string, error) {
	claims := map[string]interface{}{
		"userId":   userId,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	payload, _ := json.Marshal(claims)
	encoded := base64.RawURLEncoding.EncodeToString(payload)
	// Create HMAC signature
	key := []byte("cl_system_jwt_secret_key_2024")
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(encoded))
	sig := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
	return base64.RawURLEncoding.EncodeToString([]byte("{\"typ\":\"JWT\",\"alg\":\"HS256\"}")) + "." + encoded + "." + sig, nil
}
