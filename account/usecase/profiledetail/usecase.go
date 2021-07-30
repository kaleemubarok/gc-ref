package profiledetail

import (
	"errors"

	"github.com/kaleemubarok/gc-ref/account/model"
)

func (u *Usecase) GetUserByUserID(userID string) (model.User, error) {
	if userID == "" {
		return model.User{}, errors.New("unauthorized")
	}

	user, err := u.dbRsc.GetUserByUserID(userID)
	if err != nil {
		return user, errors.New("internal server error")
	}

	if user.ProfilePic == "" {
		user.ProfilePic = u.cfg.Get("PROFILE_PIC_HOST") + user.UserID
	}
	return model.User{
		UserID:     user.UserID,
		Password:   "",
		Salt:       "",
		UserName:   user.UserName,
		ProfilePic: user.ProfilePic,
	}, nil
}
