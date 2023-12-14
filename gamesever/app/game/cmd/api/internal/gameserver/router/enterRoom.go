package router

import (
	"fmt"
	"go-cms/app/game/cmd/api/internal/gameserver/core"
	"go-cms/app/game/cmd/api/internal/gameserver/pb"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/protobuf/proto"
)

type EnterRoomRouter struct {
	znet.BaseRouter
}

// 进入场景
func (m *EnterRoomRouter) Handle(req ziface.IRequest) {
	fmt.Println("进入场景 EnterRoomRouter")
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
	fmt.Println("进入场景 Nickname ", player.Userinfo.Nickname)
}
