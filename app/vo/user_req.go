package vo

type UserCreateReq struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password"`
	Nickname string `json:"nickname" form:"nickname" validate:"required"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Avatar   string `json:"avatar" form:"avatar"`
}

type UserUpdataReq struct {
	ID       uint   `json:"id" form:"id" validate:"required"`
	Nickname string `json:"nickname" form:"nickname"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Avatar   string `json:"avatar" form:"avatar"`
}

type UserDeleteReq struct {
	DeleteReq
}

type UserListReq struct {
	ReqPageInfo
	Username string `json:"username" form:"username"`
	Nickname string `json:"nickname" form:"nickname"`
}
