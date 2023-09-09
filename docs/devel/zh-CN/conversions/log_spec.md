## 日志规范

- 日志包为 `github.com/daz-3ux/dBlog/internal/pkg/log`

- 使用结构化的日志打印格式
  - log.Infow, log.Warnw, log.Errorw ...
  - log.Infow("Update post function called")

- 日志均以大写开头,结尾无需标点符号

- 使用过去时
  - Could not delete X (✔)
  - Cannot delete X (❌)

- 遵循日志规范级别:
  - Debug --> log.Debugw, 调试信息
  - Info --> log.Infow, 重要信息
  - Warn --> log.Warnw, 警告信息
  - Error --> log.Errorw, 错误信息
  - Panic --> log.Panicw, 恐慌错误
  - Fatal --> log.Fatalw, 致命错误\

- 日志设置:
  - 开发环境: 
    - 使用 Debug 模式
    - 格式使用 console / json
    - 开启 caller
  - 生产环境:
    - 使用 Info 模式
    - 格式使用 json
    - 开启 caller
  
- 在记录日志时,不要输出敏感信息!

- 对于上下文:
  - 在具有 context.Context 的函数中,使用 log.L(ctx).Infow() 获取日志记录