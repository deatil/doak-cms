package router

import (
    "github.com/smallnest/rpcx/server"

    "github.com/deatil/doak-cms/app/service"
)

// Rpc 注册
func Rpc(s *server.Server) {
    // 分类
    s.RegisterName("Cate", new(service.Cate), "")

    // 文章
    s.RegisterName("Art", new(service.Art), "")

    // 单页
    s.RegisterName("Page", new(service.Page), "")

    // 附件
    s.RegisterName("Attach", new(service.Attach), "")

    // 配置
    s.RegisterName("Config", new(service.Config), "")

    // 标签
    s.RegisterName("Tag", new(service.Tag), "")

    // 账号
    s.RegisterName("User", new(service.User), "")
}


