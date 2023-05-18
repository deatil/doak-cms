---
name: 从数据源加载
---

## 从数据源加载

就像之前说的，从多个数据源加载配置是基本操作。

那么，到底什么是 **数据源** 呢？

一个 **数据源** 可以是 `[]byte` 类型的原始数据，`string` 类型的文件路径或 `io.ReadCloser`。您可以加载 **任意多个** 数据源。如果您传递其它类型的数据源，则会直接返回错误。

```go
cfg, err := ini.Load(
    []byte("raw data"), // 原始数据
    "filename",         // 文件路径
    ioutil.NopCloser(bytes.NewReader([]byte("some other data"))),
)
```

或者从一个空白的文件开始：

```go
cfg := ini.Empty()
```

当您在一开始无法决定需要加载哪些数据源时，仍可以使用 [Append()](https://gowalker.org/gopkg.in/ini.v1#File_Append) 在需要的时候加载它们。

```go
err := cfg.Append("other file", []byte("other raw data"))
```

当您想要加载一系列文件，但是不能够确定其中哪些文件是不存在的，可以通过调用函数 [LooseLoad()](https://gowalker.org/gopkg.in/ini.v1#LooseLoad) 来忽略它们。

```go
cfg, err := ini.LooseLoad("filename", "filename_404")
```

更牛逼的是，当那些之前不存在的文件在重新调用 [Reload()](https://gowalker.org/gopkg.in/ini.v1#File_Reload) 方法的时候突然出现了，那么它们会被正常加载。

### 数据覆写

在加载多个数据源时，如果某一个键在一个或多个数据源中出现，则会出现数据覆写。该键从前一个数据源读取的值会被下一个数据源覆写。

举例来说，如果加载两个配置文件 `my.ini` 和 `my.ini.local`（[<i class="far fa-file-alt"></i> 开始使用](../intro/getting_started) 中的输入和输出文件），`app_mode` 的值会是 `production` 而不是 `development`。

```go
cfg, err := ini.Load("my.ini", "my.ini.local")
...

cfg.Section("").Key("app_mode").String() // production
```

数据覆写只有在一种情况下不会触发，即使用 [ShadowLoad](./work_with_keys#same-key-with-multiple-values) 加载数据源。

### 跳过无法识别的数据行

某些情况下，您的配置文件可能包含非键值对的数据行，解析器默认会报错并终止解析。如果您希望解析器能够忽略并它们完成剩余内容的解析，则可以通过如下方法实现：

```go
cfg, err := ini.LoadSources(ini.LoadOptions{
    SkipUnrecognizableLines: true,
}, "other.ini")
```

### 保存配置

终于到了这个时刻，是时候保存一下配置了。

比较原始的做法是输出配置到某个文件：

```go
// ...
err = cfg.SaveTo("my.ini")
err = cfg.SaveToIndent("my.ini", "\t")
```

另一个比较高级的做法是写入到任何实现 `io.Writer` 接口的对象中：

```go
// ...
cfg.WriteTo(writer)
cfg.WriteToIndent(writer, "\t")
```

默认情况下，空格将被用于对齐键值之间的等号以美化输出结果，以下代码可以禁用该功能：

```go
ini.PrettyFormat = false
```
