package storage

import (
	"gamingtec_exe/utils"
	"log"
	"sync"

	pb "gamingtec_exe/api/proto"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *UserStore) CreateUser(user *pb.User) *pb.User {
	var (
		now = timestamppb.Now()
	)

	s.mu.Lock()
	defer s.mu.Unlock()

	user.Id = uuid.NewString()
	user.CreatedAt = now
	user.UpdatedAt = now
	s.users[user.Id] = user

	log.Printf("User %s %s has been successfully created.", user.FirstName, user.LastName)

	return user
}

func (s *UserStore) UpdateUser(user *pb.User) (*pb.User, bool) {
	var (
		now = timestamppb.Now()
	)

	s.mu.Lock()
	defer s.mu.Unlock()

	// if the user does not exist return false
	existingUser, exists := s.users[user.Id]
	if !exists {
		return nil, false
	}

	user.CreatedAt = existingUser.CreatedAt
	user.UpdatedAt = now
	s.users[user.Id] = user

	log.Printf("User %s %s has been successfully updated.", user.FirstName, user.LastName)

	return user, true
}

func (s *UserStore) DeleteUser(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; exists {
		delete(s.users, id)
		return true
	}

	log.Printf("User with id: %s has been successfully deleted.", id)
	return false
}

func (s *UserStore) GetUser(id string) (*pb.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[id]

	if exists {
		log.Printf("User with id: %s has been found.", id)
	} else {
		log.Printf("User not found with id: %s.", id)
	}

	return user, exists
}

func (s *UserStore) ListUsers(country string, page, pageSize int) []*pb.User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// create a slice from mapper
	userSlice := utils.MapToSlice(s.users)

	// Filter by country
	filteredUsers := utils.FilterByCountry(country, userSlice)

	// Apply pagination
	paginatedUsers := utils.HandlePagination(pageSize, page, filteredUsers)

	log.Printf("Found %d users in total.", len(paginatedUsers))

	return paginatedUsers
}
