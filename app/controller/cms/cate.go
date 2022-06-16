package cms

import (
    "github.com/gofiber/fiber/v2"
)

// 首页
type Cate struct{
    Base
}

func (this *Cate) Index(ctx *fiber.Ctx) error {
    data := "data"

    return ctx.Render(this.Theme("cate"), fiber.Map{
        "Data": data,
    })
}
