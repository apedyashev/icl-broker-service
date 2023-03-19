package controller

import (
	"icl-broker/pkg/adapter/repository"
	"icl-broker/pkg/service"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostController interface {
	UserPosts(ctx echo.Context) error
	PostById(ctx echo.Context) error
	Create(ctx echo.Context) error
	AddImageToPost(ctx echo.Context) error
}

type postController struct {
	postService service.Post
}

func NewPostController() PostController {
	return &postController{
		postService: service.NewPostService(
			repository.NewPostRepository(),
			repository.NewImageRepository(),
		),
	}
}

func (c *postController) UserPosts(ctx echo.Context) error {
	posts, err := c.postService.UserPosts(ctx)
	if err != nil {
		log.Println("post:Posts error", err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
	return ctx.JSON(http.StatusOK, posts)
}

func (c *postController) Create(ctx echo.Context) error {
	post, err := c.postService.Create(ctx)
	if err != nil {
		log.Println("post:Create error", err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusCreated, post)
}

func (c *postController) AddImageToPost(ctx echo.Context) error {
	post, err := c.postService.AddImageToPost(ctx)
	if err != nil {
		log.Println("post:Create error", err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusCreated, post)
}

func (c *postController) PostById(ctx echo.Context) error {

	post, err := c.postService.PostById(ctx)
	if err != nil {
		log.Println("PostById error", err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusOK, post)
}
