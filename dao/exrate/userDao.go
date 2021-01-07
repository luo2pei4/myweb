package exratedao

import (
	"fmt"
	exratedto "myweb/dto/exrate"
)

// SelectUserInfoByUserID 通过用户ID查询用户信息
func SelectUserInfoByUserID(userID string) (userInfo *exratedto.UserInfo, err error) {

	sql := "select id, userid, userpwd, username, email, createtime from users where userid = '%v'"
	sql = fmt.Sprintf(sql, userID)

	rows, err := conn.Select(sql)

	if err != nil {
		return nil, err
	}

	userInfo = &exratedto.UserInfo{}
	rows.Next()
	err = rows.Scan(&userInfo.ID, &userInfo.UserID, &userInfo.UserPwd, &userInfo.UserName, &userInfo.Email, &userInfo.CreateTime)

	if err != nil {
		return nil, err
	}

	return
}
