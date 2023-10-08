# 认证与授权

## 认证 - Authentication - Authn
- 采用`基础认证` + `令牌认证`
  - 基础认证: 用户名+密码
  - 令牌认证: JWT
![JWT认证流程](../../../../internal/resource/JWT.png)

### JWT
- Json Web Token
- 由三部分组成: Header, Payload, Signature
- 一个示例:
```text
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ
```
- Header
  - 由两部分组成:
    - Token 的类型
    - Token 采用的加密算法
  - 例如:
```json
{
  "typ": "JWT",
  "alg": "HS256"
}
```
- Payload
  - 由多个键值对组成
    - iss: 签发者
    - sub: 主题
    - aud: 受众
    - exp: 过期时间
    - nbf: 生效时间
    - iat: 签发时间
    - jti: 唯一身份标识
    - 自定义键值对
  - 例如:
```json
{
  "id": 2,
  "username": "daz",
  "nbf": 1527931805,
  "iat": 1527931805
}
```
- Signature
  - 生成方式:
    - 用 Base64 对 Header 和 Payload 进行编码
    - 用 Secret 对编码后的 Header 和 Payload 进行加密, 生成 Signature
  - Secret 相当于一个私钥, 只有服务端知道
    - dBlog 将其存放于 config/dazBlog.yaml 中

### 认证流程
- [token](../../../../pkg/token/token.go) 签发以及解析密钥
- [Authn](../../../../internal/pkg/middleware/authn.go) 中间件进行认证
- 基于以上实现 `POST /login` 以及 `PUT /v1/users/:name/change-password` 接口
  - POST 用于创建
  - PUT 用于更新,是一个幂等操作

## 授权 - Authorization - Authz
- 使用 ACL (Access Control List) 模型进行授权
  - 基于 [casbin](./useCasbin.md) 进行开发
- TODO: 实现使用 RBAC (Role-Based Access Control) 模型进行授权