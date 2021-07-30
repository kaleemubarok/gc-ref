package profiledetail

import (
	"github.com/kaleemubarok/gc-ref/account/config"
	"github.com/kaleemubarok/gc-ref/account/model"
	"github.com/kaleemubarok/gc-ref/account/resource/account"
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
