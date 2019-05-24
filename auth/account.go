package auth

import (
	"demowebserver/dbhelper"
	"demowebserver/model"
)

func Register(data model.UserInfo) error {

	if err := checkUserInfoValidate(data); err != nil {
		return err
	} else {
		return dbhelper.Register(data)
	}
}

func checkUserInfoValidate(data model.UserInfo) error {

	return nil
}
