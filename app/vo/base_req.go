package vo

type SystemLoginReq struct {
	Username  string `json:"username" form:"username" validate:"required"`
	Password  string `json:"password" form:"password" validate:"required"`
	Captcha   string `json:"captcha" form:"captcha"`
	CaptchaId string `json:"captcha_id" form:"captcha_id"`
}
