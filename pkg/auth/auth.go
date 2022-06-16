package auth

import (
    "golang.org/x/crypto/bcrypt"
)

// 对密码进行加密
func Hash(password string) string {
    bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes)
}

// 对比明文密码和数据库的哈希值
func Check(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
