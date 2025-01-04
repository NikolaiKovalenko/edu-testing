package httpclient_mock

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser_Success(t *testing.T) {
	t.Parallel()
	// Mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		user := User{ID: 1, Name: "John Doe"}
		json.NewEncoder(w).Encode(user)
	}))
	defer mockServer.Close()

	// Create a mock HTTP client
	client := &http.Client{}

	// Use the mock server's URL
	userService := NewUserService(client, mockServer.URL)

	// Test the GetUser method
	user, err := userService.GetUser(1)

	assert.NoError(t, err, "expected no error")
	assert.NotNil(t, user, "expected a user object")
	assert.Equal(t, 1, user.ID, "expected user ID to match")
	assert.Equal(t, "John Doe", user.Name, "expected user Name to match")
}

func TestGetUser_Failure(t *testing.T) {
	t.Parallel()
	// Mock HTTP server for failure scenario
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound) // Simulating a 404 Not Found
	}))
	defer mockServer.Close()

	// Create a mock HTTP client
	client := &http.Client{}

	// Use the mock server's URL
	userService := NewUserService(client, mockServer.URL)

	// Test the GetUser method
	user, err := userService.GetUser(999)

	assert.Error(t, err, "expected an error")
	assert.Nil(t, user, "expected user to be nil")
}

// OUTPUT
//=== RUN   TestGetUser_Success
//=== PAUSE TestGetUser_Success
//=== CONT  TestGetUser_Success
//--- PASS: TestGetUser_Success (0.00s)
//=== RUN   TestGetUser_Failure
//=== PAUSE TestGetUser_Failure
//=== CONT  TestGetUser_Failure
//--- PASS: TestGetUser_Failure (0.00s)
//PASS
//
//Process finished with the exit code 0
