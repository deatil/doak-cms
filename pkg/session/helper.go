package session

import (
    "github.com/gofiber/fiber/v2"
)

// session id
func ID(ctx *fiber.Ctx) string {
    return Session(ctx).ID()
}

// 获取
func Get(ctx *fiber.Ctx, key string) any {
    return Session(ctx).Get(key)
}

// 设置
func Set(ctx *fiber.Ctx, key string, val any) error {
    sess := Session(ctx)
    sess.Set(key, val)

    return sess.Save()
}

// 删除
func Delete(ctx *fiber.Ctx, key string) error {
    sess := Session(ctx)
    sess.Delete(key)

    return sess.Save()
}

// 清空
func Destroy(ctx *fiber.Ctx) error {
    sess := Session(ctx)
    sess.Destroy()

    return sess.Save()
}

// Regenerate
func Regenerate(ctx *fiber.Ctx) error {
    sess := Session(ctx)
    sess.Regenerate()

    return sess.Save()
}
