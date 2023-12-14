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

	"go-cms/app/game/cmd/api/internal/svc"

	"github.com/aceld/zinx/zconf"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"google.golang.org/protobuf/proto"
)

type GameServer struct {
	*rest.Server
	Sev    ziface.IServer
	Conf   config.GameConf
	SvcCtx *svc.ServiceContext
}

var once_GameServer sync.Once = sync.Once{}
var GSObj *GameServer // 游戏服务器

func GetGameServer(conf config.GameConf, svcCtx *svc.ServiceContext) *GameServer {
	once_GameServer.Do(func() {
		GSObj = new(GameServer)
		GSObj.Conf = conf
		GSObj.SvcCtx = svcCtx
	})
	return GSObj
}

// Stop stops the gateway server.
func (m *GameServer) Stop() {
	m.Sev.Stop()
	logx.Error("GameServer stop ...")
	m.Server.Stop()
	os.Exit(0)
}

// Start starts the gateway server.
func (m *GameServer) Start() {
	zlog.SetLogLevel(zlog.LogInfo)
	// zlog.Debug("===> 调试Debug：debug不应该出现")
	// zlog.Info("===> 调试Debug：info应该出现")
	// zlog.Warn("===> 调试Debug：warn应该出现")
	// zlog.Error("===> 调试Debug：error应该出现")

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
	core.GetWM().SvcCtx = m.SvcCtx
}

// 建立连接
func OnConnectionAdd(conn ziface.IConnection) {
	fmt.Println("有新连接请求 %d, ip=%s", conn.GetConnID(), conn.RemoteAddrString())
	logx.Infof("有新连接请求 %d, ip=%s", conn.GetConnID(), conn.RemoteAddrString())
	//core.GetWM().ConnList[conn.GetConnID()] = conn
}

// 断开连接
func OnConnectionLost(conn ziface.IConnection) {
	fmt.Println("断开连接 %d, ip=%s", conn.GetConnID(), conn.RemoteAddrString())
	//获取当前连接的PID属性
	pID, _ := conn.GetProperty("pID")
	player := core.GetWM().GetPlayerByPID(pID.(int64))
	if player != nil {
		logx.Error("用户断开连接 pid=", pID, player.Userinfo.Username)
		core.GetWM().RemovePlayerByPID(player.Pid)
	}
	// //根据pID获取对应的玩家对象
	// player := core.GetWM().GetPlayerByPID(playerID)
	// fmt.Println("客户端断开连接 player  PID = UserId = ", pID, player.Userinfo.UserId)
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
	// fmt.Println("触发玩家下线业务 pID=", pID, "player userid=", player.Userinfo.UserId)
	// //触发玩家下线业务
	// list := core.WorldMgrObj.GetAllPlayers()
	// fmt.Println(list)
	// fmt.Println("====> Player ", playerID, " 下线 或 断开 =====")
	// fmt.Println("OnConnectionLost  剩余 用户 = ", len(list))
}

// 心跳检测
func pingMsg(conn ziface.IConnection) []byte {
	pingResp := &pb.PingResp{}
	pingResp.NowTime = time.Now().Format("2016-01-02 15:04:05")
	pingResp.Sn = 1
	data, err := proto.Marshal(pingResp)
	if err != nil {
		logx.Errorf("pingMsg error connID= %d,ip=%s,err=%s ", conn.GetConnID(), conn.RemoteAddrString(), err.Error())
	}
	return data
}

// 远程连接不存活时
func myOnRemoteNotAlive(conn ziface.IConnection) {
	logx.Errorf("myOnRemoteNotAlive不存活 connID= %d,ip=%s ", conn.GetConnID(), conn.RemoteAddrString())
	//关闭链接
	conn.Stop()
}
