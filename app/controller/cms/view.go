package cms

import (
    "github.com/gofiber/fiber/v2"
)

// 详情
type View struct{
    Base
}

func (this *View) Index(ctx *fiber.Ctx) error {
    data := "data"

    return ctx.Render(this.Theme("view"), fiber.Map{
        "Data": data,
    })
}
