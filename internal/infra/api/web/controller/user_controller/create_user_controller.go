package user_controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rodrigoachilles/auction-go/configuration/rest_err"
	"github.com/rodrigoachilles/auction-go/internal/infra/api/web/validation"
	"github.com/rodrigoachilles/auction-go/internal/usecase/user_usecase"
	"net/http"
)

type UserController struct {
	userUseCase user_usecase.UserUseCaseInterface
}

func NewUserController(userUseCase user_usecase.UserUseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var userInputDTO user_usecase.UserInputDTO

	if err := c.ShouldBindJSON(&userInputDTO); err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	userOutputDTO, err := u.userUseCase.CreateUser(context.Background(), userInputDTO)
	if err != nil {
		restErr := rest_err.ConvertError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusCreated, userOutputDTO)
}
