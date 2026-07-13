package session

import (
    "github.com/gofiber/fiber/v3"
)

// session id
func ID(ctx fiber.Ctx) string {
    return Session(ctx).ID()
}

// 获取
func Get(ctx fiber.Ctx, key string) any {
    return Session(ctx).Get(key)
}

// 设置
func Set(ctx fiber.Ctx, key string, val any) {
    sess := Session(ctx)
    sess.Set(key, val)
}

// 删除
func Delete(ctx fiber.Ctx, key string) {
    sess := Session(ctx)
    sess.Delete(key)
}

// 清空
func Destroy(ctx fiber.Ctx) error {
    sess := Session(ctx)
    return sess.Destroy()
}

// Regenerate
func Regenerate(ctx fiber.Ctx) error {
    sess := Session(ctx)
    return sess.Regenerate()
}
