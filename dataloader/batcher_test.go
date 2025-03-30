package dataloader

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

// 辅助工具：创建测试用的 Gin Context
func createTestContext() *gin.Context {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(nil)
	ctx.Request = httptest.NewRequest("GET", "/", nil)
	return ctx
}

// 测试 BatchFunc 的模拟实现
type mockBatchFunc[K comparable, V any] struct {
	mu     sync.Mutex
	called int
	result map[K]V
	err    error
	panic  bool
}

func (m *mockBatchFunc[K, V]) fn(_ *gin.Context, keys []K) (map[K]V, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.called++

	if m.panic {
		panic("mock panic")
	}

	if m.err != nil {
		return nil, m.err
	}

	result := make(map[K]V)
	for _, k := range keys {
		if v, ok := m.result[k]; ok {
			result[k] = v
		}
	}
	return result, nil
}

/******************** 测试用例实现 ********************/

// 测试1: 超时触发批处理
func TestBatchTimeout(t *testing.T) {
	ctx := createTestContext()
	ctx.Set(batchCtx, &batchFactory{batch: &sync.Map{}})

	mock := &mockBatchFunc[string, int]{
		result: map[string]int{"test": 42},
	}

	// 创建 batcher: 容量5，等待50ms
	batcher := batch[string, int](ctx, "timeout_test", mock.fn, 5, 50*time.Millisecond, NoopTracer[string, int]{})

	// 发送单个请求
	resultChan := batcher.Load("test")

	// 等待超过50ms
	time.Sleep(100 * time.Millisecond)

	select {
	case res := <-resultChan:
		assert.NoError(t, res.Error)
		assert.Equal(t, 42, res.Data)
		assert.Equal(t, 1, mock.called)
	default:
		t.Fatal("预期收到结果但通道未关闭")
	}
}

// 测试2: 容量触发批处理
func TestBatchCapacityTrigger(t *testing.T) {
	ctx := createTestContext()
	ctx.Set(batchCtx, &batchFactory{batch: &sync.Map{}})

	mock := &mockBatchFunc[int, string]{
		result: map[int]string{1: "a", 2: "b", 3: "c"},
	}

	// 容量3，等待1小时（测试不应等待）
	batcher := batch[int, string](ctx, "capacity_test", mock.fn, 3, time.Hour, NoopTracer[int, string]{})

	var wg sync.WaitGroup
	results := make([]*Result[string], 3)

	// 发送3个请求
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			results[idx] = <-batcher.Load(idx + 1)
		}(i)
	}

	wg.Wait()

	assert.Equal(t, 1, mock.called)
	for i := 0; i < 3; i++ {
		assert.NoError(t, results[i].Error)
	}
}

// 测试3: 错误处理
func TestBatchErrorHandling(t *testing.T) {
	ctx := createTestContext()
	ctx.Set(batchCtx, &batchFactory{batch: &sync.Map{}})

	expectedErr := errors.New("mock error")
	mock := &mockBatchFunc[string, bool]{err: expectedErr}

	batcher := batch[string, bool](ctx, "error_test", mock.fn, 2, 50*time.Millisecond, NoopTracer[string, bool]{})

	errChan1 := batcher.Load("key1")
	errChan2 := batcher.Load("key2")

	res1 := <-errChan1
	res2 := <-errChan2

	assert.ErrorIs(t, res1.Error, expectedErr)
	assert.ErrorIs(t, res2.Error, expectedErr)
	assert.Equal(t, 1, mock.called)
}

// 测试4: Panic处理
func TestBatchPanicHandling(t *testing.T) {
	ctx := createTestContext()
	ctx.Set(batchCtx, &batchFactory{batch: &sync.Map{}})

	// 重定向log输出到测试logger
	defer func(orig *log.Logger) {
		log.SetOutput(orig.Writer())
	}(log.Default())
	log.SetOutput(testLogger{t})

	mock := &mockBatchFunc[int, interface{}]{panic: true}
	batcher := batch[int, interface{}](ctx, "panic_test", mock.fn, 5, 10*time.Millisecond, NoopTracer[int, interface{}]{})

	result := <-batcher.Load(123)

	assert.Contains(t, result.Error.Error(), "panic received")
	assert.Equal(t, 1, mock.called)
}

// 测试5: 并发安全
func TestConcurrentSafety(t *testing.T) {
	ctx := createTestContext()
	ctx.Set(batchCtx, &batchFactory{batch: &sync.Map{}})

	const numRequests = 100
	mock := &mockBatchFunc[int, int]{
		result: make(map[int]int),
	}
	for i := 0; i < numRequests; i++ {
		mock.result[i] = i * 2
	}

	batcher := batch[int, int](ctx, "concurrency_test", mock.fn, 20, 50*time.Millisecond, NoopTracer[int, int]{})

	var wg sync.WaitGroup
	results := make([]*Result[int], numRequests)

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			results[idx] = <-batcher.Load(idx)
		}(i)
	}

	wg.Wait()

	// 验证所有结果
	for i := 0; i < numRequests; i++ {
		assert.NoError(t, results[i].Error)
		assert.Equal(t, i*2, results[i].Data)
	}

	// 验证批处理次数：100/20 = 5次（需要根据实际触发逻辑调整）
	assert.True(t, mock.called >= 5 && mock.called <= 6) // 包含可能的超时批次
}

// 测试6: 键去重
func TestKeyDeduplication(t *testing.T) {
	ctx := createTestContext()
	ctx.Set(batchCtx, &batchFactory{batch: &sync.Map{}})

	mock := &mockBatchFunc[string, int]{
		result: map[string]int{"a": 1, "b": 2},
	}

	batcher := batch[string, int](ctx, "dedup_test", mock.fn, 5, 100*time.Millisecond, NoopTracer[string, int]{})

	// 发送重复请求
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-batcher.Load("a")
			<-batcher.Load("b")
		}()
	}
	wg.Wait()

	// 应只调用一次，且传入的keys是去重后的 ["a", "b"]
	mock.mu.Lock()
	defer mock.mu.Unlock()
	assert.Equal(t, 1, mock.called)
}

/******************** 辅助工具 ********************/

// 测试专用的 logger
type testLogger struct{ t *testing.T }

func (l testLogger) Write(p []byte) (n int, err error) {
	l.t.Logf("[LOG] %s", string(p))
	return len(p), nil
}
