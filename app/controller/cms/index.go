package cms

import (
    "github.com/gofiber/fiber/v2"
)

/**
 * 首页
 *
 * @create 2022-6-19
 * @author deatil
 */
type Index struct{
    Base
}

func (this *Index) Index(ctx *fiber.Ctx) error {
    if siteOk, siteErr := this.SiteopenCheck(ctx); !siteOk {
        return siteErr
    }

    return this.View(ctx, "index", fiber.Map{})
}
