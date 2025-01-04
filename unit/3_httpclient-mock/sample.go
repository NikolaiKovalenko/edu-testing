package httpclient_mock

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserService struct {
	client  *http.Client
	baseURL string
}

func NewUserService(client *http.Client, baseURL string) *UserService {
	return &UserService{client: client, baseURL: baseURL}
}

func (us *UserService) GetUser(userID int) (*User, error) {
	url := fmt.Sprintf("%s/users/%d", us.baseURL, userID)

	resp, err := us.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch user data")
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
