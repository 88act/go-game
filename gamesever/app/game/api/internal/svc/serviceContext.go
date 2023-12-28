package svc

import (
	"context"
	"go-game/app/game/api/internal/config"
	"go-game/app/usercenter/rpc/usercenter"

	"go-game/common/myconfig"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	//BasicRpc      basic.Basic
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	logc.MustSetup(c.LogConf)
	logc.Info(context.Background(), c.Name, " 服务启动...", c.Host, " port=", c.Port)
	myconfig.HttpRoot = c.LocalRes.BaseUrl
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		//BasicRpc:      basic.NewBasic(zrpc.MustNewClient(c.BasicRpcConf)),
	}

}
