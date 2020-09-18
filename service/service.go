package service

import (
	"fmt"
	"scoremanager/response"

	_ "scoremanager/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @获取指定ID记录
// @Description get record by ID
// @Accept  json
// @Produce json
// @Param   some_id     path    int     true        "userId"
// @Success 200 {string} string	"ok"
// @Router /record/{some_id} [get]
func test(c *gin.Context) {
	fmt.Println("test")
	// panic(response.ServerError)
	data := make(map[string]interface{})
	data["j"] = "k"
	res := response.NewSuccess(data)
	c.JSON(res.StatusCode, res)
}

// func getStudent(c *gin.Context, tokenMap map[string]interface{}) {
// 	// panic(response.ServerError)
// 	c.JSON(200, gin.H{
// 		"message": "pong",
// 	})
// }

// func middle1(c *gin.Context) {
// 	defer fmt.Println("1")
// 	c.Next()
// }

// func middle2(c *gin.Context) {
// 	fmt.Println("2")
// }

// @title 测试
// @version 0.0.1
// @description  测试
// @BasePath /api/v1/
func StartService() {
	r := gin.Default()
	// r.Use(middle1)
	r.Use(Cors())
	r.Use(response.ErrorHandler)
	r.NoMethod(response.HandleNotFound)
	r.NoRoute(response.HandleNotFound)
	r.GET("/test", test)
	// url := ginSwagger.URL("http://192.168.2.89:8082/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.GET("/ping", secret.TokenServiceHandler(getStudent))
	// r.Use(middle2)
	r.Run(":8082") // listen and serve on 0.0.0.0:8080
}
