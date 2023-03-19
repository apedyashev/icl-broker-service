package service

import (
	"errors"
	"fmt"
	"icl-broker/pkg/domain"
	"icl-broker/pkg/model"

	"github.com/golang-jwt/jwt/v4"
)

type UserService interface {
	UserByClaims(ctx Context) (*model.User, error)
}

type userService struct {
	userDomainService domain.UserService
}

func NewUserService(userDomainService domain.UserService) UserService {
	return &authSerice{userDomainService}
}

func (s *authSerice) UserByClaims(ctx Context) (*model.User, error) {
	claims := ctx.Get("claims").(jwt.MapClaims)
	fmt.Printf("claims = %+v\n", claims["id"])

	if claims["id"] == nil {
		fmt.Printf("current user ID is 0")
		return nil, errors.New("current user ID is 0")
	}

	userId := claims["id"].(float64)
	user, err := s.userDomainService.UserById(uint(userId))
	if err != nil {
		return nil, err
	}

	return user, nil
}
