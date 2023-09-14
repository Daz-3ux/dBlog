## known

- 放置需要共享的 key
  - X-Request-ID 会同时被 `日志` 以及 `gin.Context` 需要
  - 所以将 key 的名字保存在共享包 `known` 中, 便于使用