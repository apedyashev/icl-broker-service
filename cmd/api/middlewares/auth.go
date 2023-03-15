package middlewares

import (
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth"
)

func IsUserEnabled(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		// TODO: check if user is enabled
		// get user from Redis cache (by email)
		// if not avalable in redis, then fetch user from the auth service and put to redis
		// validate user (e.g roles, isBlocked etc.)

		// if err != nil {
		// 	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		// 	return
		// }

		fmt.Printf("DEBUG: IsUserEnabled, user %s\n", claims["email"])
		// user is not disabled, pass it through
		next.ServeHTTP(w, r)
	})
}
