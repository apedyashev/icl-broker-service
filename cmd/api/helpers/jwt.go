package helpers

import (
	"os"

	"github.com/go-chi/jwtauth"
)

var TokenAuth *jwtauth.JWTAuth

func init() {
	jwtSecret := os.Getenv("JWT_SECRET")
	TokenAuth = jwtauth.New("HS256", []byte(jwtSecret), nil)
}

func GenerateToken(email string) (string, error) {
	// claims containing email will be enough
	_, tokenString, err := TokenAuth.Encode(map[string]interface{}{"email": email})
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
