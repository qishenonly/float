# Float Backend Docker 部署指南

完整的Float Backend微服务堆栈，包含API服务、MySQL、Redis和MinIO。

## 📋 前置要求

- Docker >= 20.10
- Docker Compose >= 2.0
- 磁盘空间 >= 5GB（MySQL + Redis + MinIO 数据存储）

## 🚀 快速开始

### 1. 配置环境变量

```bash
# 复制示例配置文件
cp .env.example .env

# 编辑 .env 文件，根据需要修改各项配置
# 关键配置项：
# - DB_PASSWORD: MySQL密码
# - JWT_SECRET: JWT密钥（必须更改）
# - MINIO_PASSWORD: MinIO密码
```

### 2. 构建镜像

```bash
# 使用Makefile（推荐）
cd ../..
make docker-build

# 或直接使用docker compose
docker compose build
```

### 3. 启动服务

```bash
# 使用Makefile
make docker-up

# 或直接使用docker compose
docker compose up -d
```

### 4. 验证服务状态

```bash
# 查看容器运行状态
docker compose ps

# 查看API日志
docker compose logs -f api

# 检查所有服务健康状态
docker compose ps --format "table {{.Names}}\t{{.Status}}"
```

## 📚 服务访问信息

| 服务 | 地址 | 默认端口 | 凭证 |
|------|------|---------|------|
| API | http://localhost:8080 | 8080 | - |
| MySQL | localhost:3306 | 3306 | root / `MYSQL_ROOT_PASSWORD` |
| Redis | localhost:6379 | 6379 | - |
| MinIO API | http://localhost:9000 | 9000 | minioadmin / minioadmin |
| MinIO Console | http://localhost:9001 | 9001 | minioadmin / minioadmin |

## 🔧 常用命令

```bash
# 启动所有服务
docker compose up -d

# 停止所有服务（保留数据）
docker compose down

# 完全清理（包括数据）
docker compose down -v

# 查看特定服务日志
docker compose logs -f [service-name]
# 例如：docker compose logs -f api

# 进入容器终端
docker compose exec [service-name] /bin/sh
# 例如：docker compose exec api /bin/sh

# 重启特定服务
docker compose restart [service-name]

# 查看服务资源使用情况
docker stats
```

## 📝 配置说明

### API 服务配置

API 服务的配置文件映射在 `../../config` 目录，包括：
- `config.yaml`: 默认配置
- `config.dev.yaml`: 开发环境配置
- `config.prod.yaml`: 生产环境配置

修改配置后需要重启API服务。

### 数据库初始化

可选：将SQL初始化脚本放在 `init-scripts` 目录下，Docker会在MySQL启动时自动执行。

```bash
# 示例
touch init-scripts/01-init.sql
# 编辑SQL文件...
# 重启MySQL服务会自动执行
```

### 数据持久化

所有数据都持久化存储在Docker卷中：
- `mysql_data`: MySQL数据
- `redis_data`: Redis数据
- `minio_data`: MinIO数据
- `api_uploads`: API上传文件
- `api_logs`: API日志

## 🔐 安全建议

1. **生产环境必须修改**：
   - `JWT_SECRET`: 使用 `openssl rand -base64 32` 生成
   - `MYSQL_ROOT_PASSWORD`: 设置强密码
   - `MINIO_PASSWORD`: 设置强密码

2. **不提交敏感信息**：
   - `.env` 文件不应上传到版本控制
   - 已在 `.gitignore` 中排除

3. **定期备份**：
   ```bash
   # 备份MySQL数据
   docker compose exec mysql mysqldump -uroot -p$MYSQL_ROOT_PASSWORD --all-databases > backup.sql
   ```

## 🐛 故障排除

### API 无法连接数据库

```bash
# 检查MySQL服务是否正常
docker compose logs mysql

# 验证数据库连接
docker compose exec mysql mysql -uroot -p$MYSQL_ROOT_PASSWORD -e "SELECT 1"
```

### Redis 连接失败

```bash
# 检查Redis是否在运行
docker compose exec redis redis-cli ping

# 如果设置了密码，使用-a参数
docker compose exec redis redis-cli -a $REDIS_PASSWORD ping
```

### MinIO 无法访问

```bash
# 检查MinIO日志
docker compose logs minio

# 访问http://localhost:9001查看Web控制台
```

### 端口冲突

如果默认端口被占用，编辑 `.env` 文件修改端口：
```env
API_PORT=8081
MYSQL_PORT=3307
REDIS_PORT=6380
MINIO_API_PORT=9002
MINIO_CONSOLE_PORT=9003
```

## 📦 生成发布镜像

打包镜像为tar文件用于离线部署：

```bash
# 使用Makefile
make docker-build-tar

# 生成的文件：float-backend-api.tar
# 在其他机器上加载：
docker load < float-backend-api.tar
```

## 🔄 更新部署

```bash
# 更新源代码后，重新构建镜像
docker compose build --no-cache

# 停止旧容器并启动新容器
docker compose down
docker compose up -d
```

## 📖 更多信息

- [Docker Compose 官方文档](https://docs.docker.com/compose/)
- [Float Backend GitHub](https://github.com/qishenonly/float)
