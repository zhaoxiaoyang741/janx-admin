package vo

type RoleCreateReq struct {
	Name    string  `json:"name" form:"name" validate:"required"`
	Keyword string  `json:"keyword" form:"keyword" validate:"required"`
	Desc    *string `json:"desc" form:"desc"`
	Status  uint    `json:"status" form:"status"`
	Sort    uint    `json:"sort" form:"sort"`
}
type RoleUpdateReq struct {
	ID      uint    `json:"id" form:"id" validate:"required"`
	Name    string  `json:"name" form:"name" validate:"required"`
	Keyword string  `json:"keyword" form:"keyword" validate:"required"`
	Desc    *string `json:"desc" form:"desc"`
	Status  uint    `json:"status" form:"status"`
	Sort    uint    `json:"sort" form:"sort"`
}

type RoleDeleteReq struct {
	DeleteReq
}

type RoleListReq struct {
	ReqPageInfo
	Name    string `json:"name" form:"name" validate:"required"`
	Keyword string `json:"keyword" form:"keyword" validate:"required"`
}
