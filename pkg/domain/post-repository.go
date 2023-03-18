package domain

import (
	"icl-broker/pkg/model"
)

type PostsFiltersDTO struct {
	// param will be supported by v5
	// UserId uint   `param:"userId"`
	Limit  uint16 `query:"limit"`
	Offset uint16 `query:"offet"`
}

type PostCreateDTO struct {
	UserId      uint   `json:"userId"`
	Description string `json:"description"`
}

type PostRepository interface {
	UserPosts(userId uint, queryParams *PostsFiltersDTO) ([]*PostCompact, error)
	PostById(postId uint) (*model.Post, error)
	Create(post *PostCreateDTO) (*model.Post, error)
	Update(post *ServicePost) (*PostDetailed, error)
	AddImage(postId string, image *model.Image) error
	// AddComment()
}
