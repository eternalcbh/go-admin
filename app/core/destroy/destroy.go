package destroy

import (
	"go-admin/app/core/event_manage"
	"go-admin/app/global/consts"
	"go-admin/app/global/variable"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	// 用于系统信号的监听
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM) // 监听可能的退出信号
		received := <-c
		variable.ZapLog.Warn(consts.ProcessKilled, zap.String("信号值", received.String()))
		(event_manage.CreateEventManageFactory()).FuzzyCall(variable.EventDestroyPrefix)
		close(c)
		os.Exit(1)
	}()
}
