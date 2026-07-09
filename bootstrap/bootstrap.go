package bootstrap

import (
    "github.com/deatil/doak-cms/app/boot"

    // rpc_boot "github.com/deatil/doak-cms/plugin/rpc/boot"
)

func Exec() {
    // http 服务
    go boot.Run()

    // rpc 服务
    // go rpc_boot.Run()

    // 阻塞退出
    select {}
}

