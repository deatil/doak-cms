package redis

import (
    "sync"

    "github.com/deatil/doak-cms/pkg/config"
)

var storage *Redis
var once sync.Once

// Redis
// Storage().Get(key string) ([]byte, error)
// Storage().Set(key string, val []byte, exp time.Duration)
// Storage().Delete(key string) error
// Storage().Reset() error
// Storage().Close()
func Storage() *Redis {
    once.Do(func() {
        storage = Manager("redis")
    })

    return storage
}

// redis 存储
func Manager(typ string) *Redis {
    // 配置
    cfg := config.Section(typ)

    storage := New(Config{
        Host:      cfg.Key("host").MustString("127.0.0.1"),
        Port:      cfg.Key("port").MustInt(6379),
        Username:  cfg.Key("username").MustString(""),
        Password:  cfg.Key("password").MustString(""),
        Database:  cfg.Key("db").MustInt(5),
        Reset:     cfg.Key("reset").MustBool(false),
        KeyPrefix: cfg.Key("key-prefix").MustString(""),
        // Example: redis://<user>:<pass>@localhost:6379/<db>
        URL:       cfg.Key("url").MustString(""),
        TLSConfig: nil,
    })

    return storage
}
