# 部署步骤

### 拉取 Redis 和 MySQL 镜像并启动容器

```bash
docker pull redis && docker pull mysql
```

### 启动redis容器

```bash
docker run -d --name redis-sectran -p 6379:6379 redis:latest
```

### 启动mysql容器

```bash
docker run -d \
  --name mysql-sectran \
  -e MYSQL_ROOT_PASSWORD=root \
  -e MYSQL_DATABASE=root \
  -p 3306:3306 \
  mysql:latest
```

### 拷贝sql到mysql容器中

```bash
docker cp sectran.sql mysql-sectran:/sectran.sql
```

### 导入sql

```bash
docker exec -it mysql-sectran /bin/bash
mysql -uroot -p < /sectran.sql
```

