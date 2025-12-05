# 用户认证和管理 API 测试指南

## 环境准备

确保已启动服务：
```bash
cd backend
make run
```

服务地址：`http://localhost:8080`

## 1. 用户注册

**POST** `/api/v1/auth/register`

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "Test123456",
    "display_name": "Test User"
  }'
```

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "user_id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "display_name": "Test User",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 900
  }
}
```

## 2. 用户登录

**POST** `/api/v1/auth/login`

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "Test123456"
  }'
```

## 3. 刷新Token

**POST** `/api/v1/auth/refresh`

```bash
curl -X POST http://localhost:8080/api/v1/auth/refresh \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "你的refresh_token"
  }'
```

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 900
  }
}
```

## 4. 获取当前用户信息

**GET** `/api/v1/users/me`

```bash
curl -X GET http://localhost:8080/api/v1/users/me \
  -H "Authorization: Bearer 你的access_token"
```

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "display_name": "Test User",
    "currency": "CNY",
    "theme": "light",
    "language": "zh-CN",
    "dark_mode": false,
    "gesture_lock": true,
    "continuous_days": 0,
    "total_records": 0,
    "verified": false,
    "membership_level": "FREE",
    "created_at": "2025-12-05T10:00:00Z"
  }
}
```

## 5. 更新用户信息

**PUT** `/api/v1/users/me`

```bash
curl -X PUT http://localhost:8080/api/v1/users/me \
  -H "Authorization: Bearer 你的access_token" \
  -H "Content-Type: application/json" \
  -d '{
    "display_name": "新昵称",
    "dark_mode": true,
    "language": "en-US"
  }'
```

## 6. 修改密码

**PUT** `/api/v1/users/me/password`

```bash
curl -X PUT http://localhost:8080/api/v1/users/me/password \
  -H "Authorization: Bearer 你的access_token" \
  -H "Content-Type: application/json" \
  -d '{
    "old_password": "Test123456",
    "new_password": "NewPass123456"
  }'
```

## 7. 获取用户统计

**GET** `/api/v1/users/me/stats`

```bash
curl -X GET http://localhost:8080/api/v1/users/me/stats \
  -H "Authorization: Bearer 你的access_token"
```

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "total_assets": 0,
    "total_debt": 0,
    "net_worth": 0,
    "total_records": 0,
    "continuous_days": 0,
    "month_income": 0,
    "month_expense": 0,
    "month_net": 0,
    "active_budgets": 0,
    "active_bills": 0,
    "active_savings": 0,
    "active_wishlists": 0
  }
}
```

## 常见错误

### 401 Unauthorized
- Token 无效或已过期
- 未提供 Authorization header
- Token 格式错误

```json
{
  "code": 401,
  "message": "token无效或已过期"
}
```

### 400 Bad Request
- 请求参数验证失败
- 用户名或邮箱已存在

```json
{
  "code": 400,
  "message": "用户名已存在"
}
```

## 测试检查清单

- [ ] 注册新用户成功
- [ ] 重复用户名注册失败
- [ ] 重复邮箱注册失败
- [ ] 登录成功并获取Token
- [ ] 错误密码登录失败
- [ ] 使用Token访问受保护接口
- [ ] 刷新Token成功
- [ ] 更新用户信息成功
- [ ] 修改密码成功
- [ ] 获取用户统计数据

## 下一步

完成用户模块后，可以继续实现：
1. 交易记录模块（Transaction）
2. 账户管理模块（Account）
3. 分类管理模块（Category）
4. 账单订阅模块（Bill）
5. 储蓄计划模块（Saving）
