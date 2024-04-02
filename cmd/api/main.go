package main

import (
	"go-admin/app/global/variable"
	"go-admin/routers"
)

// 这里可以存放门户类网入口
func main() {
	router := routers.InitApiRouter()
	_ = router.Run(variable.ConfigYml.GetString("HttpServer.Api.Port"))
}
