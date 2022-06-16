package admin

import (
    "github.com/spf13/cast"
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/db"
    "github.com/deatil/doak-cms/pkg/auth"
    "github.com/deatil/doak-cms/pkg/http"
    "github.com/deatil/doak-cms/pkg/validate"

    "github.com/deatil/doak-cms/app/model"
    appAuth "github.com/deatil/doak-cms/app/auth"
)

// 我的信息
type Profile struct{
    Base
}

// 页面
func (this *Profile) Index(ctx *fiber.Ctx) error {
    return this.View(ctx, "profile/index", fiber.Map{})
}

// 保存
func (this *Profile) Save(ctx *fiber.Ctx) error {
    nickname := cast.ToString(ctx.FormValue("nickname"))
    password := cast.ToString(ctx.FormValue("password"))
    sign := cast.ToString(ctx.FormValue("sign"))

    // 验证
    errs := validate.Validate(
        map[string]any{
            "nickname": nickname,
        },
        map[string]string{
            "nickname": "required|minLen:3",
        },
        map[string]string{
            "nickname.required": "账号不能为空",
            "nickname.minLen": "账号不能少于3位",
        },
    )

    if (errs != nil) {
        return http.Error(ctx, 1, errs.One())
    }

    // 需要更新的数据
    updateData := map[string]any{
        "nickname": nickname,
        "sign": sign,
    }

    if password != "" && len(password) > 3 {
        updateData["password"] = auth.Hash(password)
    }

    // 当前账号
    id := appAuth.GetUserInfo(ctx).Id

    var user model.User
    _, err := db.
        Engine().
        Table(user).
        Where("id = ?", id).
        Update(updateData)
    if err != nil {
        return http.Error(ctx, 1, "保存失败")
    }

    return http.Success(ctx, "保存成功", "")
}
