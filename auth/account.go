package auth

import (
	"crypto/sha256"
	"demowebserver/dbhelper"
	"demowebserver/model"
	"encoding/hex"
	"errors"
	"regexp"
)

func Register(data model.UserInfo) error {

	if !checkUserInfoValidate(data) {
		return errors.New("Input Formate Error")
	}

	if dbhelper.CheckAccountExist(data.Account) {
		return errors.New("Account Is Exist")
	}

	if dbhelper.CheckEmailExist(data.Email) {
		return errors.New("Email Is Exist")
	}

	return dbhelper.Register(data)
}

func PasswordEncrypt(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}

func checkUserInfoValidate(data model.UserInfo) bool {

	accountAndPasswordRe := regexp.MustCompile(`^[\w]{6,16}$`)
	emailRe := regexp.MustCompile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)

	contents := accountAndPasswordRe.Find([]byte(data.Account))
	if len(contents) == 0 {
		return false
	}

	contents = accountAndPasswordRe.Find([]byte(data.Password))
	if len(contents) == 0 {
		return false
	}

	if len(data.Name) < 4 || len(data.Name) > 16 {
		return false
	}

	contents = emailRe.Find([]byte(data.Email))
	if len(contents) == 0 {
		return false
	}

	return true
}
