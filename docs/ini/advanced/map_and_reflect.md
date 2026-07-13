---
name: 结构体与分区双向映射
---

### 映射到结构

想要使用更加面向对象的方式玩转 INI 吗？好主意。

```ini
Name = Unknwon
age = 21
Male = true
Born = 1993-01-01T20:17:05Z

[Note]
Content = Hi is a good man!
Cities = HangZhou, Boston
```

```go
type Note struct {
	Content string
	Cities  []string
}

type Person struct {
	Name string
	Age  int `ini:"age"`
	Male bool
	Born time.Time
	Note
	Created time.Time `ini:"-"`
}

func main() {
	cfg, err := ini.Load("path/to/ini")
	// ...
	p := new(Person)
	err = cfg.MapTo(p)
	// ...

	// 一切竟可以如此的简单。
	err = ini.MapTo(p, "path/to/ini")
	// ...

	// 嗯哼？只需要映射一个分区吗？
	n := new(Note)
	err = cfg.Section("Note").MapTo(n)
	// ...
}
```

结构的字段怎么设置默认值呢？很简单，只要在映射之前对指定字段进行赋值就可以了。如果键未找到或者类型错误，该值不会发生改变。

```go
// ...
p := &Person{
	Name: "Joe",
}
// ...
```

这样玩 INI 真的好酷啊！然而，如果不能还给我原来的配置文件，有什么卵用？

### 从结构反射

可是，我有说不能吗？

```go
type Embeded struct {
	Dates  []time.Time `delim:"|" comment:"Time data"`
	Places []string    `ini:"places,omitempty"`
	None   []int       `ini:",omitempty"`
}

type Author struct {
	Name      string `ini:"NAME"`
	Male      bool
	Age       int `comment:"Author's age"`
	GPA       float64
	NeverMind string `ini:"-"`
	*Embeded `comment:"Embeded section"`
}

func main() {
	a := &Author{"Unknwon", true, 21, 2.8, "",
		&Embeded{
			[]time.Time{time.Now(), time.Now()},
			[]string{"HangZhou", "Boston"},
			[]int{},
		}}
	cfg := ini.Empty()
	err = ini.ReflectFrom(cfg, a)
	// ...
}
```

瞧瞧，奇迹发生了。

```ini
NAME = Unknwon
Male = true
; Author's age
Age = 21
GPA = 2.8

; Embeded section
[Embeded]
; Time data
Dates = 2015-08-07T22:14:22+08:00|2015-08-07T22:14:22+08:00
places = HangZhou,Boston
```

### 配合 ShadowLoad 进行映射

如果您希望配合 [ShadowLoad](../howto/work_with_keys#same-key-with-multiple-values) 将某个分区映射到结构体，则需要指定 `allowshadow` 标签。

假设您有以下配置文件：

```ini
[IP]
value = 192.168.31.201
value = 192.168.31.211
value = 192.168.31.221
```

您应当通过如下方式定义对应的结构体：

```go
type IP struct {
   Value    []string `ini:"value,omitempty,allowshadow"`
}
```

如果您不需要前两个标签规则，可以使用 `ini:",,allowshadow"` 进行简写。

### 映射/反射的其它说明

任何嵌入的结构都会被默认认作一个不同的分区，并且不会自动产生所谓的父子分区关联：

```go
type Child struct {
	Age string
}

type Parent struct {
	Name string
	Child
}

type Config struct {
	City string
	Parent
}
```

示例配置文件：

```ini
City = Boston

[Parent]
Name = Unknwon

[Child]
Age = 21
```

很好，但是，我就是要嵌入结构也在同一个分区。好吧，你爹是李刚！

```go
type Child struct {
	Age string
}

type Parent struct {
	Name string
	Child `ini:"Parent"`
}

type Config struct {
	City string
	Parent
}
```

示例配置文件：

```ini
City = Boston

[Parent]
Name = Unknwon
Age = 21
```

另请查看 [<i class="far fa-file-alt"></i> 自定义键名和键值映射器](./name_and_value_mapper) 的有关用法。