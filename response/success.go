package response

import (
	"net/http"
)

func NewSuccess(msg string) *Response {
	return &Response{
		StatusCode: http.StatusOK,
		Msg:        msg,
		Code:       0,
	}
}
