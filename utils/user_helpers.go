package utils

import (
	pb "gamingtec_exe/api/proto"
)

const (
	defaultPageSize = 5
)

func FilterByCountry(country string, users []*pb.User) []*pb.User {
	var (
		filteredUsers []*pb.User
	)
	// If country is empty return all users
	if country == "" {
		return users
	}

	// Loop through the users and filter the ones with the given country
	for _, user := range users {
		if user.Country == country {
			filteredUsers = append(filteredUsers, user)
		}
	}

	return filteredUsers
}

func HandlePagination(pageSize int, page int, users []*pb.User) []*pb.User {

	// if there is no pagination return all users
	if page == 0 && pageSize == 0 {
		return users
	}

	// set the default page size
	if pageSize == 0 {
		pageSize = defaultPageSize
	}

	// get the start & end indexes
	start := (page - 1) * pageSize
	end := start + pageSize

	// validation checks
	if start >= len(users) || start < 0 || end < 0 {
		return []*pb.User{}
	}

	if end > len(users) {
		end = len(users)
	}

	return users[start:end]
}

func MapToSlice(userMap map[string]*pb.User) []*pb.User {
	var userSlice []*pb.User
	for _, user := range userMap {
		userSlice = append(userSlice, user)
	}
	return userSlice
}
