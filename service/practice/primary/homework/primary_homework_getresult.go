package homework

import (
	practiceconst "exercise/components/consts/practice"
	"exercise/components/errors"
	practicedomain "exercise/domain/practice"
	"exercise/service/practice"
	"git.zuoyebang.cc/pkg/golib/v2/zlog"
	"github.com/gin-gonic/gin"
)

func init() {
	action := NewPrimaryHomeworkGetresultHandler()
	practice.RegisterHandlerExecutor(
		practiceconst.ExerciseModulePrimary,
		practiceconst.ExerciseSceneHomework,
		practiceconst.ExerciseApiGetresult,
		NewPrimaryHomeworkGetresultExecutor(action),
	)
}

// PrimaryHomeworkGetresultExecutor Primary模块 executor
type PrimaryHomeworkGetresultExecutor struct {
	handler practice.ExerciseHandler[*practicedomain.PrimaryHomeworkGetresultRequest, *practicedomain.PrimaryHomeworkGetresultResponse]
}

// NewPrimaryHomeworkGetresultExecutor 创建 Primary模块 executor
func NewPrimaryHomeworkGetresultExecutor(handler practice.ExerciseHandler[*practicedomain.PrimaryHomeworkGetresultRequest, *practicedomain.PrimaryHomeworkGetresultResponse]) *PrimaryHomeworkGetresultExecutor {
	return &PrimaryHomeworkGetresultExecutor{handler: handler}
}

// Execute 执行处理逻辑
func (e *PrimaryHomeworkGetresultExecutor) Execute(ctx *gin.Context) (interface{}, error) {
	// 绑定请求参数
	var reqData = &practicedomain.PrimaryHomeworkGetresultRequest{}
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
	req := &practicedomain.ExerciseRequest[*practicedomain.PrimaryHomeworkGetresultRequest]{
		ExerciseModule: practiceconst.ExerciseModulePrimary,
		ExerciseScene:  practiceconst.ExerciseSceneHomework,
		ExerciseApi:    practiceconst.ExerciseApiGetresult,
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

// NewPrimaryHomeworkGetresultHandler 创建处理器
func NewPrimaryHomeworkGetresultHandler() practice.ExerciseHandler[*practicedomain.PrimaryHomeworkGetresultRequest, *practicedomain.PrimaryHomeworkGetresultResponse] {
	return &primaryhomeworkgetresultHandler[*practicedomain.PrimaryHomeworkGetresultRequest, *practicedomain.PrimaryHomeworkGetresultResponse]{}
}

type primaryhomeworkgetresultHandler[Req *practicedomain.PrimaryHomeworkGetresultRequest, Rsp *practicedomain.PrimaryHomeworkGetresultResponse] struct {
}

func (m primaryhomeworkgetresultHandler[Req, Rsp]) Action(ctx *gin.Context, req *practicedomain.ExerciseRequest[Req]) (*practicedomain.ExerciseResponse[Rsp], error) {
	return nil, nil
}
