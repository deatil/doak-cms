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

/**
 * 我的信息
 *
 * @create 2022-6-19
 * @author deatil
 */
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
    avatar := cast.ToString(ctx.FormValue("avatar"))
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
        return http.Error(ctx, errs.One())
    }

    // 需要更新的数据
    updateData := map[string]any{
        "nickname": nickname,
        "avatar": avatar,
        "sign": sign,
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
        return http.Error(ctx, "保存失败")
    }

    return http.Success(ctx, "保存成功", "")
}

// 密码
func (this *Profile) Password(ctx *fiber.Ctx) error {
    return this.View(ctx, "profile/password", fiber.Map{})
}

// 保存密码
func (this *Profile) PasswordSave(ctx *fiber.Ctx) error {
    oldpassword := cast.ToString(ctx.FormValue("oldpassword"))
    newpassword := cast.ToString(ctx.FormValue("newpassword"))
    newpasswordCheck := cast.ToString(ctx.FormValue("newpassword_check"))

    // 验证
    errs := validate.Validate(
        map[string]any{
            "oldpassword": oldpassword,
            "newpassword": newpassword,
            "newpassword_check": newpasswordCheck,
        },
        map[string]string{
            "oldpassword": "required",
            "newpassword": "required",
            "newpassword_check": "required",
        },
        map[string]string{
            "oldpassword.required": "旧密码不能为空",
            "newpassword.required": "新密码不能为空",
            "newpassword_check.required": "确认新密码不能为空",
        },
    )

    if (errs != nil) {
        return http.Error(ctx, errs.One())
    }

    if (newpassword != newpasswordCheck) {
        return http.Error(ctx, "两次新密码不一致")
    }

    if (oldpassword == newpassword) {
        return http.Error(ctx, "新密码不能与旧密码相同")
    }

    // 登录账号信息
    userInfo := appAuth.GetUserInfo(ctx)

    // 密码验证
    if !auth.Check(oldpassword, userInfo.Password) {
        return http.Error(ctx, "旧密码错误")
    }

    // 需要更新的数据
    updateData := map[string]any{
        "password": auth.Hash(newpassword),
    }

    // 当前账号
    id := userInfo.Id

    var user model.User
    _, err := db.
        Engine().
        Table(user).
        Where("id = ?", id).
        Update(updateData)
    if err != nil {
        return http.Error(ctx, "密码更改失败")
    }

    return http.Success(ctx, "密码更改成功", "")
}
