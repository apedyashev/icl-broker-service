package domain

import "icl-broker/pkg/model"

type TokenService interface {
	GenerateToken(u *model.User) (string, error)
}

// type tokenService struct {
// }

// func NewTokenService() TokenService {
// 	return &tokenService{}
// }

// func (s *tokenService) GenerateToken() (string, error) {
// 	return "", nil
// }
