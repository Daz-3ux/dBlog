# dazBlog
dazBlog 是一个基于 Go 语言开发的博客系统  
使用了 `Golang` + `Gin` + `MySQL` + `chatGPT` + `Docker` + `Nginx`

## Features
- 使用了简洁架构,目录结构规范清晰
- 使用众多常用 Go 包
- 具备认证 (Gin JWT) 以及授权 (Casbin) 功能
- 使用 [langchaingo](https://github.com/tmc/langchaingo) 调用 OPENAI, 使用 GPT-3.5-turbo 模型总结文章内容
- 独立封装 log, error 包
- 使用 Makefile 高效管理项目
- 静态代码检查
- 带有单元测试,性能测试,模糊测试,Mock测试
- 实现了众多的 Web 功能
  - HTTP, HTTPS, gRPC
  - 优雅关停,中间件,跨域,异常恢复
- 使用 MariaDB 存储数据
- RESTful API 设计规范以及 OpenAPI 3.0/Swagger 2.0 API 文档
- 支持 Docker 部署
- 接入腾讯公益 404 页面
- 完善的文档

## Installation

#### 自行构建
```shell
git clone git@github.com:Daz-3ux/dBlog.git

cd dBlog

make tool.verify && make ca && make

./_output/platforms/linux/amd64/dBlog -c configs/dazBlog.yaml
```

#### Dockerfile(推荐)
```shell
docker build -t dazblog-image:latest .

docker run --network=host \
-e DB_HOST=your_db_host \
-e DB_PORT=your_db_port \
-e DB_USER=your_db_user \
-e DB_PASSWORD=your_db_password \
-e DB_NAME=your_db_name \
-e OPENAI_API_KEY=sk-xxxxxxxxxxxxxxxxxx \
--restart always \
dazblog-image:latest
```

#### Docker
```shell
docker pull realdaz/dazblog

docker run --network=host \
-e DB_HOST=your_db_host \
-e DB_PORT=your_db_port \
-e DB_USER=your_db_user \
-e DB_PASSWORD=your_db_password \
-e DB_NAME=your_db_name \
-e OPENAI_API_KEY=sk-xxxxxxxxxxxxxxxxxx \
--restart always \
realdaz/dazblog
```

## 基于 Nginx 实现高可用
[Nginx+Keepalived 保证 dBlog 高可用](./docs/devel/zh-CN/conversions/Nginx.md)


## 数据库配置
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
    - 基于 MySQL 触发器会自动删除用户所有博客并清除权限
- 博客管理
  - 创建博客
  - 获取博客列表
  - 获取博客详情
  - 更新博客
  - 删除博客
  - 批量删除博客
- OpenAI 调用
  - 创建 AI 内容分析
    - 调用 OPENAI GPT-3.5-turbo 模型总结文章内容
  - 获取 AI 内容
  - 更新 AI 内容
  - 列出 AI 内容
  - 删除 AI 内容

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

### OPENAI 的调用
[langchaingo](./docs/devel/zh-CN/conversions/GPT.md)

---

## License
[MIT](https://choosealicense.com/licenses/mit/)