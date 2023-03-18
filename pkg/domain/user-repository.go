package domain

import "icl-broker/pkg/model"

type UserRepository interface {
	UserByCredentials(c *LoginCredentials) (*model.User, error)
	Create(u *RegisterRequestBody) (*model.User, error)
}
