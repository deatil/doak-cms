package response

import (
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/config"
)

// cms 主题
func CmsTheme(tpl string) string {
    // 配置
    cfg := config.Section("view")

    theme := cfg.Key("theme").MustString("doak")

    view := theme + "/" + tpl

    return view
}

// 管理端错误显示
func AdminErrorRender(ctx *fiber.Ctx, msg string) error {
    ctx.Render("admin/error/index", fiber.Map{
        "Msg": msg,
    })

    return nil
}

// cms 错误显示
func CmsErrorRender(ctx *fiber.Ctx, msg string) error {
    ctx.Render(CmsTheme("error"), fiber.Map{
        "Msg": msg,
    })

    return nil
}
