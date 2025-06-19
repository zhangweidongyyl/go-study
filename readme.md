origin  ssh://git@github.com/gyuho/learn.git (fetch)


# 练习处理器代码生成器

这个代码生成器用于快速生成练习模块的处理器代码，包括Handler和Executor的实现。

## 功能特点

- 自动生成处理器代码框架
- 支持自定义模块、场景和API名称
- 生成的代码包含完整的类型定义和错误处理
- 自动注册到HandlerExecutorRegistry

## 使用方法

1. 在项目根目录下运行命令：

```bash
go build
 ./study  generate --module primary --scene homework --api getresult
 
--- 先打包 再执行生成
go run main.go generate --module <模块名> --scene <场景名> --api <API名>
```

参数说明：
- `--module, -m`: 练习模块名称（必需），例如：primary, middle
- `--scene, -s`: 练习场景名称（必需），例如：homework
- `--api, -a`: API名称（必需），例如：getabstract

2. 示例：

```bash
# 生成小学作业摘要处理器
go run main.go generate --module primary --scene homework --api getabstract

# 生成中学作业摘要处理器
go run main.go generate --module middle --scene homework --api getabstract
```

3. 生成的文件将保存在 `service/practice/<module>/<scene>/` 目录下

## 注意事项

1. 生成代码后需要：
    - 在 `domain/practice/practice_domain.go` 中定义对应的请求和响应结构体
    - 实现 `Action` 方法中的具体业务逻辑
    - 确保相关的常量已在 `components/consts/practice` 中定义

2. 文件命名规则：
    - 生成的文件名格式为：`<module>_<scene>_<api>.go`
    - 例如：`primary_homework_getabstract.go`

3. 代码结构：
    - 生成的代码包含完整的类型定义和基础实现
    - 包含参数验证和错误处理
    - 自动注册到HandlerExecutorRegistry

## 示例代码

生成的代码将包含以下主要组件：

1. Executor结构体和实现：
```go
type PrimaryHomeworkGetAbstractExecutor struct {
    handler practice.ExerciseHandler[*practicedomain.PrimaryHomeworkGetAbstractRequest, *practicedomain.PrimaryHomeworkGetAbstractResponse]
}
```

2. Handler结构体和实现：
```go
type primaryHomeworkGetAbstractHandler[Req *practicedomain.PrimaryHomeworkGetAbstractRequest, Rsp *practicedomain.PrimaryHomeworkGetAbstractResponse] struct {
}
```

3. 自动注册：
```go
func init() {
    action := NewPrimaryHomeworkGetAbstractHandler()
    practice.RegisterHandlerExecutor(
        practiceconst.ExerciseModulePrimary,
        practiceconst.ExerciseSceneHomework,
        practiceconst.ExerciseApiGetAbstract,
        NewPrimaryHomeworkGetAbstractExecutor(action),
    )
}
```
