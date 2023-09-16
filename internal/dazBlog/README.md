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

## 中间件
- 使用 `gin` 框架的中间件

## 跨域
- 同源策略需同时满足以下三个条件
  - 协议相同
    - HTTP
    - HTTPS
  - 域名相同
  - 端口相同
- 使用 Cors 中间件解决跨域问题

## 优雅关停
- 将监听函数放在一个 goroutine 中,使用 channel 通知 goroutine 关停
- 收到信号后 10s 内关闭服务(10s 内将未处理完的请求处理完毕)