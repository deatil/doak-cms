package admin

import (
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/http"
)

// 标签
type Tag struct{
    Base
}

// 页面
func (this *Tag) Index(ctx *fiber.Ctx) error {
    data := "data"

    return this.View(ctx, "tag/index", fiber.Map{
        "Title": data,
    })
}

// 编辑
func (this *Tag) Edit(ctx *fiber.Ctx) error {
    data := "data"

    return this.View(ctx, "tag/edit", fiber.Map{
        "Title": data,
    })
}

// 编辑保存
func (this *Tag) EditSave(ctx *fiber.Ctx) error {
    return http.Success(ctx, "保存成功", "")
}

// 删除
func (this *Tag) Delete(ctx *fiber.Ctx) error {
    return http.Success(ctx, "删除成功", "")
}
