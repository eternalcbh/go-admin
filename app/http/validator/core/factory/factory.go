package factory

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/core/container"
	"go-admin/app/global/my_errors"
	"go-admin/app/global/variable"
	"go-admin/app/http/validator/core/interf"
)

// 表单参数验证器工厂（请勿修改）
func Create(key string) func(context *gin.Context) {

	if value := container.CreateContainersFactory().Get(key); value != nil {
		if val, isOk := value.(interf.ValidatorInterface); isOk {
			return val.CheckParams
		}
	}
	variable.ZapLog.Error(my_errors.ErrorsValidatorNotExists + ", 验证器模块：" + key)
	return nil
}
