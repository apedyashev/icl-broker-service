package helpers

import (
	"os"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/golang-jwt/jwt/v4"
)

var TokenAuth *jwtauth.JWTAuth

func init() {
	// jwtSecret := os.Getenv("JWT_SECRET")
	// TokenAuth = jwtauth.New("HS256", []byte(jwtSecret), nil)
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func GenerateToken(email string) (string, error) {
	// // claims containing email will be enough
	// _, tokenString, err := TokenAuth.Encode(map[string]interface{}{"email": email})
	// if err != nil {
	// 	return "", err
	// }

	// return tokenString, nil
	// Set custom claims
	claims := &jwtCustomClaims{
		"Jon Snow",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	jwtSecret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(jwtSecret))

}
