# 分类管理 API 测试指南

## 环境准备

确保已启动服务：
```bash
cd backend
make run
```

服务地址：`http://localhost:8080`

**注意**：所有分类管理接口都需要认证，请先登录获取 access_token。

---

## 1. 获取分类列表

**GET** `/api/v1/categories`

获取用户的所有分类（包含系统默认分类+用户自定义分类）

```bash
curl -X GET "http://localhost:8080/api/v1/categories" \
  -H "Authorization: Bearer 你的access_token"
```

**查询参数：**
- `type`（可选）：分类类型
  - `expense` - 仅支出分类
  - `income` - 仅收入分类
  - 不传则返回所有分类

**按类型筛选示例：**
```bash
# 仅获取支出分类
curl -X GET "http://localhost:8080/api/v1/categories?type=expense" \
  -H "Authorization: Bearer 你的access_token"

# 仅获取收入分类
curl -X GET "http://localhost:8080/api/v1/categories?type=income" \
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
      "type": "expense",
      "name": "餐饮美食",
      "icon": "fa-utensils",
      "color": "orange",
      "display_order": 1,
      "is_system": true,
      "is_active": true,
      "created_at": "2025-12-05T10:00:00Z",
      "updated_at": "2025-12-05T10:00:00Z"
    },
    {
      "id": 11,
      "type": "expense",
      "name": "宠物开销",
      "icon": "fa-paw",
      "color": "brown",
      "display_order": 0,
      "is_system": false,
      "is_active": true,
      "created_at": "2025-12-05T11:30:00Z",
      "updated_at": "2025-12-05T11:30:00Z"
    }
  ]
}
```

---

## 2. 获取分类详情

**GET** `/api/v1/categories/:id`

```bash
curl -X GET "http://localhost:8080/api/v1/categories/1" \
  -H "Authorization: Bearer 你的access_token"
```

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "type": "expense",
    "name": "餐饮美食",
    "icon": "fa-utensils",
    "color": "orange",
    "display_order": 1,
    "is_system": true,
    "is_active": true,
    "created_at": "2025-12-05T10:00:00Z",
    "updated_at": "2025-12-05T10:00:00Z"
  }
}
```

---

## 3. 创建自定义分类

**POST** `/api/v1/categories`

```bash
curl -X POST "http://localhost:8080/api/v1/categories" \
  -H "Authorization: Bearer 你的access_token" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "expense",
    "name": "宠物开销",
    "icon": "fa-paw",
    "color": "brown",
    "display_order": 10
  }'
```

**请求参数：**
| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| type | string | 是 | 分类类型：`expense` 或 `income` |
| name | string | 是 | 分类名称，最大50字符 |
| icon | string | 是 | FontAwesome图标代码，如 `fa-paw` |
| color | string | 是 | 颜色标识，如 `brown`、`orange`、`blue` |
| display_order | int | 否 | 显示排序，默认0 |

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 11,
    "type": "expense",
    "name": "宠物开销",
    "icon": "fa-paw",
    "color": "brown",
    "display_order": 10,
    "is_system": false,
    "is_active": true,
    "created_at": "2025-12-05T11:30:00Z",
    "updated_at": "2025-12-05T11:30:00Z"
  }
}
```

---

## 4. 更新分类

**PUT** `/api/v1/categories/:id`

**注意**：只能更新用户自己创建的分类，系统分类不可修改。

```bash
curl -X PUT "http://localhost:8080/api/v1/categories/11" \
  -H "Authorization: Bearer 你的access_token" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "宠物用品",
    "icon": "fa-dog",
    "color": "amber"
  }'
```

**请求参数：**
| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| name | string | 否 | 分类名称 |
| icon | string | 否 | 图标代码 |
| color | string | 否 | 颜色标识 |
| display_order | int | 否 | 显示排序 |
| is_active | bool | 否 | 是否启用 |

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

**错误示例（尝试修改系统分类）：**
```json
{
  "code": 403,
  "message": "系统分类不可修改",
  "data": null
}
```

---

## 5. 删除分类

**DELETE** `/api/v1/categories/:id`

**注意**：只能删除用户自己创建的分类，系统分类不可删除。

```bash
curl -X DELETE "http://localhost:8080/api/v1/categories/11" \
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

**错误示例（尝试删除系统分类）：**
```json
{
  "code": 403,
  "message": "系统分类不可删除",
  "data": null
}
```

---

## 6. 获取系统默认分类

**GET** `/api/v1/categories/system`

获取所有系统默认分类（不需要登录即可访问）

```bash
curl -X GET "http://localhost:8080/api/v1/categories/system"
```

**查询参数：**
- `type`（可选）：分类类型 `expense` 或 `income`

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "type": "expense",
      "name": "餐饮美食",
      "icon": "fa-utensils",
      "color": "orange",
      "display_order": 1,
      "is_system": true,
      "is_active": true
    },
    {
      "id": 2,
      "type": "expense",
      "name": "购物消费",
      "icon": "fa-bag-shopping",
      "color": "purple",
      "display_order": 2,
      "is_system": true,
      "is_active": true
    }
  ]
}
```

---

## 系统默认分类列表

### 支出分类（Expense）

| ID | 名称 | 图标 | 颜色 |
|----|------|------|------|
| 1 | 餐饮美食 | fa-utensils | orange |
| 2 | 购物消费 | fa-bag-shopping | purple |
| 3 | 交通出行 | fa-bus | blue |
| 4 | 住房物业 | fa-house | green |
| 5 | 医疗健康 | fa-heartbeat | red |
| 6 | 文化娱乐 | fa-gamepad | pink |

### 收入分类（Income）

| ID | 名称 | 图标 | 颜色 |
|----|------|------|------|
| 7 | 工资薪水 | fa-sack-dollar | indigo |
| 8 | 理财收益 | fa-chart-line | red |
| 9 | 兼职外快 | fa-briefcase | green |
| 10 | 礼金红包 | fa-gift | pink |

---

## 常见错误码

| 状态码 | 说明 |
|--------|------|
| 200 | 成功 |
| 400 | 请求参数错误（如缺少必填字段） |
| 401 | 未登录或token无效 |
| 403 | 权限不足（如尝试修改/删除系统分类） |
| 404 | 分类不存在 |
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

# 2. 获取所有分类
curl -X GET "http://localhost:8080/api/v1/categories" \
  -H "Authorization: Bearer $TOKEN"

# 3. 创建自定义分类
curl -X POST "http://localhost:8080/api/v1/categories" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "expense",
    "name": "测试分类",
    "icon": "fa-star",
    "color": "yellow"
  }'

# 4. 更新分类（假设ID为11）
curl -X PUT "http://localhost:8080/api/v1/categories/11" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"更新后的分类"}'

# 5. 删除分类
curl -X DELETE "http://localhost:8080/api/v1/categories/11" \
  -H "Authorization: Bearer $TOKEN"
```

### 2. 权限测试

```bash
# 尝试修改系统分类（应该失败）
curl -X PUT "http://localhost:8080/api/v1/categories/1" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"修改系统分类"}'

# 尝试删除系统分类（应该失败）
curl -X DELETE "http://localhost:8080/api/v1/categories/1" \
  -H "Authorization: Bearer $TOKEN"
```

---

## 数据库种子数据

如果数据库中没有系统默认分类，请运行以下SQL脚本：

```bash
cd backend/scripts
mysql -u root -p floatisland < seed_categories.sql
```

或手动插入：
```sql
USE floatisland;

-- 支出分类
INSERT INTO categories (user_id, type, name, icon, color, display_order, is_system, is_active)
VALUES 
  (0, 'expense', '餐饮美食', 'fa-utensils', 'orange', 1, TRUE, TRUE),
  (0, 'expense', '购物消费', 'fa-bag-shopping', 'purple', 2, TRUE, TRUE),
  (0, 'expense', '交通出行', 'fa-bus', 'blue', 3, TRUE, TRUE),
  (0, 'expense', '住房物业', 'fa-house', 'green', 4, TRUE, TRUE),
  (0, 'expense', '医疗健康', 'fa-heartbeat', 'red', 5, TRUE, TRUE),
  (0, 'expense', '文化娱乐', 'fa-gamepad', 'pink', 6, TRUE, TRUE);

-- 收入分类
INSERT INTO categories (user_id, type, name, icon, color, display_order, is_system, is_active)
VALUES 
  (0, 'income', '工资薪水', 'fa-sack-dollar', 'indigo', 1, TRUE, TRUE),
  (0, 'income', '理财收益', 'fa-chart-line', 'red', 2, TRUE, TRUE),
  (0, 'income', '兼职外快', 'fa-briefcase', 'green', 3, TRUE, TRUE),
  (0, 'income', '礼金红包', 'fa-gift', 'pink', 4, TRUE, TRUE);
```
