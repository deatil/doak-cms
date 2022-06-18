package middleware

import (
    "github.com/spf13/cast"
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/db"
    "github.com/deatil/doak-cms/pkg/config"
    "github.com/deatil/doak-cms/pkg/cookie"
    "github.com/deatil/doak-cms/pkg/session"

    "github.com/deatil/doak-cms/app/model"
    "github.com/deatil/doak-cms/app/response"
)

// 权限验证
func NewAuth() fiber.Handler {
    return func(ctx *fiber.Ctx) error {
        userid := session.Get(ctx, "userid")
        if userid == nil {
            cookieKey := config.Section("cookie").Key("key").MustString("doak")

            userid = cookie.Get(ctx, cookieKey)
            if userid == nil {
                return response.AdminErrorRender(ctx, "请先登录")
            }
        }

        // 账号信息
        var user model.User
        _, err := db.Engine().
            Where("id = ?", cast.ToInt64(userid)).
            Get(&user)
        if err != nil {
            session.Delete(ctx, "userid")

            return response.AdminErrorRender(ctx, "账号不存在或者被禁用")
        }

        ctx.Locals("userid", userid)
        ctx.Locals("user", user)

        return ctx.Next()

        // 无数据发送[204]
        // return ctx.SendStatus(fiber.StatusNoContent)
    }
}
