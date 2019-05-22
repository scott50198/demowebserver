package auth

import (
	"demowebserver/model"
	"demowebserver/utils"
)

func Register(data model.UserInfo) bool {

	data.CreateTime = utils.GetNowTimeMillis()
	data.UpdateTime = utils.GetNowTimeMillis()

	return true
}
