package users

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/global/consts"
	"go-admin/app/global/variable"
)

// 模拟Aop 实现对某个控制器函数的前置和后置回调
type DestoryBefore struct{}

// 前置函数必须具有返回值，这样才能控制流程是否继续往下执行
func (d *DestoryBefore) Before(context *gin.Context) bool {
	userId := context.GetFloat64(consts.ValidatorPrefix + "id")
	variable.ZapLog.Sugar().Infof("模拟 Users 删除操作，Before 回调，用户ID：%.f\n", userId)
	if userId > 10 {
		return true
	}
	return false
}
