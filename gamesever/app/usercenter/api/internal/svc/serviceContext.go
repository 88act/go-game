package svc

import (
	"context"
	"go-game/app/usercenter/api/internal/config"
	"go-game/app/usercenter/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	//RedisClient *redis.Redis
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {

	logc.MustSetup(c.LogConf)
	logc.Info(context.Background(), "UsercenterApi  服务器启动...", c.Host, " port=", c.Port)

	return &ServiceContext{
		Config:        c,
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
