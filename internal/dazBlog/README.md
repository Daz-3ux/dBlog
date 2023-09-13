## dazBlog 目录架构
- dazBlog.go
  - Cobra 运行入口
  - 管理子命令
  - web 服务入口
- helper.go
  - 使用 `viper` 读取配置文件

## web 服务框架
- 使用 `gin` 框架
- HTTP 请求处理流程
[HTTP 请求处理流程](../resource/HTTP.jpg)