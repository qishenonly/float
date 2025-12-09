# 交易记录 API 测试指南

## 环境准备

确保已启动服务：
```bash
cd backend
make run
```

服务地址：`http://localhost:8080`

**注意**：所有交易记录接口都需要认证，请先登录获取 access_token。

---

## 1. 获取交易列表

**GET** `/api/v1/transactions`

获取用户的交易列表，支持多维度过滤和分页。

```bash
curl -X GET "http://localhost:8080/api/v1/transactions?page=1&page_size=20" \
  -H "Authorization: Bearer 你的access_token"
```

**查询参数：**

| 参数 | 类型 | 说明 |
|------|------|------|
| type | string | 交易类型：`expense`(支出)、`income`(收入)、`transfer`(转账) |
| category_id | int | 分类ID |
| account_id | int | 账户ID |
| start_date | string | 开始日期，格式：YYYY-MM-DD |
| end_date | string | 结束日期，格式：YYYY-MM-DD |
| search_keyword | string | 搜索关键词（搜索标题、备注、地点） |
| sort_by | string | 排序字段：`date`(默认)、`amount` |
| sort_order | string | 排序顺序：`desc`(默认)、`asc` |
| page | int | 页码，默认1 |
| page_size | int | 每页数量，默认20，最大100 |

**高级查询示例：**

```bash
# 查询本月支出，按金额排序
curl -X GET "http://localhost:8080/api/v1/transactions?type=expense&start_date=2025-12-01&end_date=2025-12-31&sort_by=amount&sort_order=desc&page=1&page_size=20" \
  -H "Authorization: Bearer 你的access_token"

# 搜索包含"午餐"的交易
curl -X GET "http://localhost:8080/api/v1/transactions?search_keyword=午餐" \
  -H "Authorization: Bearer 你的access_token"

# 查询某个账户的所有交易
curl -X GET "http://localhost:8080/api/v1/transactions?account_id=1&page=1&page_size=50" \
  -H "Authorization: Bearer 你的access_token"

# 查询某个分类的所有交易
curl -X GET "http://localhost:8080/api/v1/transactions?category_id=10" \
  -H "Authorization: Bearer 你的access_token"
```

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "total": 100,
    "page": 1,
    "page_size": 20,
    "items": [
      {
        "id": 1,
        "user_id": 1,
        "type": "expense",
        "category_id": 10,
        "category": {
          "id": 10,
          "user_id": 1,
          "type": "expense",
          "name": "餐饮美食",
          "icon": "fa-utensils",
          "color": "orange",
          "display_order": 0,
          "is_system": true,
          "is_active": true,
          "created_at": "2025-12-05T10:00:00Z",
          "updated_at": "2025-12-05T10:00:00Z"
        },
        "account_id": 100,
        "account": {
          "id": 100,
          "user_id": 1,
          "account_type": "bank",
          "account_name": "招商银行",
          "account_number": "8888",
          "icon": "fa-credit-card",
          "color": "red",
          "balance": 9900.00,
          "initial_balance": 10000.00,
          "include_in_total": true,
          "display_order": 0,
          "is_active": true,
          "created_at": "2025-12-05T10:00:00Z",
          "updated_at": "2025-12-09T12:30:00Z"
        },
        "to_account_id": null,
        "amount": 100.00,
        "currency": "CNY",
        "title": "午餐消费",
        "description": "在食堂吃饭",
        "location": "公司食堂",
        "transaction_date": "2025-12-09",
        "transaction_time": "12:30:00",
        "bill_id": null,
        "wishlist_id": null,
        "tags": ["工作餐", "午餐"],
        "images": ["https://example.com/image1.jpg"],
        "created_at": "2025-12-09T12:30:00Z",
        "updated_at": "2025-12-09T12:30:00Z"
      }
    ]
  }
}
```

---

## 2. 获取交易详情

**GET** `/api/v1/transactions/:id`

获取单条交易的完整信息。

```bash
curl -X GET "http://localhost:8080/api/v1/transactions/1" \
  -H "Authorization: Bearer 你的access_token"
```

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "user_id": 1,
    "type": "expense",
    "category_id": 10,
    "category": {
      "id": 10,
      "user_id": 1,
      "type": "expense",
      "name": "餐饮美食",
      "icon": "fa-utensils",
      "color": "orange",
      "display_order": 0,
      "is_system": true,
      "is_active": true,
      "created_at": "2025-12-05T10:00:00Z",
      "updated_at": "2025-12-05T10:00:00Z"
    },
    "account_id": 100,
    "account": {
      "id": 100,
      "user_id": 1,
      "account_type": "bank",
      "account_name": "招商银行",
      "account_number": "8888",
      "icon": "fa-credit-card",
      "color": "red",
      "balance": 9900.00,
      "initial_balance": 10000.00,
      "include_in_total": true,
      "display_order": 0,
      "is_active": true,
      "created_at": "2025-12-05T10:00:00Z",
      "updated_at": "2025-12-09T12:30:00Z"
    },
    "amount": 100.00,
    "currency": "CNY",
    "title": "午餐消费",
    "description": "在食堂吃饭",
    "location": "公司食堂",
    "transaction_date": "2025-12-09",
    "transaction_time": "12:30:00",
    "tags": ["工作餐", "午餐"],
    "images": ["https://example.com/image1.jpg"],
    "created_at": "2025-12-09T12:30:00Z",
    "updated_at": "2025-12-09T12:30:00Z"
  }
}
```

---

## 3. 创建交易

**POST** `/api/v1/transactions`

创建一条新的交易记录。

```bash
curl -X POST "http://localhost:8080/api/v1/transactions" \
  -H "Authorization: Bearer 你的access_token" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "expense",
    "category_id": 10,
    "account_id": 100,
    "amount": 100.00,
    "currency": "CNY",
    "title": "午餐消费",
    "description": "在食堂吃饭",
    "location": "公司食堂",
    "transaction_date": "2025-12-09",
    "transaction_time": "12:30:00",
    "tags": ["工作餐", "午餐"],
    "images": ["https://example.com/image1.jpg"]
  }'
```

**请求参数：**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| type | string | 是 | 交易类型：`expense`(支出)、`income`(收入)、`transfer`(转账) |
| amount | float | 是 | 交易金额，必须 > 0 |
| transaction_date | string | 是 | 交易日期，格式：YYYY-MM-DD |
| category_id | int | 否 | 分类ID（支出/收入建议填写） |
| account_id | int | 否 | 账户ID（支出/收入必填，转账填源账户） |
| to_account_id | int | 否 | 转入账户ID（转账必填） |
| currency | string | 否 | 货币单位，默认CNY |
| title | string | 否 | 交易标题/商家名称，最大200字符 |
| description | string | 否 | 交易备注 |
| location | string | 否 | 交易地点，最大200字符 |
| transaction_time | string | 否 | 交易时间，格式：HH:MM:SS |
| bill_id | int | 否 | 关联账单ID |
| wishlist_id | int | 否 | 关联心愿单ID |
| tags | array | 否 | 标签数组 |
| images | array | 否 | 图片URL数组 |

**创建支出交易示例：**
```bash
curl -X POST "http://localhost:8080/api/v1/transactions" \
  -H "Authorization: Bearer 你的access_token" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "expense",
    "category_id": 10,
    "account_id": 100,
    "amount": 50.00,
    "title": "晚餐",
    "description": "烤肉自助",
    "transaction_date": "2025-12-09"
  }'
```

**创建收入交易示例：**
```bash
curl -X POST "http://localhost:8080/api/v1/transactions" \
  -H "Authorization: Bearer 你的access_token" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "income",
    "category_id": 20,
    "account_id": 100,
    "amount": 5000.00,
    "title": "本月工资",
    "transaction_date": "2025-12-01"
  }'
```

**创建转账交易示例：**
```bash
curl -X POST "http://localhost:8080/api/v1/transactions" \
  -H "Authorization: Bearer 你的access_token" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "transfer",
    "account_id": 100,
    "to_account_id": 101,
    "amount": 1000.00,
    "title": "转账到支付宝",
    "transaction_date": "2025-12-09"
  }'
```

**响应示例：**
```json
{
  "code": 201,
  "message": "success",
  "data": {
    "id": 1,
    "user_id": 1,
    "type": "expense",
    "category_id": 10,
    "account_id": 100,
    "amount": 100.00,
    "currency": "CNY",
    "title": "午餐消费",
    "description": "在食堂吃饭",
    "location": "公司食堂",
    "transaction_date": "2025-12-09",
    "transaction_time": "12:30:00",
    "tags": ["工作餐"],
    "images": [],
    "created_at": "2025-12-09T12:30:00Z",
    "updated_at": "2025-12-09T12:30:00Z"
  }
}
```

---

## 4. 批量创建交易

**POST** `/api/v1/transactions/batch`

一次性创建多条交易记录。

```bash
curl -X POST "http://localhost:8080/api/v1/transactions/batch" \
  -H "Authorization: Bearer 你的access_token" \
  -H "Content-Type: application/json" \
  -d '{
    "transactions": [
      {
        "type": "expense",
        "category_id": 10,
        "account_id": 100,
        "amount": 50.00,
        "title": "早餐",
        "transaction_date": "2025-12-09"
      },
      {
        "type": "expense",
        "category_id": 2,
        "account_id": 100,
        "amount": 200.00,
        "title": "买衣服",
        "transaction_date": "2025-12-09"
      },
      {
        "type": "income",
        "category_id": 20,
        "account_id": 100,
        "amount": 500.00,
        "title": "兼职收入",
        "transaction_date": "2025-12-09"
      }
    ]
  }'
```

**请求参数：**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| transactions | array | 是 | 交易记录数组，最多100条 |

**响应示例：**
```json
{
  "code": 201,
  "message": "success",
  "data": {
    "success_count": 3,
    "failure_count": 0,
    "errors": []
  }
}
```

**部分成功示例：**
```json
{
  "code": 201,
  "message": "success",
  "data": {
    "success_count": 2,
    "failure_count": 1,
    "errors": [
      "第2条：账户不存在"
    ]
  }
}
```

---

## 5. 更新交易

**PUT** `/api/v1/transactions/:id`

更新一条交易记录（所有字段都是可选的）。

```bash
curl -X PUT "http://localhost:8080/api/v1/transactions/1" \
  -H "Authorization: Bearer 你的access_token" \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 120.00,
    "title": "更新后的标题",
    "description": "更新后的描述"
  }'
```

**请求参数：** 同创建交易，但所有字段都是可选的。

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "user_id": 1,
    "type": "expense",
    "category_id": 10,
    "account_id": 100,
    "amount": 120.00,
    "currency": "CNY",
    "title": "更新后的标题",
    "description": "更新后的描述",
    "transaction_date": "2025-12-09",
    "created_at": "2025-12-09T12:30:00Z",
    "updated_at": "2025-12-09T12:40:00Z"
  }
}
```

---

## 6. 删除交易

**DELETE** `/api/v1/transactions/:id`

删除单条交易记录，会自动恢复账户余额。

```bash
curl -X DELETE "http://localhost:8080/api/v1/transactions/1" \
  -H "Authorization: Bearer 你的access_token"
```

**响应示例：**
```
HTTP/1.1 204 No Content
```

---

## 7. 批量删除交易

**DELETE** `/api/v1/transactions/batch`

批量删除多条交易记录。

```bash
curl -X DELETE "http://localhost:8080/api/v1/transactions/batch" \
  -H "Authorization: Bearer 你的access_token" \
  -H "Content-Type: application/json" \
  -d '{
    "ids": [1, 2, 3, 4, 5]
  }'
```

**请求参数：**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| ids | array | 是 | 交易ID数组，最多1000条 |

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "success_count": 5,
    "failure_count": 0,
    "errors": []
  }
}
```

---

## 8. 获取交易统计

**GET** `/api/v1/transactions/statistics`

获取指定日期范围内的交易统计信息（收入、支出、净额等）。

```bash
# 获取最近1个月的统计
curl -X GET "http://localhost:8080/api/v1/transactions/statistics" \
  -H "Authorization: Bearer 你的access_token"

# 获取自定义日期范围的统计
curl -X GET "http://localhost:8080/api/v1/transactions/statistics?start_date=2025-12-01&end_date=2025-12-31" \
  -H "Authorization: Bearer 你的access_token"
```

**查询参数：**

| 参数 | 类型 | 说明 |
|------|------|------|
| start_date | string | 开始日期，默认1个月前 |
| end_date | string | 结束日期，默认今天 |

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "total_income": 10000.00,
    "total_expense": 5000.00,
    "net_amount": 5000.00,
    "transaction_cnt": 50
  }
}
```

---

## 9. 获取月度统计

**GET** `/api/v1/transactions/monthly-statistics`

获取指定月份的交易统计。

```bash
# 获取当前月份的统计
curl -X GET "http://localhost:8080/api/v1/transactions/monthly-statistics" \
  -H "Authorization: Bearer 你的access_token"

# 获取特定月份的统计
curl -X GET "http://localhost:8080/api/v1/transactions/monthly-statistics?month=2025-12" \
  -H "Authorization: Bearer 你的access_token"
```

**查询参数：**

| 参数 | 类型 | 说明 |
|------|------|------|
| month | string | 月份，格式：YYYY-MM，默认当前月份 |

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "month": "2025-12",
      "total_income": 10000.00,
      "total_expense": 5000.00,
      "net_amount": 5000.00,
      "transaction_cnt": 50
    }
  ]
}
```

---

## 10. 获取分类统计

**GET** `/api/v1/transactions/category-statistics`

按分类统计支出情况，包含金额、比例等。

```bash
# 获取最近1个月的分类统计
curl -X GET "http://localhost:8080/api/v1/transactions/category-statistics" \
  -H "Authorization: Bearer 你的access_token"

# 获取自定义日期范围的分类统计
curl -X GET "http://localhost:8080/api/v1/transactions/category-statistics?start_date=2025-12-01&end_date=2025-12-31" \
  -H "Authorization: Bearer 你的access_token"
```

**查询参数：**

| 参数 | 类型 | 说明 |
|------|------|------|
| start_date | string | 开始日期，默认1个月前 |
| end_date | string | 结束日期，默认今天 |

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "category_id": 10,
      "category": {
        "id": 10,
        "user_id": 1,
        "type": "expense",
        "name": "餐饮美食",
        "icon": "fa-utensils",
        "color": "orange",
        "display_order": 0,
        "is_system": true,
        "is_active": true,
        "created_at": "2025-12-05T10:00:00Z",
        "updated_at": "2025-12-05T10:00:00Z"
      },
      "total_amount": 2000.00,
      "percentage": 40.0,
      "transaction_cnt": 20
    },
    {
      "category_id": 2,
      "category": {
        "id": 2,
        "user_id": 1,
        "type": "expense",
        "name": "购物消费",
        "icon": "fa-bag-shopping",
        "color": "purple",
        "display_order": 1,
        "is_system": true,
        "is_active": true,
        "created_at": "2025-12-05T10:00:00Z",
        "updated_at": "2025-12-05T10:00:00Z"
      },
      "total_amount": 1500.00,
      "percentage": 30.0,
      "transaction_cnt": 10
    }
  ]
}
```

---

## 交易类型说明

| 类型代码 | 说明 | 说明 |
|----------|------|------|
| expense | 支出 | 花钱购买商品或服务 |
| income | 收入 | 收到钱款，包括工资、奖金等 |
| transfer | 转账 | 账户间转账，不产生资产变化 |

---

## 常见错误码

| 状态码 | 说明 |
|--------|------|
| 200 | 成功 |
| 201 | 资源创建成功 |
| 204 | 请求成功，无返回内容 |
| 400 | 请求参数错误 |
| 401 | 未登录或token无效 |
| 404 | 交易不存在 |
| 500 | 服务器内部错误 |

---

## 测试流程

### 1. 基础测试流程

```bash
# 1. 登录获取token
TOKEN=$(curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"Test123456"}' \
  | jq -r '.data.access_token')

echo "Token: $TOKEN"

# 2. 创建交易
RESPONSE=$(curl -X POST "http://localhost:8080/api/v1/transactions" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "expense",
    "category_id": 10,
    "account_id": 1,
    "amount": 100.00,
    "title": "午餐",
    "transaction_date": "2025-12-09"
  }')

TRANSACTION_ID=$(echo $RESPONSE | jq -r '.data.id')
echo "Created transaction ID: $TRANSACTION_ID"

# 3. 获取交易列表
curl -X GET "http://localhost:8080/api/v1/transactions?page=1&page_size=10" \
  -H "Authorization: Bearer $TOKEN" | jq '.'

# 4. 获取交易详情
curl -X GET "http://localhost:8080/api/v1/transactions/$TRANSACTION_ID" \
  -H "Authorization: Bearer $TOKEN" | jq '.'

# 5. 更新交易
curl -X PUT "http://localhost:8080/api/v1/transactions/$TRANSACTION_ID" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"amount": 120.00, "title": "午餐和饮料"}' | jq '.'

# 6. 获取统计信息
curl -X GET "http://localhost:8080/api/v1/transactions/statistics" \
  -H "Authorization: Bearer $TOKEN" | jq '.'

# 7. 获取分类统计
curl -X GET "http://localhost:8080/api/v1/transactions/category-statistics" \
  -H "Authorization: Bearer $TOKEN" | jq '.'

# 8. 删除交易
curl -X DELETE "http://localhost:8080/api/v1/transactions/$TRANSACTION_ID" \
  -H "Authorization: Bearer $TOKEN"

echo "Test completed!"
```

### 2. 批量操作测试

```bash
# 批量创建交易
curl -X POST "http://localhost:8080/api/v1/transactions/batch" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "transactions": [
      {"type":"expense","category_id":10,"account_id":1,"amount":50,"title":"早餐","transaction_date":"2025-12-09"},
      {"type":"expense","category_id":2,"account_id":1,"amount":200,"title":"买衣服","transaction_date":"2025-12-09"},
      {"type":"income","category_id":20,"account_id":1,"amount":5000,"title":"工资","transaction_date":"2025-12-01"}
    ]
  }' | jq '.'

# 批量删除交易（需要先获取ID）
curl -X DELETE "http://localhost:8080/api/v1/transactions/batch" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "ids": [1, 2, 3]
  }' | jq '.'
```

### 3. 高级查询测试

```bash
# 按日期范围查询
curl -X GET "http://localhost:8080/api/v1/transactions?start_date=2025-12-01&end_date=2025-12-31" \
  -H "Authorization: Bearer $TOKEN" | jq '.'

# 按分类查询
curl -X GET "http://localhost:8080/api/v1/transactions?category_id=10" \
  -H "Authorization: Bearer $TOKEN" | jq '.'

# 按账户查询
curl -X GET "http://localhost:8080/api/v1/transactions?account_id=1" \
  -H "Authorization: Bearer $TOKEN" | jq '.'

# 按类型排序
curl -X GET "http://localhost:8080/api/v1/transactions?sort_by=amount&sort_order=asc" \
  -H "Authorization: Bearer $TOKEN" | jq '.'

# 搜索
curl -X GET "http://localhost:8080/api/v1/transactions?search_keyword=午餐" \
  -H "Authorization: Bearer $TOKEN" | jq '.'
```

---

## 数据库种子数据

如果需要初始化测试数据，可运行以下脚本：

```bash
cd backend/scripts
mysql -u root -p floatisland < seed_transactions.sql
```

或手动创建测试数据：

```sql
USE floatisland;

-- 创建交易记录
INSERT INTO transactions (user_id, type, category_id, account_id, amount, currency, title, description, location, transaction_date, created_at, updated_at)
VALUES
  (1, 'expense', 1, 1, 50.00, 'CNY', '早餐', '便利店', '门口', '2025-12-09', NOW(), NOW()),
  (1, 'expense', 1, 1, 100.00, 'CNY', '午餐', '食堂', '公司', '2025-12-09', NOW(), NOW()),
  (1, 'expense', 2, 1, 200.00, 'CNY', '购物', '衣服', '商场', '2025-12-08', NOW(), NOW()),
  (1, 'income', 7, 1, 5000.00, 'CNY', '工资', '本月薪资', '公司', '2025-12-01', NOW(), NOW()),
  (1, 'transfer', NULL, 1, 1000.00, 'CNY', '转账', '转账到支付宝', '在线', '2025-12-05', NOW(), NOW());
```

---

## 测试完成检查清单

- [ ] 创建交易成功
- [ ] 批量创建交易成功
- [ ] 查询交易列表成功
- [ ] 获取交易详情成功
- [ ] 更新交易成功
- [ ] 删除交易成功
- [ ] 批量删除交易成功
- [ ] 查询交易统计成功
- [ ] 查询月度统计成功
- [ ] 查询分类统计成功
- [ ] 验证账户余额自动更新
- [ ] 验证删除交易后余额恢复
- [ ] 验证分页功能
- [ ] 验证排序功能
- [ ] 验证搜索功能

---

## 性能测试

使用 `wrk` 进行性能测试：

```bash
# 获取交易列表（最常用操作）
wrk -t12 -c400 -d30s -H "Authorization: Bearer YOUR_TOKEN" \
  http://localhost:8080/api/v1/transactions

# 创建交易
wrk -t12 -c100 -d30s -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -s create_transaction.lua \
  http://localhost:8080/api/v1/transactions
```

---

最后更新：2025-12-09  
版本：1.0
