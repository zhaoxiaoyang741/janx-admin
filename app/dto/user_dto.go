package dto

import "janx-admin/app/model"

type UserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username" form:"username"`
	Nickname string `json:"nickname" form:"nickname"`
}

func (u *UserInfo) TransitUser(user *model.User) UserInfo {
	return UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
	}
}
