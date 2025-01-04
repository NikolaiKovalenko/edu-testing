package testing_main

import (
	"errors"
)

// User represents a user model.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UserRepository defines the interface for user data operations.
type UserRepository interface {
	GetUserByID(id int) (*User, error)
}

// UserService uses the UserRepository to manage user data.
type UserService struct {
	repo UserRepository
}

// NewUserService creates a new UserService with the provided repository.
func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUser retrieves a user by ID using the repository.
func (us *UserService) GetUser(id int) (*User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}
	return us.repo.GetUserByID(id)
}
