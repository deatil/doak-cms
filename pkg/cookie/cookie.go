package cookie

import (
    "time"

    "github.com/gofiber/fiber/v2"
)

// 获取
func Get(ctx *fiber.Ctx, key string) string {
    return ctx.Cookies(key)
}

// 设置
func Set(ctx *fiber.Ctx, key string, value string, path string, exp time.Time) {
    ctx.Cookie(&fiber.Cookie{
        Name:     key,
        Value:    value,
        Path:     path,
        Expires:  exp,
        HTTPOnly: true,
        // disabled, lax, strict, none
        // fiber.CookieSameSiteDisabled
        // fiber.CookieSameSiteLaxMode
        // fiber.CookieSameSiteStrictMode
        // fiber.CookieSameSiteNoneMode
        SameSite: "lax",
    })
}

// 删除
func Delete(ctx *fiber.Ctx, key string) {
    ctx.Cookie(&fiber.Cookie{
        Name:     key,
        Expires:  time.Now().Add(-(time.Hour * 2)),
        HTTPOnly: true,
        SameSite: "lax",
    })
}

// 用[fiber.Cookie]设置
func SetWithCookie(ctx *fiber.Ctx, data fiber.Cookie) {
    /*
    type fiber.Cookie struct {
        Name     string    `json:"name"`
        Value    string    `json:"value"`
        Path     string    `json:"path"`
        Domain   string    `json:"domain"`
        Expires  time.Time `json:"expires"`
        Secure   bool      `json:"secure"`
        HTTPOnly bool      `json:"http_only"`
        SameSite string    `json:"same_site"`
    }
    */

    ctx.Cookie(&data)
}
