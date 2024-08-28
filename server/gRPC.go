package server

import (
	"user-service/service"

	pb "github.com/Yuqizhoujoe/user-service-proto/proto"

	"google.golang.org/grpc"
)

func Register(grpcServer *grpc.Server, userService *service.UserService) {
	pb.RegisterUserServiceServer(grpcServer, userService)
}
