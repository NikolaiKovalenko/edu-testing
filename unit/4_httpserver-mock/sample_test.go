package httpserver_mock

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserHandler(t *testing.T) {
	t.Parallel()
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/user", nil)
	assert.NoError(t, err)

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UserHandler)

	// Perform the request
	handler.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code, "expected status code to be 200")

	// Decode the response body
	var user User
	err = json.NewDecoder(rr.Body).Decode(&user)
	assert.NoError(t, err, "expected no error when decoding JSON response")

	// Validate the response
	expectedUser := User{ID: 1, Name: "John Doe"}
	assert.Equal(t, expectedUser, user, "expected user does not match actual user")
}

// OUTPUT:
//=== RUN   TestUserHandler
//--- PASS: TestUserHandler (0.00s)
//PASS
//
//Process finished with the exit code 0
