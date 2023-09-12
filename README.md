# dazBlog

## 采用模型
- 模型层 -- 控制层 -- 业务层 -- 仓库层
- model -- controller -- biz -- store
- 存储对象结构与其方法 -- 业务路由 -- 业务逻辑处理 -- 与数据库/第三方服务进行 CRUD

## 日志系统
[基于 zap 构建可自定义的日志系统](./internal/pkg/log/README.md)

## 版本信息
[打印详细版本信息](./pkg/version/README.md)