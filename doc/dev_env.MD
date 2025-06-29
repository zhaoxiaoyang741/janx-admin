# 开发环境搭建

## 1. 安装GO
[Golang下载位置](https://golang.google.cn/dl/)下载对应系统版本安装宝后安装

## 2. 安装postgresql
docker 安装命令：   
```
docker pull postgres:17.5-alpine //拉取镜像   
docker run -d --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pg123456 -v /mnt/pgdata:/var/lib/postgresql/data -p 5432:5432 postgres:17.5-alpine //启动容器
```
## 3. 修改配置
修改配置文件`config.yaml`中的数据库连接信息
## 4. 安装依赖
```
go mod tidy
```
## 4. 启动项目
```
go run main.go
```

