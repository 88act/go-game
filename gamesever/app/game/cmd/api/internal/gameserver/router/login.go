package router

import (
	"fmt"
	"go-cms/app/game/cmd/api/internal/gameserver/core"
	"go-cms/app/game/cmd/api/internal/gameserver/pb"
	"go-cms/app/usercenter/cmd/rpc/usercenter"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/zeromicro/go-zero/core/logc"

	"google.golang.org/protobuf/proto"
)

type LoginRouter struct {
	znet.BaseRouter
}

func (m *LoginRouter) Handle(req ziface.IRequest) {
	fmt.Println("登录消息 LoginRouter")
	ctx := req.GetConnection().Context()
	logc.Infof(ctx, "有登录 conn=%d, ip=%s", req.GetConnection().GetConnID(), req.GetConnection().RemoteAddrString())
	msg := &pb.GameLogin{}
	err := proto.Unmarshal(req.GetData(), msg)
	if err != nil {
		logMsg := getLogMsg(req) + err.Error()
		logc.Error(ctx, logMsg)
		return
	}
	// 读取用户信息
	rpcReq := usercenter.GetUserInfoReq{Id: msg.UserId}
	rpcResp, err := core.GetWM().SvcCtx.UsercenterRpc.GetUserInfo(ctx, &rpcReq)
	if err != nil {
		logMsg := "读用户信息错误," + getLogMsg(req) + err.Error() + ",msg=" + msg.String()
		logc.Error(ctx, logMsg)
		return
	}
	conn := req.GetConnection()
	conn.SetProperty("pID", rpcResp.User.Id)
	player := core.NewPlayer(conn, rpcResp.User.Id)
	resp := &pb.UserInfo{}
	resp.UserId = rpcResp.User.Id
	resp.RoomId = 1000 // 默认场景id
	resp.Nickname = rpcResp.User.Nickname
	resp.Username = rpcResp.User.Username
	resp.Image = rpcResp.User.Avatar
	resp.Plat = msg.Plat
	resp.Online = 1
	//
	player.Userinfo = resp
	player.RoomId = resp.RoomId
	core.GetWM().AddPlayer(player)
	//
	core.GetWM().SendWorld(ctx, pb.S_LoginResp, resp)
}
