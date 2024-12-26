# 使用官方的 Go 镜像作为构建和运行的基础镜像
FROM golang:1.23.3

# 设置工作目录
WORKDIR /app

# 复制项目文件到容器中
COPY . .

ENV ENV=production

# 运行 build-app.sh 构建项目
RUN chmod +x build-app.sh && ./build-app.sh

# 设置容器启动时运行的命令
CMD ["./output/bootstrap.sh"]

EXPOSE 9100