package vo

type ReqPageInfo struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

type DeleteReq struct {
	Ids []uint `json:"ids" form:"ids" `
}
