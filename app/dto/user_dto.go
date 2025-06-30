package dto

import "janx-admin/app/model"

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
}

func (u *UserInfo) TransitUser(user *model.User) UserInfo {
	return UserInfo{
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}
}
