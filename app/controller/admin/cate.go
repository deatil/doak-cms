package admin

import (
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/http"
)

// 分类
type Cate struct{
    Base
}

// 列表
func (this *Cate) Index(ctx *fiber.Ctx) error {
    data := "data"

    return this.View(ctx, "cate/index", fiber.Map{
        "Title": data,
    })
}

// 添加
func (this *Cate) Add(ctx *fiber.Ctx) error {
    data := "data"

    return ctx.Render("cate/add", fiber.Map{
        "Title": data,
    })
}

// 添加保存
func (this *Cate) AddSave(ctx *fiber.Ctx) error {
    return http.Success(ctx, "添加成功", "")
}

// 编辑
func (this *Cate) Edit(ctx *fiber.Ctx) error {
    data := "data"

    return this.View(ctx, "cate/edit", fiber.Map{
        "Title": data,
    })
}

// 编辑保存
func (this *Cate) EditSave(ctx *fiber.Ctx) error {
    return http.Success(ctx, "保存成功", "")
}

// 删除
func (this *Cate) Delete(ctx *fiber.Ctx) error {
    return http.Success(ctx, "删除成功", "")
}
