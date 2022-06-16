---
name: 操作分区（Section）
---

## 操作分区（Section）

获取指定分区：

```go
sec, err := cfg.GetSection("section name")
```

如果您想要获取默认分区，则可以用空字符串代替分区名：

```go
sec, err := cfg.GetSection("")
```

相对应的，还可以使用 `ini.DEFAULT_SECTION` 来获取默认分区：

```go
sec, err := cfg.GetSection(ini.DEFAULT_SECTION)
```

当您非常确定某个分区是存在的，可以使用以下简便方法：

```go
sec := cfg.Section("section name")
```

如果不小心判断错了，要获取的分区其实是不存在的，那会发生什么呢？没事的，它会自动创建并返回一个对应的分区对象给您。

创建一个分区：

```go
err := cfg.NewSection("new section")
```

获取所有分区对象或名称：

```go
secs := cfg.Sections()
names := cfg.SectionStrings()
```

### 读取父子分区

您可以在分区名称中使用 `.` 来表示两个或多个分区之间的父子关系。如果某个键在子分区中不存在，则会去它的父分区中再次寻找，直到没有父分区为止。

```ini
NAME = ini
VERSION = v1
IMPORT_PATH = gopkg.in/%(NAME)s.%(VERSION)s

[package]
CLONE_URL = https://%(IMPORT_PATH)s

[package.sub]
```

```go
cfg.Section("package.sub").Key("CLONE_URL").String()	// https://gopkg.in/ini.v1
```

### 无法解析的分区

如果遇到一些比较特殊的分区，它们不包含常见的键值对，而是没有固定格式的纯文本，则可以使用 `LoadOptions.UnparsableSections` 进行处理：

```go
cfg, err := ini.LoadSources(ini.LoadOptions{
    UnparseableSections: []string{"COMMENTS"},
}, `[COMMENTS]
<1><L.Slide#2> This slide has the fuel listed in the wrong units <e.1>`)

body := cfg.Section("COMMENTS").Body()

/* --- start ---
<1><L.Slide#2> This slide has the fuel listed in the wrong units <e.1>
------  end  --- */
```