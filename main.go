package main

import (
	"scoremanager/service"
)

// @title 框架接口
// @version 0.0.1
// @swagger 3.0
// @description 接口文档
// @BasePath

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @contact.name API Support
// @contact.url http://www.swagger.io/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 192.168.2.89:8082
func main() {
	service.StartService()
}
