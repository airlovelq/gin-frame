package user

import (
	"scoremanager/controller"
	"scoremanager/database"
	"scoremanager/errorcode"
	"scoremanager/middleware/cache"
	"scoremanager/response"
	"scoremanager/secret"
	"scoremanager/utils"

	"gopkg.in/guregu/null.v3"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	emailSrc    string = "email"
	passwordSrc string = "password"
	phoneSrc    string = "phone"
)

const (
	registerNamespace         string = "register"
	registerDeadlineNamespace string = "register-deadline"
	resetNamespace            string = "reset"
	resetDeadlineNamespace    string = "reset-deadline"
)

func decryptPassword(passwordEncrypt string, keyFile string) (string, error) {
	rsaPrivateKey, err := secret.LoadPrivateKeyFile(keyFile)
	if err != nil {
		return "", err
	}
	passwordreal, err := secret.Decrypt(passwordEncrypt, rsaPrivateKey)
	return passwordreal, err
}

func getPasswordHash(passwordEncrypt string, keyFile string) ([]byte, error) {
	passwordreal, err := decryptPassword(passwordEncrypt, keyFile)
	if err != nil {
		return nil, err
	}
	passwordhash, err := bcrypt.GenerateFromPassword([]byte(passwordreal), bcrypt.DefaultCost)
	return passwordhash, err
}

func checkPassword(passwordEncrypt string, keyFile string, passwordHash []byte) error {
	passwordreal, err := decryptPassword(passwordEncrypt, keyFile)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(passwordHash, []byte(passwordreal))
	return err
}

type UserOp struct {
	*controller.BaseOp
	// cacheOp cache.CacheOp
	// dbOp    *database.DatabaseOp
}

func NewUserOp() *UserOp {
	userOp := UserOp{
		BaseOp: controller.NewBaseOp(),
	}
	return &userOp
}

func (op *UserOp) checkValidateCode(cacheKey string, validateCode string) {
	validateCodeInCache, err := op.CacheOp.Get(cacheKey)
	if err != nil {
		panic(errorcode.ValidateCodeExpiredError)
	}
	validateCodeInCacheStr, err := op.CacheOp.String(validateCodeInCache)
	if err != nil {
		panic(errorcode.ValidateCodeExpiredError)
	}
	if validateCodeInCacheStr != validateCode {
		panic(errorcode.ValidateCodeNotMatchError)
	}
}

func (op *UserOp) Login(logID string, passwordEncrypt string) *response.Response {
	op.BeginOp()
	defer op.EndOp()
	var user database.User
	var err error
	if utils.IsEmail(logID) {
		user, err = op.DbOp.GetUserByEmail(logID)
	} else if utils.IsPhone(logID) {
		user, err = op.DbOp.GetUserByPhone(logID)
	} else {
		user, err = op.DbOp.GetUserByUserName(logID)
	}
	if err != nil {
		panic(errorcode.UserNotExistError)
	}
	if user.Banned == 1 {
		panic(errorcode.UserAlreadyBanned)
	}
	err = checkPassword(passwordEncrypt, utils.GetEnvDefault("PRIVATE_KEY_PATH", "/home/luqin")+"/private.pem", []byte(user.Password_Hash))
	if err != nil {
		panic(errorcode.PasswordNotMatch)
	}
	token, err := secret.GenerateToken(user.ID.String(), user.User_Type)
	resdata := make(map[string]interface{})
	resdata["user_id"] = user.ID
	resdata["token"] = token
	return response.NewSuccess(resdata)
}

func (op *UserOp) SendEmailRegisterValidateCode(email string) *response.Response {
	op.BeginOp()
	defer op.EndOp()
	if _, err := op.DbOp.GetUserByEmail(email); err == nil {
		panic(errorcode.UserAlreadyExistError)
	}
	registerDeadlineCacheKey := registerDeadlineNamespace + ":email:" + email
	dl, err := op.CacheOp.Get(registerDeadlineCacheKey)
	if dl != nil {
		panic(errorcode.TooFrequentError)
	}
	validateCode := uuid.NewV4()
	registerCacheKey := registerNamespace + ":email:" + email
	err = op.CacheOp.Set(registerCacheKey, validateCode.String(), cache.WithEx(3600))
	if err != nil {
		panic(errorcode.CacheOperationError)
	}
	emailContent := "<p>Hi,</p> <p style=\"text-indent:2em;\">Welcome to register for Platform user.</p> <p style=\"text-indent:2em;\">Your validation code is </p> <p style=\"text-indent:2em;color:red\"><B>" + validateCode.String() +
		"</B></p> <p style=\"text-indent:2em;\">The validation code will expire after 1 hours. If expires, retrieve it again.</p> <p style=\"text-indent:16em;\"> ------ Platform </p>"

	err = utils.SendMail(email, "The Register Validate Code", emailContent)
	if err != nil {
		panic(errorcode.SendEmailError)
	}
	_ = op.CacheOp.Set(registerDeadlineCacheKey, "1", cache.WithEx(60))
	return response.NewSuccess(nil)
}

func (op *UserOp) SendPhoneRegisterValidateCode(email string) error {

	return nil
}

func (op *UserOp) RegisterByEmail(email string, validateCode string, passwordEncrypt string) *response.Response {
	op.BeginOp()
	defer op.EndOp()
	_, err := op.DbOp.GetUserByEmail(email)
	if err == nil {
		panic(errorcode.UserAlreadyExistError)
	}
	registerCacheKey := registerNamespace + ":email:" + email
	op.checkValidateCode(registerCacheKey, validateCode)
	passwordhash, err := getPasswordHash(passwordEncrypt, utils.GetEnvDefault("PRIVATE_KEY_PATH", "/home/luqin")+"/private.pem")
	userID, err := op.DbOp.CreateUserByEmail(email, string(passwordhash), 1)
	if err != nil {
		panic(errorcode.CreateUserError)
	}
	op.CacheOp.Delete(registerCacheKey)
	res := make(map[string]interface{})
	res["id"] = userID
	return response.NewSuccess(res)
}

func (op *UserOp) RegisterByPhone(phone string, validateCode string, passwordEncrypt string) error {
	return nil
}

func (op *UserOp) SendEmailResetPasswordValidateCode(email string) *response.Response {
	op.BeginOp()
	defer op.EndOp()
	_, err := op.DbOp.GetUserByEmail(email)
	if err != nil {
		panic(errorcode.UserNotExistError)
	}
	validateCode := uuid.NewV4().String()
	emailContent := "<p>Hi,</p> <p style=\"text-indent:2em;\">Welcome to reset password for Platform user.</p> <p style=\"text-indent:2em;\">Your validation code is </p> <p style=\"text-indent:2em;color:red\"><B>" + validateCode +
		"</B></p> <p style=\"text-indent:2em;\">The validation code will expire after 1 hours. If expires, retrieve it again.</p> <p style=\"text-indent:16em;\"> ------ Platform </p>"
	op.sendResetEmail(email, passwordSrc, email, emailContent, validateCode)
	return response.NewSuccess(nil)
}

func (op *UserOp) SendPhoneResetPasswordValidateCode(email string) error {
	return nil
}

func (op *UserOp) ResetPasswordInLoginStatus(userID string, passwordEncrypt string) *response.Response {
	op.BeginOp()
	defer op.EndOp()
	passwordhash, err := getPasswordHash(passwordEncrypt, utils.GetEnvDefault("PRIVATE_KEY_PATH", "/home/luqin")+"/private.pem")

	err = op.DbOp.ResetPasswordByID(userID, string(passwordhash))
	if err != nil {
		panic(errorcode.ResetPasswordError)
	}
	return response.NewSuccess(nil)
}

func (op *UserOp) ResetPasswordByEmailValidateCode(email string, validateCode string, passwordEncrypt string) *response.Response {
	op.BeginOp()
	defer op.EndOp()
	_, err := op.DbOp.GetUserByEmail(email)
	if err != nil {
		panic(errorcode.UserNotExistError)
	}
	resetCacheKey := resetNamespace + ":password:" + email
	op.checkValidateCode(resetCacheKey, validateCode)
	passwordhash, err := getPasswordHash(passwordEncrypt, utils.GetEnvDefault("PRIVATE_KEY_PATH", "/home/luqin")+"/private.pem")
	err = op.DbOp.ResetPasswordByEmail(email, string(passwordhash))
	if err != nil {
		panic(errorcode.ResetPasswordError)
	}
	op.CacheOp.Delete(resetCacheKey)
	return response.NewSuccess(nil)
}

func (op *UserOp) ResetPasswordByPhoneValidateCode(phone string, validateCode string, passwordEncrypt string) *response.Response {
	return nil
}

func (op *UserOp) SendResetEmailValidateCode(userID string, email string) *response.Response {
	op.BeginOp(controller.NotUseDBOption)
	defer op.EndOp()
	validateCode := uuid.NewV4().String()
	emailContent := "<p>Hi,</p> <p style=\"text-indent:2em;\">Welcome to reset email for Platform user.</p> <p style=\"text-indent:2em;\">Your validation code is </p> <p style=\"text-indent:2em;color:red\"><B>" + validateCode +
		"</B></p> <p style=\"text-indent:2em;\">The validation code will expire after 1 hours. If expires, retrieve it again.</p> <p style=\"text-indent:16em;\"> ------ Platform </p>"
	op.sendResetEmail(userID+":"+email, emailSrc, email, emailContent, validateCode)
	return response.NewSuccess(nil)
}

func (op *UserOp) ResetEmail(userID string, email string, validateCode string) *response.Response {
	op.BeginOp()
	defer op.EndOp()
	resetCacheKey := resetNamespace + ":email:" + userID + ":" + email
	op.checkValidateCode(resetCacheKey, validateCode)
	err := op.DbOp.ResetEmailByID(userID, email)
	if err != nil {
		panic(errorcode.ResetEmailError)
	}
	op.CacheOp.Delete(resetCacheKey)
	return response.NewSuccess(nil)
}

func (op *UserOp) SendResetPhoneValidateCode(userID string, email string) *response.Response {
	return nil
}

func (op *UserOp) ResetPhone(userID string, phone string, validateCode string) *response.Response {
	return nil
}

func (op *UserOp) sendResetEmail(source string, stype string, email string, emailContent string, validateCode string) {
	resetDeadlineCacheKey := resetDeadlineNamespace + ":" + stype + ":" + source
	dl, err := op.CacheOp.Get(resetDeadlineCacheKey)
	if dl != nil {
		panic(errorcode.TooFrequentError)
	}
	resetCacheKey := resetNamespace + ":" + stype + ":" + source
	err = op.CacheOp.Set(resetCacheKey, validateCode, cache.WithEx(3600))
	if err != nil {
		panic(errorcode.CacheOperationError)
	}
	err = utils.SendMail(email, "The Reset "+stype+" Validate Code", emailContent)
	if err != nil {
		panic(errorcode.SendEmailError)
	}
	_ = op.CacheOp.Set(resetDeadlineCacheKey, "1", cache.WithEx(60))
}

func (op *UserOp) sendResetPhoneMsg(source string, stype string, phone string) {

}

func (op *UserOp) EditUserInfo(userID string, sex null.Int, age null.Int, userName null.String, name null.String, info null.String) *response.Response {
	op.BeginOp()
	defer op.EndOp()
	// // phone email必须使用验证码编辑的接口
	// delete(infos, "phone")
	// delete(infos, "email")
	if err := op.DbOp.UpdateUserInfo(userID, sex, age, userName, name, info); err != nil {
		panic(errorcode.UpdateUserInfoError)
	}
	return response.NewSuccess(nil)
}

func (op *UserOp) GetUserInfo(userID string) *response.Response {
	op.BeginOp()
	defer op.EndOp()
	if user, err := op.DbOp.GetUserByUserID(userID); err != nil {
		panic(errorcode.UpdateUserInfoError)
	} else {
		return response.NewSuccess(user)
	}
}
