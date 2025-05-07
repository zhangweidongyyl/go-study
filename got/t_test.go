package got

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 定义存储接口
type Storage interface {
	GetUser(id string) (User, error)
	SaveUser(user User) error
}

// 用户结构体
type User struct {
	ID   string
	Name string
	Age  int
}

// 业务服务，依赖Storage接口
type UserService struct {
	storage Storage
}

// 构造函数：通过依赖注入传递Storage实现
func NewUserService(storage Storage) *UserService {
	return &UserService{storage: storage}
}

// 业务方法：获取用户并更新年龄
func (s *UserService) UpdateUserAge(id string, newAge int) error {
	user, err := s.storage.GetUser(id)
	if err != nil {
		return err
	}
	user.Age = newAge
	return s.storage.SaveUser(user)
}

// 真实存储实现（如MySQL、PostgreSQL）
type RealStorage struct{}

func (s *RealStorage) GetUser(id string) (User, error) {
	// 真实数据库查询逻辑
	return User{ID: id, Name: "Real User", Age: 30}, nil
}

func (s *RealStorage) SaveUser(user User) error {
	// 真实数据库保存逻辑
	return nil
}

type testStorage struct{}

func (s *testStorage) GetUser(id string) (User, error) {
	return User{ID: id, Name: "Test User", Age: 20}, nil
}
func (s *testStorage) SaveUser(user User) error {
	return nil
}

// Mock存储实现
type MockStorage struct {
	// 可记录调用参数或次数
	GetUserCalledWith  string
	SaveUserCalledWith User
}

func (m *MockStorage) GetUser(id string) (User, error) {
	m.GetUserCalledWith = id // 记录调用参数
	if id == "error" {
		return User{}, errors.New("mock get error")
	}
	return User{ID: id, Name: "Mocked User", Age: 25}, nil
}

func (m *MockStorage) SaveUser(user User) error {
	m.SaveUserCalledWith = user // 记录调用参数
	if user.ID == "error" {
		return errors.New("mock save error")
	}
	return nil
}

func TestUpdateUserAge_Success(t *testing.T) {
	// 初始化Mock
	mockStorage := &MockStorage{}
	service := NewUserService(mockStorage)

	// 测试逻辑
	err := service.UpdateUserAge("123", 40)

	// 断言结果
	assert.NoError(t, err)
	assert.Equal(t, "123", mockStorage.GetUserCalledWith)     // 验证GetUser参数
	assert.Equal(t, "123", mockStorage.SaveUserCalledWith.ID) // 验证SaveUser参数
	assert.Equal(t, 40, mockStorage.SaveUserCalledWith.Age)   // 验证年龄更新
}

func TestUpdateUserAge_GetUserError(t *testing.T) {
	mockStorage := &MockStorage{}
	service := NewUserService(mockStorage)

	// 触发错误场景
	err := service.UpdateUserAge("error", 40)

	// 验证错误
	assert.ErrorContains(t, err, "mock get error")
}

func TestUpdateUserAge_SaveError(t *testing.T) {
	mockStorage := &MockStorage{}
	service := NewUserService(mockStorage)

	// 触发保存错误
	err := service.UpdateUserAge("error", 40)

	// 注意：这里需要根据业务逻辑调整测试数据
	// 假设当ID为"error"时，SaveUser返回错误
	assert.ErrorContains(t, err, "mock save error")
}

func TestUpdateUserAge_TestStorage(t *testing.T) {
	testStorage := &testStorage{}
	service := NewUserService(testStorage)

	err := service.UpdateUserAge("test123", 30)
	assert.NoError(t, err)
}
