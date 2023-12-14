package router

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/zeromicro/go-zero/core/logc"
)

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// Ping Handle
func (m *PingRouter) Handle(req ziface.IRequest) {
	ctx := req.GetConnection().Context()
	logMsg := "ping msg," + getLogMsg(req)
	logc.Info(ctx, logMsg)
}
