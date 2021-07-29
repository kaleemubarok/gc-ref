package userauth

import "time"

type AuthClient struct {
	host    string
	timeout time.Duration
}

type User struct {
	UserID     string    `json:"user_id"`
	Username   string    `json:"username"`
	CreatedAt  time.Time `json:"created_at"`
	ProfilePic string    `json:"profile_pic"`
}

type IClient interface {
	GetUserInfo(accessToken string) *User
	GetUserByID(accessToken, userID string) *User
}

func NewClient(host string, timeout time.Duration) IClient {
	return &AuthClient{
		host:    host,
		timeout: timeout,
	}
}
