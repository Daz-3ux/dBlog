## 错误规范

### 使用语义化两级错误码
  - 语义化: 通过错误码可以直接知道错误的原因
  - 两级: 平台级.资源级
    - 平台级是固定的,客户端可根据其进行通用的错误处理
    - 资源级可以根据需要自定义错误吗,也可使用默认错误码

### 错误码示例
```
{
  "code": "InvalidParameter.BadAuthenticationData",
  "message": "Bad Authentication data."
}
```
### 平台级错误吗
| 错误码              | 错误描述               | 错误类型 |
|------------------|--------------------|------|
| InternalError    | 内部错误               | 1    |
| InvalidParameter | 参数错误(参数类型/格式/值...) | 0    |
| AuthFailure      | 认证/授权错误            | 0    |
| ResourceNotFound | 资源不存在              | 0    |
| FailedOperation  | 操作失败               | 2    |

- 0 为客户端错误
- 1 为服务端错误
- 2 代表客户端/服务端均有可能


