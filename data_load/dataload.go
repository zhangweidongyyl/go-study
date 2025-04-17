package data_load

import (
	"context"
	"github.com/gin-gonic/gin"
	"sync"
)

// BatchProcessor 通用批处理器
type BatchProcessor[T any, R any] struct {
	ctx            context.Context
	batchSize      int
	maxConcurrency int
	processFunc    func(context.Context, []T) ([]R, error)
	onSuccess      func([]T, []R)
	onError        func([]T, error)
}

// NewBatchProcessor 创建批处理器
func NewBatchProcessor[T any, R any](
	ctx context.Context,
	batchSize int,
	maxConcurrency int,
	processFunc func(context.Context, []T) ([]R, error),
	onSuccess func([]T, []R),
	onError func([]T, error),
) *BatchProcessor[T, R] {
	return &BatchProcessor[T, R]{
		ctx:            ctx,
		batchSize:      batchSize,
		maxConcurrency: maxConcurrency,
		processFunc:    processFunc,
		onSuccess:      onSuccess,
		onError:        onError,
	}
}

// Process 执行批处理
func (bp *BatchProcessor[T, R]) Process(items []T) ([]R, error) {
	if len(items) == 0 {
		return nil, nil
	}

	var (
		results      []R
		errors       []error
		mu           sync.Mutex
		errMu        sync.Mutex
		wg           sync.WaitGroup
		sem          = make(chan struct{}, bp.maxConcurrency)
		totalBatches = (len(items)-1)/bp.batchSize + 1
	)

	for i := 0; i < totalBatches; i++ {
		start := i * bp.batchSize
		end := start + bp.batchSize
		if end > len(items) {
			end = len(items)
		}
		batch := items[start:end]

		wg.Add(1)
		go func(b []T) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			res, err := bp.processFunc(bp.ctx, b)
			if err != nil {
				errMu.Lock()
				errors = append(errors, err)
				errMu.Unlock()
				if bp.onError != nil {
					bp.onError(b, err)
				}
				return
			}

			mu.Lock()
			results = append(results, res...)
			mu.Unlock()

			if bp.onSuccess != nil {
				bp.onSuccess(b, res)
			}
		}(batch)
	}

	wg.Wait()
	close(sem)

	if len(errors) > 0 {
		return results, errors[0] // 可根据需要修改错误返回逻辑
	}
	return results, nil
}
func getResourceInfoBatch(ctx *gin.Context, resourceIds []int64, needUrl bool) ([]videodownload.VideoResourceDomain, error) {
	resourceIds = operation.UniqueInt64(resourceIds)
	if len(resourceIds) == 0 {
		return nil, nil
	}

	processor := batch.NewBatchProcessor[int64, videodownload.VideoResourceDomain](
		ctx,
		GetResourceBatchSize,
		PoolRuntineSize,
		func(c context.Context, ids []int64) ([]videodownload.VideoResourceDomain, error) {
			return getResourceInfoList(c, ids, needUrl)
		},
		func(ids []int64, results []videodownload.VideoResourceDomain) {
			if len(results) > 0 {
				zlog.Infof(ctx, "batch success, params: %+v, results: %+v", ids, results)
			}
		},
		func(ids []int64, err error) {
			zlog.Errorf(ctx, "batch error: %+v, params: %+v", err, ids)
		},
	)

	return processor.Process(resourceIds)
}
