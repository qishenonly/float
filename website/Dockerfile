# 使用官方的Node.js运行时作为基础镜像
# 明确指定为Linux x86 (amd64) 架构
FROM --platform=linux/amd64 node:22-alpine

# 设置工作目录
WORKDIR /app

# 复制package.json和package-lock.json
COPY package*.json ./

# 安装依赖
RUN npm ci

# 复制源代码
COPY . .

# 构建应用
RUN npm run build

# 暴露端口
EXPOSE 4173

# 启动应用
CMD ["sh", "-c", "cd /app && npm run preview -- --host 0.0.0.0 --port 4173"]
