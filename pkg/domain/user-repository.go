package domain

import "icl-broker/pkg/model"

type UserRepository interface {
	UserById(uid uint) (*model.User, error)
	UserByCredentials(c *LoginCredentials) (*model.User, error)
	Create(u *RegisterRequestBody) (*model.User, error)
}
