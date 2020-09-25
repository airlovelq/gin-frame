package errorcode

import (
	"net/http"
	"scoremanager/response"

	"github.com/gin-gonic/gin"
)

type Error struct {
	response.Response
}

func OtherError(message string) *Error {
	return NewError(http.StatusForbidden, 100403, message)
}

func (e *Error) Error() string {
	return e.Msg
}

func NewError(statusCode, Code int, msg string) *Error {
	return &Error{
		Response: response.Response{
			StatusCode: statusCode,
			Code:       Code,
			Msg:        msg,
			Data:       nil,
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
)
