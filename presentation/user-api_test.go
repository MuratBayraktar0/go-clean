package presentation

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang-clean/application/usecase"
	"golang-clean/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestUserApi_GetAll(t *testing.T) {
	mockRepo := &mockUserRepo{}

	userID1 := uuid.New().String()
	userID2 := uuid.New().String()
	mockEntity := &[]domain.UserEntity{
		{
			ID:        userID1,
			Name:      "John Doe",
			Email:     "john@example.com",
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
		{
			ID:        userID2,
			Name:      "John Doe 2",
			Email:     "john2@example.com",
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
	}

	mockRepo.On("GetAll", mock.Anything).Return(mockEntity, nil)

	usecase := usecase.NewUserUsecase(context.Background(), mockRepo)

	// Create a Fiber app
	app := fiber.New()

	// Create an instance of the UserApi with mock usecase
	userApi := NewUserApi(usecase)

	// Register routes
	userApi.Router(context.Background(), app)

	req := httptest.NewRequest("GET", "/user", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var users []domain.UserDTO
	err := json.NewDecoder(resp.Body).Decode(&users)
	assert.NoError(t, err)

	assert.Len(t, users, 2)
	assert.Equal(t, users[0].ID, userID1)
	assert.Equal(t, users[1].ID, userID2)
	assert.Equal(t, users[0].Name, "John Doe")
	assert.Equal(t, users[1].Name, "John Doe 2")
	assert.Equal(t, users[0].Email, "john@example.com")
	assert.Equal(t, users[1].Email, "john2@example.com")
}

func TestUserApi_Get(t *testing.T) {
	mockRepo := &mockUserRepo{}

	userID := uuid.New().String()
	mockEntity := &domain.UserEntity{

		ID:        userID,
		Name:      "John Doe",
		Email:     "john@example.com",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	mockRepo.On("Get", mock.Anything, mock.AnythingOfType("string")).Return(mockEntity, nil)

	usecase := usecase.NewUserUsecase(context.Background(), mockRepo)

	// Create a Fiber app
	app := fiber.New()

	// Create an instance of the UserApi with mock usecase
	userApi := NewUserApi(usecase)

	// Register routes
	userApi.Router(context.Background(), app)

	req := httptest.NewRequest("GET", "/user/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var user domain.UserDTO
	err := json.NewDecoder(resp.Body).Decode(&user)
	assert.NoError(t, err)

	assert.Equal(t, user.ID, userID)
	assert.Equal(t, user.Name, "John Doe")
	assert.Equal(t, user.Email, "john@example.com")
}

func TestUserApi_Create(t *testing.T) {
	mockRepo := &mockUserRepo{}
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("domain.UserEntity")).Return(nil)

	usecase := usecase.NewUserUsecase(context.Background(), mockRepo)

	// Create a Fiber app
	app := fiber.New()

	// Create an instance of the UserApi with mock usecase
	userApi := NewUserApi(usecase)

	// Register routes
	userApi.Router(context.Background(), app)

	user := domain.UserDTO{Name: "John Doe"}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest("POST", "/user", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var response map[string]string
	err := json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)

	assert.Equal(t, "created", response["message"])
}
