package main

import (
	"go-admin/app/global/variable"
	"go-admin/routers"
)

// 这里可以存放后端路由
func main() {
	router := routers.InitWebRouter()
	_ = router.Run(variable.ConfigYml.GetString("HttpServer.Web.Port"))
}
