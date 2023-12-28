package main

import (
	"flag"
	"fmt"

	"go-game/app/usercenter/rpc/internal/config"
	"go-game/app/usercenter/rpc/internal/server"
	"go-game/app/usercenter/rpc/internal/svc"
	"go-game/app/usercenter/rpc/pb"
	"go-game/common/interceptor/rpcserver"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUsercenterServer(ctx)
	//logx.DisableStat()
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUsercenterServer(grpcServer, srv)
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	//rpc log
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)
	defer s.Stop()
	fmt.Printf("Starting  usercenter rpc server at %s...\n", c.ListenOn)
	s.Start()
}
