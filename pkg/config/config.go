package config

import (
    "sync"
    "gopkg.in/ini.v1"
)

var cfg *ini.File
var once sync.Once

// 操作分区
func Section(name string) *ini.Section {
    once.Do(func() {
        // 配置文件
        var file = "./config/config.ini"

        cfg = Manager(file)
    })

    return cfg.Section(name)
}

// 管理
func Manager(file string) *ini.File {
    conf, _ := ini.LoadSources(ini.LoadOptions{
        Insensitive: true,
        AllowBooleanKeys: true,
    }, file)

    return conf
}

