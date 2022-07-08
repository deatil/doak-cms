package middleware

import (
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/app/data"
    "github.com/deatil/doak-cms/app/response"
)

// 网站是否开启检测
func NewSiteopenCheck() fiber.Handler {
    return func(ctx *fiber.Ctx) error {
        if data.GetSetting("website_status") != "1" {
            return response.CmsErrorRender(ctx, "网站当前关闭调整中...")
        }

        return ctx.Next()
    }
}
