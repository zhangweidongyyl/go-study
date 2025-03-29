package dataloader

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

// User 测试用用户结构体
type User struct {
	ID   int
	Name string
}

// MockUserDB 模拟数据库实现
type MockUserDB struct {
	// 通过闭包自定义查询行为
	GetUsersByIDsFunc func(ids []int) ([]User, error)
}

func (m *MockUserDB) GetUsersByIDs(ids []int) ([]User, error) {
	return m.GetUsersByIDsFunc(ids)
}

// 测试用例1：正常批量查询
func TestBatchUserLoader_NormalCase(t *testing.T) {
	// 初始化模拟数据库（返回所有请求用户）
	mockDB := &MockUserDB{
		GetUsersByIDsFunc: func(ids []int) ([]User, error) {
			users := make([]User, 0, len(ids))
			for _, id := range ids {
				users = append(users, User{
					ID:   id,
					Name: fmt.Sprintf("User%d", id),
				})
			}
			return users, nil
		},
	}

	// 执行批量加载
	ctx := &gin.Context{}
	ids := []int{1, 2, 3}
	result, err := BatchUserLoader(ctx, mockDB, ids)

	// 验证结果
	assert.NoError(t, err)
	assert.Equal(t, 3, len(result))
	assert.Equal(t, "User2", result[2].Name)
}

// 测试用例2：数据库返回错误
func TestBatchUserLoader_DBError(t *testing.T) {
	// 模拟数据库返回错误
	expectedErr := errors.New("connection timeout")
	mockDB := &MockUserDB{
		GetUsersByIDsFunc: func(ids []int) ([]User, error) {
			return nil, expectedErr
		},
	}

	// 执行批量加载
	ctx := &gin.Context{}
	ids := []int{1, 2, 3}
	result, err := BatchUserLoader(ctx, mockDB, ids)

	// 验证错误传递
	assert.ErrorContains(t, err, "connection timeout")
	assert.Nil(t, result)
}

// 测试用例3：部分ID不存在
func TestBatchUserLoader_PartialResults(t *testing.T) {
	// 模拟只返回偶数ID用户
	mockDB := &MockUserDB{
		GetUsersByIDsFunc: func(ids []int) ([]User, error) {
			users := make([]User, 0)
			for _, id := range ids {
				if id%2 == 0 {
					users = append(users, User{ID: id})
				}
			}
			return users, nil
		},
	}

	// 执行批量加载
	ctx := &gin.Context{}
	ids := []int{1, 2, 3, 4, 5}
	result, err := BatchUserLoader(ctx, mockDB, ids)

	// 验证部分结果
	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))  // 应返回2个用户（ID 2和4）
	assert.NotContains(t, result, 1) // 确认不包含ID 1
}

// 测试用例4：处理空请求
func TestBatchUserLoader_EmptyIDs(t *testing.T) {
	called := false
	mockDB := &MockUserDB{
		GetUsersByIDsFunc: func(ids []int) ([]User, error) {
			called = true
			return []User{}, nil
		},
	}

	// 执行空ID列表查询
	ctx := &gin.Context{}
	result, err := BatchUserLoader(ctx, mockDB, []int{})

	// 验证行为
	assert.NoError(t, err)
	assert.False(t, called) // 预期不调用数据库
	assert.Empty(t, result)
}

// 测试用例5：自动去重处理
func TestBatchUserLoader_Deduplication(t *testing.T) {
	callCount := 0
	mockDB := &MockUserDB{
		GetUsersByIDsFunc: func(ids []int) ([]User, error) {
			callCount++
			// 验证接收到的ID已去重
			assert.ElementsMatch(t, []int{1, 2}, ids)
			return []User{{ID: 1}, {ID: 2}}, nil
		},
	}

	// 执行包含重复ID的查询
	ctx := &gin.Context{}
	ids := []int{1, 2, 2, 1, 1}
	result, err := BatchUserLoader(ctx, mockDB, ids)

	// 验证结果
	assert.NoError(t, err)
	assert.Equal(t, 1, callCount) // 确保只调用一次
	assert.Equal(t, 2, len(result))
}

// 测试用例6：并发安全性验证
func TestBatchUserLoader_ConcurrentAccess(t *testing.T) {
	var (
		mutex   sync.Mutex
		callIDs []int
		mockDB  = &MockUserDB{
			GetUsersByIDsFunc: func(ids []int) ([]User, error) {
				mutex.Lock()
				defer mutex.Unlock()
				callIDs = append(callIDs, ids...)
				return []User{{ID: 999}}, errors.New("mock error")
			},
		}
		wg  sync.WaitGroup
		n   = 100
		ctx = &gin.Context{}
	)

	// 启动并发请求
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(id int) {
			defer wg.Done()
			_, _ = BatchUserLoader(ctx, mockDB, []int{id})
		}(i)
	}
	wg.Wait()

	// 验证合并效果
	mutex.Lock()
	defer mutex.Unlock()
	assert.GreaterOrEqual(t, len(callIDs), n) // 至少包含所有ID
}

// 测试用例7：上下文超时处理
func TestBatchUserLoader_ContextTimeout(t *testing.T) {
	// 创建带超时的上下文
	ctx, cancel := gin.CreateTestContext(nil)
	ctx.Request = httptest.NewRequest("GET", "/", nil)
	ctx.Set("timeout", true)

	// 模拟慢查询
	mockDB := &MockUserDB{
		GetUsersByIDsFunc: func(ids []int) ([]User, error) {
			time.Sleep(200 * time.Millisecond) // 超过上下文超时时间
			return []User{{ID: 1}}, nil
		},
	}

	// 设置上下文超时
	timeoutCtx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()
	ctx.Request = ctx.Request.WithContext(timeoutCtx)

	// 执行查询
	ids := []int{1}
	_, err := BatchUserLoader(ctx, mockDB, ids)

	// 验证超时错误
	assert.ErrorContains(t, err, "context deadline exceeded")
}
