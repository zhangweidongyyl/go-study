package dataloader

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

const batchCtx = "__data_loader_batching__"

type batchFactory struct {
	batch *sync.Map
	lock  sync.Mutex
}

func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(batchCtx, &batchFactory{
			batch: &sync.Map{},
		})
		ctx.Next()
	}
}

func batch[K Key, V Value](ctx *gin.Context, key string, fn BatchFunc[K, V], cap int, wait time.Duration, tracer Tracer[K, V]) *batcher[K, V] {
	factory := ctx.MustGet(batchCtx).(*batchFactory)
	item, exits := factory.batch.Load(key)
	if exits {
		return item.(*batcher[K, V])
	}
	factory.lock.Lock()
	defer factory.lock.Unlock()
	item, exits = factory.batch.Load(key)
	if exits {
		return item.(*batcher[K, V])
	}
	item = &batcher[K, V]{
		requests: make([]*batchRequest[K, V], 0, cap),
		input:    make(chan []*batchRequest[K, V]),
		batchFn:  fn,
		tracer:   tracer,
		cap:      cap,
		wait:     wait,
		ctx:      ctx,
	}
	factory.batch.Store(key, item)
	return item.(*batcher[K, V])
}

// type used to on input channel
type batchRequest[K Key, V Value] struct {
	key     K
	channel chan *Result[V]
}

type batcher[K Key, V Value] struct {
	requests []*batchRequest[K, V]
	input    chan []*batchRequest[K, V]
	batchFn  BatchFunc[K, V]
	tracer   Tracer[K, V]
	cap      int
	wait     time.Duration
	lock     sync.Mutex
	ctx      *gin.Context
}

func (b *batcher[K, V]) collect(req *batchRequest[K, V]) {
	b.lock.Lock()
	defer b.lock.Unlock()
	if len(b.requests) == 0 {
		go b.batch()
		go b.sleep()
	}
	b.requests = append(b.requests, req)
	if len(b.requests) == b.cap {
		b.input <- b.requests
		b.requests = make([]*batchRequest[K, V], 0, b.cap)
		return
	}
}

func (b *batcher[K, V]) sleep() {
	timer := time.NewTimer(b.wait)
	defer timer.Stop()
	select {
	case <-timer.C:
		b.lock.Lock()
		defer b.lock.Unlock()
		if len(b.requests) > 0 {
			b.input <- b.requests
			b.requests = make([]*batchRequest[K, V], 0, b.cap)
		}
	}
}

// execute the batcher of all items in queue
func (b *batcher[K, V]) batch() {
	var (
		keys     = make([]K, 0)
		reqs     = make([]*batchRequest[K, V], 0)
		items    = make(map[K]V, 0)
		panicErr interface{}
	)

	for _, item := range <-b.input {
		keys = append(keys, item.key)
		reqs = append(reqs, item)
	}

	finish := b.tracer.TraceBatch(b.ctx, keys)
	defer finish()
	var err error
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicErr = r
				const size = 64 << 10
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				log.Printf("Dataloader: Panic received in batcher function:: %v\n%s", panicErr, buf)
			}
		}()

		items, err = b.batchFn(b.ctx, lo.Uniq(keys))
	}()

	if panicErr != nil {
		for _, req := range reqs {
			req.channel <- &Result[V]{Error: fmt.Errorf("panic received in batcher function: %v", panicErr)}
			close(req.channel)
		}
		return
	}

	if err != nil {
		for _, req := range reqs {
			req.channel <- &Result[V]{Error: err}
			close(req.channel)
		}
		return
	}

	for _, req := range reqs {
		data, _ := items[req.key]
		req.channel <- &Result[V]{Data: data}
		close(req.channel)
	}
}

// 新增 Load 方法（batcher结构体的成员方法）
func (b *batcher[K, V]) Load(key K) <-chan *Result[V] {
	ch := make(chan *Result[V], 1) // 缓冲通道防止阻塞
	req := &batchRequest[K, V]{
		key:     key,
		channel: ch,
	}
	b.collect(req)
	return ch
}
