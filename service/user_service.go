package service

import (
	"scoremanager/controller/user"
	"scoremanager/errorcode"
	"scoremanager/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v3"
)

type emailParams struct {
	Email null.String `json:"email"`
}

type registerOrResetPasswordByEmailParams struct {
	Email            null.String `json:"email"`
	Password_Encrypt null.String `json:"password_encrypt"`
	Validate_Code    null.String `json:"validate_code"`
}

type passwordEncryptParams struct {
	Password_Encrypt null.String `json:"password_encrypt"`
}

type loginParams struct {
	Log_ID           null.String `json:"log_id"`
	Password_Encrypt null.String `json:"password_encrypt"`
}

type resetEmailParams struct {
	Email         null.String `json:"email"`
	Validate_Code null.String `json:"validate_code"`
}

type userInfoParams struct {
	Sex       null.Int    `json:"sex,omitempty"`
	Age       null.Int    `json:"age,omitempty"`
	User_Name null.String `json:"user_name,omitempty"`
	Name      null.String `json:"name,omitempty`
	Info      null.String `json:"info,omitempty"`
}

// @Summary 邮件发送注册验证码
// @Description 邮件发送注册验证码
// @Tags 用户管理
// @Accept  json
// @Produce json
// @Param   params     body    emailParams     true        "邮箱"
// @Success 200 {object}  response.Response	"Success data为null"
// @Failure 500  {object}  response.Response  "内部错误 data为null"
// @Failure 401  {object}  response.Response  "token错误 data为null"
// @Failure 404  {object}  response.Response  "无法访问 data为null"
// @Router /user/register/email/validate [post]
func sendEmailRegisterValidateCode(c *gin.Context) {
	var jsonParams emailParams
	err := c.BindJSON(&jsonParams)
	if err != nil {
		panic(errorcode.ParamError)
	}
	paramOk := utils.CheckEmailMustExist(jsonParams.Email)
	if !paramOk {
		panic(errorcode.ParamError)
	}

	userOp := user.NewUserOp()
	res := userOp.SendEmailRegisterValidateCode(jsonParams.Email.String)
	c.JSON(res.StatusCode, res)
}

// @Summary 邮件发送重置密码验证码
// @Description 邮件发送重置密码验证码
// @Tags 用户管理
// @Accept  json
// @Produce json
// @Param   params     body    emailParams     true        "邮箱"
// @Success 200 {object}  response.Response	"Success data为null"
// @Failure 500  {object}  response.Response  "内部错误 data为null"
// @Failure 401  {object}  response.Response  "token错误 data为null"
// @Failure 404  {object}  response.Response  "无法访问 data为null"
// @Router /user/password/reset/email/validate [post]
func sendEmailResetPasswordValidateCode(c *gin.Context) {
	var jsonParams emailParams
	err := c.BindJSON(&jsonParams)
	if err != nil {
		panic(errorcode.ParamError)
	}
	paramOk := utils.CheckEmailMustExist(jsonParams.Email)
	if !paramOk {
		panic(errorcode.ParamError)
	}

	userOp := user.NewUserOp()
	res := userOp.SendEmailResetPasswordValidateCode(jsonParams.Email.String)
	c.JSON(res.StatusCode, res)
}

// @Summary 邮件注册
// @Description 邮件注册
// @Tags 用户管理
// @Accept  json
// @Produce json
// @Param   params     body    registerOrResetPasswordByEmailParams     true        "邮箱"
// @Success 200  {object}  response.Response  "Success data为{\"id": user id}"
// @Failure 500  {object}  response.Response  "内部错误 data为null"
// @Failure 401  {object}  response.Response  "token错误 data为null"
// @Failure 404  {object}  response.Response  "无法访问 data为null"
// @Router /user/register/email [post]
func registerByEmail(c *gin.Context) {
	var jsonParams registerOrResetPasswordByEmailParams
	err := c.BindJSON(&jsonParams)
	if err != nil {
		panic(errorcode.ParamError)
	}
	// email, _ := c.Get("email")
	// validateCode, _ := c.Get("validate_code")
	// passwordEncrypt, _ := c.Get("password_encrypt")
	paramOk := utils.CheckEmailMustExist(jsonParams.Email) && utils.CheckStringMustExist(jsonParams.Validate_Code, jsonParams.Password_Encrypt)
	if !paramOk {
		panic(errorcode.ParamError)
	}
	userOp := user.NewUserOp()
	res := userOp.RegisterByEmail(jsonParams.Email.String, jsonParams.Validate_Code.String, jsonParams.Password_Encrypt.String)
	c.JSON(res.StatusCode, res)
}

// @Summary 重置密码（已登录状态）
// @Description 重置密码（已登录状态）
// @Tags 用户管理
// @Accept  json
// @Produce json
// @security ApiKeyAuth
// @Param   params  body passwordEncryptParams true "密码（加密过的）"
// @Success 200  {object}  response.Response  "Success data为null"
// @Failure 500  {object}  response.Response  "内部错误 data为null"
// @Failure 401  {object}  response.Response  "token错误 data为null"
// @Failure 404  {object}  response.Response  "无法访问 data为null"
// @Router /user/password/reset [post]
func resetPasswordInLoginStatus(c *gin.Context, tokenMap map[string]interface{}) {
	var jsonParams passwordEncryptParams
	err := c.BindJSON(&jsonParams)
	if err != nil {
		panic(errorcode.ParamError)
	}
	paramOk := utils.CheckUUID4MustExist(tokenMap["user_id"]) && utils.CheckStringMustExist(jsonParams.Password_Encrypt)
	if !paramOk {
		panic(errorcode.ParamError)
	}
	userOp := user.NewUserOp()
	res := userOp.ResetPasswordInLoginStatus(tokenMap["user_id"].(string), jsonParams.Password_Encrypt.String)
	c.JSON(res.StatusCode, res)
}

// @Summary 登录
// @Description 登录接口，通过email，手机号或用户名
// @Tags 用户管理
// @Accept  json
// @Produce json
// @Param   params  body loginParams true "log_id:email，手机号或用户名 password_encrypt:加密后的密码"
// @Success 200  {object}  response.Response  "Success data为{\"token\":token, \"user_id\":user id}"
// @Failure 500  {object}  response.Response  "内部错误 data为null"
// @Failure 401  {object}  response.Response  "token错误 data为null"
// @Failure 404  {object}  response.Response  "无法访问 data为null"
// @Router /user/login [post]
func login(c *gin.Context) {
	var jsonParams loginParams
	err := c.BindJSON(&jsonParams)
	if err != nil {
		panic(errorcode.ParamError)
	}
	paramOk := utils.CheckStringMustExist(jsonParams.Log_ID, jsonParams.Password_Encrypt)
	if !paramOk {
		panic(errorcode.ParamError)
	}
	userOp := user.NewUserOp()
	res := userOp.Login(jsonParams.Log_ID.String, jsonParams.Password_Encrypt.String)
	c.JSON(res.StatusCode, res)
}

// @Summary 重置密码（发送邮件方式）
// @Description 重置密码（发送邮件方式），根据邮件中验证码
// @Tags 用户管理
// @Accept  json
// @Produce json
// @Param   params  body registerOrResetPasswordByEmailParams true "email"
// @Success 200  {object}  response.Response  "Success data为null"
// @Failure 500  {object}  response.Response  "内部错误 data为null"
// @Failure 401  {object}  response.Response  "token错误 data为null"
// @Failure 404  {object}  response.Response  "无法访问 data为null"
// @Router /user/password/reset/email [post]
func resetPasswordByEmail(c *gin.Context) {
	var jsonParams registerOrResetPasswordByEmailParams
	err := c.BindJSON(&jsonParams)
	if err != nil {
		panic(errorcode.ParamError)
	}
	paramOk := utils.CheckEmailMustExist(jsonParams.Email) && utils.CheckStringMustExist(jsonParams.Validate_Code, jsonParams.Password_Encrypt)
	if !paramOk {
		panic(errorcode.ParamError)
	}
	userOp := user.NewUserOp()
	res := userOp.ResetPasswordByEmailValidateCode(jsonParams.Email.String, jsonParams.Validate_Code.String, jsonParams.Password_Encrypt.String)
	c.JSON(res.StatusCode, res)
}

// @Summary 重置邮箱
// @Description 重置邮箱，根据邮件中验证码
// @Tags 用户管理
// @Accept  json
// @Produce json
// @security ApiKeyAuth
// @Param   params  body resetEmailParams true "email,validate_code"
// @Success 200  {object}  response.Response  "Success data为null"
// @Failure 500  {object}  response.Response  "内部错误 data为null"
// @Failure 401  {object}  response.Response  "token错误 data为null"
// @Failure 404  {object}  response.Response  "无法访问 data为null"
// @Router /user/email/reset [post]
func resetEmail(c *gin.Context, tokenMap map[string]interface{}) {
	var jsonParams resetEmailParams
	err := c.BindJSON(&jsonParams)
	if err != nil {
		panic(errorcode.ParamError)
	}
	paramOk := utils.CheckUUID4MustExist(tokenMap["user_id"]) && utils.CheckEmailMustExist(jsonParams.Email) && utils.CheckStringMustExist(jsonParams.Validate_Code)
	if !paramOk {
		panic(errorcode.ParamError)
	}
	userOp := user.NewUserOp()
	res := userOp.ResetEmail(tokenMap["user_id"].(string), jsonParams.Email.String, jsonParams.Validate_Code.String)
	c.JSON(res.StatusCode, res)
}

// @Summary 发送重置邮箱验证码邮件
// @Description 发送重置邮箱验证码邮件
// @Tags 用户管理
// @Accept  json
// @Produce json
// @security ApiKeyAuth
// @Param   params  body emailParams true "email"
// @Success 200  {object}  response.Response  "Success data为null"
// @Failure 500  {object}  response.Response  "内部错误 data为null"
// @Failure 401  {object}  response.Response  "token错误 data为null"
// @Failure 404  {object}  response.Response  "无法访问 data为null"
// @Router /user/email/reset/validate [post]
func sendResetEmailValidateCode(c *gin.Context, tokenMap map[string]interface{}) {
	var jsonParams emailParams
	err := c.BindJSON(&jsonParams)
	if err != nil {
		panic(errorcode.ParamError)
	}
	paramOk := utils.CheckUUID4MustExist(tokenMap["user_id"]) && utils.CheckEmailMustExist(jsonParams.Email)
	if !paramOk {
		panic(errorcode.ParamError)
	}
	userOp := user.NewUserOp()
	res := userOp.SendResetEmailValidateCode(tokenMap["user_id"].(string), jsonParams.Email.String)
	c.JSON(res.StatusCode, res)
}

// @Summary 编辑用户信息
// @Description 编辑用户信息
// @Tags 用户管理
// @Accept  json
// @Produce json
// @security ApiKeyAuth
// @Param   params  body userInfoParams true "sex:0男，1女，2保密 age:年龄 user_name:用户名，昵称 name:实际姓名 info:其他信息"
// @Success 200  {object}  response.Response  "调用成功 data为null"
// @Failure 500  {object}  response.Response  "内部错误 data为null"
// @Failure 401  {object}  response.Response  "token错误 data为null"
// @Failure 404  {object}  response.Response  "无法访问 data为null"
// @Router /user/info [post]
func editUserInfo(c *gin.Context, tokenMap map[string]interface{}) {
	var jsonParams userInfoParams
	err := c.BindJSON(&jsonParams)
	if err != nil {
		panic(errorcode.ParamError)
	}
	paramOk := utils.CheckUUID4MustExist(tokenMap["user_id"]) && utils.CheckIntIfExist(jsonParams.Sex, jsonParams.Age) && utils.CheckStringIfExist(jsonParams.User_Name, jsonParams.Name, jsonParams.Info)
	if !paramOk {
		panic(errorcode.ParamError)
	}
	userOp := user.NewUserOp()
	res := userOp.EditUserInfo(tokenMap["user_id"].(string), jsonParams.Sex, jsonParams.Age, jsonParams.User_Name, jsonParams.Name, jsonParams.Info)
	c.JSON(res.StatusCode, res)
}

// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags 用户管理
// @Accept  json
// @Produce json
// @security ApiKeyAuth
// @Success 200  {object}  response.Response  "调用成功 data为null"
// @Failure 500  {object}  response.Response  "内部错误 data为null"
// @Failure 401  {object}  response.Response  "token错误 data为null"
// @Failure 404  {object}  response.Response  "无法访问 data为null"
// @Router /user/info [get]
func getUserInfo(c *gin.Context, tokenMap map[string]interface{}) {
	paramOk := utils.CheckUUID4MustExist(tokenMap["user_id"])
	if !paramOk {
		panic(errorcode.ParamError)
	}
	userOp := user.NewUserOp()
	res := userOp.GetUserInfo(tokenMap["user_id"].(string))
	c.JSON(res.StatusCode, res)
}
