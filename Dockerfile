FROM golang:alpine

MAINTAINER jinke<geekestli@163.com>

# 环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

# 源码转移
WORKDIR /var/apps/echo
COPY . .

# 编译
RUN sh build.sh

# 启动服务
EXPOSE 3000
CMD ./server