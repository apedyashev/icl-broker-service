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

type AuthController interface {
	Login(ctx echo.Context) error
	Register(ctx echo.Context) error
}

type authController struct {
	authService service.Auth
}

func NewAuthController(authGrpcClients *grpc.GRPCAuthClients) AuthController {
	return &authController{
		authService: service.NewAuthService(
			domain.NewUserService(
				repository.NewUserRepository(authGrpcClients.User),
				infraService.NewJwtService(),
			),
		),
	}
}

func (c *authController) Login(ctx echo.Context) error {
	response, err := c.authService.Login(ctx)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}
	return ctx.JSON(http.StatusNotImplemented, response)
}

func (c *authController) Register(ctx echo.Context) error {
	response, err := c.authService.Register(ctx)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}
	return ctx.JSON(http.StatusNotImplemented, response)
}
