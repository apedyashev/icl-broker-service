package repository

import (
	"context"
	"fmt"
	"icl-broker/pkg/adapter/grpc/pb"
	infraService "icl-broker/pkg/adapter/infra-service"
	"icl-broker/pkg/domain"
	"icl-broker/pkg/model"
)

type userRepository struct {
	httpClient     infraService.HttpClient
	authGrpcClient pb.UserServiceClient
}

func NewUserRepository(authGrpcClient pb.UserServiceClient) domain.UserRepository {
	return &userRepository{
		httpClient:     infraService.NewHttpClient(),
		authGrpcClient: authGrpcClient,
	}
}

func (r *userRepository) UserByCredentials(cred *domain.LoginCredentials) (*model.User, error) {
	respUser, err := r.authGrpcClient.GetByCredentials(context.Background(), &pb.GetByCredentialsRequest{
		Email:    cred.Email,
		Password: cred.Password,
	})
	if err != nil {
		fmt.Println("error calling auth auth:GetByCredentials via gRPC", err)
		return nil, err
	}

	// TODO: create a mapper to convert pb.User to models.User
	user := model.User{
		ID:        uint(respUser.ID),
		Name:      respUser.Name,
		Username:  respUser.Username,
		Email:     respUser.Email,
		CreatedAt: respUser.CreatedAt.AsTime(),
		UpdatedAt: respUser.UpdatedAt.AsTime(),
	}
	return &user, nil
}

func (r *userRepository) Create(u *domain.RegisterRequestBody) (*model.User, error) {
	respUser, err := r.authGrpcClient.Create(context.TODO(), &pb.CreateRequest{
		Name:     u.Name,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	})
	if err != nil {
		fmt.Println("error calling auth auth:Create via gRPC", err)
		return nil, err
	}

	// TODO: create a mapper to convert pb.User to models.User
	user := model.User{
		ID:        uint(respUser.ID),
		Name:      respUser.Name,
		Username:  respUser.Username,
		Email:     respUser.Email,
		CreatedAt: respUser.CreatedAt.AsTime(),
		UpdatedAt: respUser.UpdatedAt.AsTime(),
	}
	return &user, nil
}

func (r *userRepository) UserById(uid uint) (*model.User, error) {
	respUser, err := r.authGrpcClient.GetById(context.Background(), &pb.GetByIdRequest{
		ID: uint64(uid),
	})
	if err != nil {
		fmt.Println("error calling auth auth:GetById via gRPC", err)
		return nil, err
	}

	// TODO: create a mapper to convert pb.User to models.User
	user := model.User{
		ID:        uint(respUser.ID),
		Name:      respUser.Name,
		Username:  respUser.Username,
		Email:     respUser.Email,
		CreatedAt: respUser.CreatedAt.AsTime(),
		UpdatedAt: respUser.UpdatedAt.AsTime(),
	}
	return &user, nil
}
