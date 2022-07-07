package cms

import (
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/app/data"
    "github.com/deatil/doak-cms/app/response"
)

/**
 * Base
 *
 * @create 2022-6-19
 * @author deatil
 */
type Base struct{}

// 主题
func (this *Base) Theme(tpl string) string {
    return response.CmsTheme(tpl)
}

// 视图
func (this *Base) View(ctx *fiber.Ctx, tpl string, info fiber.Map) error {
    // 设置数据
    info["settings"] = data.GetSettings()

    return ctx.Render(this.Theme(tpl), info)
}

// 检测网站是否关闭
func (this *Base) SiteopenCheck(ctx *fiber.Ctx) (bool, error) {
    if data.GetSetting("website_status") != "1" {
        return false, response.CmsErrorRender(ctx, "网站当前关闭调整中...")
    }

    return true, nil
}
