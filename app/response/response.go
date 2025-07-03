package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, httpStatus int, code int, data gin.H, message string) {
	c.JSON(httpStatus, gin.H{"code": code, "data": data, "message": message})
}

// 返回前端-成功
func Success(c *gin.Context, data gin.H, message string) {
	Response(c, http.StatusOK, 200, data, message)
}

// 返回前端-失败
func Fail(c *gin.Context, data gin.H, message string) {
	Response(c, http.StatusBadRequest, 400, data, message)
}

func FailWithoutData(c *gin.Context, message string) {
	Response(c, http.StatusBadRequest, 400, nil, message)
}

type ResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseDataWithData struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
}

type ResListData struct {
	Total int         `json:"total" example:"100" description:"总记录数"`
	List  interface{} `json:"list" swaggerignore:"true" description:"数据列表"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
}
