#!/bin/bash

# Float Website Docker Build Script
# 为linux/amd64平台构建Docker镜像并打包为tar文件

set -e

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 配置
IMAGE_NAME="float-website"
IMAGE_TAG="latest"
TAR_OUTPUT="float-website-docker-image.tar.gz"
REGISTRY="${REGISTRY:-}"

# 添加registry前缀（如果指定）
if [ -n "$REGISTRY" ]; then
    IMAGE_FULL_NAME="${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG}"
else
    IMAGE_FULL_NAME="${IMAGE_NAME}:${IMAGE_TAG}"
fi

echo -e "${YELLOW}=== Float Website Docker Build ===${NC}"
echo -e "Image: ${YELLOW}${IMAGE_FULL_NAME}${NC}"
echo -e "Platform: ${YELLOW}linux/amd64${NC}"
echo -e "Output TAR: ${YELLOW}${TAR_OUTPUT}${NC}"
echo ""

# 检查Docker是否安装
if ! command -v docker &> /dev/null; then
    echo -e "${RED}❌ Docker未安装${NC}"
    exit 1
fi

# 构建Docker镜像
echo -e "${YELLOW}[1/3] 构建Docker镜像...${NC}"
docker build --platform linux/amd64 \
    -t "${IMAGE_FULL_NAME}" \
    -f Dockerfile \
    .

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ 镜像构建成功${NC}"
else
    echo -e "${RED}❌ 镜像构建失败${NC}"
    exit 1
fi

# 保存镜像为tar文件
echo -e "${YELLOW}[2/3] 保存镜像为TAR文件...${NC}"
docker save "${IMAGE_FULL_NAME}" | gzip > "${TAR_OUTPUT}"

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ 镜像保存成功${NC}"
    TAR_SIZE=$(du -h "${TAR_OUTPUT}" | cut -f1)
    echo -e "文件大小: ${YELLOW}${TAR_SIZE}${NC}"
else
    echo -e "${RED}❌ 镜像保存失败${NC}"
    exit 1
fi

# 显示镜像信息
echo -e "${YELLOW}[3/3] 镜像信息${NC}"
docker image inspect "${IMAGE_FULL_NAME}" --format='
镜像ID: {{.ID}}
大小: {{.Size | printf "%.2f MB"}}
创建时间: {{.Created}}
平台: {{.Architecture}}/{{.Os}}
' | numfmt --to=iec 2>/dev/null || docker image inspect "${IMAGE_FULL_NAME}" --format='
镜像ID: {{.ID}}
创建时间: {{.Created}}
'

echo ""
echo -e "${GREEN}=== 构建完成 ===${NC}"
echo -e "镜像已保存到: ${YELLOW}${TAR_OUTPUT}${NC}"
echo ""
echo "使用镜像:"
echo -e "  加载镜像: ${YELLOW}docker load -i ${TAR_OUTPUT}${NC}"
echo -e "  运行容器: ${YELLOW}docker run -p 4173:4173 ${IMAGE_FULL_NAME}${NC}"
echo ""
