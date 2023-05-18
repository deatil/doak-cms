package http

import (
    "github.com/gofiber/fiber/v2"
)

type RespData struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Data    any    `json:"data"`
}

// 输出
func Response(ctx *fiber.Ctx, code int, msg string, data any) error {
    return ctx.JSON(&RespData{
        Code:    code,
        Message: msg,
        Data:    data,
    })
}

// 正常
func Success(ctx *fiber.Ctx, msg string, data any) error {
    return Response(ctx, 0, msg, data)
}

// 错误
func Error(ctx *fiber.Ctx, msg string) error {
    return Response(ctx, 1, msg, "")
}
