# HTTPS 的使用

- HTTPS 基于 HTTP 协议加入了 SSL 安全层，通过加密的通道进行数据传输，安全性更优

## 使用方式

- 使用 HTTPS 的数据传输加密能力，开启 HTTPS 单向认证，验证服务端合法性
- 服务端对客户端的认证则采用 Bearer 认证

## CA 证书

- CA
  - Certicate Authority
- CA 证书
  - 证书颁发机构的证书
  - 用于验证服务端的合法性
  - 证书 = 公钥（服务端生成密码对中的公钥）+ 申请者与颁发者信息 + 签名（用 CA 机构生成的密码对的私钥进行签名)
  - 颁发证书其实就是使用 CA 的私钥对证书请求签名文件进行签名
- 使用 `X.509` 证书格式
  - 采用 PEM (Privacy Enhanced Mail) 编码
- 使用 `.csr` 作为后缀
  - Certificate Signing Request, 即证书签名请求
  - 并不是`证书`，而是证书请求文件
  - 核心内容是公钥和申请者信息
  - 在生成该申请时,会产生一个私钥,私钥用于后续的证书签名,不需要提供给 CA 机构
- 使用 HTTPS 的数据传输加密能力，开启 HTTPS 单向认证，验证服务端合法性
- 服务端对客户端的认证则采用 JWT Token 认证
- 功能实现流程分以下 3 步
  1. 生成根证书, 服务端证书, 服务端私钥
  	- 在 Makefile 中使用命令自动生成
  2. 配置 `tls` 证书文件和 `HTTPS` 服务监听端口
     - 在 `configs/dazBlog.yaml` 中进行配置
  3. `tls` 证书文件, 并启动 `HTTPS` 服务器
     - 在 `internal/dazBlog/dazBlog` 中进行配置