package with_testify_mock

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetUser_Success(t *testing.T) {
	t.Parallel()
	// Create an instance of the mock repository
	mockRepo := new(MockUserRepository)

	// Create an instance of UserService with the mock repository
	userService := NewUserService(mockRepo)

	// Setup expectations
	call := mockRepo.On("GetUserByID", 1).Return(&User{ID: 1, Name: "John Doe"}, nil).Once()

	// Call the method to test
	user, err := userService.GetUser(1)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "John Doe", user.Name)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)

	// Assert that mock function called only once
	call.Times(1)

	// Same
	mockRepo.AssertCalled(t, "GetUserByID", 1)

	// unset mock behaviour
	call.Unset()
}

func TestGetUser_NotFound(t *testing.T) {
	// Create an instance of the mock repository
	mockRepo := new(MockUserRepository)

	// Create an instance of UserService with the mock repository
	userService := NewUserService(mockRepo)

	// Setup expectations for user not found
	mockRepo.On("GetUserByID", 999).Return(new(User), errors.New("user not found")).Once()

	// Call the method to test
	_, err := userService.GetUser(999)

	// Assertions
	require.Error(t, err)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)

	// Ensure that GetUserByID was called exactly once with the specified argument
	mockRepo.AssertCalled(t, "GetUserByID", 999)
}

func TestGetUser_InvalidID(t *testing.T) {
	// Create an instance of the mock repository
	mockRepo := new(MockUserRepository)

	// Create an instance of UserService with the mock repository
	userService := NewUserService(mockRepo)

	// Call the method to test with an invalid user ID
	_, err := userService.GetUser(0)

	// Assertions
	assert.Error(t, err)

	// Ensure that GetUserByID was not called
	mockRepo.AssertNotCalled(t, "GetUserByID", mock.Anything)
}

// OUTPUT
//=== RUN   TestGetUser_Success
//=== PAUSE TestGetUser_Success
//=== RUN   TestGetUser_NotFound
//--- PASS: TestGetUser_NotFound (0.00s)
//=== RUN   TestGetUser_InvalidID
//--- PASS: TestGetUser_InvalidID (0.00s)
//=== CONT  TestGetUser_Success
//--- PASS: TestGetUser_Success (0.00s)
//PASS
