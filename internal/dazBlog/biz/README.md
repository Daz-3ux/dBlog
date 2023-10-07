## biz layer
- 业务逻辑层
- business logic layer

## 开发流程
- [biz.go](./biz.go) 中存放需要在 biz 层实现的模块
  - [user.go](./user/user.go) 是 user 模块在 biz 层实现具体方法
  - [post.go](./post/post.go) 是 post 模块在 biz 层实现具体方法
- 在构造 `model.UserM` 时, 使用 [copier](https://github.com/jinzhu/copier) 简化代码量
- 将 `CreateUserRequest` 结构体的定义文件 `user.go` 放置在 `pkg/api/dazBlog/v1` 目录下
  - CreateUserRequest 对用户暴露, 作为 `POST /v1/users` 接口的请求 Body,将其放置在 pkg
  - 其是专门用来做请求参数的结构体,所以将其放至 `pkg/api`
  - 考虑后续新加服务,将其放至 `pkg/api/dazBlog`
  - 考虑 API 版本更新,将其放至 `pkg/api/dazBlog/v1`