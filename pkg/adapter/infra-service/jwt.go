package infraService

import (
	"icl-broker/pkg/domain"
	"icl-broker/pkg/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtCustomClaims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	// model.Claims
	jwt.RegisteredClaims
}

// var TokenAuth *jwtauth.JWTAuth

// func init() {
// 	jwtSecret := os.Getenv("JWT_SECRET")
// 	TokenAuth = jwtauth.New("HS256", []byte(jwtSecret), nil)
// }

type jwtService struct{}

func NewJwtService() domain.TokenService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(u *model.User) (string, error) {
	claims := &jwtCustomClaims{
		ID:    u.ID,
		Email: u.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	jwtSecret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(jwtSecret))

	// claims containing email will be enough
	// _, tokenString, err := TokenAuth.Encode(map[string]interface{}{
	// 	"id":    u.ID,
	// 	"email": u.Email,
	// })
	// if err != nil {
	// 	return "", err
	// }

	// return tokenString, nil
}
