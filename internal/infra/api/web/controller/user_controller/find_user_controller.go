package user_controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rodrigoachilles/auction-go/configuration/rest_err"
	"net/http"
)

func (u *UserController) FindUsers(c *gin.Context) {
	users, err := u.userUseCase.FindUsers(context.Background())
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (u *UserController) FindUserById(c *gin.Context) {
	userId := c.Param("id")

	if err := uuid.Validate(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "userId",
			Message: "Invalid UUID value",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	userData, err := u.userUseCase.FindUserById(context.Background(), userId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, userData)
}
