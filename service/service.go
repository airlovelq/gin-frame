package service

import (
	"fmt"
	"scoremanager/response"

	"github.com/gin-gonic/gin"
)

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
func StartService() {
	r := gin.Default()
	// r.Use(middle1)
	r.Use(Cors())
	r.Use(response.ErrorHandler)
	r.NoMethod(response.HandleNotFound)
	r.NoRoute(response.HandleNotFound)
	r.GET("/test", test)
	// r.GET("/ping", secret.TokenServiceHandler(getStudent))
	// r.Use(middle2)
	r.Run(":8082") // listen and serve on 0.0.0.0:8080
}
