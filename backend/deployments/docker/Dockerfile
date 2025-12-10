# 构建阶段
# 使用 --platform 指令支持多架构
FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder

WORKDIR /app

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 编译为目标平台
ARG TARGETOS=linux
ARG TARGETARCH=amd64
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o server ./cmd/server

# 运行阶段
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# 从构建阶段复制文件
COPY --from=builder /app/server .
COPY --from=builder /app/config ./config

EXPOSE 8080

CMD ["./server"]
