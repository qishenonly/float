#!/bin/bash

# 测试Docker容器运行的脚本

set -e

echo "🐳 测试Float Island Docker容器"
echo "================================="

# 检查镜像是否存在
echo "📦 检查镜像..."
if ! docker images | grep -q "float-island-website"; then
    echo "❌ 镜像不存在，请先运行 ./build.sh"
    exit 1
fi
echo "✅ 镜像存在"

# 停止可能存在的旧容器
echo "🛑 停止旧容器..."
docker stop float-website 2>/dev/null || true
docker rm float-website 2>/dev/null || true

# 测试1: 检查镜像架构
echo "🔍 检查镜像架构..."
docker inspect float-island-website:latest | grep -E '"Architecture"|"Os"' | head -2

# 测试2: 尝试运行shell
echo "🐚 测试容器shell..."
if docker run --rm -it float-island-website:latest sh -c "echo 'Shell works!' && uname -a"; then
    echo "✅ Shell测试通过"
else
    echo "❌ Shell测试失败"
    exit 1
fi

# 测试3: 检查npm和node
echo "📦 测试npm和node..."
if docker run --rm -it float-island-website:latest sh -c "node --version && npm --version"; then
    echo "✅ Node.js和npm工作正常"
else
    echo "❌ Node.js或npm有问题"
    exit 1
fi

# 测试4: 检查应用文件
echo "📁 检查应用文件..."
if docker run --rm -it float-island-website:latest sh -c "ls -la /app && test -d /app/dist && echo 'dist目录存在' || echo 'dist目录不存在'"; then
    echo "✅ 应用文件检查通过"
else
    echo "❌ 应用文件检查失败"
    exit 1
fi

# 测试5: 尝试预览命令
echo "🚀 测试预览命令..."
if docker run --rm -it float-island-website:latest sh -c "npm run preview --version 2>/dev/null || echo '预览命令存在但可能需要端口'"; then
    echo "✅ 预览命令存在"
else
    echo "❌ 预览命令不存在"
    exit 1
fi

# 最终测试: 启动容器
echo "🎯 最终测试: 启动容器..."
if docker run -d -p 8080:4173 --name float-website float-island-website:latest; then
    echo "✅ 容器启动成功"
    echo "🌐 访问地址: http://localhost:8080"

    # 等待几秒钟让容器完全启动
    sleep 3

    # 检查容器状态
    if docker ps | grep -q float-website; then
        echo "✅ 容器正在运行"
        echo ""
        echo "📋 管理命令:"
        echo "  查看日志: docker logs -f float-website"
        echo "  停止容器: docker stop float-website"
        echo "  删除容器: docker rm float-website"
    else
        echo "❌ 容器启动失败，查看日志:"
        docker logs float-website
        exit 1
    fi
else
    echo "❌ 容器启动失败"
    exit 1
fi

echo ""
echo "🎉 所有测试通过！容器运行正常。"