package admin

import (
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/http"
)

// 文章
type Art struct{
    Base
}

// 列表
func (this *Art) Index(ctx *fiber.Ctx) error {
    data := "data"

    return this.View(ctx, "art/index", fiber.Map{
        "Title": data,
    })
}

// 添加
func (this *Art) Add(ctx *fiber.Ctx) error {
    data := "data"

    return this.View(ctx, "art/add", fiber.Map{
        "Title": data,
    })
}

// 添加保存
func (this *Art) AddSave(ctx *fiber.Ctx) error {
    return http.Success(ctx, "添加成功", "")
}

// 编辑
func (this *Art) Edit(ctx *fiber.Ctx) error {
    data := "data"

    return this.View(ctx, "art/edit", fiber.Map{
        "Title": data,
    })
}

// 编辑保存
func (this *Art) EditSave(ctx *fiber.Ctx) error {
    return http.Success(ctx, "保存成功", "")
}

// 删除
func (this *Art) Delete(ctx *fiber.Ctx) error {
    return http.Success(ctx, "删除成功", "")
}
