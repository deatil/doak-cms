package log

import (
    "time"
    "sync"

    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "github.com/natefinch/lumberjack"

    "github.com/deatil/doak-cms/pkg/config"
)

var log *zap.Logger
var once sync.Once

// 日志
// Log().Info("hello info")
// Log().Debug("hello debug")
// Log().Error("hello error")
// Log().Warn("hello Warn")
func Log() *zap.Logger {
    once.Do(func() {
        log = Manager("log")
    })

    return log
}

// 日志管理
func Manager(typ string) *zap.Logger {
    // 配置
    cfg := config.Section(typ)

    filename := cfg.Key("filename").String()
    maxSize := cfg.Key("max-size").MustInt(1)
    maxBackups := cfg.Key("max-backups").MustInt(5)
    maxAge := cfg.Key("max-age").MustInt(30)
    compress := cfg.Key("compress").MustBool(false)

    // 文件writeSyncer
    fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
        Filename:   filename,   // 日志文件存放目录
        MaxSize:    maxSize,    // 文件大小限制,单位MB
        MaxBackups: maxBackups, // 最大保留日志文件数量
        MaxAge:     maxAge,     // 日志文件保留天数
        Compress:   compress,   // 是否压缩处理
    })

    //日志级别
    highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool{  //error级别
        return lev >= zap.InfoLevel
    })

    // 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
    fileCore := zapcore.NewCore(
        getEncoder(typ),
        fileWriteSyncer,
        highPriority,
    )

    // AddCaller() 为显示文件名和行号
    logger := zap.New(fileCore, zap.AddCaller())

    return logger
}

// 获取 zapcore.EncoderConfig
func getEncoderConfig(typ string) zapcore.EncoderConfig {
    // 配置
    cfg := config.Section(typ)

    stacktraceKey := cfg.Key("stacktrace-key").String()
    encodeLevel := cfg.Key("encode-level").String()

    conf := zapcore.EncoderConfig{
        MessageKey:     "message",
        LevelKey:       "level",
        TimeKey:        "time",
        NameKey:        "logger",
        CallerKey:      "caller",
        StacktraceKey:  stacktraceKey,
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,
        EncodeTime:     formatTimeEncoder(typ),
        EncodeDuration: zapcore.SecondsDurationEncoder,
        EncodeCaller:   zapcore.FullCallerEncoder,
    }

    switch {
        case encodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
            conf.EncodeLevel = zapcore.LowercaseLevelEncoder
        case encodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
            conf.EncodeLevel = zapcore.LowercaseColorLevelEncoder
        case encodeLevel == "CapitalLevelEncoder": // 大写编码器
            conf.EncodeLevel = zapcore.CapitalLevelEncoder
        case encodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
            conf.EncodeLevel = zapcore.CapitalColorLevelEncoder
        default:
            conf.EncodeLevel = zapcore.LowercaseLevelEncoder
    }

    return conf
}

// 获取zapcore.Encoder
func getEncoder(typ string) zapcore.Encoder {
    // 格式化方式
    format := config.Section(typ).Key("format").String()

    if format == "json" {
        return zapcore.NewJSONEncoder(getEncoderConfig(typ))
    }

    return zapcore.NewConsoleEncoder(getEncoderConfig(typ))
}

// 自定义日志输出时间格式
func formatTimeEncoder(typ string) func(time.Time, zapcore.PrimitiveArrayEncoder) {
    return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
        // 前缀
        prefix := config.Section(typ).Key("prefix").String()

        enc.AppendString(t.Format(prefix + "2006/01/02 - 15:04:05"))
    }
}
