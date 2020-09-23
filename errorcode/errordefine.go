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
	ServerError               = NewError(http.StatusInternalServerError, 200500, "Server Error")
	NotFound                  = NewError(http.StatusNotFound, 200404, http.StatusText(http.StatusNotFound))
	TokenError                = NewError(http.StatusUnauthorized, 200801, "Token Error")
	ParamError                = NewError(http.StatusInternalServerError, 200802, "Param Error")
	JSONOperationError        = NewError(http.StatusInternalServerError, 200803, "Json Operation Error")
	SendEmailError            = NewError(http.StatusInternalServerError, 200901, "Send Email Error")
	CacheOperationError       = NewError(http.StatusInternalServerError, 200902, "Cache Operation Error")
	TooFrequentError          = NewError(http.StatusInternalServerError, 200902, "Operate Too Frequently. Retry After Waiting For Some Seconds")
	ValidateCodeNotMatchError = NewError(http.StatusInternalServerError, 200903, "Validate Code Not Match Error")
	ValidateCodeExpiredError  = NewError(http.StatusInternalServerError, 200903, "Validate Code Expired Error")
	LoadPrivateKeyError       = NewError(http.StatusInternalServerError, 200904, "Load Private Key Error")
	DecryptError              = NewError(http.StatusInternalServerError, 200905, "Decrypt Error")
	GeneratePasswordHashError = NewError(http.StatusInternalServerError, 200906, "Generate Password Hash Error")
	UserAlreadyExistError     = NewError(http.StatusInternalServerError, 200907, "User Already Exist Error")
	UserNotExistError         = NewError(http.StatusInternalServerError, 200908, "User Not Exist Error")
	UserAlreadyBanned         = NewError(http.StatusInternalServerError, 200909, "User Already Banned Error")
	PasswordNotMatch          = NewError(http.StatusInternalServerError, 200910, "Password Not Match Error")
	CreateUserError           = NewError(http.StatusInternalServerError, 200918, "Create User Error")
	ResetPasswordError        = NewError(http.StatusInternalServerError, 200919, "Reset Password Error")
	ResetEmailError           = NewError(http.StatusInternalServerError, 200920, "Reset Email Error")
	UpdateUserInfoError       = NewError(http.StatusInternalServerError, 200921, "Update User Info Error")
)
