package utils

import (
    "io"
    "os"
    "fmt"
    "bufio"
    "errors"
    "net/http"
    "crypto/md5"
    "crypto/rand"
    "encoding/hex"

    "github.com/deatil/doak-cms/pkg/time"
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

    data := string(b) + time.Now().DateTimeString()

    return MD5(data)
}

// 格式化数据大小
func FormatSize(size int64) string {
    units := []string{" B", " KB", " MB", " GB", " TB", " PB"}

    s := float64(size)

    i := 0
    for ; s >= 1024 && i < 5; i++ {
        s /= 1024
    }

    return fmt.Sprintf("%.2f%s", s, units[i])
}

// 文件是否存在
func FileExists(path string) bool {
    _, err := os.Stat(path)

    return err == nil || os.IsExist(err)
}

// 文件删除
func FileDelete(path string) error {
    return os.Remove(path)
}

// 创建文件夹
func MkDir(directory string) error {
    return os.MkdirAll(directory, 0666)
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

// 列出文件
func ListFiles(directory string) []string {
    if !FileExists(directory) {
        return []string{}
    }

    fs, err := os.ReadDir(directory)
    if err != nil {
        return []string{}
    }

    sz := len(fs)
    if sz == 0 {
        return []string{}
    }

    ret := make([]string, 0, sz)
    for i := 0; i < sz; i++ {
        if !fs[i].IsDir() {
            ret = append(ret, fs[i].Name())
        }
    }

    return ret
}

// 列出文件
func ListEmbedFiles(f http.File) []string {
    fs, err := f.Readdir(-1)
    if err != nil {
        return []string{}
    }

    sz := len(fs)
    if sz == 0 {
        return []string{}
    }

    ret := make([]string, 0, sz)
    for i, n := 0, sz; i < n; i++ {
        if !fs[i].IsDir() {
            ret = append(ret, fs[i].Name())
        }
    }

    return ret
}
