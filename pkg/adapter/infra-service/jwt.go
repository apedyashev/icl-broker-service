package infraService

import (
	"icl-broker/pkg/domain"
	"icl-broker/pkg/model"
	"os"

	"github.com/go-chi/jwtauth"
)

var TokenAuth *jwtauth.JWTAuth

func init() {
	jwtSecret := os.Getenv("JWT_SECRET")
	TokenAuth = jwtauth.New("HS256", []byte(jwtSecret), nil)
}

type jwtService struct{}

func NewJwtService() domain.TokenService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(u *model.User) (string, error) {
	// claims containing email will be enough
	_, tokenString, err := TokenAuth.Encode(map[string]interface{}{
		"id":    u.ID,
		"email": u.Email,
	})
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
