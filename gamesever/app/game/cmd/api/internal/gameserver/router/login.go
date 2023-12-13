package router

import (
	"fmt"

	"go-cms/app/game/cmd/api/internal/gameserver/core"
	"go-cms/app/game/cmd/api/internal/gameserver/pb"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
	"github.com/gogf/gf/util/gconv"

	"google.golang.org/protobuf/proto"
)

type LoginRouter struct {
	znet.BaseRouter
}

// Ping Handle
func (m *LoginRouter) Handle(request ziface.IRequest) {

	msg := &pb.Login{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		zlog.Errorf()
		return
	}

	pID, err := request.GetConnection().GetProperty("pID")
	if err != nil {
		fmt.Println("GetProperty pID error", err)
		request.GetConnection().Stop()
		return
	}

	//fmt.Println("接收到 login  3 pID=", pID)
	//3. 根据pID得到player对象
	player := core.WorldMgrObj.GetPlayerByPID(pID.(int64))
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
	resp.UserId = msg.UserId                              //memUser.Id
	resp.RoomId = 1000                                    // 默认房间号
	resp.Nickname = "Realname" + gconv.String(msg.UserId) // memUser.Realname
	resp.Username = msg.Username                          //"username" + gconv.String(pID)

	fmt.Println("接收到 login  6 pID=", resp)
	resp.Webrtc = "https://210.0.0.1" //zconf.GlobalObject.Webrtc + "webrtc/push/room/" + cu.Guid + "/" + gconv.String(resp.RoomId) + "/" + gconv.String(resp.UserId)
	resp.Image = ""                   // memUser.Avatar
	resp.Plat = msg.Plat
	resp.Online = 1
	player.Userinfo = resp
	// fmt.Printf("登录 resp Plat %s ", resp.Plat)
	// fmt.Println("接收到 login  5 pID=", resp)
	//发送给所有人
	//fmt.Println("LoginRouter 发送给所有人")
	core.WorldMgrObj.SendAll(pb.S_UserInfo, resp)
	player.SendMsg(pb.S_UserInfo, resp)

}
