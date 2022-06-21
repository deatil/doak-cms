package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/jet"
    "github.com/smallnest/rpcx/server"

    "github.com/deatil/doak-cms/app/view"
    "github.com/deatil/doak-cms/app/router"
    "github.com/deatil/doak-cms/bootstrap"
)

func main() {
    // http 服务
    go bootstrap.HttpServer(
        func(engine *jet.Engine) {
            // 添加方法
            view.SetViewFuncs(engine)
        },
        func(app *fiber.App) {
            // 添加路由
            router.Http(app)
        },
    )

    // rpc 部分
    go bootstrap.RpcServer(func(s *server.Server) {
        router.Rpc(s)
    })

    // 阻塞退出
    select {}
}

