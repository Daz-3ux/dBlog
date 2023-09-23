# store layer
- 仓库层

## 开发流程
- `使用单例工厂模式`
- 创建一个结构体: 用来创建 store 层实例
  - store 建立在 `*gorm.DB` 之上
- 创建一个 New 函数: 初始化 store 层实例
- 创建一个包级别的 store 层变量: 为了方便调用 store 层实例
  - `S *datastore`
- 创建一个 sync.Once 变量: 为了避免实例被多次创建,
  - `once sync.Once`

## 设计思想
- 在 [store.go](./store.go) 中定义 store 层需要实现的模块
  - 在 [user.go](./user.go) 中定义 user 模块需要在 store 层所实现的方法, 并实现
- 从物理上分割多表的代码实现,使代码更易阅读及维护
- store 层依赖 `*gorm.DB` 对象, 因为其是项目无关的,可供第三方引用的动作,所以将其以包的形式放置在 `pkg`
  - [db.go](../../../pkg/db/db.go)
- 将 `initStore` 具体实现放置在 [helper.go](../helper.go) 之中, 使 main 保持简洁