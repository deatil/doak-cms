package config

import (
    "embed"
)

//go:embed config.ini
var Config embed.FS

// 读取配置
func ReadConfig(name string) ([]byte, error) {
    return Config.ReadFile(name)
}
