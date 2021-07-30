package account

import (
	"github.com/jmoiron/sqlx"
	"github.com/kaleemubarok/gc-ref/account/model"
)

type DBResource struct {
	db *sqlx.DB
}

type IDB interface {
	Register(uid, username, password, salt string) error
	GetUserByUserID(userID string) (model.User, error)
	GetUserByUserName(userName string) (model.User, error)
	UpdateProfilePic(userID, profilePic string) error
	UpdatePassword(userID, password string) error
}

func NewDBResource(dbParam *sqlx.DB) IDB {
	return &DBResource{
		db: dbParam,
	}
}
