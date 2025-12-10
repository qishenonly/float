.PHONY: help build run test docker-build docker-build-tar docker-up docker-down docker-logs clean

help: ## 显示帮助信息
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## 编译项目
	go build -o bin/server cmd/server/main.go

run: ## 运行项目
	go run cmd/server/main.go

test: ## 运行测试
	go test -v ./...

deps: ## 安装依赖
	go mod download
	go mod tidy

docker-build: ## 构建 Docker 镜像
	cd deployments/docker && docker-compose build

docker-build-tar: ## 构建 Docker 镜像并打包为 tar 文件 (支持版本号: make docker-build-tar VERSION=v1.0.0)
	@echo "开始构建镜像..."; \
	cd deployments/docker && docker-compose build; \
	echo "镜像构建完成，开始打包..."; \
	VERSION=$${VERSION:-latest}; \
	LATEST_IMAGE=float-backend-api:latest; \
	TARGET_IMAGE=float-backend-api:$$VERSION; \
	if [ "$$VERSION" != "latest" ]; then \
		echo "标签化镜像: $$LATEST_IMAGE -> $$TARGET_IMAGE"; \
		docker tag $$LATEST_IMAGE $$TARGET_IMAGE; \
	fi; \
	if docker images --quiet $$TARGET_IMAGE > /dev/null 2>&1; then \
		echo "📦 检测到镜像: $$TARGET_IMAGE"; \
		TAR_FILE=float-backend-api-$$VERSION.tar; \
		docker save -o $$TAR_FILE $$TARGET_IMAGE; \
		echo "✓ 镜像已打包为 $$TAR_FILE"; \
		ls -lh $$TAR_FILE; \
	else \
		echo "❌ 错误：未找到镜像 $$TARGET_IMAGE"; \
		echo "可用镜像列表:"; \
		docker images | grep -E "float|api|backend" || docker images | head -10; \
		exit 1; \
	fi

docker-up: ## 启动 Docker 容器
	cd deployments/docker && docker-compose up -d

docker-down: ## 停止 Docker 容器
	cd deployments/docker && docker-compose down

docker-logs: ## 查看 Docker 日志
	cd deployments/docker && docker-compose logs -f api

clean: ## 清理编译文件
	rm -rf bin/
	go clean

.DEFAULT_GOAL := help
