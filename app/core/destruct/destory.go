package destruct

import (
	"go-admin/app/core/event"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	// 用于系统信号监听
	go func() {
		c := make(chan os.Signal)
		// 监听器
		signal.Notify(c)
		received := <-c
		switch received {
		case os.Interrupt, os.Kill, syscall.SIGQUIT:
			event.CreateEventManageFactory().Dispatch("")
		}
		os.Exit(1)
	}()
}
