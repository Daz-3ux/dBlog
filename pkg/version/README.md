## 版本功能
- 通过`go build`时的`-ldflags`参数，将版本信息注入到二进制文件中

### 版本信息输出示例:
```shell
gitVersion: 9440e49                                 
gitCommit: 9440e4978a24f7151c0e39e6aec40d04ed8965ca
gitTreeState: clean                                   
buildDate: 2023-09-11T11:10:22Z                    
goVersion: go1.21.1                                
compiler: gc                                      
platform: linux/amd64
``` 

### 实现细节
  - 在 Makefile 中通过 `go build -ldflags $(GO_LDFLAGS)` 获取部分版本信息
    - 使用 `git describe --tags --always --match='v*'` 获取版本号
    - 使用 `date -u +'%Y-%m-%dT%H:%M:%SZ'` 获取构建时间
    - 使用 `git rev-parse HEAD` 获取构建是 commit ID
    - 通过 `git status --porcelain 2` 获取 git tree state
  - 使用标准库 `runtime` 获取部分信息
    - 使用 `runtime.Version()` 获取 Go 版本信息
    - 使用 `runtime.Compiler` 获取编译器信息
    - 使用 `fmt.Sprintf( "%s/%s" , runtime.GOOS, runtime.GOARCH)` 获取编译环境信息

### --version 选项实现
- 调用 `pflag` 库将 --version 选项添加到 flag 中
- `verflag.AddFlags(cmd.PersistentFlags())`将 `--version` 标志持久化,使其可以在子命令中使用