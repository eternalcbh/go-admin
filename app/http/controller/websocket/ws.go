package websocket

import (
	"github.com/gin-gonic/gin"
	serviceWs "go-admin/app/service/websocket"
)

type Ws struct {
}

// OnOpen 主要解决握手+协议升级
func (w *Ws) OnOpen(context *gin.Context) (*serviceWs.Ws, bool) {
	return (&serviceWs.Ws{}).OnOpen(context)
}

// OnMessage 处理业务消息
func (w *Ws) OnMessage(serviceWs *serviceWs.Ws, context *gin.Context) {
	serviceWs.OnMessage(context)
}
