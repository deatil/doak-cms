package view

import (
    "github.com/uniplaces/carbon"
    "github.com/gofiber/template/jet"

    "github.com/deatil/doak-cms/pkg/config"
)

func JetEngine(typ string) *jet.Engine {
    // 配置
    cfg := config.Section(typ)

    views := cfg.Key("views").String()
    ext := cfg.Key("ext").String()

    engine := jet.New(views, ext)

    engine.Delims("{{", "}}")
    engine.AddFunc("createTime", carbon.CreateFromTimestamp)

    // engine.Layout(key)
    // engine.Reload(false)
    // engine.Debug(false)

    return engine
}
