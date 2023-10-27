# 为 dBlog 创建数据库
- 使用 Docker + MariaDB 创建数据库
- sqldump 数据相关详细信息查看:
[sql](/internal/pkg/model/README.md)

## 创建数据库
- 启动数据库实例
```shell
docker run -p 3306:3306 --name dazblogDB -e MARIADB_ROOT_PASSWORD=passwd -d mariadb:lts
```

### 使用 mysql 还原 dBlog.sql (推荐)
- [dBlog.sql位置](../../../../configs/dBlog.sql)

### 手动方式 (不推荐)
- 初始化数据库
```sql
CREATE DATABASE `dazblog`;
USE `dazblog`;
```

- 初始化用户表
```sql
CREATE TABLE `users` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `postcount` int(11) NOT NULL DEFAULT '0',
    `username` varchar(30) NOT NULL,
    `password` varchar(255) NOT NULL,
    `nickname` varchar(30) NOT NULL,
    `email` varchar(320) NOT NULL,
    `gender` ENUM('Male', 'Female', 'Other') DEFAULT NULL,
    `phone` varchar(16),
    `qq` varchar(16),
    `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `updatedAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP(),
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`),
    UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB, AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
```

- 初始化博客表
```sql
CREATE TABLE `posts` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(30) NOT NULL,
    `postID` varchar(100) NOT NULL,
    `title` varchar(150) NOT NULL,
    `content` longtext NOT NULL,
    `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `updatedAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP(),
    PRIMARY KEY (`id`),
    UNIQUE KEY `postID` (`postID`),
    KEY `idx_username` (`username`),
    CONSTRAINT fk_post_username FOREIGN KEY (`username`) REFERENCES `users`(`username`) ON DELETE CASCADE
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
```

- 初始化 AI 内容生成表
```sql
CREATE TABLE `ai` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `username` varchar(30) NOT NULL,
    `postID` varchar(100) NOT NULL,
    `content` longtext NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `postID` (`postID`),
    CONSTRAINT fk_post_id FOREIGN KEY (postID) REFERENCES posts(postID) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
```
