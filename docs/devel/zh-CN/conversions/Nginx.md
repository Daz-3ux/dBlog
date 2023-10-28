## 通过 Nginx 实现 dBlog 的高可用

### 安装并启动 Nginx
```shell
pacman -S nginx
systemctl start nginx
systemctl enable nginx
# 查看 Nginx 状态
systemctl status nginx
```

### 设置`反向代理`
- 反向代理（Reverse Proxy）是指以代理服务器来接收 Internet 上的连接请求，然后将请求转发给内部网络上的服务器
- 并将从服务器上得到的结果返回给 Internet 上请求连接的客户端，此时代理服务器对外就表现为一个反向代理服务器
- 配置文件：
  - 假设 API 服务器的域名为 dazblog.com
  - 在 /etc/nginx/conf.d/ 目录下创建 dazblog.conf 文件
  - server 部分配置如下：
```text
    server {                     
        listen      80;           
        server_name  dazblog.com;    
        client_max_body_size 1024M;
                                 
        location / {             
            proxy_set_header Host $http_host;
            proxy_set_header X-Forwarded-Host $http_host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_pass  http://127.0.0.1:8081/;
            client_max_body_size 100m;
        }
    }
````
- 
	- 完整配置文件如下：
```text
# 运行用户
user  daz;

# 进程数
worker_processes  2;

# 错误日志存放位置
error_log  /var/log/nginx/error.log warn;
# PID 文件: 记录 Nginx 的主进程 ID
pid        /var/run/nginx.pid;

# 工作模式及连接数上限
events {
    worker_connections  1024; # 单个 work_process 进程的最大并发连接数
}

# HTTP 请求相关的配置
http {
    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;
    
    ## 日志格式
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    access_log  /var/log/nginx/access.log  main;

    sendfile        on;
    # tcp_nopush     on;
    
    ## 连接超时时间
    keepalive_timeout  60;
    tcp_nodelay        on;

    #gzip  on;

    include /etc/nginx/conf.d/*.conf;

    server {  
        listen      80;                                                        
        server_name  dazblog.com;                                              
        client_max_body_size 1024M;

        location / {
            proxy_set_header Host $http_host;
            proxy_set_header X-Forwarded-Host $http_host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_pass  http://127.0.0.1:8081/;
            client_max_body_size 5m;
        }
    } 
}
```

- 配置完成后，重启 Nginx
- 在 `etc/hosts` 中添加 `127.0.0.1 ``dazblog.com`

## 配置 Nginx 作为负载均衡
- 负载均衡（Load Balance）是指将请求分发到多个服务器上进行处理，从而共同完成工作任务
- 使用 Nginx 的默认轮询算法，将请求分发到多个服务器上进行处理
  - 为什么不使用随机?轮询天生就比随机平滑
- 配置文件:
  - 修改 /etc/nginx/conf.d/dazblog.conf 文件
```text
	upstream dazblog.com {
		server 127.0.0.1:8081;
		server 127.0.0.1:8082;
		server 127.0.0.1:8083;
	}	
```
 - 在 server 字段中将 `proxy_pass  http://127.0.0.1:8081/;` 修改为 `proxy_pass http://dazblog.com/;`
- 重启 Nginx

## 使用 Keepalived 实现 Nginx 的高可用
- Nginx 可以保证后端 API 服务的高可用
- Nginx 本身是单点的,需要使用 Keepalived 实现高可用,防止单点故障
  - 使用 Keepalived 的抢占模式: Master 从故障中恢复后，会将 VIP 从 Backup 节点中抢占过来
- 配置方案
  - 两台服务器
  - 一台作为主服务器,一台作为备份服务器
  - 主服务器宕机后,备份服务器接管。主服务器恢复后,备份服务器退回
  - 两台服务器的 Nginx 配置文件一致

![VIP](../../../../internal/resource/vip.jpg)