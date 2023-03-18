package domain

import (
	"icl-broker/pkg/model"
	"time"
)

type PostDetailed model.Post

type PostCompact struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UserId      uint      `json:"userId"`
	Description string    `json:"description"`
	Images      []string  `json:"images"` // just image IDs
	LikesCount  int       `json:"likesCount"`
}

type ServicePost struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UserId      uint      `json:"userId"`
	Description string    `json:"description"`
	Likers      []uint    `json:"likers"`
	Images      []string  `json:"images"`
	// Comments    []Comment      `json:"comments"`
}

// User
type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequestBody struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
