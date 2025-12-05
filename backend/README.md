# FloatIsland Backend

FloatIsland 个人财务管理应用后端服务

## 技术栈

- **框架**: Go 1.21+ with Gin
- **数据库**: MySQL 8.0+
- **缓存**: Redis 7+
- **ORM**: GORM
- **认证**: JWT
- **日志**: Zap
- **配置**: Viper

## 项目结构

```
backend/
├── cmd/server/          # 应用入口
├── internal/            # 内部代码
│   ├── api/            # API 层
│   │   ├── handlers/   # HTTP 处理函数
│   │   ├── middlewares/# 中间件
│   │   └── routes/     # 路由定义
│   ├── service/        # 业务逻辑层
│   ├── repository/     # 数据访问层
│   ├── models/         # 数据模型
│   ├── dto/            # 数据传输对象
│   └── utils/          # 工具函数
├── pkg/                # 公共库
│   ├── cache/          # Redis 封装
│   ├── database/       # 数据库封装
│   ├── logger/         # 日志封装
│   └── storage/        # 对象存储封装
├── config/             # 配置文件
├── migrations/         # 数据库迁移
├── deployments/        # 部署配置
│   └── docker/         # Docker 配置
├── docs/               # 文档
└── tests/              # 测试
```

## 快速开始

### 前置要求

- Go 1.21+
- MySQL 8.0+
- Redis 7+
- Docker (可选)

### 本地开发

1. **安装依赖**

```bash
make deps
```

2. **配置环境**

复制配置文件并修改：
```bash
cp config/config.yaml config/config.local.yaml
# 编辑 config/config.local.yaml 配置数据库等信息
```

3. **运行服务**

```bash
make run
```

服务将在 `http://localhost:8080` 启动

### Docker 部署

1. **构建并启动所有服务**

```bash
make docker-up
```

这将启动：
- API 服务 (端口 8080)
- MySQL (端口 3306)
- Redis (端口 6379)
- MinIO (端口 9000, 9001)

2. **查看日志**

```bash
make docker-logs
```

3. **停止服务**

```bash
make docker-down
```

## API 文档

详细的 API 文档请查看：
- [API 设计文档](../design/api/api-spec.md)
- [架构设计文档](../design/architecture/architecture.md)

### 主要端点

- `GET /health` - 健康检查
- `POST /api/v1/auth/register` - 用户注册
- `POST /api/v1/auth/login` - 用户登录
- `GET /api/v1/users/me` - 获取当前用户信息
- `GET /api/v1/transactions` - 获取交易列表
- `POST /api/v1/transactions` - 创建交易

## 开发指南

### 添加新的 API 端点

1. 在 `internal/api/handlers/` 添加处理函数
2. 在 `internal/api/routes/router.go` 注册路由
3. 在 `internal/service/` 实现业务逻辑
4. 在 `internal/repository/` 实现数据访问

### 数据库迁移

使用 golang-migrate 管理数据库迁移：

```bash
# 创建迁移文件
migrate create -ext sql -dir migrations -seq create_users_table

# 执行迁移
migrate -path migrations -database "mysql://user:pass@tcp(localhost:3306)/float_db" up

# 回滚迁移
migrate -path migrations -database "mysql://user:pass@tcp(localhost:3306)/float_db" down
```

## 测试

```bash
# 运行所有测试
make test

# 运行特定测试
go test -v ./internal/service/...
```

## 配置说明

主要配置项（`config/config.yaml`）：

- `server.port`: 服务端口
- `database.*`: MySQL 数据库配置
- `redis.*`: Redis 配置
- `jwt.secret`: JWT 密钥
- `storage.*`: 对象存储配置

## 贡献指南

1. Fork 项目
2. 创建功能分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## 许可证

MIT License

---

**版本**: 1.0.0  
**更新日期**: 2025-12-05
