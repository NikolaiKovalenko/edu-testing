package with_testify_mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser_Success(t *testing.T) {
	t.Parallel()
	// Create an instance of the mock repository
	mockRepo := new(MockUserRepository)

	// Create an instance of UserService with the mock repository
	userService := NewUserService(mockRepo)

	// Setup expectations
	call := mockRepo.On("GetUserByID", 1).Return(&User{ID: 1, Name: "John Doe"}, nil)

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
