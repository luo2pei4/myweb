package service

import (
	exratedao "myweb/dao/exrate"
	exratedto "myweb/dto/exrate"
)

// GetUserInfo 获取用户信息
func GetUserInfo(userID string) (userInfo *exratedto.UserInfo, err error) {

	return exratedao.SelectUserInfoByUserID(userID)
}
