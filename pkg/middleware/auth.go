package middleware

import (
	"errors"
	"janx-admin/global"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// InitAuth 初始化jwt中间件
func InitAuth() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           global.Conf.Jwt.Realm,                                 // jwt标识
		Key:             []byte(global.Conf.Jwt.Key),                           // 服务端密钥
		Timeout:         time.Hour * time.Duration(global.Conf.Jwt.Timeout),    // token过期时间
		MaxRefresh:      time.Hour * time.Duration(global.Conf.Jwt.MaxRefresh), // token最大刷新时间(RefreshToken过期时间=Timeout+MaxRefresh)
		PayloadFunc:     payloadFunc,                                           // 有效载荷处理
		IdentityHandler: identityHandler,                                       // 解析Claims
		Authenticator:   login,                                                 // 校验token的正确性, 处理登录逻辑
		Authorizator:    authorizator,                                          // 用户登录校验成功处理
		Unauthorized:    unauthorized,                                          // 用户登录校验失败处理
		LoginResponse:   loginResponse,                                         // 登录成功后的响应
		LogoutResponse:  logoutResponse,                                        // 登出后的响应
		RefreshResponse: refreshResponse,                                       // 刷新token后的响应
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",    // 自动在这几个地方寻找请求中的token
		TokenHeadName:   "Bearer",                                              // header名称
		TimeFunc:        time.Now,
	})
	return authMiddleware, err
}

// 有效载荷处理
func payloadFunc(data interface{}) jwt.MapClaims {
	// if v, ok := data.(map[string]interface{}); ok {
	// 	var user model.User
	// 	// 将用户json转为结构体
	// 	util.JsonI2Struct(v["user"], &user)
	// 	return jwt.MapClaims{
	// 		jwt.IdentityKey: user.ID,
	// 		"user":          v["user"],
	// 	}
	// }
	return jwt.MapClaims{}
}

// 解析Claims
func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	// 此处返回值类型map[string]interface{}与payloadFunc和authorizator的data类型必须一致, 否则会导致授权失败还不容易找到原因
	return map[string]interface{}{
		"IdentityKey": claims[jwt.IdentityKey],
		"user":        claims["user"],
	}
}

// 校验token的正确性, 处理登录逻辑
func login(c *gin.Context) (interface{}, error) {
	return nil, errors.New("未实现")
	// var req vo.RegisterAndLoginRequest
	// // 请求json绑定
	// if err := c.ShouldBind(&req); err != nil {
	// 	return "", err
	// }

	// u := &model.User{
	// 	Username: req.Username,
	// 	Password: req.Password,
	// }
	// ok := util.Verify(req.CaptchaId, req.Captcha)
	// if !ok {
	// 	return nil, errors.New("captcha verify fail")
	// }

	// // 密码校验
	// userRepository := repository.NewUserRepository()
	// user, err := userRepository.Login(u)
	// if err != nil {
	// 	return nil, err
	// }

	// // 将用户以json格式写入, payloadFunc/authorizator会使用到
	// return map[string]interface{}{
	// 	"user": util.Struct2Json(user),
	// }, nil
}

// 用户登录校验成功处理
func authorizator(data interface{}, c *gin.Context) bool {
	// if v, ok := data.(map[string]interface{}); ok {
	// 	userStr := v["user"].(string)
	// 	var user model.User
	// 	// 将用户json转为结构体
	// 	util.Json2Struct(userStr, &user)
	// 	// 将用户保存到context, api调用时取数据方便
	// 	c.Set("user", user)
	// 	return true
	// }
	// return false
	return false
}

// 用户登录校验失败处理
func unauthorized(c *gin.Context, code int, message string) {
	// common.Log.Debugf("JWT认证失败, 错误码: %d, 错误信息: %s", code, message)
	// response.Response(c, code, code, nil, fmt.Sprintf("JWT认证失败, 错误码: %d, 错误信息: %s", code, message))
}

// 登录成功后的响应
func loginResponse(c *gin.Context, code int, token string, expires time.Time) {
	// response.Response(c, code, code,
	// 	gin.H{
	// 		"token":   token,
	// 		"expires": expires.Format("2006-01-02 15:04:05"),
	// 	},
	// 	"登录成功")
}

// 登出后的响应
func logoutResponse(c *gin.Context, code int) {
	// response.Success(c, nil, "退出成功")
}

// 刷新token后的响应
func refreshResponse(c *gin.Context, code int, token string, expires time.Time) {
	// response.Response(c, code, code,
	// 	gin.H{
	// 		"token":   token,
	// 		"expires": expires,
	// 	},
	// 	"刷新token成功")
}
