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

## docker构建与运行

```
    docker build -t bookinfo:v.0.1 .
    docker run -p:8082:8080 -d bookinfo:v0.1
    docker images | grep bookinfo 
    
```

## 镜像打tag 推送
```
    docker tag bookinfo:v0.1 harbor.test.com/test/bookinfo:v0.1
   docker login harbor.test.com -u xxxx -p xxxxx
   # 拷贝证书到本地目录，否则推送时会报错
   cp /etc/docker/certs.d/harbor.test.com/ca.crt ~/.docker/
   docker push harbor.test.com/test/bookinfo:v0.1
```

## 容器变成镜像
```
docker commit 4908f71c1f48 bookinfo:v0.2
docker run -p:8083:8080 -d bookinfo:v0.2
```

## 1.24.3-alpine3.21容器编译 

```
cd /root/work/BookInfo
# 指定容器编译
docker run -v ./:/server --rm -itd --name compile golang:1.24.3-alpine3.21
docker exec -it compile /bin/sh
go env -w GOPROXY=https://goproxy.cn,direct
go build -o bk
```

## alpine:3.21
- 最小容器运行
```
    docker build -t bookinfo:base-image-lab -f deploy/Dockerfile-baseimage-lab .
    docker images | grep lab
    docker tag bookinfo:base-image-lab harbor.test.com/test/bookinfo:base-image-lab
    docker push harbor.test.com/test/bookinfo:base-image-lab
    # 运行容器 --rm 停掉容器后，会自动删除容器
    docker run --rm --name bookinfo  -p:8084:8080 -d harbor.test.com/test/bookinfo:base-image-lab
```

## 构建多阶段编译
```
    docker build -t bookinfo:multi-compile-lab -f deploy/Dockerfile.muticompile-lab .
```