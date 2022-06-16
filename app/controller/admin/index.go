package admin

import (
    "github.com/gofiber/fiber/v2"
)

// 首页
type Index struct{
    Base
}

func (this *Index) Index(ctx *fiber.Ctx) error {

    return this.View(ctx, "index/index", fiber.Map{
        "Title": "Title",
    })
}
