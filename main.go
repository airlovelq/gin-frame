package main

import (
	"scoremanager/service"
)

// @title 框架接口
// @version 0.0.1
// @description 接口文档
// @BasePath
// @contact.name API Support
// @contact.url http://www.swagger.io/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:9009
func main() {
	service.StartService()
}
