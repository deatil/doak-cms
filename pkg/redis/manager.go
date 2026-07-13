package redis

import (
    "fmt"
    "time"
    "context"

    "github.com/go-redis/redis/v8"
)

/**
 * redis
 *
 * @create 2022-7-6
 * @author deatil
 */
type Redis struct {
    // 上下文
    ctx context.Context

    // 客户端
    client *redis.Client

    // key 前缀
    keyPrefix string
}

// 构造函数
func New(config ...Config) *Redis {
    // 合并默认
    cfg := configDefault(config...)

    var options *redis.Options
    var err error

    if cfg.URL != "" {
        options, err = redis.ParseURL(cfg.URL)
        options.TLSConfig = cfg.TLSConfig
        if err != nil {
            panic(err)
        }
    } else {
        options = &redis.Options{
            Addr:      fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
            DB:        cfg.Database,
            Username:  cfg.Username,
            Password:  cfg.Password,
            TLSConfig: cfg.TLSConfig,
        }
    }

    // 上下文
    ctx := context.Background()

    // 客户端
    client := redis.NewClient(options)

    // 尝试连接
    if err := client.Ping(ctx).Err(); err != nil {
        panic(err)
    }

    // 测试
    if cfg.Reset {
        if err := client.FlushDB(ctx).Err(); err != nil {
            panic(err)
        }
    }

    // 创建
    return &Redis{
        ctx:       ctx,
        client:    client,
        keyPrefix: cfg.KeyPrefix,
    }
}

// 获取
func (this *Redis) Get(key string) ([]byte, error) {
    if len(key) <= 0 {
        return nil, nil
    }

    key = this.FormatKey(key)

    val, err := this.client.Get(this.ctx, key).Bytes()
    if err == redis.Nil {
        return nil, nil
    }

    return val, err
}

// 设置
func (this *Redis) Set(key string, val []byte, exp time.Duration) error {
    if len(key) <= 0 || len(val) <= 0 {
        return nil
    }

    key = this.FormatKey(key)

    return this.client.Set(this.ctx, key, val, exp).Err()
}

// 删除
func (this *Redis) Delete(key string) error {
    if len(key) <= 0 {
        return nil
    }

    key = this.FormatKey(key)

    return this.client.Del(this.ctx, key).Err()
}

// 重设
func (this *Redis) Reset() error {
    return this.client.FlushDB(this.ctx).Err()
}

// 关闭
func (this *Redis) Close() error {
    return this.client.Close()
}

// =================

// 获取上下文
func (this *Redis) GetCtx() context.Context {
    return this.ctx
}

// 获取客户端
func (this *Redis) GetClient() *redis.Client {
    return this.client
}

// 格式化 key
func (this *Redis) FormatKey(key string) string {
    if this.keyPrefix == "" {
        return key
    }

    return this.keyPrefix + ":" + key
}
