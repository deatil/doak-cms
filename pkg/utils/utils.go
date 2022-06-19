package utils

import (
    "io"
    "crypto/md5"
    "crypto/rand"
    "encoding/hex"
)

// MD5 哈希值
func MD5(data string) string {
    sum := md5.Sum([]byte(data))
    return hex.EncodeToString(sum[:])
}

// 32位 唯一ID
func Uniqueid() string {
    b := make([]byte, 48)

    if _, err := io.ReadFull(rand.Reader, b); err != nil {
        return ""
    }

    return MD5(string(b))
}
