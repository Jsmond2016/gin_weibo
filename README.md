# 微博项目实战

> 课件：https://www.chaindesk.cn/witbook/19/379


## 项目实战

上述课件中，全部内容已完成，


项目启动说明：

- 配置好 docker 和数据库配置文件

```bash
docker pull mysql:latest

docker images

# 账号密码 root/123456
docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql
```

新建数据库：

```sql
create database gin_weibo;

use gin_web;
```


- 项目启动，直接运行

```bash
go mod tidy

gowatch main.go
```


推荐资料：

- https://www.liwenzhou.com/posts/Go/go_mysql/  Go语言操作MySQL——李文周的博客
