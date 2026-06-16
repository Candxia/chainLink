package utility

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

type uPassword struct{}

// Password 密码工具实例
var Password = new(uPassword)

// generateSalt 生成指定长度的随机盐值
func (u *uPassword) generateSalt(length int) string {
	bytes := make([]byte, length)
	_, _ = rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// Encrypt 使用 HMAC-SHA256 对密码进行加盐哈希加密
// 格式: salt$hash，存储时包含盐值
func (u *uPassword) Encrypt(password string) (string, error) {
	salt := u.generateSalt(16)
	mac := hmac.New(sha256.New, []byte(salt))
	mac.Write([]byte(password))
	hash := hex.EncodeToString(mac.Sum(nil))
	return salt + "$" + hash, nil
}

// Equal 比对明文密码与已存储的加盐密文是否一致
func (u *uPassword) Equal(password, stored string) bool {
	parts := strings.SplitN(stored, "$", 2)
	if len(parts) != 2 {
		return false
	}
	salt := parts[0]
	expected := parts[1]
	mac := hmac.New(sha256.New, []byte(salt))
	mac.Write([]byte(password))
	hash := hex.EncodeToString(mac.Sum(nil))
	return hash == expected
}
