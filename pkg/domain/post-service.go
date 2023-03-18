package domain

import "icl-broker/pkg/model"

type PostService interface {
	UserPosts(userId uint, queryParams *PostsFiltersDTO) ([]*PostCompact, error)
	Create(post *PostCreateDTO) (*model.Post, error)
	Update(post *model.Post) (*PostDetailed, error)
	AddImage(postId string, image *model.Image) error
	PostById(postId uint) (*model.Post, error)
}

type postService struct {
	postRepository PostRepository
}

func NewPostService(postRepository PostRepository) PostService {
	return &postService{postRepository}
}

func (s *postService) UserPosts(userId uint, queryParams *PostsFiltersDTO) ([]*PostCompact, error) {
	return s.postRepository.UserPosts(userId, queryParams)
}

func (s *postService) Create(post *PostCreateDTO) (*model.Post, error) {
	return s.postRepository.Create(post)
}

func (s *postService) Update(post *model.Post) (*PostDetailed, error) {
	servicePost := DomainToService(post)
	return s.postRepository.Update(servicePost)
}

func (s *postService) AddImage(postId string, image *model.Image) error {
	return nil
}

func (s *postService) PostById(postId uint) (*model.Post, error) {
	return s.postRepository.PostById(postId)
}
