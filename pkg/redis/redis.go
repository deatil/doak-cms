package redis

import (
    "sync"

    "github.com/gofiber/storage/redis"

    "github.com/deatil/doak-cms/pkg/config"
)

var store *redis.Storage
var once sync.Once

// Redis
// Redis().Get(key string) ([]byte, error)
// Redis().Set(key string, val []byte, exp time.Duration)
// Redis().Delete(key string) error
// Redis().Reset() error
// Redis().Close()
func Redis() *redis.Storage {
    once.Do(func() {
        store = Manager("redis")
    })

    return sess
}

// redis 存储
func Manager(typ string) *redis.Storage {
    // 配置
    cfg := config.Section(typ)

    storage := redis.New(redis.Config{
        Host:      cfg.Key("host").MustString("127.0.0.1"),
        Port:      cfg.Key("port").MustInt(6379),
        Username:  cfg.Key("username").MustString(""),
        Password:  cfg.Key("password").MustString(""),
        Database:  cfg.Key("db").MustInt(5),
        Reset:     cfg.Key("reset").MustBool(false),
        // Example: redis://<user>:<pass>@localhost:6379/<db>
        URL:       cfg.Key("url").MustString(""),
        TLSConfig: nil,
    })

    return storage
}
