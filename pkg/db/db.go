package db

import (
    "os"
    "time"
    "sync"

    _ "github.com/go-sql-driver/mysql"
    "xorm.io/xorm"
    "xorm.io/xorm/log"
    "xorm.io/xorm/names"

    "github.com/deatil/doak-cms/pkg/config"
)

var enginer *xorm.Engine
var once sync.Once

// 数据库引擎
// dbsession := Engine().NewSession()
func Engine() *xorm.Engine {
    once.Do(func() {
        enginer = Manager("db")
    })

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

    // 在控制台打印出生成的SQL语句
    showSql := cfg.Key("show-sql").MustBool(false)
    if showSql {
        engine.ShowSQL(true)
    }

    // 记录等级
    logLevel := cfg.Key("log-level").MustString("info")
    switch logLevel {
        case "debug":
            engine.Logger().SetLevel(log.LOG_DEBUG)
        case "info":
            engine.Logger().SetLevel(log.LOG_INFO)
        case "warning":
            engine.Logger().SetLevel(log.LOG_WARNING)
        case "err":
            engine.Logger().SetLevel(log.LOG_ERR)
    }

    // 日志文件
    logFile := cfg.Key("log-file").MustString("")

    // 日志
    f, err := os.Create(logFile)
    if err != nil {
        panic(err.Error())
    }

    logger := log.NewSimpleLogger(f)

    // 记录 sql 日志
    recordSql := cfg.Key("record-sql").MustBool(false)
    if recordSql {
        logger.ShowSQL(true)
    }

    engine.SetLogger(logger)

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

    if engine.Ping() != nil {
        panic("连接数据库失败")
    }

    return engine
}
