package controller

import (
	infraService "icl-broker/pkg/adapter/infra-service"
	"icl-broker/pkg/adapter/repository"
	"icl-broker/pkg/domain"
	"icl-broker/pkg/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	CurrentUser(ctx echo.Context) error
}

type userController struct {
	userService service.UserService
}

func NewUserController() UserController {
	return &userController{
		userService: service.NewUserService(
			domain.NewUserService(
				repository.NewUserRepository(),
				infraService.NewJwtService(),
			),
		),
	}
}

func (c *userController) CurrentUser(ctx echo.Context) error {
	user, err := c.userService.UserByClaims(ctx)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}
	return ctx.JSON(http.StatusOK, user)
}
