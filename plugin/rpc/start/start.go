package start

import (
    "time"
    golog "log"

    "github.com/smallnest/rpcx/server"
    "github.com/rpcxio/rpcx-etcd/serverplugin"

    "github.com/deatil/doak-cms/pkg/log"
    "github.com/deatil/doak-cms/pkg/config"
)

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
