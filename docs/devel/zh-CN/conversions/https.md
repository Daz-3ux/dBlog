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
- 使用双向认证
  - 服务端验证客户端的合法性
  - 客户端验证服务端的合法性
	![https](https://raw.githubusercontent.com/Daz-Bot/Img-hosting/master/host/202310082207340.png)
- 使用自签证书
  - 签发根证书私钥
	- 生成根证书私钥
	  - `openssl genrsa -out ca.key 4096`
	- 生成请求文件
	  - `openssl req -new -key ca.key -out ca.csr -subj "/C=CN/ST=ShannXi/L=Xi'an/O=devops/OU=XiyouLUG/CN=127.0.0.
		1/emailAddress=daz-3ux@proton.me"`
    - 生成根证书
      - `openssl x509 -req -in ca.csr -signkey ca.key -out ca.crt -days 3650`
  - 生成服务端证书
  	- 生成服务端私钥
  	  - `openssl genrsa -out server.key 4096`
  	- 生成服务端公钥
      - `openssl rsa -in server.key -pubout -out server.pem`
    - 生成服务端向 CA 申请签名的 CSR
      - `openssl req -new -key server.key -out server.csr -subj "/C=CN/ST=ShannXi/L=Xi'an/O=devops/OU=XiyouLUG/CN=127.0.0.1/mailAddress=daz-3ux@proton.me"`
    - 生成服务端带有 CA 签名的证书
      - `openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -in server.csr -out server.crt 
		-days 3650`
  - 生成客户端证书
    - 生成客户端私钥
	  - `openssl genrsa -out client.key 4096`
    - 生成客户端公钥
	  - `openssl rsa -in client.key -pubout -out client.pem`
    - 生成客户端向 CA 申请签名的 CSR
      - `openssl req -new -key client.key -out client.csr -subj "/C=CN/ST=ShannXi/L=Xi'an/O=devops/OU=XiyouLUG/CN=127.0.0.1/mailAddress=daz-3ux@proton.me"`
      - 生成客户端带有 CA 签名的证书
		- `openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -in client.csr -out client.crt -days 3650`