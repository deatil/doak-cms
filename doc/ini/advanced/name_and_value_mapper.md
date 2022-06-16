---
name: 自定义键名和键值映射器
---

### 键名映射器（Name Mapper）

为了节省您的时间并简化代码，本库支持类型为 [`NameMapper`](https://gowalker.org/gopkg.in/ini.v1#NameMapper) 的名称映射器，该映射器负责结构字段名与分区名和键名之间的映射。

目前有 2 款内置的映射器：

- `AllCapsUnderscore`：该映射器将字段名转换至格式 `ALL_CAPS_UNDERSCORE` 后再去匹配分区名和键名。
- `TitleUnderscore`：该映射器将字段名转换至格式 `title_underscore` 后再去匹配分区名和键名。

使用方法：

```go
type Info struct {
	PackageName string
}

func main() {
	err = ini.MapToWithMapper(&Info{}, ini.TitleUnderscore, []byte("package_name=ini"))
	// ...

	cfg, err := ini.Load([]byte("PACKAGE_NAME=ini"))
	// ...
	info := new(Info)
	cfg.NameMapper = ini.AllCapsUnderscore
	err = cfg.MapTo(info)
	// ...
}
```

使用函数 `ini.ReflectFromWithMapper` 时也可应用相同的规则。

### 键值映射器（Value Mapper）

值映射器允许使用一个自定义函数自动展开值的具体内容，例如在运行时获取环境变量：

```go
type Env struct {
	Foo string `ini:"foo"`
}

func main() {
	cfg, err := ini.Load([]byte("[env]\nfoo = ${MY_VAR}\n"))
	cfg.ValueMapper = os.ExpandEnv
	// ...
	env := &Env{}
	err = cfg.Section("env").MapTo(env)
}
```

本例中，`env.Foo` 将会是运行时所获取到环境变量 `MY_VAR` 的值。
