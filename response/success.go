package response

import (
	"net/http"
)

func NewSuccess(data interface{}) *Response {
	return &Response{
		StatusCode: http.StatusOK,
		Msg:        "Success",
		Code:       0,
		Data:       data,
	}
}
