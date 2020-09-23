package service

import (
	"scoremanager/controller/user"
	"scoremanager/errorcode"
	"scoremanager/utils"

	"github.com/gin-gonic/gin"
)

// @Summary 邮件发送注册验证码
// @Description 邮件发送注册验证码
// @Accept  json
// @Produce json
// @Param   email     query    string     true        "邮箱"
// @Success 200 {object}  response.Response	"Success"
// @Failure 500 {object}  response.Response "Failure"
// @Router /user/register/email/validate [post]
func sendEmailRegisterValidateCode(c *gin.Context) {
	jsonParams := make(map[string]interface{})
	err := c.BindJSON(&jsonParams)
	if err != nil {
		panic(errorcode.ParamError)
	}
	paramOk := utils.Check_Email(jsonParams["email"])
	if !paramOk {
		panic(errorcode.ParamError)
	}

	userOp := user.NewUserOp()
	res := userOp.SendEmailRegisterValidateCode(jsonParams["email"].(string))
	c.JSON(res.StatusCode, res)
}

// @Summary 邮件发送重置密码验证码
// @Description 邮件发送重置密码验证码
// @Accept  json
// @Produce json
// @Param   email     query    string     true        "邮箱"
// @Success 200 {object}  response.Response	"Success"
// @Failure 500 {object}  response.Response "Failure"
// @Router /user/password/reset/email/validate [post]
func sendEmailResetPasswordValidateCode(c *gin.Context) {
	jsonParams := make(map[string]interface{})
	err := c.BindJSON(&jsonParams)
	if err != nil {
		panic(errorcode.ParamError)
	}
	paramOk := utils.Check_Email(jsonParams["email"])
	if !paramOk {
		panic(errorcode.ParamError)
	}

	userOp := user.NewUserOp()
	res := userOp.SendEmailResetPasswordValidateCode(jsonParams["email"].(string))
	c.JSON(res.StatusCode, res)
}

// @Summary 邮件注册
// @Description 邮件注册
// @Accept  json
// @Produce json
// @Param   email     body    string     true        "邮箱"
// @Param   validate_code     body    string     true        "验证码"
// @Param   password_encrypt  body string true "密码（加密过的）"
// @Success 200  {object}  response.Response  "Success"
// @Failure 500  {object}  response.Response  "Failure"
// @Router /user/register/email [post]
func registerByEmail(c *gin.Context) {

	jsonParams := make(map[string]interface{})
	err := c.BindJSON(&jsonParams)
	if err != nil {
		panic(errorcode.ParamError)
	}
	// email, _ := c.Get("email")
	// validateCode, _ := c.Get("validate_code")
	// passwordEncrypt, _ := c.Get("password_encrypt")
	paramOk := utils.Check_Email(jsonParams["email"]) && utils.Check_String(jsonParams["validate_code"], jsonParams["password_encrypt"])
	if !paramOk {
		panic(errorcode.ParamError)
	}
	userOp := user.NewUserOp()
	res := userOp.RegisterByEmail(jsonParams["email"].(string), jsonParams["validate_code"].(string), jsonParams["password_encrypt"].(string))
	c.JSON(res.StatusCode, res)
}

// @Summary 重置密码（已登录状态）
// @Description 重置密码（已登录状态）
// @Accept  json
// @Produce json
// @Param   password_encrypt  body string true "密码（加密过的）"
// @Success 200  {object}  response.Response  "Success"
// @Failure 500  {object}  response.Response  "Failure"
// @Router /user/password/reset [post]
func resetPasswordInLoginStatus(c *gin.Context, tokenMap map[string]interface{}) {
	jsonParams := make(map[string]interface{})
	err := c.BindJSON(&jsonParams)
	if err != nil {
		panic(errorcode.ParamError)
	}
	paramOk := utils.Check_UUID4(tokenMap["user_id"]) && utils.Check_String(jsonParams["password_encrypt"])
	if !paramOk {
		panic(errorcode.ParamError)
	}
	userOp := user.NewUserOp()
	res := userOp.ResetPasswordInLoginStatus(tokenMap["user_id"].(string), jsonParams["password_encrypt"].(string))
	c.JSON(res.StatusCode, res)
}

// @Summary 登录
// @Description 登录接口，通过email，手机号或用户名
// @Accept  json
// @Produce json
// @Param   log_id  body string true "email，手机号或用户名"
// @Param   password_encrypt  body string true "密码（加密过的）"
// @Success 200  {object}  response.Response  {"code":200,"data":{"token":"token"},"msg":""}
// @Failure 500  {object}  response.Response  "Failure"
// @Router /user/login [post]
func login(c *gin.Context) {
	jsonParams := make(map[string]interface{})
	err := c.BindJSON(&jsonParams)
	if err != nil {
		panic(errorcode.ParamError)
	}
	paramOk := utils.Check_String(jsonParams["log_id"], jsonParams["password_encrypt"])
	if !paramOk {
		panic(errorcode.ParamError)
	}
	userOp := user.NewUserOp()
	res := userOp.Login(jsonParams["log_id"].(string), jsonParams["password_encrypt"].(string))
	c.JSON(res.StatusCode, res)
}

// @Summary 重置密码（发送邮件方式）
// @Description 重置密码（发送邮件方式），根据邮件中验证码
// @Accept  json
// @Produce json
// @Param   email  body string true "email"
// @Param   password_encrypt  body string true "密码（加密过的）"
// @Param   validate_code  body string true "验证码"
// @Success 200  {object}  response.Response  "Success"
// @Failure 500  {object}  response.Response  "Failure"
// @Router /user/password/reset/email [post]
func resetPasswordByEmail(c *gin.Context) {
	jsonParams := make(map[string]interface{})
	err := c.BindJSON(&jsonParams)
	if err != nil {
		panic(errorcode.ParamError)
	}
	paramOk := utils.Check_Email(jsonParams["email"]) && utils.Check_String(jsonParams["validate_code"], jsonParams["password_encrypt"])
	if !paramOk {
		panic(errorcode.ParamError)
	}
	userOp := user.NewUserOp()
	res := userOp.ResetPasswordByEmailValidateCode(jsonParams["email"].(string), jsonParams["validate_code"].(string), jsonParams["password_encrypt"].(string))
	c.JSON(res.StatusCode, res)
}

// @Summary 重置邮箱
// @Description 重置邮箱，根据邮件中验证码
// @Accept  json
// @Produce json
// @Param   email  body string true "email"
// @Param   validate_code  body string true "验证码"
// @Success 200  {object}  response.Response  "Success"
// @Failure 500  {object}  response.Response  "Failure"
// @Router /user/email/reset [post]
func resetEmail(c *gin.Context, tokenMap map[string]interface{}) {
	jsonParams := make(map[string]interface{})
	err := c.BindJSON(&jsonParams)
	if err != nil {
		panic(errorcode.ParamError)
	}
	paramOk := utils.Check_UUID4(tokenMap["user_id"]) && utils.Check_Email(jsonParams["email"]) && utils.Check_String(jsonParams["validate_code"])
	if !paramOk {
		panic(errorcode.ParamError)
	}
	userOp := user.NewUserOp()
	res := userOp.ResetEmail(tokenMap["user_id"].(string), jsonParams["email"].(string), jsonParams["validate_code"].(string))
	c.JSON(res.StatusCode, res)
}

// @Summary 发送重置邮箱验证码邮件
// @Description 发送重置邮箱验证码邮件
// @Accept  json
// @Produce json
// @Param   email  body string true "email"
// @Success 200  {object}  response.Response  "Success"
// @Failure 500  {object}  response.Response  "Failure"
// @Router /user/email/reset/validate [post]
func sendResetEmailValidateCode(c *gin.Context, tokenMap map[string]interface{}) {
	jsonParams := make(map[string]interface{})
	err := c.BindJSON(&jsonParams)
	if err != nil {
		panic(errorcode.ParamError)
	}
	paramOk := utils.Check_UUID4(tokenMap["user_id"]) && utils.Check_Email(jsonParams["email"])
	if !paramOk {
		panic(errorcode.ParamError)
	}
	userOp := user.NewUserOp()
	res := userOp.SendResetEmailValidateCode(tokenMap["user_id"].(string), jsonParams["email"].(string))
	c.JSON(res.StatusCode, res)
}

// @Summary 编辑用户信息
// @Description 编辑用户信息
// @Accept  json
// @Produce json
// @Param   sex  body int false "0"
// @Param   age  body int false "20"
// @Param   user_name  body string false "haha"
// @Param   name  body string false "haha"
// @Param   info  body string false "other info"
// @Success 200  {object}  response.Response  "Success"
// @Failure 500  {object}  response.Response  "Failure"
// @Router /user/info [post]
func editUserInfo(c *gin.Context, tokenMap map[string]interface{}) {
	jsonParams := make(map[string]interface{})
	err := c.BindJSON(&jsonParams)
	if err != nil {
		panic(errorcode.ParamError)
	}
	paramOk := utils.Check_UUID4(tokenMap["user_id"]) && utils.CheckNumIfExist(jsonParams["sex"], jsonParams["age"]) && utils.CheckStringIfExist(jsonParams["user_name"], jsonParams["name"], jsonParams["info"]) && len(jsonParams) >= 1
	if !paramOk {
		panic(errorcode.ParamError)
	}
	userOp := user.NewUserOp()
	res := userOp.EditUserInfo(tokenMap["user_id"].(string), jsonParams)
	c.JSON(res.StatusCode, res)
}
