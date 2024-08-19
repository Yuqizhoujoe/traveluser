package server

import (
	"user-service/service"
	pb "user-service/user-service/proto"

	"google.golang.org/grpc"
)

func Register(grpcServer *grpc.Server, userService *service.UserService) {
	pb.RegisterUserServiceServer(grpcServer, userService)
}
