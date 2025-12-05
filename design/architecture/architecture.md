# FloatIsland - 后端架构设计方案

基于前端需求和数据库设计，为 FloatIsland 个人财务管理应用制定完整的后端架构方案。

### 1. 技术栈选型

#### 1.1 后端框架

**推荐：Go + Gin 框架**

理由：
- ✅ 高性能，低内存占用，适合移动端后端服务
- ✅ 并发支持优秀，能高效处理大量用户请求
- ✅ 静态类型安全，减少运行时错误
- ✅ 部署简单，编译成单一二进制文件
- ✅ Gin 生态成熟，中间件丰富

#### 1.2 数据库

**主数据库：MySQL 8.0+**

理由：
- ✅ 成熟稳定，广泛使用，社区和云服务支持完善
- ✅ 优秀的 JSON 字段支持（用于 tags, images, changelog 等字段）
- ✅ 强大的索引和查询优化能力
- ✅ 完善的事务支持（InnoDB 引擎）
- ✅ 云服务支持广泛（阿里云RDS、腾讯云TencentDB、AWS RDS等）
- ✅ GORM 与 MySQL 集成成熟

**缓存层：Redis 7+**

用途：
- Session 存储
- 验证码缓存
- 热点数据缓存（用户余额、月度统计等）
- 分布式锁
- 消息队列（账单提醒、推送通知）

#### 1.3 对象存储

**云服务推荐：**
- **阿里云 OSS**（推荐国内部署）
- **腾讯云 COS**（推荐国内部署）
- **AWS S3**（推荐海外部署）

**自建方案：**
- **MinIO** - 开源对象存储，兼容 S3 API

**用途：**
- 用户头像存储
- 交易凭证图片
- 心愿单商品图片
- 数据导出文件临时存储
- 应用更新包托管

#### 1.4 云部署支持

**支持的云平台：**
- ✅ **阿里云**：ECS + RDS MySQL + OSS + Redis
- ✅ **腾讯云**：CVM + TencentDB MySQL + COS + Redis
- ✅ **AWS**：EC2 + RDS MySQL + S3 + ElastiCache
- ✅ **华为云**：ECS + RDS MySQL + OBS + Redis

#### 1.5 其他组件

| 组件 | 选型 | 用途 |
|------|------|------|
| ORM | GORM | MySQL 数据库操作 |
| 数据库驱动 | gorm.io/driver/mysql | MySQL驱动 |
| 日志 | Zap | 高性能结构化日志 |
| 配置管理 | Viper | 多格式配置文件解析 |
| API文档 | Swaggo | 自动生成 Swagger 文档 |
| 参数验证 | go-playground/validator | 请求参数验证 |
| 任务调度 | robfig/cron | 定时任务调度 |
| 消息推送 | Firebase Cloud Messaging (FCM) | 移动端推送通知 |
| 邮件服务 | SMTP / SendGrid / 阿里云邮件 | 邮件通知 |
| 文件上传 | 对应云厂商 SDK | OSS/COS/S3 操作 |

---

### 2. 系统架构设计

#### 2.1 整体架构图

```mermaid
graph TB
    subgraph "客户端层"
        A1[Android App<br/>Capacitor]
        A2[iOS App<br/>Capacitor]
        A3[Web App<br/>Vue 3]
    end
    
    subgraph "负载均衡层"
        LB[Nginx / ALB]
    end
    
    subgraph "API网关层"
        GW[API Gateway<br/>认证 / 限流 / 路由]
    end
    
    subgraph "应用服务层"
        S1[用户服务<br/>User Service]
        S2[交易服务<br/>Transaction Service]
        S3[账户服务<br/>Account Service]
        S4[账单服务<br/>Bill Service]
        S5[储蓄服务<br/>Savings Service]
        S6[通知服务<br/>Notification Service]
        S7[导出服务<br/>Export Service]
    end
    
    subgraph "数据层"
        DB[(MySQL<br/>主数据库)]
        CACHE[(Redis<br/>缓存)]
        OSS[对象存储<br/>MinIO/OSS]
    end
    
    subgraph "基础设施层"
        CRON[定时任务<br/>Cron Jobs]
        MQ[消息队列<br/>Redis]
        LOG[日志系统<br/>ELK/Loki]
        MONITOR[监控告警<br/>Prometheus]
    end
    
    A1 --> LB
    A2 --> LB
    A3 --> LB
    LB --> GW
    GW --> S1
    GW --> S2
    GW --> S3
    GW --> S4
    GW --> S5
    GW --> S6
    GW --> S7
    
    S1 --> DB
    S2 --> DB
    S3 --> DB
    S4 --> DB
    S5 --> DB
    S6 --> DB
    S7 --> DB
    
    S1 --> CACHE
    S2 --> CACHE
    S3 --> CACHE
    S6 --> CACHE
    
    S1 --> OSS
    S2 --> OSS
    S7 --> OSS
    
    CRON --> S4
    CRON --> S6
    S6 --> MQ
```

#### 2.2 分层架构

```
┌─────────────────────────────────────────┐
│           API Layer (路由层)              │
│  - HTTP 路由定义                          │
│  - 请求验证                                │
│  - 响应格式化                              │
└─────────────────────────────────────────┘
                  ↓
┌─────────────────────────────────────────┐
│        Middleware Layer (中间件层)        │
│  - 认证鉴权 (JWT)                         │
│  - 日志记录                                │
│  - 错误处理                                │
│  - 限流熔断                                │
│  - CORS                                   │
└─────────────────────────────────────────┘
                  ↓
┌─────────────────────────────────────────┐
│        Service Layer (业务逻辑层)         │
│  - 业务规则实现                            │
│  - 数据组装                                │
│  - 事务管理                                │
└─────────────────────────────────────────┘
                  ↓
┌─────────────────────────────────────────┐
│      Repository Layer (数据访问层)        │
│  - CRUD 操作封装                          │
│  - 查询构建                                │
│  - 缓存策略                                │
└─────────────────────────────────────────┘
                  ↓
┌─────────────────────────────────────────┐
│         Model Layer (数据模型层)          │
│  - 数据库模型                              │
│  - DTO 对象                               │
│  - 数据验证                                │
└─────────────────────────────────────────┘
```

---

### 3. 目录结构设计

推荐的 Go 项目目录结构（遵循 Go 标准布局）：

```
float-backend/
├── cmd/
│   └── server/
│       └── main.go                 # 应用入口
├── internal/
│   ├── api/
│   │   ├── handlers/              # HTTP 处理函数
│   │   │   ├── user.go
│   │   │   ├── transaction.go
│   │   │   ├── account.go
│   │   │   ├── bill.go
│   │   │   ├── savings.go
│   │   │   ├── wishlist.go
│   │   │   └── export.go
│   │   ├── middlewares/           # 中间件
│   │   │   ├── auth.go
│   │   │   ├── logger.go
│   │   │   ├── cors.go
│   │   │   └── ratelimit.go
│   │   └── routes/                # 路由定义
│   │       └── router.go
│   ├── service/                   # 业务逻辑层
│   │   ├── user_service.go
│   │   ├── transaction_service.go
│   │   ├── account_service.go
│   │   ├── bill_service.go
│   │   ├── savings_service.go
│   │   ├── wishlist_service.go
│   │   ├── notification_service.go
│   │   └── export_service.go
│   ├── repository/                # 数据访问层
│   │   ├── user_repo.go
│   │   ├── transaction_repo.go
│   │   ├── account_repo.go
│   │   ├── bill_repo.go
│   │   └── ...
│   ├── models/                    # 数据模型
│   │   ├── user.go
│   │   ├── transaction.go
│   │   ├── account.go
│   │   └── ...
│   ├── dto/                       # 数据传输对象
│   │   ├── request/
│   │   │   ├── user_request.go
│   │   │   └── transaction_request.go
│   │   └── response/
│   │       ├── user_response.go
│   │       └── transaction_response.go
│   └── utils/                     # 工具函数
│       ├── jwt.go
│       ├── password.go
│       ├── validator.go
│       └── response.go
├── pkg/                           # 公共库（可被外部引用）
│   ├── cache/
│   │   └── redis.go
│   ├── database/
│   │   └── mysql.go
│   ├── storage/
│   │   └── oss.go
│   └── logger/
│       └── logger.go
├── config/                        # 配置文件
│   ├── config.yaml
│   ├── config.dev.yaml
│   └── config.prod.yaml
├── migrations/                    # 数据库迁移文件
│   ├── 000001_create_users_table.up.sql
│   ├── 000001_create_users_table.down.sql
│   └── ...
├── scripts/                       # 脚本文件
│   ├── init_db.sh
│   └── seed_data.sql
├── docs/                          # 文档
│   ├── api/                       # API 文档
│   └── architecture.md
├── tests/                         # 测试
│   ├── unit/
│   └── integration/
├── deployments/                   # 部署配置
│   └── docker/
│       ├── Dockerfile
│       ├── docker-compose.yml
│       ├── docker-compose.prod.yml
│       └── .env.example
├── .env.example                   # 环境变量示例
├── .gitignore
├── go.mod
├── go.sum
├── Makefile                       # 构建脚本
└── README.md
```

---

### 4. API 设计规范

#### 4.1 RESTful API 设计原则

| 资源 | GET (查询) | POST (创建) | PUT (更新) | DELETE (删除) |
|------|-----------|------------|-----------|--------------|
| /users | 获取用户列表 | 创建用户 | 更新用户 | 删除用户 |
| /transactions | 获取交易列表 | 创建交易 | 更新交易 | 删除交易 |
| /accounts | 获取账户列表 | 创建账户 | 更新账户 | 删除账户 |
| /bills | 获取账单列表 | 创建账单 | 更新账单 | 删除账单 |
| /savings | 获取储蓄计划 | 创建计划 | 更新计划 | 删除计划 |
| /wishlists | 获取心愿单 | 创建心愿 | 更新心愿 | 删除心愿 |

#### 4.2 API 版本控制

采用 URL 路径版本化：

```
/api/v1/users
/api/v1/transactions
/api/v2/transactions  # 新版本API
```

#### 4.3 统一响应格式

**成功响应：**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "username": "john_doe"
  }
}
```

**分页响应：**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "items": [...],
    "pagination": {
      "page": 1,
      "page_size": 20,
      "total": 100,
      "total_pages": 5
    }
  }
}
```

**错误响应：**

```json
{
  "code": 400,
  "message": "Invalid request parameters",
  "errors": [
    {
      "field": "email",
      "message": "Invalid email format"
    }
  ]
}
```

#### 4.4 核心 API 端点设计

##### 用户认证模块

```
POST   /api/v1/auth/register           # 用户注册
POST   /api/v1/auth/login              # 用户登录
POST   /api/v1/auth/logout             # 用户登出
POST   /api/v1/auth/refresh            # 刷新 Token
POST   /api/v1/auth/forgot-password    # 忘记密码
POST   /api/v1/auth/reset-password     # 重置密码
```

##### 用户管理模块

```
GET    /api/v1/users/me                # 获取当前用户信息
PUT    /api/v1/users/me                # 更新用户信息
PUT    /api/v1/users/me/avatar         # 更新头像
PUT    /api/v1/users/me/password       # 修改密码
GET    /api/v1/users/me/stats          # 用户统计信息
```

##### 交易记录模块

```
GET    /api/v1/transactions            # 获取交易列表
POST   /api/v1/transactions            # 创建交易
GET    /api/v1/transactions/:id        # 获取交易详情
PUT    /api/v1/transactions/:id        # 更新交易
DELETE /api/v1/transactions/:id        # 删除交易
GET    /api/v1/transactions/stats      # 交易统计
GET    /api/v1/transactions/summary    # 月度汇总
```

##### 账户管理模块

```
GET    /api/v1/accounts                # 获取账户列表
POST   /api/v1/accounts                # 创建账户
GET    /api/v1/accounts/:id            # 获取账户详情
PUT    /api/v1/accounts/:id            # 更新账户
DELETE /api/v1/accounts/:id            # 删除账户
GET    /api/v1/accounts/balance        # 账户余额汇总
```

##### 信用账户模块

```
GET    /api/v1/credit-accounts         # 获取信用账户列表
POST   /api/v1/credit-accounts         # 创建信用账户
GET    /api/v1/credit-accounts/:id     # 获取信用账户详情
PUT    /api/v1/credit-accounts/:id     # 更新信用账户
DELETE /api/v1/credit-accounts/:id     # 删除信用账户
```

##### 分类管理模块

```
GET    /api/v1/categories              # 获取分类列表
POST   /api/v1/categories              # 创建自定义分类
PUT    /api/v1/categories/:id          # 更新分类
DELETE /api/v1/categories/:id          # 删除分类
GET    /api/v1/categories/system       # 获取系统默认分类
```

##### 账单订阅模块

```
GET    /api/v1/bills                   # 获取账单列表
POST   /api/v1/bills                   # 创建账单
GET    /api/v1/bills/:id               # 获取账单详情
PUT    /api/v1/bills/:id               # 更新账单
DELETE /api/v1/bills/:id               # 删除账单
GET    /api/v1/bills/:id/history       # 获取账单历史
GET    /api/v1/bills/upcoming          # 获取即将到期的账单
```

##### 储蓄计划模块

```
GET    /api/v1/savings                 # 获取储蓄计划列表
POST   /api/v1/savings                 # 创建储蓄计划
GET    /api/v1/savings/:id             # 获取计划详情
PUT    /api/v1/savings/:id             # 更新计划
DELETE /api/v1/savings/:id             # 删除计划
POST   /api/v1/savings/:id/deposit     # 添加存款记录
GET    /api/v1/savings/:id/records     # 获取存款记录
```

##### 心愿单模块

```
GET    /api/v1/wishlists               # 获取心愿单列表
POST   /api/v1/wishlists               # 创建心愿
GET    /api/v1/wishlists/:id           # 获取心愿详情
PUT    /api/v1/wishlists/:id           # 更新心愿
DELETE /api/v1/wishlists/:id           # 删除心愿
POST   /api/v1/wishlists/:id/deposit   # 添加储蓄
GET    /api/v1/wishlists/:id/deposits  # 获取储蓄记录
```

##### 预算管理模块

```
GET    /api/v1/budgets                 # 获取预算列表
POST   /api/v1/budgets                 # 创建预算
GET    /api/v1/budgets/:id             # 获取预算详情
PUT    /api/v1/budgets/:id             # 更新预算
DELETE /api/v1/budgets/:id             # 删除预算
GET    /api/v1/budgets/current         # 获取当前月预算
```

##### 通知模块

```
GET    /api/v1/notifications           # 获取通知列表
PUT    /api/v1/notifications/:id/read  # 标记已读
DELETE /api/v1/notifications/:id       # 删除通知
PUT    /api/v1/notifications/read-all  # 全部已读
GET    /api/v1/notifications/unread    # 获取未读数量
```

##### 数据导出模块

```
POST   /api/v1/export                  # 创建导出任务
GET    /api/v1/export/:id              # 获取导出任务状态
GET    /api/v1/export/:id/download     # 下载文件
GET    /api/v1/export/history          # 导出历史
```

##### 软件更新模块

```
GET    /api/v1/app-updates/check       # 检查更新
GET    /api/v1/app-updates/latest      # 获取最新版本
GET    /api/v1/app-updates/:id         # 获取版本详情
GET    /api/v1/app-updates/history     # 更新历史
```

---

### 5. 认证与授权设计

#### 5.1 认证方案：JWT (JSON Web Token)

**Token 结构：**

- **Access Token**: 短期有效（15分钟），用于API访问
- **Refresh Token**: 长期有效（7天），用于刷新 Access Token

**Token Payload 示例：**

```json
{
  "user_id": 123,
  "username": "john_doe",
  "email": "john@example.com",
  "exp": 1733456789,
  "iat": 1733455889
}
```

#### 5.2 认证流程

```mermaid
sequenceDiagram
    participant C as 客户端
    participant API as API服务
    participant Redis as Redis
    
    C->>API: POST /auth/login (username, password)
    API->>API: 验证密码
    API->>Redis: 存储 Refresh Token
    API->>C: 返回 Access Token + Refresh Token
    
    C->>API: GET /users/me (携带 Access Token)
    API->>API: 验证 Access Token
    API->>C: 返回用户信息
    
    Note over C,API: Access Token 过期
    
    C->>API: POST /auth/refresh (Refresh Token)
    API->>Redis: 验证 Refresh Token
    API->>API: 生成新的 Access Token
    API->>C: 返回新的 Access Token
```

#### 5.3 权限控制

目前为个人应用，暂时不需要复杂的 RBAC (角色权限控制)，但需要确保：

- 用户只能访问自己的数据
- 所有 API 除了登录/注册外都需要认证
- 在 Repository 层强制添加 `user_id` 过滤

**权限中间件示例逻辑：**

```go
// 伪代码
func AuthMiddleware(c *gin.Context) {
    token := c.GetHeader("Authorization")
    claims, err := ValidateToken(token)
    if err != nil {
        c.JSON(401, "Unauthorized")
        return
    }
    c.Set("user_id", claims.UserID)
    c.Next()
}
```

---

### 6. 核心业务逻辑设计

#### 6.1 交易记录处理逻辑

**创建交易时的业务规则：**

1. 验证账户归属（account_id 必须属于当前用户）
2. 验证分类归属（category_id 必须属于当前用户或为系统分类）
3. 根据交易类型更新账户余额：
   - `expense`: 减少账户余额
   - `income`: 增加账户余额
   - `transfer`: 转出账户减少，转入账户增加
4. 使用**数据库事务**确保数据一致性
5. 更新用户统计（总记录数、连续记账天数）
6. 如果关联账单或心愿单，更新相应记录

**伪代码示例：**

```go
func (s *TransactionService) CreateTransaction(req *dto.CreateTransactionRequest, userID int64) error {
    // 开启事务
    tx := s.db.Begin()
    defer tx.Rollback()
    
    // 1. 验证账户归属
    account, err := s.accountRepo.FindByIDAndUser(req.AccountID, userID)
    if err != nil {
        return errors.New("账户不存在")
    }
    
    // 2. 创建交易记录
    transaction := &models.Transaction{
        UserID: userID,
        Type: req.Type,
        Amount: req.Amount,
        // ... 其他字段
    }
    if err := s.transactionRepo.Create(tx, transaction); err != nil {
        return err
    }
    
    // 3. 更新账户余额
    if req.Type == "expense" {
        account.Balance -= req.Amount
    } else if req.Type == "income" {
        account.Balance += req.Amount
    }
    if err := s.accountRepo.Update(tx, account); err != nil {
        return err
    }
    
    // 4. 提交事务
    return tx.Commit().Error
}
```

#### 6.2 账单提醒定时任务

**Cron 任务设计：**

- 每天凌晨 00:00 执行
- 查询 `next_billing_date` 在未来 3 天内的账单
- 生成通知记录
- 发送推送通知

```go
func (job *BillReminderJob) Run() {
    // 1. 查询即将到期的账单
    bills := billRepo.FindUpcoming(3) // 3天内
    
    // 2. 创建通知
    for _, bill := range bills {
        notification := &models.Notification{
            UserID: bill.UserID,
            Type: "bill_reminder",
            Title: fmt.Sprintf("账单提醒：%s", bill.BillName),
            Content: fmt.Sprintf("您的%s将在%s扣费", bill.BillName, bill.NextBillingDate),
        }
        notificationRepo.Create(notification)
        
        // 3. 发送推送
        pushService.Send(bill.UserID, notification)
    }
}
```

#### 6.3 预算预警逻辑

**实时监控预算使用情况：**

- 每次创建支出交易后触发
- 计算当前月/周期的支出总额
- 与预算金额对比
- 超过阈值（默认80%）时创建通知

```go
func (s *TransactionService) checkBudgetAlert(tx *gorm.DB, transaction *models.Transaction) {
    if transaction.Type != "expense" {
        return
    }
    
    // 查询相关预算
    budgets := s.budgetRepo.FindActiveByCategory(transaction.UserID, transaction.CategoryID)
    
    for _, budget := range budgets {
        usage := budget.SpentAmount / budget.BudgetAmount
        if usage >= budget.AlertThreshold {
            // 创建预警通知
            s.notificationService.CreateBudgetAlert(budget)
        }
    }
}
```

#### 6.4 数据导出逻辑

**异步任务处理：**

1. 用户发起导出请求，创建 `export_jobs` 记录，状态为 `pending`
2. 后台任务队列（使用 Redis）处理导出任务
3. 根据导出格式生成文件（Excel/PDF/CSV）
4. 上传到对象存储
5. 更新任务状态为 `completed`，填充 `file_url`
6. 发送邮件通知用户或推送消息

```go
func (worker *ExportWorker) Process(jobID int64) {
    // 1. 更新状态为 processing
    job := exportRepo.FindByID(jobID)
    job.Status = "processing"
    exportRepo.Update(job)
    
    // 2. 查询数据
    transactions := transactionRepo.FindByDateRange(job.UserID, job.StartDate, job.EndDate)
    
    // 3. 生成文件
    var fileData []byte
    switch job.ExportType {
    case "excel":
        fileData = generateExcel(transactions)
    case "pdf":
        fileData = generatePDF(transactions)
    case "csv":
        fileData = generateCSV(transactions)
    }
    
    // 4. 上传到对象存储
    url, err := storageService.Upload(fileData, "exports/...")
    
    // 5. 更新任务
    job.Status = "completed"
    job.FileURL = url
    job.FileSize = len(fileData)
    exportRepo.Update(job)
    
    // 6. 发送通知
    notificationService.SendExportComplete(job.UserID, url)
}
```

---

### 7. 数据库设计补充

#### 7.1 数据库迁移策略

使用 [golang-migrate/migrate](https://github.com/golang-migrate/migrate) 管理数据库迁移：

```bash
migrate create -ext sql -dir migrations -seq create_users_table
```

**迁移文件示例：**

`000001_create_users_table.up.sql`:
```sql
CREATE TABLE users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    -- ...
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

`000001_create_users_table.down.sql`:
```sql
DROP TABLE IF EXISTS users;
```

#### 7.2 MySQL 数据库索引优化

在[数据库设计文档](file:///Users/qiuhaonan/Developer/projects/Float/design/db/db.md)基础上，额外建议的性能优化索引：

```sql
-- 交易记录全文搜索索引（MySQL 8.0+）
CREATE FULLTEXT INDEX idx_transactions_search 
ON transactions(title, description) WITH PARSER ngram;

-- 分类统计优化
CREATE INDEX idx_transactions_category_date 
ON transactions(category_id, transaction_date);

-- 账单到期查询优化
CREATE INDEX idx_bills_next_billing 
ON bills(next_billing_date, is_active);

-- 用户日期复合索引
CREATE INDEX idx_user_date_type 
ON transactions(user_id, transaction_date, type);
```

---

### 8. 安全设计

#### 8.1 安全措施

| 威胁 | 防护措施 |
|------|---------|
| SQL注入 | 使用 ORM 参数化查询 |
| XSS攻击 | 前端输入验证 + 后端HTML转义 |
| CSRF | 使用 JWT，无需 CSRF Token |
| 暴力破解 | 登录失败限制（Redis计数）+ 验证码 |
| 敏感数据泄露 | HTTPS 加密传输 + 数据库加密存储 |
| 未授权访问 | JWT 认证 + 用户数据隔离 |

#### 8.2 密码安全

```go
import "golang.org/x/crypto/bcrypt"

// 密码哈希
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// 密码验证
func CheckPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

#### 8.3 API 限流

使用中间件实现基于 IP/用户的限流：

```go
// 基于 Redis 的限流中间件
func RateLimitMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := c.GetInt64("user_id")
        key := fmt.Sprintf("rate_limit:user:%d", userID)
        
        count, _ := redisClient.Incr(key).Result()
        if count == 1 {
            redisClient.Expire(key, time.Minute)
        }
        
        if count > 100 { // 每分钟100次请求
            c.JSON(429, gin.H{"error": "Too many requests"})
            c.Abort()
            return
        }
        
        c.Next()
    }
}
```

---

### 9. 部署与运维

#### 9.1 容器化部署

**Dockerfile 示例：**

```dockerfile
# 构建阶段
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# 运行阶段
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/
COPY --from=builder /app/server .
COPY --from=builder /app/config ./config
EXPOSE 8080
CMD ["./server"]
```

**docker-compose.yml:**

```yaml
version: '3.8'

services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DB_HOST=mysql
      - REDIS_HOST=redis
    depends_on:
      - mysql
      - redis
    restart: unless-stopped
  
  mysql:
    image: mysql:8.0
    command: --default-authentication-plugin=mysql_native_password
    environment:
      MYSQL_ROOT_PASSWORD: ${MYSQL_ROOT_PASSWORD}
      MYSQL_DATABASE: float_db
      MYSQL_USER: float_user
      MYSQL_PASSWORD: ${DB_PASSWORD}
    volumes:
      - mysql_data:/var/lib/mysql
    restart: unless-stopped
  
  redis:
    image: redis:7-alpine
    restart: unless-stopped
  
  minio:
    image: minio/minio
    command: server /data --console-address ":9001"
    environment:
      MINIO_ROOT_USER: minioadmin
      MINIO_ROOT_PASSWORD: ${MINIO_PASSWORD}
    volumes:
      - minio_data:/data
    restart: unless-stopped

volumes:
  mysql_data:
  minio_data:
```

#### 9.2 Kubernetes 部署（可选）

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: float-backend
spec:
  replicas: 3
  selector:
    matchLabels:
      app: float-backend
  template:
    metadata:
      labels:
        app: float-backend
    spec:
      containers:
      - name: api
        image: float-backend:latest
        ports:
        - containerPort: 8080
        env:
        - name: DB_HOST
          value: "mysql-service"
        - name: REDIS_HOST
          value: "redis-service"
        resources:
          requests:
            memory: "128Mi"
            cpu: "100m"
          limits:
            memory: "512Mi"
            cpu: "500m"
```

#### 9.3 监控与日志

**监控指标（Prometheus）：**

- API 请求数、延迟、错误率
- 数据库连接池状态
- Redis 缓存命中率
- 业务指标（活跃用户数、交易数等）

**日志收集（ELK / Loki）：**

- 结构化日志输出（JSON 格式）
- 日志级别：DEBUG, INFO, WARN, ERROR
- 日志轮转和归档

**健康检查端点：**

```go
r.GET("/health", func(c *gin.Context) {
    // 检查数据库连接
    if err := db.DB().Ping(); err != nil {
        c.JSON(503, gin.H{"status": "unhealthy", "database": "down"})
        return
    }
    
    // 检查 Redis 连接
    if err := redisClient.Ping().Err(); err != nil {
        c.JSON(503, gin.H{"status": "unhealthy", "redis": "down"})
        return
    }
    
    c.JSON(200, gin.H{"status": "healthy"})
})
```

---

### 10. 性能优化策略

#### 10.1 数据库优化

1. **连接池配置：**
   ```go
   db.SetMaxOpenConns(25)
   db.SetMaxIdleConns(5)
   db.SetConnMaxLifetime(5 * time.Minute)
   ```

2. **N+1 查询优化 - 使用预加载：**
   ```go
   db.Preload("Category").Preload("Account").Find(&transactions)
   ```

3. **批量插入：**
   ```go
   db.CreateInBatches(transactions, 100)
   ```

#### 10.2 缓存策略

**多层缓存架构：**

1. **应用层缓存**（本地缓存，如 `go-cache`）
   - 系统分类列表
   - 配置信息

2. **Redis 缓存**
   - 用户信息（TTL: 1小时）
   - 账户余额（TTL: 5分钟）
   - 月度统计（TTL: 30分钟）

**缓存更新策略：**

- 写操作同时更新缓存（Write-Through）
- 缓存失效时重新加载（Cache-Aside）

```go
func (s *AccountService) GetBalance(userID int64) (float64, error) {
    // 1. 尝试从缓存获取
    cacheKey := fmt.Sprintf("account:balance:%d", userID)
    if val, err := cache.Get(cacheKey); err == nil {
        return val.(float64), nil
    }
    
    // 2. 从数据库查询
    balance, err := s.accountRepo.GetTotalBalance(userID)
    if err != nil {
        return 0, err
    }
    
    // 3. 写入缓存
    cache.Set(cacheKey, balance, 5*time.Minute)
    
    return balance, nil
}
```

#### 10.3 异步处理

使用 Redis 作为消息队列，处理耗时任务：

- 数据导出
- 批量通知发送
- 统计数据预计算

---

### 11. 测试策略

#### 11.1 单元测试

使用 `testify` 框架编写单元测试：

```go
func TestTransactionService_CreateTransaction(t *testing.T) {
    // Setup
    mockRepo := new(MockTransactionRepo)
    service := NewTransactionService(mockRepo)
    
    // Given
    req := &dto.CreateTransactionRequest{
        Type: "expense",
        Amount: 100.00,
    }
    
    // When
    err := service.CreateTransaction(req, 1)
    
    // Then
    assert.NoError(t, err)
    mockRepo.AssertCalled(t, "Create")
}
```

#### 11.2 集成测试

测试 API 端到端流程：

```go
func TestTransactionAPI_Integration(t *testing.T) {
    // Setup test database
    db := setupTestDB()
    router := setupRouter(db)
    
    // Create test user and login
    token := loginTestUser(router)
    
    // Test create transaction
    w := httptest.NewRecorder()
    body := `{"type":"expense","amount":100}`
    req, _ := http.NewRequest("POST", "/api/v1/transactions", strings.NewReader(body))
    req.Header.Set("Authorization", "Bearer "+token)
    router.ServeHTTP(w, req)
    
    assert.Equal(t, 200, w.Code)
}
```

#### 11.3 性能测试

使用工具进行压力测试：
- **wrk** - HTTP 基准测试
- **Apache JMeter** - 复杂场景测试

```bash
wrk -t12 -c400 -d30s http://localhost:8080/api/v1/transactions
```

---

## Verification Plan

### Automated Tests

1. **单元测试**
   ```bash
   go test ./internal/service/... -v -cover
   ```

2. **API 集成测试**
   ```bash
   go test ./tests/integration/... -v
   ```

### Manual Verification

需要用户协助验证的部分：

1. **技术栈选型确认**
   - 是否认可 Go + Gin 作为后端框架？
   - 是否可以使用 PostgreSQL 作为主数据库？
   - 是否需要支持云部署（阿里云/腾讯云/AWS）？

2. **API 设计评审**
   - 查看 API 端点设计章节，确认接口设计是否符合前端需求
   - 是否有遗漏的业务功能？

3. **部署方案确认**
   - 预期的部署方式（Docker / Kubernetes / 云主机）
   - 是否需要 CI/CD 自动化部署？

---

## 附录

### A. 开发环境搭建

```bash
# 1. 安装 Go 1.21+
brew install go

# 2. 初始化项目
mkdir float-backend && cd float-backend
go mod init github.com/yourusername/float-backend

# 3. 安装依赖
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/golang-jwt/jwt/v5
go get -u github.com/go-redis/redis/v8
go get -u github.com/spf13/viper

# 4. 启动本地数据库（Docker）
docker-compose up -d mysql redis

# 5. 运行迁移
migrate -path migrations -database "mysql://float_user:password@tcp(localhost:3306)/float_db" up

# 6. 启动服务
go run cmd/server/main.go
```

### B. 推荐的开发工具

- **IDE**: GoLand / VS Code (with Go extension)
- **API 测试**: Postman / Insomnia
- **数据库管理**: DBeaver / TablePlus
- **Git 客户端**: SourceTree / GitKraken

### C. 参考资料

- [Gin Web Framework](https://gin-gonic.com/)
- [GORM 文档](https://gorm.io/)
- [Go 标准项目布局](https://github.com/golang-standards/project-layout)
- [JWT 最佳实践](https://jwt.io/introduction)

---

**文档版本**: v1.0  
**创建日期**: 2025-12-05  
**作者**: Antigravity AI
