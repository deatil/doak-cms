package utils

import (
    "crypto/md5"
    "encoding/hex"
)

// MD5 哈希值
func MD5(data string) string {
    sum := md5.Sum([]byte(data))
    return hex.EncodeToString(sum[:])
}
