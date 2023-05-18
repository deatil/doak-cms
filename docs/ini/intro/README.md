---
name: 简介
---

# INI [![Build Status](https://travis-ci.org/go-ini/ini.svg?branch=master)](https://travis-ci.org/go-ini/ini)  [![](https://sourcegraph.com/github.com/go-ini/ini/-/badge.svg)](https://sourcegraph.com/github.com/go-ini/ini?badge)

![](https://avatars0.githubusercontent.com/u/10216035?v=3&s=200)

本包提供了 Go 语言中读写 INI 文件的功能。

### 功能特性

- 支持覆盖加载多个数据源（file, `[]byte`, `io.Reader` and `io.ReadCloser`）
- 支持递归读取键值
- 支持读取父子分区
- 支持读取自增键名
- 支持读取多行的键值
- 支持大量辅助方法
- 支持在读取时直接转换为 Go 语言类型
- 支持读取和 **写入** 分区和键的注释
- 轻松操作分区、键值和注释
- 在保存文件时分区和键值会保持原有的顺序

### 下载安装

最低要求安装 Go 语言版本为 **1.6**。

```sh
$ go get gopkg.in/ini.v1
```

如需更新请添加 `-u` 选项。

### 获取帮助

- [API 文档](https://gowalker.org/gopkg.in/ini.v1)
- [创建工单](https://github.com/go-ini/ini/issues/new)
