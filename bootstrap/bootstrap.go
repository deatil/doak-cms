package bootstrap

import (
    "time"
    golog "log"

    "github.com/gofiber/fiber/v2"
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

    // 设置模板驱动
    jetEngine := view.JetEngine("view")
    jetFunc(jetEngine)

    app := fiber.New(fiber.Config{
        Views: jetEngine,
    })

    debug := cfg.Key("debug").MustBool(false)

    // 中间件
    if debug {
        // 日志
        app.Use(logger.New(logger.Config{
            Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
            TimeFormat: "2006-01-02 15:04:05",
            TimeZone:   "Asia/Chongqing",
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

    // 添加路由
    appFunc(app)

    // 运行端口
    address := cfg.Key("address").String()
    golog.Fatal(app.Listen(address))
}

// rpc 部分
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
