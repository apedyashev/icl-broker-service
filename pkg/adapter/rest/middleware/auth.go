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

// TODO: implement this!
// func IsUserEnabled(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		_, claims, _ := jwtauth.FromContext(r.Context())
// 		// TODO: check if user is enabled
// 		// get user from Redis cache (by email)
// 		// if not avalable in redis, then fetch user from the auth service and put to redis
// 		// validate user (e.g roles, isBlocked etc.)

// 		// if err != nil {
// 		// 	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
// 		// 	return
// 		// }

// 		fmt.Printf("DEBUG: IsUserEnabled, user %s\n", claims["email"])
// 		// user is not disabled, pass it through
// 		next.ServeHTTP(w, r)
// 	})
// }
