package service

import (
	"fmt"
	_ "scoremanager/docs"
	"scoremanager/errorcode"
	"scoremanager/response"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary 测试接口
// @Description Test Swagger
// @Accept  json
// @Produce json
// @Param   some_id     path    int     true        "userId"
// @Success 200 {object} response.Response	"ok"
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
	// url := ginSwagger.URL("http://192.168.2.89:8082/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.GET("/ping", secret.TokenServiceHandler(getStudent))
	// r.Use(middle2)
	r.Run(":8082") // listen and serve on 0.0.0.0:8080
}
