package auth

import (
	"demowebserver/dbhelper"
	"demowebserver/model"
)

func Register(data model.UserInfo) error {

	return dbhelper.Register(data)
}
