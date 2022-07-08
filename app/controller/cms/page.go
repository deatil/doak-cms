package cms

import (
    "github.com/spf13/cast"
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/db"
    "github.com/deatil/doak-cms/pkg/validate"

    "github.com/deatil/doak-cms/app/model"
    "github.com/deatil/doak-cms/app/response"
)

/**
 * 单页
 *
 * @create 2022-6-19
 * @author deatil
 */
type Page struct{
    Base
}

// 详情
func (this *Page) Index(ctx *fiber.Ctx) error {
    name := cast.ToString(ctx.Params("name"))

    // 验证
    errs := validate.Validate(
        map[string]any{
            "name": name,
        },
        map[string]string{
            "name": "required|minLen:3|isAlphaDash",
        },
        map[string]string{
            "name.required": "分类标识不能为空",
            "name.minLen": "分类标识不能少于3位",
            "name.isAlphaDash": "分类标识错误",
        },
    )

    if (errs != nil) {
        return response.CmsErrorRender(ctx, "页面不存在")
    }

    // 信息
    var data model.Page
    _, err := db.Engine().
        Where("slug = ?", name).
        Get(&data)
    if err != nil || data.Id == 0 {
        return response.CmsErrorRender(ctx, "页面不存在")
    }

    // 模板
    tpl := data.Tpl
    if tpl == "" {
        return response.CmsErrorRender(ctx, "页面不存在")
    }

    return this.View(ctx, tpl, fiber.Map{
        "data": data,
    })
}
