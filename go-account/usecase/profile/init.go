package profile

import (
	"github.com/kaleemubarok/gc-ref/go-account/config"
	"github.com/kaleemubarok/gc-ref/go-account/model"
	"github.com/kaleemubarok/gc-ref/go-account/resource/account"
)

type Usecase struct {
	dbRsc account.IDB
	cfg   config.Config
}

type IUsecase interface {
	GetUserByUserID(userID string) (model.User, error)
}

func NewUsecase(dbRsc account.IDB, cfg config.Config) IUsecase {
	return &Usecase{
		dbRsc: dbRsc,
		cfg:   cfg,
	}
}
