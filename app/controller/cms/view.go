package cms

import (
    "github.com/spf13/cast"
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/db"

    "github.com/deatil/doak-cms/app/model"
    "github.com/deatil/doak-cms/app/response"
)

// 详情
type View struct{
    Base
}

func (this *View) Index(ctx *fiber.Ctx) error {
    id := cast.ToInt64(ctx.Params("id"))
    if id == 0 {
        return response.CmsErrorRender(ctx, "数据不存在")
    }

    // 分类信息
    var data model.Art
    _, err := db.Engine().
        Where("id = ?", id).
        Get(&data)
    if err != nil {
        return response.CmsErrorRender(ctx, "数据不存在")
    }

    return ctx.Render(this.Theme("view"), fiber.Map{
        "data": data,
    })
}
