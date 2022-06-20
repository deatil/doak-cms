package cms

import (
    "github.com/gofiber/fiber/v2"
)

// 单页
type Page struct{
    Base
}

func (this *Page) Index(ctx *fiber.Ctx) error {
    data := "data"

    tpl := "about-page"

    return ctx.Render(this.Theme(tpl), fiber.Map{
        "Data": data,
    })
}
