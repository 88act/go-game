package svc

import (
	"context"
	"go-game/app/basic/api/internal/config"
	"go-game/app/basic/rpc/basic"
	"go-game/app/usercenter/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	BasicRpc      basic.Basic
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	logc.MustSetup(c.LogConf)
	logc.Info(context.Background(), c.Name, " 服务启动...", c.Host, " port=", c.Port)

	return &ServiceContext{
		Config:        c,
		BasicRpc:      basic.NewBasic(zrpc.MustNewClient(c.BasicRpcConf)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
