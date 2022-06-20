package utils

import (
    "io"
    "os"
    "fmt"
    "bufio"
    "errors"
    "crypto/md5"
    "crypto/rand"
    "encoding/hex"
)

type File interface {
    io.Reader
}

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

// 大文件 Md5
func FileMD5(filename string) (string, error) {
    if info, err := os.Stat(filename); err != nil {
        return "", err
    } else if info.IsDir() {
        return "", errors.New("不是文件无法计算")
    }

    openfile, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer openfile.Close()

    return FileMD5WithStream(openfile)
}

// 大文件 Md5
func FileMD5WithStream(openfile File) (string, error) {
    const bufferSize = 65536

    hash := md5.New()
    for buf, reader := make([]byte, bufferSize), bufio.NewReader(openfile); ; {
        n, err := reader.Read(buf)
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err
        }

        hash.Write(buf[:n])
    }

    checksum := fmt.Sprintf("%x", hash.Sum(nil))
    return checksum, nil
}
