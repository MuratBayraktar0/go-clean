package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"golang-clean/domain"
)

// Mock repository for user
type mockUserRepo struct {
	mock.Mock
}

func (m *mockUserRepo) Get(ctx context.Context, id string) (*domain.UserEntity, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.UserEntity), args.Error(1)
}

func (m *mockUserRepo) GetAll(ctx context.Context) (*[]domain.UserEntity, error) {
	args := m.Called(ctx)
	return args.Get(0).(*[]domain.UserEntity), args.Error(1)
}

func (m *mockUserRepo) Create(ctx context.Context, user domain.UserEntity) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *mockUserRepo) Update(ctx context.Context, user domain.UserEntity) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

// ... Implement other repository methods ...

func TestUsecase_Get(t *testing.T) {
	mockRepo := &mockUserRepo{}

	userID := uuid.New().String()
	mockEntity := &domain.UserEntity{
		ID:        userID,
		Name:      "John Doe",
		Email:     "john@example.com",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	mockRepo.On("Get", mock.Anything, userID).Return(mockEntity, nil)

	usecase := NewUserUsecase(mockRepo, time.Second*5)
	ctx := context.Background()

	result, err := usecase.Get(ctx, userID)

	assert.NoError(t, err)
	assert.Equal(t, mockEntity.Name, result.Name)
	assert.Equal(t, mockEntity.Email, result.Email)
	mockRepo.AssertExpectations(t)
}

func TestUsecase_GetAll(t *testing.T) {
	mockRepo := &mockUserRepo{}

	mockUsers := []domain.UserEntity{
		{ID: "1", Name: "John Doe", Email: "john@example.com", UpdatedAt: time.Now(), CreatedAt: time.Now()},
		{ID: "2", Name: "Jane Doe", Email: "jane@example.com", UpdatedAt: time.Now(), CreatedAt: time.Now()},
	}

	mockRepo.On("GetAll", mock.Anything).Return(&mockUsers, nil)

	usecase := NewUserUsecase(mockRepo, time.Second*5)
	ctx := context.Background()

	results, err := usecase.GetAll(ctx)

	assert.NoError(t, err)
	assert.Len(t, *results, len(mockUsers))

	for i, user := range *results {
		assert.Equal(t, mockUsers[i].Name, user.Name)
		assert.Equal(t, mockUsers[i].Email, user.Email)
	}

	mockRepo.AssertExpectations(t)
}

func TestUsecase_Create(t *testing.T) {
	mockRepo := &mockUserRepo{}
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("domain.UserEntity")).Return(nil)

	usecase := NewUserUsecase(mockRepo, time.Second*5)
	ctx := context.Background()

	userDTO := domain.UserDTO{
		Name:  "Jane Doe",
		Email: "jane@example.com",
	}

	err := usecase.Create(ctx, userDTO)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

}

func TestUsecase_Update(t *testing.T) {
	mockRepo := &mockUserRepo{}

	userID := uuid.New().String()
	userDTO := domain.UserDTO{
		Name:  "Updated Name",
		Email: "updated@example.com",
	}

	mockRepo.On("Update", mock.Anything, mock.AnythingOfType("domain.UserEntity")).Return(nil)

	usecase := NewUserUsecase(mockRepo, time.Second*5)
	ctx := context.Background()

	err := usecase.Update(ctx, userID, userDTO)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
