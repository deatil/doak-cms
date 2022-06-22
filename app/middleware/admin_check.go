package middleware

import (
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/app/auth"
    "github.com/deatil/doak-cms/app/response"
)

// admin 检测
func NewAdminCheck() fiber.Handler {
    return func(ctx *fiber.Ctx) error {
        if !auth.IsAdmin(ctx) {
            return response.AdminErrorRender(ctx, "你没有权限访问该页面")
        }

        return ctx.Next()
    }
}
