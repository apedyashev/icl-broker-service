package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	infraService "icl-broker/pkg/adapter/infra-service"
	"icl-broker/pkg/domain"
	"icl-broker/pkg/model"
	"log"
	"net/http"
)

type userRepository struct {
	httpClient infraService.HttpClient
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{
		httpClient: infraService.NewHttpClient(),
	}
}

func (r *userRepository) UserByCredentials(cred *domain.LoginCredentials) (*model.User, error) {
	OnResponse := func(response *http.Response) error {
		fmt.Println("status code", response.StatusCode)
		if response.StatusCode != http.StatusOK {
			return errors.New(
				fmt.Sprintf("Response to auth-service/auth/login failed: status=%d", response.StatusCode),
			)
		}
		return nil
	}

	var postUrl = "http://auth-service/api/auth/login"
	log.Println("calling", postUrl)

	body, err := json.Marshal(cred)
	if err != nil {
		return nil, err
	}

	var user model.User
	err = r.httpClient.Post(postUrl, &infraService.PostConfig{
		OnResponse: OnResponse,
		Body:       body,
	}, &user)

	if err != nil {
		return nil, err
	}

	fmt.Printf("user %+v", user)
	return &user, nil
}

func (r *userRepository) Create(u *domain.RegisterRequestBody) (*model.User, error) {
	OnResponse := func(response *http.Response) error {
		fmt.Println("status code", response.StatusCode)
		if response.StatusCode != http.StatusCreated {
			return errors.New(
				fmt.Sprintf("Response to auth-service/auth/register failed: status=%d", response.StatusCode),
			)
		}
		return nil
	}

	var postUrl = "http://auth-service/api/auth/register"
	log.Printf("calling %s with %+v", postUrl, u)

	body, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	var user model.User
	err = r.httpClient.Post(postUrl, &infraService.PostConfig{
		OnResponse: OnResponse,
		Body:       body,
	}, &user)

	if err != nil {
		fmt.Println("error after Post", err)
		return nil, err
	}

	fmt.Printf("user %+v", user)
	return &user, nil
}
