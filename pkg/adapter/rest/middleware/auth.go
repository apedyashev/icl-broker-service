package middleware

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func AddClaims() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return errors.New("JWT token missing or invalid")
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return errors.New("failed to cast claims as jwt.MapClaims")
			}

			c.Set("claims", claims)
			return next(c)
		}
	}
}
