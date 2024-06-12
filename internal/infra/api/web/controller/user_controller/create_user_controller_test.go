package user_controller

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rodrigoachilles/auction-go/internal/internal_error"
	"github.com/rodrigoachilles/auction-go/internal/usecase/user_usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) CreateUser(ctx context.Context, userInput user_usecase.UserInputDTO) (*user_usecase.UserOutputDTO, *internal_error.InternalError) {
	args := m.Called(ctx, userInput)
	if args.Get(0) != nil {
		return args.Get(0).(*user_usecase.UserOutputDTO), args.Get(1).(*internal_error.InternalError)
	}
	return nil, args.Get(0).(*internal_error.InternalError)
}

func (m *MockUserUseCase) FindUsers(ctx context.Context) ([]user_usecase.UserOutputDTO, *internal_error.InternalError) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		return args.Get(0).([]user_usecase.UserOutputDTO), args.Get(1).(*internal_error.InternalError)
	}
	return nil, args.Get(0).(*internal_error.InternalError)
}

func (m *MockUserUseCase) FindUserById(ctx context.Context, id string) (*user_usecase.UserOutputDTO, *internal_error.InternalError) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*user_usecase.UserOutputDTO), args.Get(1).(*internal_error.InternalError)
	}
	return nil, args.Get(0).(*internal_error.InternalError)
}

type UserControllerTestSuite struct {
	suite.Suite
	userController *UserController
	mockUseCase    *MockUserUseCase
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.mockUseCase = new(MockUserUseCase)
	suite.userController = NewUserController(suite.mockUseCase)
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}

func (suite *UserControllerTestSuite) TestCreateUser_Success() {
	userInputDTO := user_usecase.UserInputDTO{
		Name: "User Test",
	}

	userOutputDTO := user_usecase.UserOutputDTO{
		Id:   "1",
		Name: "User Test",
	}

	suite.mockUseCase.On("CreateUser", mock.Anything, userInputDTO).Return(userOutputDTO).Once()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/user", bytes.NewBuffer([]byte(`{ "name": "User Test" }`)))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.userController.CreateUser(c)

	c.Writer.WriteHeaderNow()

	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestCreateUser_ValidationError() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/user", bytes.NewBuffer([]byte(`{ "name": "" }`)))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.userController.CreateUser(c)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}
