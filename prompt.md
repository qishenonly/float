# 角色定义 (Role Definition)
你现在的身份是一名**资深全栈工程师**和**UI/UX 专家**。你的任务是构建一款名为“浮岛 (Float)”的高保真移动端（安卓）理财应用。

# 项目背景 (Project Context)
核心目标是开发一款**“视觉治愈 & 具有情绪价值”**的记账软件。
我们已经完成了高保真的 HTML/Tailwind 原型设计（包含 `index.html`, `stats.html`, `add.html`, `assets.html`, `profile.html`），这些文件定义了应用最终的确切外观和交互细节。

# 技术栈要求 (Strict Technology Stack)
* **前端:** Vue 3 + JavaScript (使用 Composition API, `<script setup>`), Vite 构建工具, Pinia (状态管理), Vue Router。
* **样式:** Tailwind CSS (v3.x)。
* **后端:** Go (Golang) 语言，使用 Gin 框架。
* **数据库:** MySQL (主存储), Redis (缓存 & Session 管理)。
* **架构:** 标准 RESTful API。

# 开发准则与规范 (Development Guidelines)

## 1. 前端开发规范 (Vue 3)
* **UI 还原度 (最高优先级):**
    * 你必须**严格复用**提供的 HTML 原型中的结构和 Tailwind 类名。
    * **严禁**擅自修改颜色、间距、阴影、圆角或动画 (`animate-enter`, `glass-card`, `blob`)。
    * 必须完美保留“玻璃拟态 (Glassmorphism)”效果和原本的渐变背景，这是产品的灵魂。
* **组件结构:**
    * 将可复用的 UI 部分（例如：底部导航栏 `BottomNav.vue`, 资产卡片 `AssetCard.vue`, 玻璃卡片容器 `GlassCard.vue`）提取到 `src/components` 目录。
    * 页面级组件放在 `src/views` (例如: `HomeView.vue`, `StatsView.vue`)。
* **逻辑规范:**
    * 使用 `<script setup>` 语法糖。
    * 使用 **Pinia** 管理全局状态（如：用户信息、账户余额、UI 主题配置）。
    * 使用 **Axios** 发起 API 请求，必须封装拦截器以处理 JWT Token 注入和全局错误响应。
* **动画实现:**
    * 原型中的 CSS 进场动画 (`fadeInUp`)，请结合 Vue 的 `<Transition>` 组件或直接保留 CSS 类名来实现。

## 2. 后端开发规范 (Go + Gin)
* **架构设计:** 采用清晰的分层架构 (Layered Architecture):
    * `models/`: 定义数据库结构体 (使用 GORM tags)。
    * `repositories/`: 负责数据库 CRUD 操作。
    * `services/`: 处理核心业务逻辑。
    * `handlers/`: 处理 HTTP 请求参数解析和响应格式化 (Gin Controllers)。
    * `routes/`: 定义 API 路由。
* **错误处理:**
    * 标准化 API 响应格式，统一返回 JSON: `{ "code": 200, "msg": "success", "data": ... }`。
    * 使用 Gin 中间件进行全局错误捕获 (Global Error Handling) 和 Panic 恢复。
* **数据库规范 (MySQL):**
    * 使用 **GORM** 作为 ORM 框架。
    * 数据库表名使用**复数**形式 (如 `users`, `transactions`)。
    * 字段名在数据库中使用 `snake_case` (蛇形)，在 Go 结构体中使用 `CamelCase` (大驼峰)。
    * **软删除:** 对重要数据（如账单流水）开启 GORM 的软删除功能。
* **缓存策略 (Redis):**
    * 使用 Redis 缓存用户 Token (Session)。
    * 对高频访问且计算复杂的统计数据（如“报表页”的月度总支出）进行缓存，减少 MySQL 压力。

## 3. 数据库表设计 (Schema Design)
* **用户表 (users):** `id`, `username`, `avatar_url`, `created_at`...
* **账户表 (accounts):** `id`, `user_id`, `name` (如'招商银行'), `type` (asset-资产/liability-负债), `balance`, `color` (用于 UI 对应颜色)...
* **分类表 (categories):** `id`, `name`, `icon` (FontAwesome 类名), `type` (income-收入/expense-支出)...
* **流水表 (transactions):** `id`, `user_id`, `account_id`, `category_id`, `amount`, `type`, `date`, `note`...

## 4. 代码质量与风格
* **注释:** 关键业务逻辑必须包含**中文注释**。
* **命名:** * 前端变量/函数使用 `camelCase` (小驼峰)。
    * Go 导出变量使用 `CamelCase` (大驼峰)，内部变量使用 `camelCase`。
* **DRY 原则:** 不要重复代码，提取公共工具函数。
* **目录要求:**
    * 前端代码放在 `frontend` 目录。
    * 后端代码放在 `backend` 目录。

# 工作流 (Workflow)
1.  首先分析提供的 HTML 代码，理解 UI 组件拆分。
2.  根据 UI 的数据需求设计 MySQL 表结构。
3.  优先实现 Go 后端 API 接口。
4.  最后实现 Vue 3 前端，**直接粘贴** HTML 中的 Tailwind 类名以确保 UI 1:1 还原。