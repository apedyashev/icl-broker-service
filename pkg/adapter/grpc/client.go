package grpc

import (
	"fmt"
	"icl-broker/pkg/adapter/grpc/pb"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCAuthClients struct {
	User pb.UserServiceClient
}

func NewAuthClients() *GRPCAuthClients {
	fmt.Println("NewAuthClients")
	authGrpcAddr := os.Getenv("AUTH_SRVS_GRPC_ADDR")
	if authGrpcAddr == "" {
		log.Fatal("AUTH_SRVS_GRPC_ADDR is not set")
	}
	fmt.Println("NewAuthClients1", authGrpcAddr)

	conn, err := grpc.Dial(
		authGrpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	fmt.Println("NewAuthClients---")
	if err != nil {
		fmt.Println("NewAuthClients err")
		log.Fatal("cannot connect to AUTH server via gRPC: ", err)
	}
	log.Printf("Auth server (%s) connected", authGrpcAddr)

	fmt.Println("NewAuthClients end")
	return &GRPCAuthClients{
		User: pb.NewUserServiceClient(conn),
	}

	// defer conn.Close()

	// c := logs.NewLogServiceClient(conn)
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
}
