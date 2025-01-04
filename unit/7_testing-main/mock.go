package testing_main

import "github.com/stretchr/testify/mock"

// MockUserRepository is a mock implementation of the UserRepository interface.
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserByID(id int) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}
