# 账户管理 API 测试指南

## 环境准备

确保已启动服务：
```bash
cd backend
make run
```

服务地址：`http://localhost:8080`

**注意**：所有账户管理接口都需要认证，请先登录获取 access_token。

---

## 1. 获取账户列表

**GET** `/api/v1/accounts`

获取用户的所有资金账户。

```bash
curl -X GET "http://localhost:8080/api/v1/accounts" \
  -H "Authorization: Bearer 你的access_token"
```

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "account_type": "bank",
      "account_name": "招商银行",
      "account_number": "8888",
      "icon": "fa-credit-card",
      "color": "red",
      "balance": 10000.00,
      "initial_balance": 10000.00,
      "include_in_total": true,
      "display_order": 0,
      "is_active": true,
      "created_at": "2025-12-05T10:00:00Z",
      "updated_at": "2025-12-05T10:00:00Z"
    },
    {
      "id": 2,
      "account_type": "alipay",
      "account_name": "支付宝",
      "account_number": "test@alipay.com",
      "icon": "fa-alipay",
      "color": "blue",
      "balance": 500.50,
      "initial_balance": 500.50,
      "include_in_total": true,
      "display_order": 1,
      "is_active": true,
      "created_at": "2025-12-05T10:05:00Z",
      "updated_at": "2025-12-05T10:05:00Z"
    }
  ]
}
```

---

## 2. 获取账户详情

**GET** `/api/v1/accounts/:id`

```bash
curl -X GET "http://localhost:8080/api/v1/accounts/1" \
  -H "Authorization: Bearer 你的access_token"
```

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "account_type": "bank",
    "account_name": "招商银行",
    "account_number": "8888",
    "icon": "fa-credit-card",
    "color": "red",
    "balance": 10000.00,
    "initial_balance": 10000.00,
    "include_in_total": true,
    "display_order": 0,
    "is_active": true,
    "created_at": "2025-12-05T10:00:00Z",
    "updated_at": "2025-12-05T10:00:00Z"
  }
}
```

---

## 3. 创建账户

**POST** `/api/v1/accounts`

```bash
curl -X POST "http://localhost:8080/api/v1/accounts" \
  -H "Authorization: Bearer 你的access_token" \
  -H "Content-Type: application/json" \
  -d '{
    "account_type": "bank",
    "account_name": "工商银行",
    "account_number": "6666",
    "icon": "fa-credit-card",
    "color": "orange",
    "initial_balance": 5000.00,
    "include_in_total": true
  }'
```

**请求参数：**
| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| account_type | string | 是 | 账户类型：`bank`, `alipay`, `wechat`, `cash`, `other` |
| account_name | string | 是 | 账户名称，最大100字符 |
| account_number | string | 否 | 账号/卡号后四位 |
| icon | string | 否 | FontAwesome图标代码 |
| color | string | 否 | 颜色标识 |
| initial_balance | float | 否 | 初始余额，默认0 |
| include_in_total | bool | 否 | 是否计入总资产，默认true |
| display_order | int | 否 | 显示排序，默认0 |

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 3,
    "account_type": "bank",
    "account_name": "工商银行",
    "account_number": "6666",
    "icon": "fa-credit-card",
    "color": "orange",
    "balance": 5000.00,
    "initial_balance": 5000.00,
    "include_in_total": true,
    "display_order": 0,
    "is_active": true,
    "created_at": "2025-12-05T11:00:00Z",
    "updated_at": "2025-12-05T11:00:00Z"
  }
}
```

---

## 4. 更新账户

**PUT** `/api/v1/accounts/:id`

```bash
curl -X PUT "http://localhost:8080/api/v1/accounts/3" \
  -H "Authorization: Bearer 你的access_token" \
  -H "Content-Type: application/json" \
  -d '{
    "account_name": "工商银行储蓄卡",
    "color": "red"
  }'
```

**请求参数：**
| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| account_name | string | 否 | 账户名称 |
| account_number | string | 否 | 账号 |
| icon | string | 否 | 图标 |
| color | string | 否 | 颜色 |
| include_in_total | bool | 否 | 是否计入总资产 |
| display_order | int | 否 | 显示排序 |
| is_active | bool | 否 | 是否启用 |

**注意**：此接口不用于更新余额，余额变动应通过交易记录产生。

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

---

## 5. 删除账户

**DELETE** `/api/v1/accounts/:id`

```bash
curl -X DELETE "http://localhost:8080/api/v1/accounts/3" \
  -H "Authorization: Bearer 你的access_token"
```

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

---

## 6. 获取账户余额汇总

**GET** `/api/v1/accounts/balance`

获取所有有效账户的总资产（仅统计 `include_in_total=true` 的账户）。

```bash
curl -X GET "http://localhost:8080/api/v1/accounts/balance" \
  -H "Authorization: Bearer 你的access_token"
```

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "total_balance": 15500.50,
    "asset_balance": 15500.50,
    "debt_balance": 0
  }
}
```

---

## 账户类型说明

| 类型代码 | 说明 | 推荐图标 |
|----------|------|----------|
| bank | 银行卡 | fa-credit-card |
| alipay | 支付宝 | fa-alipay |
| wechat | 微信钱包 | fa-weixin |
| cash | 现金 | fa-money-bill |
| other | 其他 | fa-wallet |

---

## 常见错误码

| 状态码 | 说明 |
|--------|------|
| 200 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未登录或token无效 |
| 403 | 权限不足（操作他人账户） |
| 404 | 账户不存在 |
| 500 | 服务器内部错误 |
