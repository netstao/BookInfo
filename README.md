# 初始化项目结构

[cobra](https://github.com/spf13/cobra) 是一个用于生成命令行应用程序的工具。它允许你快速地创建具有子命令的复杂 CLI 应用，并且自动处理帮助信息和标志解析等常见任务。

```bash
    go mod init cloudnative.test.com
    go get -u github.com/spf13/cobra@latest
    go install github.com/spf13/cobra-cli@latest
    cobra-cli init

```

## 初始化web模块 Gin
[Gin](https://github.com/gin-gonic/gin)
```
    go get -u github.com/gin-gonic/gin@latest
    cobra-cli.exe add server
    go run main.go serve
```

## 功能开发
- [X] 首页(展示服务间的调用关系)
- [] 图书详情页

## docker构建

```
    docker build -t bookinfo:v.0.1 .
```