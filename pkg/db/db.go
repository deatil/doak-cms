package db

import (
    "os"
    "time"

    _ "github.com/go-sql-driver/mysql"
    "xorm.io/xorm"
    "xorm.io/xorm/log"
    "xorm.io/xorm/names"

    "github.com/deatil/doak-cms/pkg/config"
)

var enginer *xorm.Engine

// 数据库引擎
func Engine() *xorm.Engine {
    if enginer == nil {
        enginer = Manager("db")
    }

    return enginer
}

// 管理
func Manager(typ string) *xorm.Engine {
    // 配置
    cfg := config.Section(typ)

    engine, err := xorm.NewEngine(
        cfg.Key("type").MustString(""),
        cfg.Key("dsn").MustString(""),
    )

    if err != nil {
        panic("连接数据库失败")
    }

    if engine.Ping() != nil {
        panic("连接数据库失败")
    }

    // 日志
    f, err := os.Create("runtime/log/db-sql.log")
    if err != nil {
        panic(err.Error())
    }

    engine.SetLogger(log.NewSimpleLogger(f))

    // engine.SetMapper(names.SnakeMapper{})

    // 前缀映射
    tbPrefix := cfg.Key("prefix").MustString("")
    tbMapper := names.NewPrefixMapper(names.SnakeMapper{}, tbPrefix)
    engine.SetTableMapper(tbMapper)

    // 连接池
    engine.SetMaxIdleConns(cfg.Key("max-idle-conns").MustInt(10))
    engine.SetMaxOpenConns(cfg.Key("max-open-conns").MustInt(128))
    engine.SetConnMaxLifetime(cfg.Key("conn-max-lifetime").MustDuration())

    // 时区
    loc := cfg.Key("loc").MustString("Asia/Shanghai")
    engine.TZLocation, _ = time.LoadLocation(loc)

    return engine
}
