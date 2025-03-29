package dataloader

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	lru "github.com/hashicorp/golang-lru/v2/expirable"
)

type Cache[K Key, V Value] interface {
	Get(ctx *gin.Context, key K) (V, bool)
	Set(ctx *gin.Context, key K, value V)
	Delete(*gin.Context, K) bool
	Clear()
}

type NoopCache[K Key, V any] struct{}

func (c *NoopCache[K, V]) Get(context.Context, K) (v V, exits bool) { return }

func (c *NoopCache[K, V]) Set(context.Context, K, V) { return }

func (c *NoopCache[K, V]) Delete(context.Context, K) bool { return false }

func (c *NoopCache[K, V]) Clear() { return }

func NewPerRequestCache[K Key, V Value](function string) *PerRequestCache[K, V] {
	return &PerRequestCache[K, V]{
		function: function,
	}
}

type PerRequestCache[K Key, V Value] struct {
	function string
	mu       sync.Mutex
}

func (c *PerRequestCache[K, V]) cache(ctx *gin.Context) *MapOf[K, V] {
	key := fmt.Sprintf("per-request-cache>%s", c.function)
	if cache, exits := ctx.Get(key); exits {
		return cache.(*MapOf[K, V])
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if cache, exits := ctx.Get(key); exits {
		return cache.(*MapOf[K, V])
	}
	cache := &MapOf[K, V]{}
	ctx.Set(key, cache)
	return cache
}

func (c *PerRequestCache[K, V]) Set(ctx *gin.Context, key K, value V) {
	c.cache(ctx).Store(key, value)
}

func (c *PerRequestCache[K, V]) Get(ctx *gin.Context, key K) (V, bool) {
	return c.cache(ctx).Load(key)
}

func (c *PerRequestCache[K, V]) Delete(ctx *gin.Context, key K) bool {
	if _, found := c.Get(ctx, key); found {
		c.cache(ctx).Delete(key)
		return true
	}
	return false
}

func (c *PerRequestCache[K, V]) Clear() {
	return
}

var memCaches = map[string]any{}
var mcMu sync.Mutex

func useCache(ctx *gin.Context) bool {
	use := ctx.GetHeader("data-loader-mem-cache")
	if use == "on" {
		return true
	} else {
		return false
	}

	// return env.GetEnv().Cluster != env.TEST || env.GetShipEnvName() != "base"
}

type CacheConfig[K Key, V Value] struct {
	Expire   time.Duration
	Capacity int
	OnEvict  func(key K, value V)
}

func MemCachedFrom[K Key, V Value](batchFunc BatchFunc[K, V], config CacheConfig[K, V], opts ...Opt[K, V]) *Loader[K, V] {
	loader := From(batchFunc)
	opts = append(opts, WithCache(NewMemCache(loader.Function, config)))
	return loader.WithOpt(opts...)
}

func NewMemCache[K Key, V Value](name string, config CacheConfig[K, V]) *MemCache[K, V] {
	if cache, exits := memCaches[name]; exits {
		return cache.(*MemCache[K, V])
	}
	mcMu.Lock()
	defer mcMu.Unlock()
	if cache, exits := memCaches[name]; exits {
		return cache.(*MemCache[K, V])
	}
	cache := &MemCache[K, V]{
		cache: lru.NewLRU(config.Capacity, config.OnEvict, config.Expire),
	}
	memCaches[name] = cache
	return cache
}

type MemCache[K Key, V Value] struct {
	cache *lru.LRU[K, V]
}

func (c *MemCache[K, V]) Get(ctx *gin.Context, key K) (V, bool) {
	if !useCache(ctx) {
		var v V
		return v, false
	}
	value, exits := c.cache.Get(key)
	return value, exits
}

func (c *MemCache[K, V]) Set(_ *gin.Context, key K, value V) {
	c.cache.Add(key, value)
}

func (c *MemCache[K, V]) Delete(_ *gin.Context, key K) bool {
	c.cache.Remove(key)
	return true
}

func (c *MemCache[K, V]) Clear() {
	c.cache.Purge()
}
