package router

import (
	"fmt"

	"go-cms/app/game/cmd/api/internal/gameserver/core"
	"go-cms/app/game/cmd/api/internal/gameserver/pb"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/gogf/gf/util/gconv"
	"github.com/zeromicro/go-zero/core/logc"

	"google.golang.org/protobuf/proto"
)

type LoginRouter struct {
	znet.BaseRouter
}

// Ping Handle
func (m *LoginRouter) Handle(req ziface.IRequest) {
	ctx := req.GetConnection().Context()

	msg := &pb.Login{}
	err := proto.Unmarshal(req.GetData(), msg)
	if err != nil {
		logMsg := getLogMsg(req) + err.Error()
		logc.Error(ctx, logMsg)
		return
	}
	// 读取数据库信息 登录用户
	//core.GetWM().ConnList[req.GetConnection()]
	conn := req.GetConnection()
	conn.SetProperty("pID", 1)
	player := core.NewPlayer(conn)
	core.GetWM().AddPlayer(player)

	// pID, err := req.GetConnection().GetProperty("pID")
	// if err != nil {
	// 	req.GetConnection().Stop()
	// 	logMsg := getLogMsg(req) + err.Error()
	// 	logc.Error(ctx, logMsg)
	// 	return
	// }

	//创建一个玩家
	//player := core.NewPlayer(conn)
	//core.WorldMgrObj.AddPlayer(player)
	//将该连接绑定属性PID
	//conn.SetProperty("ConnId", conn.GetConnID())
	//同步周边玩家上线信息，与现实周边玩家信息
	//player.SyncSurrounding()
	//同步当前玩家的初始化坐标信息给客户端，走MsgID:200消息
	//player.BroadCastStartPosition()
	//fmt.Println("=====> Player pIDID = ", player.Pid, " arrived ====")

	//fmt.Println("接收到 login  3 pID=", pID)
	//3. 根据pID得到player对象
	//player := core.WorldMgrObj.GetPlayerByPID(pID.(int64))
	// // 读取数据库
	// memUser, err := service.GetMemUserSev().Get(context.Background(), msg.UserId, "")
	// if err != nil {
	// 	zlog.Error("数据库查询用户失败 Username "+msg.Username, err.Error())
	// 	request.GetConnection().Stop()
	// 	return
	// }
	// cu, err := service.GetJqCustomerSev().Get(context.Background(), memUser.CuId, "")
	// if err != nil {
	// 	zlog.Errorf("查询客户失败CuId=%d, err=%s", memUser.CuId, err.Error())
	// 	request.GetConnection().Stop()
	// 	return
	// }

	// core.WorldMgrObj.KillSameUser(memUser.Id, pID.(int64))
	resp := &pb.UserInfo{}
	resp.UserId = 1                                       //memUser.Id
	resp.RoomId = 1000                                    // 默认房间号
	resp.Nickname = "Realname" + gconv.String(msg.UserId) // memUser.Realname
	//resp.Username = msg.Username                          //"username" + gconv.String(pID)

	fmt.Println("接收到 login  6 pID=", resp)
	//resp.Webrtc = "https://210.0.0.1" //zconf.GlobalObject.Webrtc + "webrtc/push/room/" + cu.Guid + "/" + gconv.String(resp.RoomId) + "/" + gconv.String(resp.UserId)
	resp.Image = "" // memUser.Avatar
	resp.Plat = msg.Plat
	resp.Online = 1
	player.Userinfo = resp
	player.RoomId = resp.RoomId
	core.GetWM().SendWorld(ctx, pb.S_LoginResp, resp)
	// fmt.Printf("登录 resp Plat %s ", resp.Plat)
	// fmt.Println("接收到 login  5 pID=", resp)
	//发送给所有人
	//fmt.Println("LoginRouter 发送给所有人")

	//player.SendMsg(pb.S_UserInfo, resp)

}
