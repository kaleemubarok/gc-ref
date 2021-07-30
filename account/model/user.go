package model

import "time"

type User struct {
	UserID     string    `json:"user_id"`
	UserName   string    `json:"username"`
	Password   string    `json:"password"`
	Salt       string    `json:"salt"`
	CreatedAt  time.Time `json:"created_at"`
	ProfilePic string    `json:"profile_pic"`
}
