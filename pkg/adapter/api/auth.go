package api

import (
	infraService "icl-broker/pkg/adapter/infra-service"
	"icl-broker/pkg/adapter/repository"
	"icl-broker/pkg/domain"
	"icl-broker/pkg/service"
	"net/http"

	"github.com/labstack/echo"
)

type AuthController interface {
	Login(ctx echo.Context) error
	Register(ctx echo.Context) error
}

type authController struct {
	authService service.Auth
}

func NewAuthController() AuthController {
	return &authController{
		authService: service.NewAuthService(
			domain.NewUserService(
				repository.NewUserRepository(),
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
