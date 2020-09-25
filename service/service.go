package service

import (
	"fmt"
	_ "scoremanager/docs"
	"scoremanager/errorcode"
	"scoremanager/response"
	"scoremanager/secret"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary 测试接口
// @Description Test Swagger
// @Accept  json
// @Produce json
// @Param   some_id     path    int     true        "userId"
// @Success 200 {string} string	"ok"
// @Router /test [get]
func test(c *gin.Context) {
	fmt.Println("test")
	// panic(response.ServerError)
	data := make(map[string]interface{})
	data["j"] = "k"
	res := response.NewSuccess(data)
	c.JSON(res.StatusCode, res)
}

func StartService() {
	r := gin.Default()
	// r.Use(middle1)
	r.Use(Cors())
	r.Use(errorcode.ErrorHandler)
	r.NoMethod(errorcode.HandleNotFound)
	r.NoRoute(errorcode.HandleNotFound)
	r.GET("/test", test)
	user := r.Group("/user")
	{
		user.POST("/login", login)
		user.POST("/register/email/validate", sendEmailRegisterValidateCode)
		user.POST("/register/email", registerByEmail)
		user.POST("/password/reset", secret.TokenServiceHandler(resetPasswordInLoginStatus))
		user.POST("/password/reset/email/validate", sendEmailResetPasswordValidateCode)
		user.POST("/password/reset/email", resetPasswordByEmail)
		user.POST("/email/reset", secret.TokenServiceHandler(resetEmail))
		user.POST("/email/reset/validate", secret.TokenServiceHandler(sendResetEmailValidateCode))
		user.POST("/info", secret.TokenServiceHandler(editUserInfo))
		user.GET("/info", secret.TokenServiceHandler(getUserInfo))
	}

	// url := ginSwagger.URL("http://192.168.2.89:8082/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.GET("/ping", secret.TokenServiceHandler(getStudent))
	// r.Use(middle2)
	r.Run(":8082") // listen and serve on 0.0.0.0:8080
}
