package errorcode

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			var Err *Error
			if e, ok := err.(*Error); ok {
				Err = e
			} else if e, ok := err.(error); ok {
				Err = OtherError(e.Error())
			} else {
				Err = ServerError
			}
			// 记录一个错误的日志
			c.JSON(Err.StatusCode, Err)
			return
		}
	}()
	// c.Abort()
	c.Next()
}
