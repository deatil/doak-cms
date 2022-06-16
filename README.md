## doak-cms 管理系统

doak-cms 是基于 fiber 和 Xorm 的 go 文章管理系统


### 项目介绍

*  cms 系统
*  使用 Fiber Xorm Rpcx Etcd Redis 等开发


### 环境要求

 - Go >= 1.18
 - Myql
 - Etcd
 - Redis


### 安装步骤

1. 首先克隆项目到本地

```
git clone https://github.com/deatil/lakego-admin.git
```

2. 然后配置数据库等相关配置，配置位置

```
/config
```

3. 最后导入 sql 数据到数据库

```go
data.sql
```

4. 运行测试

```go
go run main.go
```

6. 后台登录账号及密码：`admin` / `123456`


### 特别鸣谢

感谢以下的项目,排名不分先后

 - github.com/gofiber/fiber

 - xorm.io/xorm

 - github.com/smallnest/rpcx


### 开源协议

*  `doak-cms` 遵循 `Apache2` 开源协议发布，在保留本系统版权的情况下提供个人及商业免费使用。


### 版权

*  该系统所属版权归 deatil(https://github.com/deatil) 所有。
