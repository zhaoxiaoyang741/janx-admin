package utils

import (
	"github.com/mojocn/base64Captcha"
)

// 定义验证码存储
var store = base64Captcha.DefaultMemStore

// CaptchaResult 验证码结果
type CaptchaResult struct {
	CaptchaId string `json:"captchaId"` // 验证码ID
	ImageUrl  string `json:"imageUrl"`  // 验证码图片Base64
}

// CreateCaptcha 生成验证码
func CreateCaptcha() (*CaptchaResult, error) {
	// 配置验证码参数：字符,公式,验证码配置
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)

	// 创建验证码
	captcha := base64Captcha.NewCaptcha(driver, store)

	// 生成验证码
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		return nil, err
	}

	return &CaptchaResult{
		CaptchaId: id,
		ImageUrl:  b64s,
	}, nil
}

// VerifyCaptcha 验证验证码
func VerifyCaptcha(captchaId string, captcha string) bool {
	// 如果参数为空，直接返回验证失败
	if captchaId == "" || captcha == "" {
		return false
	}

	// 使用存储验证验证码是否正确，第二个参数表示清除
	return store.Verify(captchaId, captcha, true)
}
