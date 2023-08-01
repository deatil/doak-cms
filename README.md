## doak-cms 文章管理系统

doak-cms 是使用 gofiber 和 Xorm 的 go 文章管理系统


### 项目介绍

*  使用 go 开发的通用文章管理系统
*  核心使用 Fiber、Xorm 及 Rpcx 等开发
*  默认打包配置文件，模板文件及静态文件，一个文件即可部署
*  不使用 rpc 可将 `main.go` 里 rpc 部分注释掉


### 环境要求

 - Go >= 1.18
 - Myql
 - Redis
 - Etcd


### 截图预览

<table>
    <tr>
        <td width="50%">
            <center>
                <img alt="登录" src="https://user-images.githubusercontent.com/24578855/174941901-3f5dbcf7-2d0a-40dc-8dc1-14a92e4a8604.png" />
            </center>
        </td>
        <td width="50%">
            <center>
                <img alt="控制台" src="https://user-images.githubusercontent.com/24578855/174941661-795ae01a-c8b4-4275-8fbb-d50a7927a200.png" />
            </center>
        </td>
    </tr>
    <tr>
        <td width="50%">
            <center>
                <img alt="文章管理" src="https://user-images.githubusercontent.com/24578855/174941958-b7127708-7c8c-4dac-98d8-efc4db0a5c8a.png" />
            </center>
        </td>
        <td width="50%">
            <center>
                <img alt="分类管理" src="https://user-images.githubusercontent.com/24578855/174942001-1d4f7890-8b5a-4396-8c22-ce09e8f9f8d6.png" />
            </center>
        </td>
    </tr>
</table>

更多截图
[doak-cms 截图](https://github.com/deatil/doak-cms/issues/1)


### 安装步骤

1. 首先克隆项目到本地

```
git clone https://github.com/deatil/doak-cms.git
```

2. 然后配置数据库等相关配置，配置位置

```
/config/config.ini
```

3. 最后导入 sql 数据到数据库

```
/doc/doak_cms.sql
```

4. 运行测试

```go
go run main.go
```

6. 后台登录账号及密码：`admin` / `123456`, 登录地址: `/sys`


### 特别鸣谢

感谢以下的项目,排名不分先后

 - github.com/gofiber/fiber

 - xorm.io/xorm

 - github.com/smallnest/rpcx

 - github.com/spf13/cast


### 开源协议

*  `doak-cms` 遵循 `Apache2` 开源协议发布，在保留本系统版权的情况下提供个人及商业免费使用。


### 版权

*  该系统所属版权归 deatil(https://github.com/deatil) 所有。
