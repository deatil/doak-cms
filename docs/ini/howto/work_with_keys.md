---
name: 操作键（Key）
---

## 操作键（Key）

获取某个分区下的键：

```go
key, err := cfg.Section("").GetKey("key name")
```

和分区一样，您也可以直接获取键而忽略错误处理：

```go
key := cfg.Section("").Key("key name")
```

判断某个键是否存在：

```go
yes := cfg.Section("").HasKey("key name")
```

创建一个新的键：

```go
err := cfg.Section("").NewKey("name", "value")
```

获取分区下的所有键或键名：

```go
keys := cfg.Section("").Keys()
names := cfg.Section("").KeyStrings()
```

获取分区下的所有键值对的克隆：

```go
hash := cfg.Section("").KeysHash()
```

#### 忽略键名的大小写

有时候分区和键的名称大小写混合非常烦人，这个时候就可以通过 [InsensitiveLoad](https://gowalker.org/gopkg.in/ini.v1#InsensitiveLoad) 将所有分区和键名在读取里强制转换为小写：

```go
cfg, err := ini.InsensitiveLoad("filename")
//...

// sec1 和 sec2 指向同一个分区对象
sec1, err := cfg.GetSection("Section")
sec2, err := cfg.GetSection("SecTIOn")

// key1 和 key2 指向同一个键对象
key1, err := sec1.GetKey("Key")
key2, err := sec2.GetKey("KeY")
```

#### 类似 MySQL 配置中的布尔值键

MySQL 的配置文件中会出现没有具体值的布尔类型的键：

```ini
[mysqld]
...
skip-host-cache
skip-name-resolve
```

默认情况下这被认为是缺失值而无法完成解析，但可以通过高级的加载选项对它们进行处理：

```go
cfg, err := ini.LoadSources(ini.LoadOptions{
    AllowBooleanKeys: true,
}, "my.cnf")
```

这些键的值永远为 `true`，且在保存到文件时也只会输出键名。

如果您想要通过程序来生成此类键，则可以使用 `NewBooleanKey`：

```go
key, err := sec.NewBooleanKey("skip-host-cache")
```

### 同个键名包含多个值

你是否也曾被下面的配置文件所困扰？

```ini
[remote "origin"]
url = https://github.com/Antergone/test1.git
url = https://github.com/Antergone/test2.git
fetch = +refs/heads/*:refs/remotes/origin/*
```

没错！默认情况下，只有最后一次出现的值会被保存到 `url` 中，可我就是想要保留所有的值怎么办啊？不要紧，用 [ShadowLoad](https://gowalker.org/gopkg.in/ini.v1#ShadowLoad) 轻松解决你的烦恼：

```go
cfg, err := ini.ShadowLoad(".gitconfig")
// ...

f.Section(`remote "origin"`).Key("url").String() 
// Result: https://github.com/Antergone/test1.git

f.Section(`remote "origin"`).Key("url").ValueWithShadows()
// Result:  []string{
//              "https://github.com/Antergone/test1.git",
//              "https://github.com/Antergone/test2.git",
//          }
```

### 读取自增键名

如果数据源中的键名为 `-`，则认为该键使用了自增键名的特殊语法。计数器从 1 开始，并且分区之间是相互独立的。

```ini
[features]
-: Support read/write comments of keys and sections
-: Support auto-increment of key names
-: Support load multiple files to overwrite key values
```

```go
cfg.Section("features").KeyStrings()	// []{"#1", "#2", "#3"}
```

### 获取上级父分区下的所有键名

```go
cfg.Section("package.sub").ParentKeys() // ["CLONE_URL"]
```