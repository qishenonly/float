# Float 后端日志指南

## 概述

已为后端添加了全面的日志系统，包括：
1. **Handler 层** - HTTP 请求处理日志
2. **Service 层** - 业务逻辑执行日志
3. **性能监控** - 执行时间统计
4. **敏感信息过滤** - 自动隐藏敏感数据

## 日志等级

| 等级 | 用途 | 示例 |
|------|------|------|
| **Info** | 正常操作流程 | 请求开始、操作成功 |
| **Warn** | 警告信息 | 验证失败、数据冲突 |
| **Error** | 错误信息 | 异常、操作失败 |
| **Debug** | 调试信息 | 中间步骤（可选） |

## 日志格式

所有日志采用统一格式：
```
[层级][操作名称] 具体信息 | 参数1: 值1 | 参数2: 值2
```

### 日志层级标记

- `[Handler]` - HTTP 处理器层
- `[Service]` - 业务服务层
- `[性能]` - 性能监控日志

## 敏感信息过滤

### 自动隐藏的敏感字段

系统会自动隐藏以下信息：

| 字段类型 | 原始值 | 隐藏后 |
|---------|-------|--------|
| 邮箱 | user@example.com | u***@example.com |
| 电话 | 13800138000 | 138****8000 |
| 密码 | mypassword123 | [密码已隐藏:13位] |
| 令牌 | eyJhbGc... | [TOKEN已隐藏] |
| 验证码 | 123456 | [CODE已隐藏] |
| 卡号 | 6228888888888888 | 6228****8888 |

### 日志中的敏感信息隐藏示例

**注册请求日志：**
```
[Handler][注册] 注册请求 | 邮箱: u***@example.com | 用户名: john_doe
[Service][注册] 开始注册 | 邮箱: u***@example.com | 用户名: john_doe
[Service][注册] 注册成功 | 用户ID: 123 | 邮箱: u***@example.com | 用户名: john_doe
```

**登录请求日志：**
```
[Handler][登录] 登录请求 | 邮箱: u***@example.com
[Service][登录] 开始登录 | 邮箱: u***@example.com
[Service][登录] 登录成功 | 用户ID: 456 | 用户名: jane_doe
```

**响应中的敏感信息：**
```
AccessToken: [TOKEN已隐藏]
RefreshToken: [TOKEN已隐藏]
```

## 性能监控

### 执行时间阈值

系统会根据不同操作设置性能警告阈值：

| 操作 | 阈值 | 说明 |
|------|------|------|
| 发送验证码 | 500ms | 邮件发送等I/O操作 |
| 用户注册 | 1000ms | 包含多个数据库操作 |
| 用户登录 | 500ms | 密码验证和token生成 |
| 刷新令牌 | 300ms | 令牌处理 |
| 修改密码 | 800ms | 密码哈希和数据库更新 |
| 获取账户列表 | 300ms | 查询操作 |
| 创建账户 | 500ms | 数据库创建操作 |

### 性能日志示例

**正常执行：**
```log
[性能][用户注册] 耗时: 234.56ms
```

**超过阈值（警告）：**
```log
[性能][用户注册] 耗时过长: 1256.78ms (阈值: 1000ms)
```

**执行失败：**
```log
[性能][用户注册] 创建用户失败 | 耗时: 45.23ms
```

## 完整日志流程示例

### 用户注册完整日志序列

```log
# 1. HTTP Handler 层
[HTTP] POST /api/v1/auth/send-verification-code 200 234ms
[Handler][发送验证码] 请求邮箱: u***@example.com
[Handler][发送验证码] 发送成功 | 邮箱: u***@example.com

# 2. Service 层
[Service][验证码] 发送请求 | 邮箱: u***@example.com
[Service][验证码] 发送成功 | 邮箱: u***@example.com
[性能][发送验证码] 耗时: 234.56ms

# 3. 发送验证码成功，用户收到邮件

# 4. 用户注册请求
[HTTP] POST /api/v1/auth/register 200 856ms
[Handler][注册] 注册请求 | 邮箱: u***@example.com | 用户名: john_doe
[Service][注册] 开始注册 | 邮箱: u***@example.com | 用户名: john_doe
[Service][注册] 用户创建成功 | 用户ID: 123
[Service][注册] 注册成功 | 用户ID: 123 | 邮箱: u***@example.com | 用户名: john_doe
[性能][用户注册] 耗时: 856.23ms
[Handler][注册] 注册成功 | 邮箱: u***@example.com | 用户ID: 123
```

### 登录失败完整日志序列

```log
[HTTP] POST /api/v1/auth/login 400 145ms
[Handler][登录] 登录请求 | 邮箱: u***@example.com

[Service][登录] 开始登录 | 邮箱: u***@example.com
[Service][登录] 密码错误 | 用户ID: 123 | 邮箱: u***@example.com
[性能][用户登录] 密码验证失败 | 耗时: 145.67ms

[Handler][登录] 登录失败 | 邮箱: u***@example.com | 错误: 邮箱或密码错误
```

## 敏感信息过滤 API

### 手动隐藏敏感信息

可以在日志代码中使用工具函数手动隐藏敏感信息：

```go
import "github.com/qiuhaonan/float-backend/pkg/logger"

// 隐藏邮箱
sanitizedEmail := logger.SanitizeEmail("user@example.com")
// 输出: u***@example.com

// 隐藏电话
sanitizedPhone := logger.SanitizePhone("13800138000")
// 输出: 138****8000

// 隐藏密码
sanitizedPwd := logger.SanitizePassword("mypassword123")
// 输出: [密码已隐藏:13位]

// 隐藏卡号
sanitizedCard := logger.SanitizeCardNumber("6228888888888888")
// 输出: 6228****8888

// 隐藏JSON字符串中的敏感字段
jsonStr := `{"email":"user@example.com","password":"secret123","token":"abc123"}`
sanitized := logger.SanitizeJSON(jsonStr)
// 输出: {"email":"u***@example.com","password":"[密码已隐藏]","token":"[TOKEN已隐藏]"}
```

## 性能监控 API

### 使用计时器

```go
import "github.com/qiuhaonan/float-backend/pkg/logger"

// 创建计时器
timer := logger.NewTimer("操作名称")

// 执行业务逻辑...

// 记录执行时间（Info 级别）
timer.Log()

// 记录执行时间，如果超过阈值则警告
timer.LogSlow(500 * time.Millisecond)

// 记录执行时间，带自定义消息
timer.LogWithMsg("info", "操作完成")

// 记录执行时间，单位ms的警告阈值
timer.LogSlowWithThreshold("注册完成", 1000) // 1000ms 阈值

// 获取执行时间
elapsed := timer.Elapsed()       // time.Duration
elapsedMs := timer.ElapsedMs()   // int64
elapsedStr := timer.ElapsedString() // "234.56ms"
```

## 日志调试技巧

### 查找特定用户的日志

```bash
# 查找用户ID为123的所有日志
grep "用户ID: 123" app.log

# 查找特定邮箱的所有日志
grep "u***@example.com" app.log
```

### 查找性能问题

```bash
# 查找所有超时的操作
grep "耗时过长\|警告阈值" app.log

# 查找执行时间超过1秒的操作
grep -E "[0-9]{4,}\.[0-9]{2}ms|[0-9]+\.[0-9]{2}s" app.log
```

### 查找错误

```bash
# 查找所有错误日志
grep "\[错误\]\|\[ERROR\]" app.log

# 查找特定操作的失败
grep "\[Service\]\[注册\].*失败" app.log
```

## 日志配置最佳实践

### 生产环境

- 设置日志级别为 `INFO`
- 启用日志轮转（按大小或时间）
- 定期清理过期日志（超过30天）
- 监控日志文件大小

### 开发环境

- 设置日志级别为 `DEBUG`
- 输出到控制台和文件
- 启用详细的堆栈跟踪
- 实时查看日志

### 测试环境

- 设置日志级别为 `INFO`
- 启用性能监控告警
- 记录所有网络请求
- 保留日志用于分析

## 常见日志查询

### 查看注册流程

```bash
# 完整的注册流程日志
grep -A 20 "\[Handler\]\[注册\].*注册请求" app.log
```

### 查看登录失败原因

```bash
# 登录失败的所有原因
grep "\[Handler\]\[登录\].*失败\|\[Service\]\[登录\].*错误" app.log
```

### 查看账户操作

```bash
# 查看所有账户相关操作
grep "\[Service\]\[账户\]" app.log
```

### 统计各操作的执行时间

```bash
# 统计注册操作的平均执行时间
grep "\[性能\]\[用户注册\]" app.log | awk -F'[: ms]' '{sum+=$NF; count++} END {print "平均耗时: "sum/count"ms"}'
```

## 总结

✅ **完整的日志系统**
- Handler 层：捕捉所有HTTP请求
- Service 层：记录业务逻辑执行
- 性能监控：自动检测性能问题
- 敏感信息过滤：保护用户隐私

✅ **易于调试**
- 统一的日志格式
- 详细的错误信息
- 完整的执行链路追踪

✅ **安全可靠**
- 自动隐藏敏感数据
- 性能异常告警
- 完整的操作审计
