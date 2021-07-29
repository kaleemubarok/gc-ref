package userauth

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"math/rand"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/twharmon/gouid"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (u *Usecase) Register(username, password, confirmPassword string) error {
	if confirmPassword != password {
		return errors.New("confirm password is mismatched")
	}

	salt := RandStringBytes(32)
	password += salt

	h := sha256.New()
	h.Write([]byte(password))
	password = fmt.Sprintf("%x", h.Sum(nil))

	uid := gouid.String(8, gouid.LowerCaseAlphaNum)
	if uid == "" {
		return errors.New("error on generating uid")
	}

	err := u.dbRsc.Register(uid, username, password, salt)
	if err != nil {
		return err
	}

	return nil
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (u *Usecase) Login(username, password string) (string, error) {
	user, err := u.dbRsc.GetUserByUserName(username)
	if err != nil {
		return "", errors.New("wrong username")
	}

	password += user.Salt
	h := sha256.New()
	h.Write([]byte(password))
	hashedPassword := fmt.Sprintf("%x", h.Sum(nil))

	if user.Password != hashedPassword {
		return "", errors.New("password incorrect")
	}

	user.Password, user.Salt = "", ""

	//set default generated profile pic
	if user.ProfilePic == "" {
		user.ProfilePic = u.cfg.Get("PROFILE_PIC_HOST") + user.UserID
	}
	fmt.Println("DEBUG: PROFILE PIC ->", user.ProfilePic)

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	tokenClaim := jwt.MapClaims{}
	tokenClaim["user_id"] = user.UserID
	tokenClaim["profile_pic"] = user.ProfilePic
	token.Claims = tokenClaim

	tokenString, err := token.SignedString(u.signingKey)
	if err != nil {
		return "", errors.New("internal server error")
	}
	return tokenString, nil
}

func (u *Usecase) ValidateSession(accessToken string) (string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return u.signingKey, nil
	})
	if err != nil {
		return "", errors.New("invalid token")
	}

	userID := claims["user_id"].(string)
	return userID, nil
}

func (u *Usecase) GenerateJWT(userID string, profilePic string) (string, error) {
	if profilePic == "" {
		profilePic = u.cfg.Get("PROFILE_PIC_HOST") + userID
	}
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	tokenClaim := jwt.MapClaims{}
	tokenClaim["user_id"] = userID
	tokenClaim["profile_pic"] = profilePic
	token.Claims = tokenClaim

	tokenString, err := token.SignedString(u.signingKey)
	if err != nil {
		log.Println(err)
		return "", errors.New("internal server error")
	}
	return tokenString, nil
}
