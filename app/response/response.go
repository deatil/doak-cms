package response

import (
    "github.com/spf13/cast"
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/config"
)

// 管理端错误显示
func AdminErrorRender(ctx *fiber.Ctx, msg string, extra ...any) error {
    url := ""
    if len(extra) > 0 {
        url = cast.ToString(extra[0])
    }

    wait := 5
    if len(extra) > 1 {
        wait = cast.ToInt(extra[1])
    }

    ctx.Render("admin/error/index", fiber.Map{
        "Msg": msg,
        "Url": url,
        "Wait": wait,
    })

    return nil
}

// ===================

// cms 主题
func CmsTheme(tpl string) string {
    // 配置
    cfg := config.Section("view")

    theme := cfg.Key("theme").MustString("doak")

    view := "cms/" + theme + "/" + tpl

    return view
}

// cms 错误显示
func CmsErrorRender(ctx *fiber.Ctx, msg string, extra ...any) error {
    url := ""
    if len(extra) > 0 {
        url = cast.ToString(extra[0])
    }

    wait := 5
    if len(extra) > 1 {
        wait = cast.ToInt(extra[1])
    }

    ctx.Render(CmsTheme("error"), fiber.Map{
        "Msg": msg,
        "Url": url,
        "Wait": wait,
    })

    return nil
}
