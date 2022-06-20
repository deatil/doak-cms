package main

import (
    "github.com/uniplaces/carbon"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/jet"
    "github.com/smallnest/rpcx/server"

    "github.com/deatil/doak-cms/app/url"
    "github.com/deatil/doak-cms/app/router"
    "github.com/deatil/doak-cms/bootstrap"
)

func main() {
    // http 服务
    go bootstrap.HttpServer(
        func(engine *jet.Engine) {
            // 添加额外方法
            engine.AddFunc("createTime", carbon.CreateFromTimestamp)
            engine.AddFunc("cateUrl", url.CateUrl)
            engine.AddFunc("artUrl", url.ArtUrl)
            engine.AddFunc("pageUrl", url.PageUrl)
            engine.AddFunc("attachUrl", url.AttachUrl)
            engine.AddFunc("attachPath", url.AttachPath)
            engine.AddFunc("attachUrlWithId", url.AttachUrlWithId)
            engine.AddFunc("attachPathWithId", url.AttachPathWithId)
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

