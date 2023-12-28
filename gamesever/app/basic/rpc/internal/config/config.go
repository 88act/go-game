package config

import (
	"go-game/common/config"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB config.DbConf
	//WxMiniConf        config.WxMiniConf
	LocalRes          config.LocalRes
	LogConf           logc.LogConf
	UsercenterRpcConf zrpc.RpcClientConf
}
