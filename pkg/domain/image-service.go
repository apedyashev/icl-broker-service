package domain

import "icl-broker/pkg/model"

type ImageService interface {
	Create(imageCreateDto *ImageCreateDTO) (*model.Image, error)
	ImagesByPostId(postId uint) ([]*model.Image, error)
}
type imageService struct {
	imageRepository ImageRepository
}

func NewImageService(imageRepository ImageRepository) ImageService {
	return &imageService{imageRepository}
}

func (s *imageService) Create(imageCreateDto *ImageCreateDTO) (*model.Image, error) {
	return s.imageRepository.Create(imageCreateDto)
}

func (s *imageService) ImagesByPostId(postId uint) ([]*model.Image, error) {
	return s.imageRepository.ImagesByPostId(postId)
}
