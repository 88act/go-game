package config

import (
	"go-game/common"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	DB common.DbConf
	//	Redis      redis.RedisConf
	WxMiniConf WxMiniConf
	LocalRes   struct {
		BaseUrl  string
		BasePath string
		Path     string
		PathUser string
	}
	LogConf logc.LogConf
}
