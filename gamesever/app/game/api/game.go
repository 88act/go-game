package main

import (
	"flag"
	"fmt"
	"go-game/app/game/api/internal/config"
	"go-game/app/game/api/internal/gameserver"
	"go-game/app/game/api/internal/handler"
	"go-game/app/game/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/game.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	handler.RegisterHandlers(server, ctx)
	gameServer := gameserver.GetGameServer(c.GameConf, ctx)
	//defer server.Stop()
	//server.Start()
	group := service.NewServiceGroup()
	group.Add(server)
	group.Add(gameServer)
	defer group.Stop()
	fmt.Printf("启动 game api  server at %s:%d...\n", c.Host, c.Port)
	fmt.Printf("启动 gameServer  at %s:%d...\n", c.GameConf.Host, c.GameConf.WsPort)
	group.Start()
}
