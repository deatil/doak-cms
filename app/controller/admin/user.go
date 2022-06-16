package admin

import (
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/http"
)

// 账号
type User struct{
    Base
}

// 登录
func (this *User) Index(ctx *fiber.Ctx) error {
    data := "data"

    return this.View(ctx, "user/index", fiber.Map{
        "Title": data,
    })
}

// 添加
func (this *User) Add(ctx *fiber.Ctx) error {
    data := "data"

    // affected, err := db.Engine().Insert(user)

    return this.View(ctx, "user/add", fiber.Map{
        "Title": data,
    })
}

// 添加保存
func (this *User) AddSave(ctx *fiber.Ctx) error {
    return http.Success(ctx, "添加成功", "")
}

// 编辑
func (this *User) Edit(ctx *fiber.Ctx) error {
    data := "data"

    return this.View(ctx, "user/edit", fiber.Map{
        "Title": data,
    })
}

// 编辑保存
func (this *User) EditSave(ctx *fiber.Ctx) error {
    return http.Success(ctx, "保存成功", "")
}

// 删除
func (this *User) Delete(ctx *fiber.Ctx) error {
    return http.Success(ctx, "删除成功", "")
}
