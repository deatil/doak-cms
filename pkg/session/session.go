package session

import (
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/session"
    "github.com/gofiber/storage/redis"

    "github.com/deatil/doak-cms/pkg/config"
)

var store *session.Store

// Session
func Session(ctx *fiber.Ctx) *session.Session {
    if store == nil {
        store = Store("session")
    }

    // 从 storage 获取 session
    sess, err := store.Get(ctx)
    if err != nil {
        panic(err)
    }

    return sess
}

// 存储
func Store(typ string) *session.Store {
    // 配置
    cfg := config.Section(typ + ".redis")

    storage := redis.New(redis.Config{
        Host:      cfg.Key("host").MustString("127.0.0.1"),
        Port:      cfg.Key("port").MustInt(6379),
        Username:  cfg.Key("username").MustString(""),
        Password:  cfg.Key("password").MustString(""),
        URL:       "",
        Database:  cfg.Key("db").MustInt(5),
        Reset:     cfg.Key("reset").MustBool(false),
        TLSConfig: nil,
    })

    // 配置
    sessCfg := config.Section(typ)

    store := session.New(session.Config{
        Storage:    storage,
        KeyLookup:  sessCfg.Key("name").MustString(""),
        Expiration: sessCfg.Key("exp").MustDuration() * time.Hour,
    })

    return store
}
