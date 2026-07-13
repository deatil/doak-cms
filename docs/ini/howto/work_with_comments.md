---
name: 操作注释（Comment）
---

## 操作注释（Comment）

下述几种情况的内容将被视为注释：

1. 所有以 `#` 或 `;` 开头的行
2. 所有在 `#` 或 `;` 之后的内容
3. 分区标签后的文字 (即 `[分区名]` 之后的内容)

如果你希望使用包含 `#` 或 `;` 的值，请使用 ``` ` ``` 或 ``` """ ``` 进行包覆。

除此之外，您还可以通过 [LoadOptions](https://gowalker.org/gopkg.in/ini.v1#LoadOptions) 完全忽略行内注释：

```go
cfg, err := ini.LoadSources(ini.LoadOptions{
    IgnoreInlineComment: true,
}, "app.ini")
```

或要求注释符号前必须带有一个空格：

```go
cfg, err := ini.LoadSources(ini.LoadOptions{
    SpaceBeforeInlineComment: true,
}, "app.ini")
```