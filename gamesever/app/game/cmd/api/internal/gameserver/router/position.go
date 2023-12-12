package router

import (
	"fmt"

	"go-cms/app/game/cmd/api/internal/gameserver/pb"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
	"google.golang.org/protobuf/proto"
)

type PositionZinxRouter struct {
	znet.BaseRouter
}

// HelloZinxRouter Handle
func (m *PositionZinxRouter) Handle(request ziface.IRequest) {
	msg := &pb.Position{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		fmt.Println("位置 Position Unmarshal error ", err, " data = ", request.GetData())
		return
	}

	fmt.Printf(" 位置 recv from client : msgId=%+v, data=%+v\n", request.GetMsgID(), msg)

	msg.X += 1
	msg.Y += 1
	msg.Z += 1
	msg.V += 1

	data, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println("位置  proto Marshal error = ", err, " msg = ", msg)
		return
	}

	err = request.GetConnection().SendMsg(101, data)

	if err != nil {
		zlog.Error(err)
	}
}
