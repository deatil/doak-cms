package boot

import (
    "github.com/gofiber/fiber/v3"
    "github.com/gofiber/template/jet/v3"

    "github.com/deatil/doak-cms/app/view"
    "github.com/deatil/doak-cms/app/start"
    "github.com/deatil/doak-cms/app/router"
)

func Run() {
    // http 服务
    start.HttpServer(
        func(engine *jet.Engine) {
            // 添加方法
            view.SetViewFuncs(engine)
        },
        func(app *fiber.App) {
            // 添加路由
            router.Http(app)
        },
    )
}

