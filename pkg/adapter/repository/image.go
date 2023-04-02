package repository

import (
	"fmt"
	"icl-broker/pkg/domain"
	"icl-broker/pkg/model"
	"log"
	"net/rpc"
	"time"
)

type imageRepository struct {
}

func NewImageRepository() domain.ImageRepository {
	return &imageRepository{}
}

// TODO: rename ImageBody to Content
type imageUploadRPCPayload struct {
	PostId    uint
	ImageBody string
}

// type imageUploadRPCResponse struct {
// 	Id    string
// 	Error bool
// }

type imageSrvs struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	PostId    uint      `bson:"postId,omitempty" json:"postId,omitempty"`
	ImageBody string    `bson:"imageBody,omitempty" json:"imageBody,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

func (r *imageRepository) Create(imageCreateDto *domain.ImageCreateDTO) (*model.Image, error) {
	imageRpcClient, err := rpc.Dial("tcp", "images-service:5001")
	if err != nil {
		return nil, err
	}
	defer imageRpcClient.Close()

	rpcPayload := imageUploadRPCPayload{
		PostId:    imageCreateDto.PostId,
		ImageBody: imageCreateDto.Content,
	}
	log.Println("Log via RPC", rpcPayload)

	var result imageSrvs
	// RPCServer - is a struct created in the logger service
	// LogInfo MUST start with a capital letter (i.e it must be exported)
	err = imageRpcClient.Call("RPCServer.SaveImage", rpcPayload, &result)
	log.Printf("Created image from RPC %+v\n", result)
	if err != nil {
		fmt.Println("error calling RPCServer.SaveImage", err)
		return nil, err
	}

	return &model.Image{Id: result.ID}, nil
}

type GetPostImagesPayload struct {
	PostId uint
}

type GetPostImageResponse struct {
	Images []*imageSrvs
}

func (r *imageRepository) ImagesByPostId(postId uint) ([]*model.Image, error) {
	imageRpcClient, err := rpc.Dial("tcp", "images-service:5001")
	if err != nil {
		return nil, err
	}

	rpcPayload := GetPostImagesPayload{
		PostId: postId,
	}
	log.Println("Log via RPC", rpcPayload)

	var result GetPostImageResponse
	// RPCServer - is a struct created in the logger service
	// LogInfo MUST start with a capital letter (i.e it must be exported)
	err = imageRpcClient.Call("RPCServer.GetPostImages", rpcPayload, &result)
	if err != nil {
		fmt.Println("error calling RPCServer.GetPostImages", err)
		return nil, err
	}

	var images []*model.Image
	log.Printf("images received via RPC %+v\n", result)
	for _, image := range result.Images {
		images = append(images, &model.Image{Id: image.ID, PostId: image.PostId, Content: image.ImageBody})
	}
	return images, nil
}
