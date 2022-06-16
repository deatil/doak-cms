package cms

import (
    "github.com/deatil/doak-cms/pkg/config"
)

// Base
type Base struct{}

func (this *Base) Theme(tpl string) string {
    // 配置
    cfg := config.Section("view")

    theme := cfg.Key("theme").MustString("doak")

    view := theme + "/" + tpl

    return view
}
