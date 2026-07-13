package redis

import "crypto/tls"

// 配置
// https://pkg.go.dev/github.com/go-redis/redis/v8#Options
type Config struct {
    Host      string
    Port      int
    Username  string
    Password  string
    Database  int
    URL       string
    Reset     bool
    TLSConfig *tls.Config
    KeyPrefix string
}

// 默认配置
var ConfigDefault = Config{
    Host:      "127.0.0.1",
    Port:      6379,
    Username:  "",
    Password:  "",
    URL:       "",
    Database:  0,
    Reset:     false,
    TLSConfig: nil,
    KeyPrefix: "",
}

// 补充未设置默认
func configDefault(config ...Config) Config {
    // 返回默认
    if len(config) < 1 {
        return ConfigDefault
    }

    cfg := config[0]

    if cfg.Host == "" {
        cfg.Host = ConfigDefault.Host
    }

    if cfg.Port <= 0 {
        cfg.Port = ConfigDefault.Port
    }

    return cfg
}
