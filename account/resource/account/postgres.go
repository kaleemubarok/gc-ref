package account

import (
	"time"

	"github.com/kaleemubarok/gc-ref/account/model"
)

var user UserDB

func (dbr *DBResource) Register(uid, username, password, salt string) error {
	query := `
		INSERT INTO
			account
		(
			user_id,
			username,
			password,
			salt,
			created_at,
			profile_pic
		)
		VALUES
		(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		)	
	`

	_, err := dbr.db.Exec(query, uid, username, password, salt, time.Now(), "")
	if err != nil {
		return err
	}

	return nil
}
func (dbr *DBResource) GetUserByUserID(userID string) (model.User, error) {
	query := `
	SELECT 
		user_id,
		username,
		password,
		salt,
		created_at,
		profile_pic
	FROM
		account
	WHERE
		user_id = $1	
	`

	err := dbr.db.Get(&user, query, userID)
	if err != nil {
		return model.User{}, err
	}
	return model.User{
		UserID:     user.UserID.String,
		UserName:   user.UserName.String,
		Password:   user.Password.String,
		Salt:       user.Salt.String,
		ProfilePic: user.ProfilePic.String,
		CreatedAt:  user.CreatedAt,
	}, nil
}
func (dbr *DBResource) GetUserByUserName(userName string) (model.User, error) {
	query := `
	SELECT 
		user_id,
		username,
		password,
		salt,
		created_at,
		profile_pic
	FROM
		account
	WHERE
		username = $1`

	err := dbr.db.Get(&user, query, userName)
	if err != nil {
		return model.User{}, err
	}
	return model.User{
		UserID:     user.UserID.String,
		UserName:   user.UserName.String,
		Password:   user.Password.String,
		Salt:       user.Salt.String,
		ProfilePic: user.ProfilePic.String,
		CreatedAt:  user.CreatedAt,
	}, nil
}
func (dbr *DBResource) UpdateProfilePic(userID, profilePic string) error {
	return nil
}
func (dbr *DBResource) UpdatePassword(userID, password string) error {
	return nil
}
