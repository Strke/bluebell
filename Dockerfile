FROM golang:alpine AS builder

# 设置环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录
WORKDIR /build

# 复制go.mod和go.sum并下载依赖
COPY go.mod .
COPY go.sum .
RUN go mod download

# 复制代码到容器中
COPY . .

#将我们的代码编译成二进制可执行文件bluebell
RUN go build -o bluebell_app .

# 创建一个小镜象
FROM debian:stretch-slim

COPY ./wait-for.sh /
COPY ./templates /templates
COPY ./static /static
COPY ./conf /conf
COPY ./sources.list /etc/apt/sources.list

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/bluebell_app /

RUN set -eux; \
	apt-get update; \
	apt-get install -y \
		--no-install-recommends \
		netcat; \
        chmod 755 wait-for.sh

EXPOSE 8081
# 需要运行的命令
# ENTRYPOINT ["/bluebell_app", "conf/config.yaml"]
