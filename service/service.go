package service

import (
	"context"

	pb "gamingtec_exe/api/proto"
	store "gamingtec_exe/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserServiceServer struct {
	store *store.UserStore
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer(store *store.UserStore) *UserServiceServer {
	return &UserServiceServer{store: store}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	return s.store.AddUser(req.User), nil
}

func (s *UserServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	user, ok := s.store.UpdateUser(req.User)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	return user, nil
}

func (s *UserServiceServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	if !s.store.DeleteUser(req.Id) {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	return &emptypb.Empty{}, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	user, ok := s.store.GetUser(req.Id)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	return user, nil
}

func (s *UserServiceServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users := s.store.ListUsers(req.Country, int(req.Page), int(req.PageSize))
	return &pb.ListUsersResponse{Users: users}, nil
}
