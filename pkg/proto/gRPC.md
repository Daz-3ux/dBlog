# gRPC 服务器
- 适用于 dazBlog 内部通信
- 模拟场景:
  - dazBlog 配套一个运营系统,运营系统需要通过接口获取所有的用户,进行注册用户统计
  - 这种内部调用最适合使用 gRPC

## gRPC 创建流程
- gRPC 文件头定义
  - 使用 `.proto` 格式的 Protobuf 定义文件来定义 gRPC 服务
- 服务定义
- 接口参数定义