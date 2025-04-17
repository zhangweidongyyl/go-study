package data_load

import (
	"context"
	"github.com/gin-gonic/gin"
)

// SequentialPager 串行分页处理器
type SequentialPager[R any] struct {
	ctx             context.Context             // 上下文
	batchSize       int                         // 每批次大小
	processFunc     func(int, int) ([]R, error) // 分页处理函数 (offset, limit)
	onSuccess       func(int, []R)              // 成功回调
	onError         func(int, error)            // 失败回调
	continueOnError bool                        // 遇到错误是否继续
}

// NewSequentialPager 创建分页处理器
func NewSequentialPager[R any](
	ctx context.Context,
	batchSize int,
	processFunc func(int, int) ([]R, error),
	onSuccess func(int, []R),
	onError func(int, error),
	continueOnError bool,
) *SequentialPager[R] {
	return &SequentialPager[R]{
		ctx:             ctx,
		batchSize:       batchSize,
		processFunc:     processFunc,
		onSuccess:       onSuccess,
		onError:         onError,
		continueOnError: continueOnError,
	}
}

// Process 执行分页处理
func (p *SequentialPager[R]) Process() ([]R, error) {
	var (
		allResults []R
		offset     int
	)

	for {
		// 执行分页查询
		results, err := p.processFunc(offset, p.batchSize)
		if err != nil {
			if p.onError != nil {
				p.onError(offset, err)
			}

			if !p.continueOnError {
				return allResults, err
			}

			// 继续处理后续批次
			offset += p.batchSize
			continue
		}

		// 处理空结果（终止条件）
		if len(results) == 0 {
			break
		}

		// 收集结果
		allResults = append(allResults, results...)

		// 执行成功回调
		if p.onSuccess != nil {
			p.onSuccess(offset, results)
		}

		// 更新偏移量
		offset += p.batchSize

		// 优化终止条件：当返回结果小于请求数量时提前终止
		if len(results) < p.batchSize {
			break
		}
	}

	return allResults, nil
}

func getResourceInfoBatch(ctx *gin.Context, needUrl bool) ([]videodownload.VideoResourceDomain, error) {
	pager := batch.NewSequentialPager[videodownload.VideoResourceDomain](
		ctx,
		GetResourceBatchSize,
		func(offset, limit int) ([]videodownload.VideoResourceDomain, error) {
			// 实现具体分页逻辑
			return getResourceInfoByPage(ctx, offset, limit, needUrl)
		},
		func(offset int, results []videodownload.VideoResourceDomain) {
			if len(results) > 0 {
				zlog.Infof(ctx, "Page success [offset:%d limit:%d], results:%d",
					offset, GetResourceBatchSize, len(results))
			}
		},
		func(offset int, err error) {
			zlog.Errorf(ctx, "Page error [offset:%d limit:%d]: %v",
				offset, GetResourceBatchSize, err)
		},
		false,
	)

	return pager.Process()
}
