package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 错误处理的结构体
type Response struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

type Error struct {
	Response
}

func OtherError(message string) *Error {
	return NewError(http.StatusForbidden, 100403, message)
}

func (e *Error) Error() string {
	return e.Msg
}

func NewError(statusCode, Code int, msg string) *Error {
	return &Error{
		Response: Response{
			StatusCode: statusCode,
			Code:       Code,
			Msg:        msg,
		},
	}
}

// 404处理
func HandleNotFound(c *gin.Context) {
	err := NotFound
	c.JSON(err.StatusCode, err)
	return
}

var (
	// Success     = NewError(http.StatusOK, 0, "success")
	ServerError = NewError(http.StatusInternalServerError, 200500, "Server Error")
	NotFound    = NewError(http.StatusNotFound, 200404, http.StatusText(http.StatusNotFound))
	TokenError  = NewError(http.StatusForbidden, 200801, "Token Error")
)
