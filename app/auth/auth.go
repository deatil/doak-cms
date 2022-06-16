package auth

import (
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/app/model"
)

// 获取账号ID
func GetUserId(ctx *fiber.Ctx) int64 {
    userid := ctx.Locals("userid")

    return userid.(int64)
}

// 获取账号信息
func GetUserInfo(ctx *fiber.Ctx) model.User {
    info := ctx.Locals("user")

    return info.(model.User)
}
