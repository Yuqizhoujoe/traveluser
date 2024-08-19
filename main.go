package main

import (
	"context"
	"log"
	"net"
	"user-service/server"
	"user-service/service"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ctx := context.Background()
	grpcServer := grpc.NewServer()

	userService, err := service.NewUserService(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize user service: %v", err)
	}
	defer userService.Close()

	server.Register(grpcServer, userService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("User service is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
