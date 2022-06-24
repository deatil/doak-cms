package admin

import (
    "time"

    "github.com/spf13/cast"
    "github.com/dchest/captcha"
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/db"
    "github.com/deatil/doak-cms/pkg/auth"
    "github.com/deatil/doak-cms/pkg/http"
    "github.com/deatil/doak-cms/pkg/config"
    "github.com/deatil/doak-cms/pkg/cookie"
    "github.com/deatil/doak-cms/pkg/session"
    "github.com/deatil/doak-cms/pkg/validate"

    "github.com/deatil/doak-cms/app/model"
    "github.com/deatil/doak-cms/app/response"
)

/**
 * 账号
 *
 * @create 2022-6-19
 * @author deatil
 */
type Auth struct{
    Base
}

// 验证码
func (this *Auth) Captcha(ctx *fiber.Ctx) error {
    captchaId := captcha.NewLen(4)

    // 保存验证码ID
    if err := session.Set(ctx, "captchaid", captchaId); err != nil {
        ctx.SendStatus(404)
        return nil
    }

    if !captcha.Reload(captchaId) {
        ctx.SendStatus(404)
        return nil
    }

    ctx.Set("Content-Type", "image/png")
    err := captcha.WriteImage(ctx.Response().BodyWriter(), captchaId, 300, 58)
    return err
}

// 登录
func (this *Auth) Login(ctx *fiber.Ctx) error {
    userid := session.Get(ctx, "userid")
    if userid != nil {
        return ctx.Redirect("/sys")
    }

    return ctx.Render(this.Theme("auth/login"), nil)
}

// 验证登录
func (this *Auth) LoginCheck(ctx *fiber.Ctx) error {
    username := cast.ToString(ctx.FormValue("username"))
    password := cast.ToString(ctx.FormValue("password"))
    captchaData := cast.ToString(ctx.FormValue("captcha"))
    rememberme := cast.ToString(ctx.FormValue("rememberme"))

    // 验证
    errs := validate.Validate(
        map[string]any{
            "username": username,
            "password": password,
            "captcha": captchaData,
        },
        map[string]string{
            "username": "required|minLen:3",
            "password": "required|minLen:6",
            "captcha": "required|len:4",
        },
        map[string]string{
            "username.required": "账号不能为空",
            "username.minLen": "账号不能少于3位",
            "password.required": "密码不能为空",
            "password.minLen": "密码不能少于6位",
            "captcha.required": "验证码不能为空",
            "captcha.len": "验证码长度错误",
        },
    )

    if (errs != nil) {
        return http.Error(ctx, errs.One())
    }

    // 验证码ID
    captchaid := session.Get(ctx, "captchaid")

    // 验证码验证
    if ok := captcha.VerifyString(captchaid.(string), captchaData); !ok {
        return http.Error(ctx, "验证码错误")
    }

    // 账号信息
    var user model.User
    _, err2 := db.Engine().Where("username = ?", username).Get(&user)
    if err2 != nil {
        return http.Error(ctx, "账号不存在或者被禁用")
    }

    // 密码验证
    if !auth.Check(password, user.Password) {
        return http.Error(ctx, "账号或者密码错误")
    }

    // 存储登录信息
    if err := session.Set(ctx, "userid", user.Id); err != nil {
        return http.Error(ctx, "登录失败")
    }

    if rememberme == "1" {
        cookieCfg := config.Section("cookie")

        cookieKey := cookieCfg.Key("key").MustString("doak")
        cookieExp := cookieCfg.Key("exp").MustDuration()
        cookiePath := cookieCfg.Key("path").MustString("/")
        cookie.Set(ctx, cookieKey, cast.ToString(user.Id), cookiePath, time.Now().Add(cookieExp))
    }

    return http.Success(ctx, "登录成功", "")
}

// 退出
func (this *Auth) Logout(ctx *fiber.Ctx) error {
    // 未登录
    userid := session.Get(ctx, "userid")
    if userid == nil {
        return response.AdminErrorRender(ctx, "请先登录")
    }

    // 删除登录信息
    if err := session.Delete(ctx, "userid"); err != nil {
        return response.AdminErrorRender(ctx, "退出登录失败")
    }

    // 删除 cookie 信息
    cookieKey := config.Section("cookie").Key("key").MustString("doak")
    cookie.Delete(ctx, cookieKey)

    return ctx.Redirect("/sys/login")
}
