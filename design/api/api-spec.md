# FloatIsland API 详细设计文档

**版本**: v1.0  
**更新日期**: 2025-12-05  
**维护者**: FloatIsland Team

本文档提供 FloatIsland 后端 API 的详细说明，包含完整的请求/响应示例，供前端开发对接使用。

---

## 目录

1. [基础说明](#基础说明)
2. [认证模块](#认证模块)
3. [用户管理](#用户管理)
4. [交易记录](#交易记录)
5. [账户管理](#账户管理)
6. [信用账户](#信用账户)
7. [分类管理](#分类管理)
8. [账单订阅](#账单订阅)
9. [储蓄计划](#储蓄计划)
10. [心愿单](#心愿单)
11. [预算管理](#预算管理)
12. [通知](#通知)
13. [数据导出](#数据导出)
14. [软件更新](#软件更新)

---

## 基础说明

### API 基础 URL

- 开发环境: `http://localhost:8080/api/v1`
- 测试环境: `https://test-api.floatisland.com/api/v1`
- 生产环境: `https://api.floatisland.com/api/v1`

### 认证方式

除登录/注册接口外，所有接口都需要在 Header 中携带 Access Token：

```
Authorization: Bearer {access_token}
```

### 统一响应格式

**成功响应：**
```json
{
  "code": 200,
  "message": "success",
  "data": {...}
}
```

**错误响应：**
```json
{
  "code": 400,
  "message": "错误描述",
  "errors": [
    {
      "field": "字段名",
      "message": "错误详情"
    }
  ]
}
```

### HTTP 状态码

| 状态码 | 说明 |
|-------|------|
| 200 | 请求成功 |
| 201 | 创建成功 |
| 400 | 参数错误 |
| 401 | 未认证 |
| 403 | 无权限 |
| 404 | 资源不存在 |
| 429 | 请求过于频繁 |
| 500 | 服务器错误 |

---

## 认证模块

### POST /auth/register

用户注册

**请求体：**
```json
{
  "username": "john_doe",
  "email": "john@example.com",
  "phone": "+8613800138000",
  "password": "SecurePass123!",
  "display_name": "John Doe"
}
```

**响应：**
```json
{
  "code": 200,
  "message": "注册成功",
  "data": {
    "user_id": 123,
    "username": "john_doe",
    "email": "john@example.com",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 900
  }
}
```

### POST /auth/login

用户登录

**请求体：**
```json
{
  "username": "john_doe",
  "password": "SecurePass123!"
}
```

**响应：**
```json
{
  "code": 200,
  "message": "登录成功",
  "data": {
    "user_id": 123,
    "username": "john_doe",
    "email": "john@example.com",
    "display_name": "John Doe",
    "avatar_url": "https://oss.example.com/avatars/123.jpg",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 900
  }
}
```

### POST /auth/refresh

刷新访问令牌

**请求体：**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**响应：**
```json
{
  "code": 200,
  "message": "Token刷新成功",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 900
  }
}
```

### POST /auth/logout

用户登出

**请求体：**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**响应：**
```json
{
  "code": 200,
  "message": "登出成功"
}
```

---

## 用户管理

### GET /users/me

获取当前用户信息

**响应：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 123,
    "username": "john_doe",
    "email": "john@example.com",
    "phone": "+8613800138000",
    "display_name": "John Doe",
    "avatar_url": "https://oss.example.com/avatars/123.jpg",
    "verified": false,
    "currency": "CNY",
    "theme": "light",
    "language": "zh-CN",
    "dark_mode": false,
    "gesture_lock": true,
    "continuous_days": 15,
    "total_records": 230,
    "total_badges": 3,
    "membership_level": "FREE",
    "created_at": "2025-01-01T00:00:00Z",
    "last_login_at": "2025-12-05T09:00:00Z"
  }
}
```

### PUT /users/me

更新用户信息

**请求体：**
```json
{
  "display_name": "John D.",
  "phone": "+8613900139000",
  "currency": "CNY",
  "dark_mode": true,
  "language": "zh-CN"
}
```

**响应：**
```json
{
  "code": 200,
  "message": "用户信息更新成功",
  "data": {
    "id": 123,
    "updated_at": "2025-12-05T10:00:00Z"
  }
}
```

### GET /users/me/stats

用户统计信息

**响应：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "total_assets": 25751.30,
    "total_debt": 8500.00,
    "net_worth": 17251.30,
    "total_records": 230,
    "continuous_days": 15,
    "month_income": 15000.00,
    "month_expense": 8500.00,
    "month_net": 6500.00,
    "active_budgets": 3,
    "active_bills": 5,
    "active_savings": 2,
    "active_wishlists": 3
  }
}
```

---

## 交易记录

### GET /transactions

获取交易列表

**查询参数：**
- `page`: 页码（默认: 1）
- `page_size`: 每页数量（默认: 20，最大: 100）
- `type`: 交易类型（expense/income/transfer）
- `category_id`: 分类ID
- `account_id`: 账户ID
- `start_date`: 开始日期（YYYY-MM-DD）
- `end_date`: 结束日期（YYYY-MM-DD）
- `keyword`: 搜索关键词
- `sort`: 排序字段（transaction_date/amount/created_at）
- `order`: 排序方式（asc/desc），默认desc

**请求示例：**
```
GET /transactions?page=1&page_size=20&type=expense&start_date=2025-12-01
```

**响应：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "items": [
      {
        "id": 1,
        "type": "expense",
        "category": {
          "id": 5,
          "name": "餐饮美食",
          "icon": "fa-utensils",
          "color": "#FF6B6B"
        },
        "account": {
          "id": 2,
          "name": "招商银行",
          "icon": "fa-credit-card"
        },
        "amount": 85.50,
        "currency": "CNY",
        "title": "午餐",
        "description": "公司楼下日料店",
        "location": "北京市朝阳区",
        "transaction_date": "2025-12-05",
        "transaction_time": "12:30:00",
        "images": ["https://oss.example.com/receipts/123.jpg"],
        "tags": ["工作餐", "日料"],
        "created_at": "2025-12-05T12:35:00Z"
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 20,
      "total": 150,
      "total_pages": 8
    },
    "summary": {
      "total_income": 0,
      "total_expense": 2450.00,
      "count": 20
    }
  }
}
```

### POST /transactions

创建交易记录

**请求体：**
```json
{
  "type": "expense",
  "category_id": 5,
  "account_id": 2,
  "amount": 85.50,
  "currency": "CNY",
  "title": "午餐",
  "description": "公司楼下日料店",
  "location": "北京市朝阳区",
  "transaction_date": "2025-12-05",
  "transaction_time": "12:30:00",
  "tags": ["工作餐", "日料"],
  "images": ["https://oss.example.com/receipts/123.jpg"]
}
```

**响应：**
```json
{
  "code": 200,
  "message": "交易创建成功",
  "data": {
    "id": 1,
    "type": "expense",
    "amount": 85.50,
    "account_balance": 12414.50,
    "created_at": "2025-12-05T12:35:00Z"
  }
}
```

### GET /transactions/:id

获取交易详情

**响应：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "type": "expense",
    "category": {
      "id": 5,
      "name": "餐饮美食",
      "icon": "fa-utensils",
      "color": "#FF6B6B"
    },
    "account": {
      "id": 2,
      "name": "招商银行",
      "type": "bank",
      "balance": 12414.50
    },
    "amount": 85.50,
    "currency": "CNY",
    "title": "午餐",
    "description": "公司楼下日料店",
    "location": "北京市朝阳区",
    "transaction_date": "2025-12-05",
    "transaction_time": "12:30:00",
    "images": ["https://oss.example.com/receipts/123.jpg"],
    "tags": ["工作餐", "日料"],
    "bill_id": null,
    "wishlist_id": null,
    "created_at": "2025-12-05T12:35:00Z",
    "updated_at": "2025-12-05T12:35:00Z"
  }
}
```

### PUT /transactions/:id

更新交易

**请求体：**
```json
{
  "amount": 90.00,
  "title": "午餐（含饮料）",
  "description": "公司楼下日料店，加了一杯饮料"
}
```

**响应：**
```json
{
  "code": 200,
  "message": "交易更新成功",
  "data": {
    "id": 1,
    "updated_at": "2025-12-05T13:00:00Z"
  }
}
```

### DELETE /transactions/:id

删除交易

**响应：**
```json
{
  "code": 200,
  "message": "交易删除成功"
}
```

### GET /transactions/stats

交易统计

**查询参数：**
- `period`: 周期（day/week/month/year），默认month
- `date`: 指定日期（YYYY-MM-DD），默认当前日期

**响应：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "period": "month",
    "date": "2025-12",
    "total_income": 15000.00,
    "total_expense": 8500.00,
    "net_amount": 6500.00,
    "transaction_count": 45,
    "avg_expense": 188.89,
    "category_breakdown": [
      {
        "category_id": 5,
        "category_name": "餐饮美食",
        "category_icon": "fa-utensils",
        "category_color": "#FF6B6B",
        "amount": 2850.00,
        "percentage": 33.5,
        "count": 18
      }
    ],
    "daily_trend": [
      {"date": "2025-12-01", "income": 0, "expense": 350.00},
      {"date": "2025-12-02", "income": 0, "expense": 280.00}
    ]
  }
}
```

---

*由于完整API文档过长，其他模块（账户管理、账单、储蓄、心愿单、预算、通知、导出、软件更新）的详细设计请参考主架构文档中的API端点列表。每个端点都遵循相同的格式规范。*

### 通用规范

#### 分页

所有列表API都支持分页，响应格式为：
```json
{
  "items": [...],
  "pagination": {
    "page": 1,
    "page_size": 20,
    "total": 150,
    "total_pages": 8,
    "has_next": true,
    "has_prev": false
  }
}
```

#### 文件上传

**POST /upload**

```
Content-Type: multipart/form-data

file: [文件内容]
type: avatar/receipt/wishlist
```

响应：
```json
{
  "code": 200,
  "message": "上传成功",
  "data": {
    "url": "https://oss.example.com/uploads/20251205/abc123.jpg",
    "filename": "abc123.jpg",
    "size": 102400,
    "mime_type": "image/jpeg"
  }
}
```

---

**文档持续更新中...**
