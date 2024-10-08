package service

import (
	"context"
	"user-service/models"
	db "user-service/package"

	pb "github.com/Yuqizhoujoe/user-service-proto/proto"

	firestore "cloud.google.com/go/firestore"
)

type UserService struct {
	pb.UnimplementedUserServiceServer // TODO: why need this one here
	firestoreClient                   *firestore.Client
}

func NewUserService(ctx context.Context) (*UserService, error) {
	client, err := db.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &UserService{
		firestoreClient: client,
	}, nil
}

func (s *UserService) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	docRef := s.firestoreClient.Collection("users").Doc(req.Email)

	// fetch the document snapshot
	doc, err := docRef.Get(ctx)
	if err != nil {
		return nil, err
	}

	// check if the document exists
	if !doc.Exists() {
		// user does not exist, create a new document
		userData := models.User{
			Rooms: []string{},
			Posts: []string{},
		}

		_, err = docRef.Set(ctx, userData)
		if err != nil {
			return nil, err
		}
		return &pb.AddUserResponse{Success: true}, nil
	}

	return &pb.AddUserResponse{Success: true}, nil
}

func (s *UserService) AddRoom(ctx context.Context, req *pb.AddRoomRequest) (*pb.AddRoomResponse, error) {
	docRef := s.firestoreClient.Collection("users").Doc(req.Email)

	// fetch the current document
	doc, err := docRef.Get(ctx)
	if err != nil {
		return nil, err
	}

	// var userData map[string]interface{}
	var userData models.User
	if err := doc.DataTo(&userData); err != nil {
		return nil, err
	}

	// update the rooms slice
	updatedRooms := append(userData.Rooms, req.RoomId)
	// rooms, ok := userData["rooms"].([]interface{})
	// if !ok {
	// 	rooms = []interface{}{}
	// }
	// rooms = append(rooms, req.RoomId)

	// write the updated rooms slice back to firestore
	// _, err = docRef.Set(ctx, map[string]interface{}{
	// 	"rooms": rooms,
	// }, firestore.MergeAll)
	// options := firestore.Merge([]string{"rooms"})
	// _, err = docRef.Set(ctx, map[string]interface{} {
	// 	"rooms": rooms,
	// }, options)
	updatedData := []firestore.Update{
		{Path: "rooms", Value: updatedRooms},
	}
	_, err = docRef.Update(ctx, updatedData)

	if err != nil {
		return nil, err
	}

	return &pb.AddRoomResponse{Success: true}, nil
}

func (s *UserService) AddPost(ctx context.Context, req *pb.AddPostRequest) (*pb.AddPostResponse, error) {
	docRef := s.firestoreClient.Collection("users").Doc(req.Email)

	// fetch the current document
	doc, err := docRef.Get(ctx)
	if err != nil {
		return nil, err
	}

	// retrieve current posts slice
	// var userData map[string]interface{}
	var userData models.User
	if err := doc.DataTo(&userData); err != nil {
		return nil, err
	}

	// update the posts field
	updatedPosts := append(userData.Posts, req.PostId)

	// posts, ok := userData["posts"].([]interface{})
	// if !ok {
	// 	posts = []interface{}{}
	// }
	// posts = append(posts, req.PostId)

	// write the updated posts slice back to Firestore
	// _, err = docRef.Set(ctx, map[string]interface{}{
	// 	"posts": posts,
	// }, firestore.MergeAll)
	_, err = docRef.Update(ctx, []firestore.Update{
		{Path: "posts", Value: updatedPosts},
	})

	if err != nil {
		return nil, err
	}

	return &pb.AddPostResponse{Success: true}, nil
}

func (s *UserService) Close() error {
	return s.firestoreClient.Close()
}
