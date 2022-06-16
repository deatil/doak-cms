package router

import (
    "github.com/smallnest/rpcx/server"

    "github.com/deatil/doak-cms/app/service"
)

// Rpc
func Rpc(s *server.Server) {
    // 注册服务
    s.RegisterName("Arith", new(service.Arith), "")
}


