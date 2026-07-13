package config

import (
    "sync"
    "gopkg.in/ini.v1"

    "github.com/deatil/doak-cms/config"
)

var cfg *ini.File
var once sync.Once

// 使用打包方式
var IsUseEmbed bool = true

// 操作分区
func Section(name string) *ini.Section {
    once.Do(func() {
        if IsUseEmbed {
            data, _ := config.ReadConfig("config.ini")
            cfg = Manager(data)
        } else {
            // 配置文件
            var file = "./config/config.ini"
            cfg = Manager(file)
        }
    })

    return cfg.Section(name)
}

// 管理
func Manager(source any) *ini.File {
    conf, _ := ini.LoadSources(ini.LoadOptions{
        Insensitive:      true,
        AllowBooleanKeys: true,
    }, source)

    return conf
}
