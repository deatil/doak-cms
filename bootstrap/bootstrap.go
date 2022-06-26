package bootstrap

import (
    "time"
    golog "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/utils"
    "github.com/gofiber/fiber/v2/middleware/csrf"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/recover"
    "github.com/gofiber/fiber/v2/middleware/encryptcookie"
    "github.com/gofiber/template/jet"

    "github.com/smallnest/rpcx/server"
    "github.com/rpcxio/rpcx-etcd/serverplugin"

    "github.com/deatil/doak-cms/pkg/log"
    "github.com/deatil/doak-cms/pkg/view"
    "github.com/deatil/doak-cms/pkg/config"
)

// http 服务
func HttpServer(jetFunc func(*jet.Engine), appFunc func(*fiber.App)) {
    cfg := config.Section("app")

    debug := cfg.Key("debug").MustBool(false)

    // 设置模板驱动
    jetEngine := view.JetEngine("view")
    jetFunc(jetEngine)

    // 启动 app
    app := fiber.New(fiber.Config{
        Views: jetEngine,
        // 默认的错误处理程序
        // fiber.DefaultErrorHandler
        // 覆盖默认的错误处理程序
        ErrorHandler: func(ctx *fiber.Ctx, err error) error {
            // 状态代码默认为500
            code := fiber.StatusInternalServerError

            // 如果是fiber.*Error，则检索自定义状态代码。
            if e, ok := err.(*fiber.Error); ok {
                code = e.Code
            }

            // 默认错误信息
            errorMsg := cfg.Key("error-msg").String()

            // 调试的时候
            if debug {
                errorMsg = err.Error()
            }

            // 页面地址
            errorHtml := cfg.Key("error-html").String()

            // 发送自定义错误页面
            err = ctx.Status(code).Render(errorHtml, fiber.Map{
                "code": code,
                "message": errorMsg,
            })
            if err != nil {
                // 失败处理
               return ctx.Status(fiber.StatusInternalServerError).
                        SendString("Internal Server Error")
            }

             // 从处理程序返回
             return nil
        },
    })

    // 中间件
    if debug {
        location := config.Section("time").Key("loc").MustString("")

        // 日志
        app.Use(logger.New(logger.Config{
            Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
            TimeFormat: "2006-01-02 15:04:05",
            TimeZone:   location,
        }))
    }

    // 拦截 panic 报错
    app.Use(recover.New(recover.Config{
        EnableStackTrace: debug,
    }))

    // Cookie 加密
    // 生成密钥 encryptcookie.GenerateKey()
    app.Use(encryptcookie.New(encryptcookie.Config{
        Key: config.Section("cookie").Key("encrypt-key").MustString(""),
    }))

    // 跨站请求伪造 CSRF
    app.Use(csrf.New(csrf.Config{
        KeyLookup:      "cookie:doakcsrf",
        CookieName:     "doakcsrf",
        CookieSameSite: "Strict",
        Expiration:     1 * time.Hour,
        KeyGenerator:   utils.UUID,
    }))

    // 添加路由
    appFunc(app)

    // 运行端口
    address := cfg.Key("address").String()
    golog.Fatal(app.Listen(address))
}

// rpc 部分
// 当前项目使用的是 json-rpc 模式
func RpcServer(rpcFunc func(*server.Server)) {
    s := server.NewServer()

    // cors
    opt := server.AllowAllCORSOptions()
    s.SetCORS(opt)

    // 添加插件
    addRpcRegistryPlugin(s)

    // s.Register(new(Arith), "")
    // s.RegisterName("Arith", new(Arith), "")

    rpcFunc(s)

    // 地址
    addr := config.Section("rpc").Key("address").String()
    err := s.Serve("tcp", addr)
    if err != nil {
        golog.Println(err)

        // 记录日志
        log.Log().Error(err.Error())
    }
}

func addRpcRegistryPlugin(s *server.Server) {
    cfg := config.Section("rpc")

    addr := cfg.Key("address").String()
    etcdAddrs := cfg.Key("etcd-address").Strings(",")
    basePath := cfg.Key("base-path").String()
    updateInterval := cfg.Key("update-interval").MustDuration()

    r := &serverplugin.EtcdV3RegisterPlugin{
        ServiceAddress: "tcp@" + addr,
        EtcdServers:    etcdAddrs,
        BasePath:       basePath,
        UpdateInterval: updateInterval * time.Minute,
    }

    err := r.Start()
    if err != nil {
        golog.Fatal(err)
    }

    s.Plugins.Add(r)
}
