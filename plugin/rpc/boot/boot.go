package boot

import (
    "github.com/smallnest/rpcx/server"

    "github.com/deatil/doak-cms/plugin/rpc/start"
    "github.com/deatil/doak-cms/plugin/rpc/router"
)

func Run() {
    // rpc 部分
    start.RpcServer(func(s *server.Server) {
        router.Rpc(s)
    })
}

