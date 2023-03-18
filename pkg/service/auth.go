package service

import (
	"fmt"
	"icl-broker/pkg/domain"
	"icl-broker/pkg/model"
)

type Auth interface {
	Login(ctx Context) (*LoginResponse, error)
	Register(ctx Context) (*RegisterResponseBody, error)
}

type authSerice struct {
	userDomainService domain.UserService
}

func NewAuthService(userDomainService domain.UserService) Auth {
	return &authSerice{userDomainService}
}

type LoginResponse struct {
	Token string      `json:"token"`
	User  *model.User `json:"user"`
}

func (s *authSerice) Login(ctx Context) (*LoginResponse, error) {
	var credentials domain.LoginCredentials
	err := ctx.Bind(&credentials)
	if err != nil {
		fmt.Printf("credentials bind erro %+v", err)
		return nil, err
	}
	fmt.Printf("credentials %+v", credentials)

	user, err := s.userDomainService.UserByCredentials(&credentials)
	if err != nil {
		return nil, err
	}

	token, err := s.userDomainService.UserToken(user)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{token, user}, nil
}

type RegisterResponseBody struct {
	User *model.User `json:"user"`
}

func (s *authSerice) Register(ctx Context) (*RegisterResponseBody, error) {
	var regRequest domain.RegisterRequestBody
	err := ctx.Bind(&regRequest)
	if err != nil {
		return nil, err
	}

	newUser, err := s.userDomainService.RegisterUser(&regRequest)
	if err != nil {
		return nil, err
	}

	return &RegisterResponseBody{User: newUser}, err
}
