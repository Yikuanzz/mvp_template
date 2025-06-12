package response

import "github.com/gin-gonic/gin"

// response 响应结构体
type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data;omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

// Error 错误响应
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, response{
		Code:    code,
		Message: message,
	})
}

// ErrorWithData 错误响应并携带数据
func ErrorWithData(code int, message string, data interface{}) response {
	return response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
