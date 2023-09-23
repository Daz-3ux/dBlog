# model layer
- model 层就是数据库表字段的 Go 结构体映射
- 根据业务需求,用 dBlog 数据库中的表来创建 Model

## 表
[user 表代码](./user.go)
[post 表代码](./post.go)
- 借助工具自动生成结构体代码
  - 创建数据库和数据库表
  - 根据数据库表生成 Model 文件
    - [db2struct](https://github.com/Shelnutt2/db2struct)
  - 优化代码

## SQL 语句存储
- 为了以后的部署,将创建数据库以及数据库表的 SQL 语句存储在了 [dBlog.sql](../../../configs/dBlog.sql) 文件中
```shell
mysqldump -h127.0.0.1 -udazblog --databases dazblog -p'passwd' --add-drop-database --add-drop-table --add-drop-trigger --add-locks --no-data > configs/dBblog.sql
```
