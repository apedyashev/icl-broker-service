package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"icl-broker/pkg/adapter/clients"
	"icl-broker/pkg/domain"
	"icl-broker/pkg/model"
	"log"

	"net/http"
	"time"
)

type postRepository struct {
	httpClient clients.HttpClient
}

func NewPostRepository() domain.PostRepository {
	return &postRepository{
		httpClient: clients.NewHttpClient(),
	}
}

type PostCreateResponseDTO struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UserId      uint      `json:"userId"`
	Description string    `json:"description"`
}

func (r *postRepository) UserPosts(userId uint, qp *domain.PostsFiltersDTO) ([]*domain.PostCompact, error) {
	BeforeDecode := func(response *http.Response) error {
		fmt.Println("status code", response.StatusCode)
		if response.StatusCode != http.StatusOK {
			return errors.New(fmt.Sprintf("Response to posts-service/posts failed: status=%d", response.StatusCode))
		}
		return nil
	}

	var postsFromSrvs []domain.ServicePost
	var postUrl = "http://posts-service/posts"
	q := domain.ListPostParamsToUrlValues(qp)
	q.Add("user_id", fmt.Sprint(userId))

	err := r.httpClient.Get(postUrl, &clients.GetConfig{
		BeforeDecode: BeforeDecode,
		QueryParams:  q,
	}, &postsFromSrvs)
	if err != nil {
		return nil, err
	}

	fmt.Printf("posts %+v", postsFromSrvs)
	return domain.ServicePostsToCompactPosts(postsFromSrvs), nil
}

func (r *postRepository) Update(post *domain.ServicePost) (*domain.PostDetailed, error) {
	OnResponse := func(response *http.Response) error {
		fmt.Println("Update status code", response.StatusCode)
		if response.StatusCode != http.StatusOK {
			return errors.New(
				fmt.Sprintf("Response to posts-service/posts/%d failed: status=%d", post.ID, response.StatusCode),
			)
		}
		return nil
	}

	var postsFromSrvs domain.ServicePost
	var postUrl = fmt.Sprintf("http://posts-service/posts/%d", post.ID)
	log.Println("calling", postUrl)

	changes, err := json.Marshal(post)
	if err != nil {
		return nil, err
	}

	err = r.httpClient.Put(postUrl, &clients.PutConfig{
		OnResponse: OnResponse,
		Body:       changes,
	}, &postsFromSrvs)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Update posts %+v", postsFromSrvs)
	return domain.ServicePostToDetailed(postsFromSrvs), nil
}

func (r *postRepository) Create(post *domain.PostCreateDTO) (*model.Post, error) {
	jsonData, _ := json.Marshal(post)

	// call the service
	request, err := http.NewRequest("POST", "http://posts-service/posts", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode != http.StatusCreated {
		return nil, errors.New(fmt.Sprintf("Response to posts-service/posts failed: status=%d", response.StatusCode))
	}
	// create variable we will read responseBody into
	var jsonFromService PostCreateResponseDTO

	// decode the response from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		return nil, err
	}

	var createdPost model.Post
	createdPost.ID = jsonFromService.ID
	createdPost.Description = jsonFromService.Description
	createdPost.UserId = jsonFromService.UserId
	createdPost.CreatedAt = jsonFromService.CreatedAt
	createdPost.UpdatedAt = jsonFromService.UpdatedAt

	return &createdPost, nil
}

func (r *postRepository) AddImage(postId string, image *model.Image) error {
	return nil
}

func (r *postRepository) PostById(postId uint) (*model.Post, error) {
	BeforeDecode := func(response *http.Response) error {
		if response.StatusCode != http.StatusOK {
			return errors.New(fmt.Sprintf("Response to posts-service/posts failed: status=%d", response.StatusCode))
		}
		return nil
	}
	var post model.Post
	var postUrl = fmt.Sprintf("http://posts-service/posts/%d", postId)
	err := r.httpClient.Get(postUrl, &clients.GetConfig{
		BeforeDecode: BeforeDecode,
	}, &post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
