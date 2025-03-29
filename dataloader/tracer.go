package dataloader

import (
	"github.com/gin-gonic/gin"
)

type TraceLoadFinishFunc[V Value] func(result *Result[V], cached bool)
type TraceLoadManyFinishFunc[V Value] func(*ResultMany[V])
type TraceBatchFinishFunc func()

type Tracer[K Key, V Value] interface {
	TraceLoad(ctx *gin.Context, key K) TraceLoadFinishFunc[V]
	// TraceLoadMany will trace the calls to LoadMany
	TraceLoadMany(ctx *gin.Context, keys []K) TraceLoadManyFinishFunc[V]
	// TraceBatch will trace data loader batches
	TraceBatch(ctx *gin.Context, keys []K) TraceBatchFinishFunc
}

// NoopTracer is the default (noop) tracer
type NoopTracer[K Key, V Value] struct{}

// TraceLoad is a noop function
func (NoopTracer[K, V]) TraceLoad(ctx *gin.Context, key K) TraceLoadFinishFunc[V] {
	return func(*Result[V], bool) {}
}

// TraceLoadMany is a noop function
func (NoopTracer[K, V]) TraceLoadMany(ctx *gin.Context, keys []K) TraceLoadManyFinishFunc[V] {
	return func(*ResultMany[V]) {}
}

// TraceBatch is a noop function
func (NoopTracer[K, V]) TraceBatch(ctx *gin.Context, keys []K) TraceBatchFinishFunc {
	return func() {}
}
