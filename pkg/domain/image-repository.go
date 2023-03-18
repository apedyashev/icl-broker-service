package domain

import "icl-broker/pkg/model"

type ImageCreateDTO struct {
	PostId  uint   `json:"postId"`
	Content string `json:"content"`
}

type ImageRepository interface {
	Create(imageCreateDto *ImageCreateDTO) (*model.Image, error)
	ImagesByPostId(postId uint) ([]*model.Image, error)
}
