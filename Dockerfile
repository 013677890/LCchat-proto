# 使用稳定的 Go 1.25 版本
FROM golang:1.25

WORKDIR /app

# 1. 设置国内代理
ENV GOPROXY=https://goproxy.cn,direct

# 2. 拷贝 go.mod 下载依赖 (当前没有 go.sum，先只拷 go.mod)
COPY go.mod ./
# 如果未来有了依赖，记得生成 go.sum 并取消下面这行的注释
# COPY go.sum ./
RUN go mod download

# 3. 拷贝源码
COPY . .

# 4. 运行应用
CMD ["go", "run", "./main.go"]
