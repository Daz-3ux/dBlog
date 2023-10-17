# dazBlog
dazBlog 是一个基于 Go 语言开发的博客系统，其目的是为了学习 Go 语言，同时也是为了学习如何开发一个完整的项目。  

## Features
- 使用了简洁架构,目录结构规范清晰
- 使用众多常用 Go 包
- 具备认证 (JWT) 以及授权 (Casbin) 功能
- 独立封装 log, error 包
- 使用 Makefile 高效管理项目
- 静态代码检查
- 带有单元测试,性能测试,模糊测试,Mock测试
- 实现了众多的 Web 功能
  - HTTP, HTTPS, gRPC
  - 优雅关停,中间件,跨域,异常恢复
- 使用 MariaDB 存储数据
- RESTful API 设计规范以及 OpenAPI 3.0/Swagger 2.0 API 文档
- 使用 Docker 部署项目
- 完善的文档

## Installation
#### 自行构建
```shell
git clone git@github.com:Daz-3ux/dBlog.git

cd dBlog

make tool.verify && make ca && make

./_output/platforms/linux/amd64/dBlog -c configs/dazBlog.yaml
```
#### Dockerfile
```shell
docker build -t dazblog-image:latest .

docker run --network=host \
-e DB_HOST=your_db_host \
-e DB_PORT=your_db_port \
-e DB_USER=your_db_user \
-e DB_PASSWORD=your_db_password \
-e DB_NAME=your_db_name \
--restart always \
dazblog-image:latest
```

#### 数据库配置
[初始化数据库](./docs/devel/zh-CN/conversions/DB.md)

## Documentation
### 实现功能
[openAPI文档](api/openapi/openapi.yaml)
[Postman文档](https://documenter.getpostman.com/view/30435589/2s9YR83t3M)
- 用户管理
  - 用户注册
  - 用户登录
  - 获取用户列表
  - 获取用户详情
  - 更新用户信息
  - 修改用户密码
  - 注销用户
- 博客管理
  - 创建博客
  - 获取博客列表
  - 获取博客详情
  - 更新博客
  - 删除博客
  - 批量删除博客

### 业务架构模型
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

### 日志系统
[基于 zap 构建可自定义的日志系统](./internal/pkg/log/README.md)

### 版本信息
[打印详细版本信息](./pkg/version/README.md)

### 认证授权系统
[dBlog的认证与授权](./docs/devel/zh-CN/conversions/auth.md)

### HTTPS 的使用
[使用HTTPS](./docs/devel/zh-CN/conversions/https.md)

## License
[MIT](https://choosealicense.com/licenses/mit/)