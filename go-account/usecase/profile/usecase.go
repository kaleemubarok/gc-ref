package profile

import (
	"errors"

	"github.com/kaleemubarok/gc-ref/go-account/model"
)

func (u *Usecase) GetUserByUserID(userID string) (model.User, error) {
	if userID == "" {
		return model.User{}, errors.New("unauthorized")
	}

	user, err := u.dbRsc.GetUserByUserID(userID)
	if err != nil {
		return user, errors.New("internal server error")
	}

	return model.User{
		UserID:     user.UserID,
		Password:   "",
		Salt:       "",
		UserName:   user.UserName,
		ProfilePic: u.cfg.Get("PROFILE_PIC_HOST") + user.UserID,
	}, nil
}
