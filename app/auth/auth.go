package auth

import (
    "time"

    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/config"
    "github.com/deatil/doak-cms/pkg/cookie"

    "github.com/deatil/doak-cms/app/model"
)

// 获取账号ID
func GetUserId(ctx *fiber.Ctx) int64 {
    userid := ctx.Locals("userid")
    if userid == nil {
        return 0
    }

    return userid.(int64)
}

// 获取账号信息
func GetUserInfo(ctx *fiber.Ctx) model.User {
    info := ctx.Locals("user")
    if info == nil {
        return model.User{}
    }

    return info.(model.User)
}

// 是否为管理员
func IsAdmin(ctx *fiber.Ctx) bool {
    id := GetUserId(ctx)

    adminid := config.Section("auth").Key("adminid").MustInt64(0)
    if id != 0 && adminid == id {
        return true
    }

    return false
}

// =====================

// 设置 Cookie
func SetCookie(ctx *fiber.Ctx, data string) {
    cookieCfg := config.Section("cookie")

    cookieKey := cookieCfg.Key("key").MustString("doak")
    cookiePath := cookieCfg.Key("path").MustString("/")
    cookieExp := cookieCfg.Key("exp").MustDuration()
    cookieTime := time.Now().Add(cookieExp)

    cookie.Set(ctx, cookieKey, data, cookiePath, cookieTime)
}

// 删除 Cookie
func DeleteCookie(ctx *fiber.Ctx) {
    // 删除 cookie 信息
    cookieKey := config.Section("cookie").Key("key").MustString("doak")
    cookie.Delete(ctx, cookieKey)
}
