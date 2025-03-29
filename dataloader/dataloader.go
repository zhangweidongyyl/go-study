package dataloader

import (
	"reflect"
	"runtime"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Key interface {
	comparable
	uint64 | string | []byte | byte | int | int32 | uint32 | int64
}

type Value any

type Result[V Value] struct {
	Data  V
	Error error
}

type ResultMany[V Value] struct {
	Data  []V
	Error []error
}

type Thunk[V Value] func() (V, error)

type ThunkMany[V Value] func() ([]V, []error)

type Opt[K Key, V Value] func(loader *Loader[K, V])

func WithCache[K Key, V Value](cache Cache[K, V]) Opt[K, V] {
	return func(loader *Loader[K, V]) {
		loader.cache = cache
	}
}

func WithTracer[K Key, V Value](tracer Tracer[K, V]) Opt[K, V] {
	return func(loader *Loader[K, V]) {
		loader.tracer = tracer
	}
}

func Capacity[K Key, V Value](cap int) Opt[K, V] {
	return func(loader *Loader[K, V]) {
		loader.batchCap = cap
	}
}

func From[K Key, V Value](fn BatchFunc[K, V], opts ...Opt[K, V]) *Loader[K, V] {
	function := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	loader := &Loader[K, V]{
		batchFn:  fn,
		batchCap: 50,
		wait:     time.Millisecond,
		tracer:   &NoopTracer[K, V]{},
		cache:    NewPerRequestCache[K, V](function),
		Function: function,
	}

	return loader.WithOpt(opts...)
}

type BatchFunc[K Key, V Value] func(*gin.Context, []K) (map[K]V, error)

type Option[K Key, V Value] func(*Loader[K, V])

type Loader[K Key, V Value] struct {
	wait      time.Duration
	tracer    Tracer[K, V]
	cache     Cache[K, V]
	batchFn   BatchFunc[K, V]
	batchLock sync.Mutex
	Function  string
	batchCap  int
}

func (l *Loader[K, V]) WithOpt(opts ...Opt[K, V]) *Loader[K, V] {
	for _, opt := range opts {
		opt(l)
	}
	return l
}

func (l *Loader[K, V]) Cache() Cache[K, V] {
	return l.cache
}

func (l *Loader[K, V]) Load(ctx *gin.Context, key K) Thunk[V] {
	finish := l.tracer.TraceLoad(ctx, key)
	if value, found := l.cache.Get(ctx, key); found {
		// zlog.Debug(components.GinContext(l.ctx), fmt.Sprintf("%T hit cache. key: %s", l.cache, key))
		return func() (V, error) {
			defer finish(&Result[V]{Data: value}, found)
			return value, nil
		}
	}

	c := make(chan *Result[V], 1)
	var result struct {
		mu    sync.RWMutex
		value *Result[V]
	}

	thunk := func() (V, error) {
		defer finish(result.value, false)
		result.mu.RLock()
		resultNotSet := result.value == nil
		result.mu.RUnlock()
		if resultNotSet {
			result.mu.Lock()
			if v, ok := <-c; ok {
				result.value = v
			}
			result.mu.Unlock()
		}

		result.mu.RLock()
		defer result.mu.RUnlock()
		if result.value.Error == nil {
			l.cache.Set(ctx, key, result.value.Data)
		}
		return result.value.Data, result.value.Error
	}

	batched := batch(ctx, l.Function, l.batchFn, l.batchCap, l.wait, l.tracer)
	batched.collect(&batchRequest[K, V]{key, c})

	return thunk
}

func (l *Loader[K, V]) LoadMany(ctx *gin.Context, keys []K) ThunkMany[V] {
	finish := l.tracer.TraceLoadMany(ctx, keys)

	var (
		length = len(keys)
		data   = make([]V, length)
		errors = make([]error, length)
		c      = make(chan *ResultMany[V], 1)
		wg     sync.WaitGroup
	)

	resolve := func(i int) {
		defer wg.Done()
		thunk := l.Load(ctx, keys[i])
		result, err := thunk()
		data[i] = result
		errors[i] = err
	}

	wg.Add(length)
	for i := range keys {
		go resolve(i)
	}

	go func() {
		wg.Wait()

		c <- &ResultMany[V]{Data: data, Error: errors}
		close(c)
	}()

	var result struct {
		mu    sync.RWMutex
		value *ResultMany[V]
	}

	thunkMany := func() ([]V, []error) {
		defer finish(result.value)
		result.mu.RLock()
		resultNotSet := result.value == nil
		result.mu.RUnlock()

		if resultNotSet {
			result.mu.Lock()
			if v, ok := <-c; ok {
				result.value = v
			}
			result.mu.Unlock()
		}
		result.mu.RLock()
		defer result.mu.RUnlock()
		return result.value.Data, result.value.Error
	}

	return thunkMany
}
