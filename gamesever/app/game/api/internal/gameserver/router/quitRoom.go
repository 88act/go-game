package router

import (
	"fmt"
	"go-game/app/game/api/internal/gameserver/core"
	"go-game/app/game/api/internal/gameserver/pb"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/protobuf/proto"
)

type QuitRoomRouter struct {
	znet.BaseRouter
}

// 退出场景
func (m *QuitRoomRouter) Handle(req ziface.IRequest) {
	fmt.Println("退出场景 QuitRoomRouter")
	ctx := req.GetConnection().Context()
	pID, err := req.GetConnection().GetProperty("pID")
	if err != nil {
		req.GetConnection().Stop()
		logMsg := getLogMsg(req) + err.Error()
		logc.Error(ctx, logMsg)
		return
	}
	player := core.WorldMgrObj.GetPlayerByPID(pID.(int64))
	msg := &pb.EnterRoom{}
	err = proto.Unmarshal(req.GetData(), msg)
	if err != nil {
		logMsg := getLogMsg(req) + err.Error()
		logc.Error(ctx, logMsg)
		return
	}
	fmt.Println("退出场景 Nickname ", player.Userinfo.Nickname)
}
