package service

import (
	"errors"
	"fmt"
	"icl-broker/pkg/domain"
	"icl-broker/pkg/model"
	"log"
	"strconv"
)

type Post interface {
	UserPosts(ctx Context) ([]*domain.PostCompact, error)
	Create(ctx Context) (*model.Post, error)
	AddImageToPost(ctx Context) (*domain.PostDetailed, error)
	PostById(ctx Context) (*model.Post, error)
}

type postService struct {
	imageDomainService domain.ImageService
	postDomainService  domain.PostService
}

func NewPostService(postRepository domain.PostRepository, imageRepository domain.ImageRepository) Post {
	return &postService{
		postDomainService:  domain.NewPostService(postRepository),
		imageDomainService: domain.NewImageService(imageRepository),
	}
}

func (s *postService) UserPosts(ctx Context) ([]*domain.PostCompact, error) {
	userIdStr := ctx.Param("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil || userId == 0 {
		log.Println("PostById error", err)
		return nil, errors.New("bad user id")
	}

	var queryParams domain.PostsFiltersDTO
	if err := ctx.Bind(&queryParams); err != nil {
		return nil, err
	}
	fmt.Printf("queryParams %+v\n", queryParams)

	return s.postDomainService.UserPosts(uint(userId), &queryParams)
}

func (s *postService) Create(ctx Context) (*model.Post, error) {
	var createPostDto domain.PostCreateDTO
	if err := ctx.Bind(&createPostDto); err != nil {
		return nil, err
	}
	post, err := s.postDomainService.Create(&createPostDto)
	if err != nil {
		return nil, err
	}

	return post, nil
}

type AddImageToPostRequestDTO struct {
	Content string `json:"content"`
}

func (s *postService) AddImageToPost(ctx Context) (*domain.PostDetailed, error) {
	var requestBody AddImageToPostRequestDTO
	if err := ctx.Bind(&requestBody); err != nil {
		return nil, err
	}

	postIdStr := ctx.Param("id")
	var post *model.Post
	if postIdStr != "" {
		fmt.Println("Post ID is NOT empty")
		postId, err := strconv.Atoi(postIdStr)
		if err != nil {
			log.Println("PostById error", err)
			return nil, errors.New("bad user id")
		}

		post, err = s.postDomainService.PostById(uint(postId))
		fmt.Println("post", post)
		if err != nil {
			return nil, err
		}
	} else {
		fmt.Println("Post ID is empty")
		var emptyPost domain.PostCreateDTO
		var err error
		post, err = s.postDomainService.Create(&emptyPost)
		if err != nil {
			return nil, err
		}
	}

	image, err := s.imageDomainService.Create(&domain.ImageCreateDTO{
		PostId:  post.ID,
		Content: requestBody.Content,
	})
	if err != nil {
		return nil, err
	}

	post.Images = append(post.Images, image.Id)

	responsePost, err := s.postDomainService.Update(post)
	if err != nil {
		log.Printf("post update error %+v\n", err)
		return nil, err
	}

	return responsePost, nil
}

func (s *postService) PostById(ctx Context) (*model.Post, error) {
	postIdStr := ctx.Param("id")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		log.Println("PostById error", err)
		return nil, errors.New("PostById error")
	}
	post, err := s.postDomainService.PostById(uint(postId))
	if err != nil {
		return nil, err
	}

	images, err := s.imageDomainService.ImagesByPostId(post.ID)
	if err != nil {
		log.Println("error from image service", err)
		// return nil, err
	}

	for _, image := range images {
		post.Images = append(post.Images, image.Id)
	}

	return post, nil
}
