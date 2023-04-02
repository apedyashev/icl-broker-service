package controller

import (
	"icl-broker/pkg/adapter/grpc"
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

func NewUserController(authGrpcClients *grpc.GRPCAuthClients) UserController {
	return &userController{
		userService: service.NewUserService(
			domain.NewUserService(
				repository.NewUserRepository(authGrpcClients.User),
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
