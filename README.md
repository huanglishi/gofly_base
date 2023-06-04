# GoFly系列介绍
用 Go 语言基于Gin、Vue、Typescript、vite、Arco Design、MySQL等框架精心打造的一款模块化、插件化、高度自动化、高性能的前后端分离架构敏捷开发框架，可快速搭建前后端分离后台管理系统。本着简化开发、提升开发效率的初衷，框架自研了一套个性化的组件，实现了可插拔的组件式开发方式，同时为了敏捷快速开发，框架特地集成了代码生成器一键生成管理后台的前后端代码，快速开发或生成api接口代码和接口文档、测试api接口等编辑，完全自主研发了自定义GO后端服务模板和前端Vue自定义模板，可以根据已建好的表结构，可以快速的一键生成整个模块的所有代码和增删改查等等功能业务，几秒即可生成应用端（小程序，app）调用的api接口。真正实现了低代码开发方式，极大的节省了人力成本的同时提高了开发效率，缩短了研发周期，是一款真正意义上实现组件化、可插拔式的敏捷开发框架。保证同等项目时间，您用GoFly的框架后，后台和接口可以减少7成时间，是否真实您自己体验，我们自己人不能崔牛了！
## 亲身体验！
框架官网地址：[GoFly全栈开发者社区](https://goflys.cn/home)
框架源码下载地址：[GoFly 后端快速开发框架源码](https://goflys.cn/prdetail?id=6)
框架演示地址：[GoFly快速开的应用演示](https://sg.goflys.cn/webbusiness/#/developer/devapi)
框架还有总管理后台轻松实现saas开发：[Go体验总后台](https://sg.goflys.cn/webadmin/)

# 基础版使用说明
## 安装fresh 热更新-边开发边编译
go install github.com/pilu/fresh@latest

## 初始化mod
go mod tidy

# 热编译运行
bee run 或 fresh 
# 打包
go build main.go
# 打包（此时会打包成Linux上可运行的二进制文件，不带后缀名的文件）
```
SET GOOS=linux
SET GOARCH=amd64
go build
```
## widows
```
// 配置环境变量
SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64

go build main.go

// 编译命令
```
### 编译成Linux环境可执行文件
```

// 配置参数
SET CGO_ENABLED=0 
SET GOOS=linux 
SET GOARCH=amd64 

go build main.go

// 编译命令
```
