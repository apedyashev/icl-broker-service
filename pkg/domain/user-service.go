package domain

import "icl-broker/pkg/model"

type UserService interface {
	UserByCredentials(c *LoginCredentials) (*model.User, error)
	UserById(uid uint) (*model.User, error)
	UserToken(u *model.User) (string, error)
	RegisterUser(u *RegisterRequestBody) (*model.User, error)
}

type userService struct {
	userRepository UserRepository
	tokenService   TokenService
}

func NewUserService(userRepository UserRepository, tokenService TokenService) UserService {
	return &userService{userRepository, tokenService}
}

func (s *userService) UserByCredentials(c *LoginCredentials) (*model.User, error) {
	return s.userRepository.UserByCredentials(c)
}

func (s *userService) UserById(uid uint) (*model.User, error) {
	return s.userRepository.UserById(uid)
}

func (s *userService) UserToken(u *model.User) (string, error) {
	return s.tokenService.GenerateToken(u)
}

func (s *userService) RegisterUser(u *RegisterRequestBody) (*model.User, error) {
	return s.userRepository.Create(u)
}
