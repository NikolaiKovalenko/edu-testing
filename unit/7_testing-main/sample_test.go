package testing_main

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMain sets up and tears down the test environment.
func TestMain(m *testing.M) {
	// Setup: Initialize resources here if needed
	// Example: Open database connections, create test files, etc.

	// Run tests
	code := m.Run()

	// Teardown: Clean up resources here if needed
	// Example: Close database connections, remove test files, etc.
	fmt.Printf("Tests finished with code %d\n", code)
	os.Exit(code)
}

func TestGetUser_Success(t *testing.T) {
	t.Parallel()
	// Create an instance of the mock repository
	mockRepo := new(MockUserRepository)

	// Create an instance of UserService with the mock repository
	userService := NewUserService(mockRepo)

	// Setup expectations
	mockRepo.On("GetUserByID", 1).Return(&User{ID: 1, Name: "John Doe"}, nil).Once()

	// Call the method to test
	user, err := userService.GetUser(1)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "John Doe", user.Name)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestGetUser_NotFound(t *testing.T) {
	t.Parallel()
	// Create an instance of the mock repository
	mockRepo := new(MockUserRepository)

	// Create an instance of UserService with the mock repository
	userService := NewUserService(mockRepo)

	// Setup expectations for user not found
	mockRepo.On("GetUserByID", 999).Return(new(User), errors.New("user not found")).Once()

	// Call the method to test
	_, err := userService.GetUser(999)

	// Assertions
	assert.Error(t, err)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

// OUTPUT
//=== RUN   TestGetUser_Success
//=== PAUSE TestGetUser_Success
//=== RUN   TestGetUser_NotFound
//=== PAUSE TestGetUser_NotFound
//=== CONT  TestGetUser_Success
//=== CONT  TestGetUser_NotFound
//--- PASS: TestGetUser_Success (0.00s)
//--- PASS: TestGetUser_NotFound (0.00s)
//PASS
//Tests finished with code 0
