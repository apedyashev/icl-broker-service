package model

import (
	"time"
)

type Post struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UserId      uint      `json:"userId"`
	Description string    `json:"description"`
	Images      []string  `json:"images"`
	Likes       []uint    `json:"likes"`
	// Images      []Image   `json:"images"`
	// Comments    []Comment     `json:"comments"`
}
