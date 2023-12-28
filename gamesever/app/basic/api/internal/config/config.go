package config

import (
	"go-game/common/config"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	LocalRes          config.LocalRes
	LogConf           logc.LogConf
	BasicRpcConf      zrpc.RpcClientConf
	UsercenterRpcConf zrpc.RpcClientConf
	JwtAuth           config.JwtAuth
	//DB                common.DbConf
	//WxMiniConf config.WxMiniConf
}
