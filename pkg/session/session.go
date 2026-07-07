package session

import (
    "time"

    "github.com/gofiber/fiber/v3"
    "github.com/gofiber/fiber/v3/middleware/session"
    "github.com/gofiber/fiber/v3/extractors"
    "github.com/gofiber/utils/v2"
    "github.com/gofiber/storage/redis/v3"

    "github.com/deatil/doak-cms/pkg/log"
    "github.com/deatil/doak-cms/pkg/config"
)

// Session
func Session(ctx fiber.Ctx) *session.Middleware {
    sess := session.FromContext(ctx)

    return sess
}

// 存储
func Store(typ string) fiber.Handler {
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
        Storage:      storage,
        Extractor:    extractors.FromCookie(sessCfg.Key("name").MustString("")),
        KeyGenerator: utils.SecureToken,

        IdleTimeout:     30 * time.Minute,  // Inactivity timeout
        // AbsoluteTimeout:   24 * time.Hour,    // Maximum session duration    
        AbsoluteTimeout: sessCfg.Key("exp").MustDuration() * time.Hour,

        ErrorHandler: func(c fiber.Ctx, err error) {
            log.Log().Error("Session error: " + err.Error())
        },
    })

    return store
}
