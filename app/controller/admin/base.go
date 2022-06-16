package admin

import (
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/app/auth"
)

// 基类
type Base struct{}

// 主题
func (this *Base) Theme(tpl string) string {
    tpl = "admin/" + tpl

    return tpl
}

// 视图
func (this *Base) View(ctx *fiber.Ctx, tpl string, data fiber.Map) error {
    // 添加登录数据
    user := auth.GetUserInfo(ctx)
    data["user"] = user

    return ctx.Render(this.Theme(tpl), data)
}
