package main

// HandlerTemplate 处理器代码模板
const HandlerTemplate = `package {{.Scene}}

import (
	practiceconst "exercise/components/consts/practice"
	"exercise/components/errors"
	practicedomain "exercise/domain/practice"
	"exercise/service/practice"
	"git.zuoyebang.cc/pkg/golib/v2/zlog"
	"github.com/gin-gonic/gin"
)

func init() {
	action := New{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Handler()
	practice.RegisterHandlerExecutor(
		practiceconst.ExerciseModule{{.ModuleUpper}},
		practiceconst.ExerciseScene{{.SceneUpper}},
		practiceconst.ExerciseApi{{.ApiUpper}},
		New{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Executor(action),
	)
}

// {{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Executor {{.ModuleUpper}}模块 executor
type {{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Executor struct {
	handler practice.ExerciseHandler[*practicedomain.{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Request, *practicedomain.{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Response]
}

// New{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Executor 创建 {{.ModuleUpper}}模块 executor
func New{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Executor(handler practice.ExerciseHandler[*practicedomain.{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Request, *practicedomain.{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Response]) *{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Executor {
	return &{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Executor{handler: handler}
}

// Execute 执行处理逻辑
func (e *{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Executor) Execute(ctx *gin.Context) (interface{}, error) {
	// 绑定请求参数
	var reqData = &practicedomain.{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Request{}
	if err := ctx.ShouldBind(&reqData); err != nil {
		zlog.Errorf(ctx, "bind request failed: %v", err)
		return nil, errors.InvalidParamError("参数错误")
	}

	// 参数校验
	if err := reqData.Validate(); err != nil {
		zlog.Errorf(ctx, "validate request failed: %v", err)
		return nil, errors.InvalidParamError("参数错误")
	}

	// 创建请求
	req := &practicedomain.ExerciseRequest[*practicedomain.{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Request]{
		ExerciseModule: practiceconst.ExerciseModule{{.ModuleUpper}},
		ExerciseScene:  practiceconst.ExerciseScene{{.SceneUpper}},
		ExerciseApi:    practiceconst.ExerciseApi{{.ApiUpper}},
		RequestData:    reqData,
	}

	// 执行处理
	rsp, err := e.handler.Action(ctx, req)
	if err != nil {
		zlog.Errorf(ctx, "handle request failed: %v", err)
		return nil, err
	}

	return rsp.ResponseData, nil
}

// New{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Handler 创建处理器
func New{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Handler() practice.ExerciseHandler[*practicedomain.{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Request, *practicedomain.{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Response] {
	return &{{.Module}}{{.Scene}}{{.Api}}Handler[*practicedomain.{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Request, *practicedomain.{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Response]{}
}

type {{.Module}}{{.Scene}}{{.Api}}Handler[Req *practicedomain.{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Request, Rsp *practicedomain.{{.ModuleUpper}}{{.SceneUpper}}{{.ApiUpper}}Response] struct {
}

func (m {{.Module}}{{.Scene}}{{.Api}}Handler[Req, Rsp]) Action(ctx *gin.Context, req *practicedomain.ExerciseRequest[Req]) (*practicedomain.ExerciseResponse[Rsp], error) {
	return nil, nil
}
`
