package storage

import (
	"sync"

	pb "gamingtec_exe/api/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/google/uuid"
)

type UserStore struct {
	mu    sync.RWMutex
	users map[string]*pb.User
}

func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[string]*pb.User),
	}
}

func (s *UserStore) AddUser(user *pb.User) *pb.User {
	var (
		now = timestamppb.Now()
	)

	s.mu.Lock()
	defer s.mu.Unlock()

	user.Id = uuid.NewString()
	user.CreatedAt = now
	user.UpdatedAt = now
	s.users[user.Id] = user

	return user
}

func (s *UserStore) UpdateUser(user *pb.User) (*pb.User, bool) {
	var (
		now = timestamppb.Now()
	)

	s.mu.Lock()
	defer s.mu.Unlock()

	existingUser, exists := s.users[user.Id]
	if !exists {
		return nil, false
	}

	user.CreatedAt = existingUser.CreatedAt
	user.UpdatedAt = now
	s.users[user.Id] = user
	return user, true
}

func (s *UserStore) DeleteUser(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; exists {
		delete(s.users, id)
		return true
	}
	return false
}

func (s *UserStore) GetUser(id string) (*pb.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[id]
	return user, exists
}

func (s *UserStore) ListUsers(country string, page, pageSize int) []*pb.User {
	var (
		defaultPageSize = 5
	)
	s.mu.RLock()
	defer s.mu.RUnlock()

	var users []*pb.User
	for _, user := range s.users {
		if country == "" || user.Country == country {
			users = append(users, user)
		}
		if country == "" {
			users = append(users, user)
		}
	}

	// if there is no pagination return all users
	if page == 0 && pageSize == 0 {
		return users
	}

	if pageSize == 0 {
		pageSize = defaultPageSize
	}

	start := (page - 1) * pageSize
	end := start + pageSize

	if start >= len(users) {
		return []*pb.User{}
	}

	if end > len(users) {
		end = len(users)
	}

	return users[start:end]
}
