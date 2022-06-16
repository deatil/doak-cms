package cms

import (
    "github.com/gofiber/fiber/v2"
)

// 首页
type Index struct{
    Base
}

func (this *Index) Index(ctx *fiber.Ctx) error {
    data := "data"

    return ctx.Render(this.Theme("index"), fiber.Map{
        "Title": data,
    })
}
