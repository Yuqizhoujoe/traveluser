package main

import (
	"context"
	"log"
	"net"
	"os"
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

	PORT := os.Getenv("PORT")
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("User service is running on port: ", PORT)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
