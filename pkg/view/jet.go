package view

import (
    "github.com/gofiber/template/jet"

    "github.com/deatil/doak-cms/resources"
    "github.com/deatil/doak-cms/pkg/config"
)

func JetEngine(typ string) *jet.Engine {
    // 配置
    cfg := config.Section(typ)

    views := cfg.Key("views").String()
    ext := cfg.Key("ext").String()

    var engine *jet.Engine
    if config.IsUseEmbed {
        engine = jet.NewFileSystem(resources.ViewFileSystem(), ext)
    } else {
        engine = jet.New(views, ext)
    }

    engine.Delims("{{", "}}")

    // engine.AddFunc("func", useFunc)

    // engine.Layout(key)
    // engine.Reload(false)
    // engine.Debug(false)

    return engine
}
