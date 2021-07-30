package userauth

import (
	"github.com/kaleemubarok/gc-ref/account/config"
	"github.com/kaleemubarok/gc-ref/account/resource/account"
)

type Usecase struct {
	dbRsc      account.IDB
	signingKey []byte
	cfg        config.Config
}

type IUsecase interface {
	Register(username, password, confirmPassword string) error
	Login(username, password string) (string, error)
	ValidateSession(accessToken string) (string, error)
	GenerateJWT(userID string, profilePic string) (string, error)
}

func NewUsecase(dbRsc account.IDB, signingKey string, config config.Config) IUsecase {
	return &Usecase{
		dbRsc:      dbRsc,
		signingKey: []byte(signingKey),
		cfg:        config,
	}
}
