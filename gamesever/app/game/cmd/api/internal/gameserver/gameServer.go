package gameserver

import (
	"fmt"
	"os"
	"sync"
	"time"

	"go-cms/app/game/cmd/api/internal/config"
	"go-cms/app/game/cmd/api/internal/gameserver/core"
	"go-cms/app/game/cmd/api/internal/gameserver/pb"
	"go-cms/app/game/cmd/api/internal/gameserver/router"

	"github.com/aceld/zinx/zconf"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"google.golang.org/protobuf/proto"
)

type GameServer struct {
	*rest.Server
	Sev  ziface.IServer
	Conf config.GameConf
}

var once_GameServer sync.Once = sync.Once{}
var obj_GameServer *GameServer

func GetGameServer(conf config.GameConf) *GameServer {
	once_GameServer.Do(func() {
		obj_GameServer = new(GameServer)
		obj_GameServer.Conf = conf
	})
	return obj_GameServer
}

// Stop stops the gateway server.
func (m *GameServer) Stop() {
	fmt.Println("Stop=====1=")
	m.Sev.Stop()
	logx.Error("game server stop ...")
	fmt.Println("Stop=====2=")
	m.Server.Stop()
	os.Exit(0)
}

// Start starts the gateway server.
func (m *GameServer) Start() {
	zconf.GlobalObject.Mode = m.Conf.Mode
	zconf.GlobalObject.Name = m.Conf.Name
	zconf.GlobalObject.Host = m.Conf.Host
	zconf.GlobalObject.MaxConn = m.Conf.MaxConn
	zconf.GlobalObject.LogDir = m.Conf.LogDir
	zconf.GlobalObject.LogIsolationLevel = zconf.GlobalObject.LogIsolationLevel
	zconf.GlobalObject.LogFile = m.Conf.LogFile
	zconf.GlobalObject.WsPort = m.Conf.WsPort
	zconf.GlobalObject.HeartbeatMax = m.Conf.HeartbeatMax
	zconf.GlobalObject.WorkerPoolSize = m.Conf.WorkerPoolSize
	zconf.GlobalObject.MaxConn = m.Conf.MaxConn

	m.Sev = znet.NewServer()

	//s.AddRouter(pb.C_Ping, &router.PingRouter{})
	m.Sev.AddRouter(pb.C_Login, &router.LoginRouter{})
	m.Sev.AddRouter(pb.C_EnterRoom, &router.EnterRoomRouter{})
	m.Sev.AddRouter(pb.C_QuitRoom, &router.QuitRoomRouter{})
	m.Sev.AddRouter(pb.C_ChatMsg, &router.ChatMsgRouter{})
	m.Sev.AddRouter(101, &router.PositionZinxRouter{})
	m.Sev.SetOnConnStart(OnConnectionAdd)
	m.Sev.SetOnConnStop(OnConnectionLost)

	// Start heartbeating detection. (启动心跳检测)
	//s.StartHeartBeat(10 * time.Second)
	m.Sev.StartHeartBeatWithOption(10*time.Second, &ziface.HeartBeatOption{
		MakeMsg:          pingMsg,
		OnRemoteNotAlive: myOnRemoteNotAlive,
		Router:           &router.PingRouter{},
		HeartBeatMsgID:   uint32(2000),
	})
	logx.Error("game server start ...")
	m.Sev.Start()
}

// 当客户端建立连接的时候的hook函数
func OnConnectionAdd(conn ziface.IConnection) {
	fmt.Println("=====> OnConnecionAdd is Called ...")
	//创建一个玩家
	player := core.NewPlayer(conn)
	//同步当前的PlayerID给客户端， 走MsgID:1 消息
	//player.SyncPID()
	//将当前新上线玩家添加到worldManager中
	core.WorldMgrObj.AddPlayer(player)
	//将该连接绑定属性PID
	conn.SetProperty("pID", player.PID)
	//同步周边玩家上线信息，与现实周边玩家信息
	//player.SyncSurrounding()
	//同步当前玩家的初始化坐标信息给客户端，走MsgID:200消息
	//player.BroadCastStartPosition()
	fmt.Println("=====> Player pIDID = ", player.PID, " arrived ====")
}

// 当客户端断开连接的时候的hook函数
func OnConnectionLost(conn ziface.IConnection) {
	//获取当前连接的PID属性
	pID, _ := conn.GetProperty("pID")
	var playerID int64
	if pID != nil {
		playerID = pID.(int64)
	}

	//根据pID获取对应的玩家对象
	player := core.WorldMgrObj.GetPlayerByPID(pID.(int64))
	fmt.Println("客户端断开连接 player  PID = UserId = ", pID, player.Userinfo.UserId)
	// 退出房间
	// if player.Userinfo.RoomId > 0 {
	// 	roomPlayer := core.WorldMgrObj.GetRoomPlayers(player.Userinfo.RoomId)
	// 	resp := &pb.QuitRoomResp{}
	// 	resp.UserInfo = player.Userinfo
	// 	//3 向所有房间人发送  不包括自己
	// 	for _, p := range roomPlayer {
	// 		fmt.Println("====2222=p.PID ", p.PID, playerID)
	// 		if p.PID != playerID {
	// 			fmt.Println("====2222= ")
	// 			p.SendMsg(pb.S_QuitRoomResp, resp)
	// 		}
	// 	}
	// }
	fmt.Println("触发玩家下线业务 pID=", pID, "player userid=", player.Userinfo.UserId)
	//触发玩家下线业务

	list := core.WorldMgrObj.GetAllPlayers()
	fmt.Println(list)
	fmt.Println("====> Player ", playerID, " 下线 或 断开 =====")
	fmt.Println("OnConnectionLost  剩余 用户 = ", len(list))

}

// ==================================
// type myHeartBeatRouter struct {
// 	znet.BaseRouter
// }

// // Handle -
// func (r *myHeartBeatRouter) Handle(req ziface.IRequest) {
// 	zlog.Ins().InfoF("Recv Heartbeat from %s, MsgID = %+v, Data = %s",
// 		req.GetConnection().RemoteAddr(), req.GetMsgID(), string(req.GetData()))
// }

// 用户自定义的心跳检测消息处理方法
func pingMsg(conn ziface.IConnection) []byte {
	// return []byte("heartbeat, I am server, I am alive")
	pingResp := &pb.PingResp{}
	pingResp.NowTime = time.Now().Format("2016-01-02 15:04:05")
	pingResp.Sn = 1
	data, err := proto.Marshal(pingResp)
	if err != nil {
		fmt.Println("err ==== ", err.Error())
	}
	return data
}

// 用户自定义的远程连接不存活时的处理方法
func myOnRemoteNotAlive(conn ziface.IConnection) {
	fmt.Println(" 不存活 myOnRemoteNotAlive is Called, connID=", conn.GetConnID(), "remoteAddr = ", conn.RemoteAddr())
	//关闭链接
	conn.Stop()
}
