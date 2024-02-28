package bootstrap

import (
	"go-admin/app/global/errors"
	"go-admin/app/global/variable"
	"log"
	"os"
)

func init() {
	// 初始化程序根目录
	if path, err := os.Getwd(); err == nil {
		variable.BASE_PATH = path
	} else {
		log.Fatal(errors.Errors_BasePath, err)
	}
}
