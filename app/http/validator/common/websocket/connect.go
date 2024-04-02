package websocket

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/global/consts"
	"go-admin/app/global/variable"
	controllerWs "go-admin/app/http/controller/websocket"
	"go-admin/app/http/validator/core/data_transfer"
	"go.uber.org/zap"
)

// 将验证器成员(字段)绑定到数据传输上下文，方便控制器获取
/**
本函数参数说明：
validatorInterface 实现了验证器接口的结构体
extra_add_data_prefix  验证器绑定参数传递给控制器的数据前缀
context  gin上下文
*/

type Connect struct {
	Token string `form:"token" json:"token" binding:"required,min=10"`
}

func (c Connect) CheckParams(context *gin.Context) {
	// 1. 首先检查是否开启websocket服务配置（在配置项中开启）
	if variable.ConfigYml.GetInt("Websocket.Start") != 1 {
		variable.ZapLog.Error(consts.WsServerNotStartMsg)
		return
	}
	// 2. 基本的验证规则没有通过
	if err := context.ShouldBind(&c); err != nil {
		variable.ZapLog.Error("客户端上线参数不合格", zap.Error(err))
		return
	}
	extraAddBindDataContext := data_transfer.DataAddContext(c, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		variable.ZapLog.Error("websocket-Connect 表单验证器json化失败")
		context.Abort()
		return
	} else {
		if serviceWs, ok := (&controllerWs.Ws{}).OnOpen(extraAddBindDataContext); ok == false {
			variable.ZapLog.Error(consts.WsOpenFailMsg)
		} else {
			(&controllerWs.Ws{}).OnMessage(serviceWs, extraAddBindDataContext) // 注意这里传递的service_ws必须是调用open返回的，必须保证的ws对象的一致性
		}
	}
}
