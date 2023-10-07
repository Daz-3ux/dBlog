# dazBlog

## 业务架构模型
- 模型层 -- 控制层 -- 业务层 -- 仓库层
- model -- controller -- biz -- store
- 存储对象结构与其方法 -- 业务路由 -- 业务逻辑处理 -- 与数据库/第三方服务进行 CRUD
![架构图](./internal/resource/arch)
- 开发顺序:
  - Model -> Store -> Biz -> Controller
  - 从下到上，优先开发依赖少的组件
  - 一次性开发完一整条链路, 从而保证整个链路的可用性
  - [Model](./internal/pkg/model/README.md)
  - [Store](./internal/dazBlog/store/README.md)
  - [Biz](./internal/dazBlog/biz/README.md)
  - [Controller](./internal/dazBlog/controller/README.md)

## 日志系统
[基于 zap 构建可自定义的日志系统](./internal/pkg/log/README.md)

## 版本信息
[打印详细版本信息](./pkg/version/README.md)

## 认证授权系统
[dBlog的认证与授权](./docs/devel/zh-CN/conversions/auth.md)