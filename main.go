package main

import (
    "github.com/deatil/doak-cms/bootstrap"

    // rpc_boot "github.com/deatil/doak-cms/plugin/rpc/boot"
)

func main() {
    // app 服务
    go bootstrap.Exec()

    // rpc 服务
    // go rpc_boot.Run()

    // 阻塞退出
    select {}
}


