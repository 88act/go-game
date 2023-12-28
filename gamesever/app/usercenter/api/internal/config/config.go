package config

import (
	"go-game/common/config"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth           config.JwtAuth
	DB                config.DbConf
	Redis             redis.RedisConf
	WxMiniConf        config.WxMiniConf
	LocalRes          config.LocalRes
	LogConf           logc.LogConf
	UsercenterRpcConf zrpc.RpcClientConf
	//BasicRpcConf      zrpc.RpcClientConf
	//Cache             cache.CacheConf
}
