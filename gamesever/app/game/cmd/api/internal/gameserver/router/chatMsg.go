package router

import (
	"fmt"
	"time"

	"go-cms/app/game/cmd/api/internal/gameserver/core"
	"go-cms/app/game/cmd/api/internal/gameserver/pb"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"google.golang.org/protobuf/proto"
)

type ChatMsgRouter struct {
	znet.BaseRouter
}

// Ping Handle
func (m *ChatMsgRouter) Handle(request ziface.IRequest) {

	//2. 得知当前的消息是从哪个玩家传递来的,从连接属性pID中获取
	pID, err := request.GetConnection().GetProperty("pID")
	if err != nil {
		fmt.Println("GetProperty pID error", err)
		request.GetConnection().Stop()
		return
	}
	player := core.WorldMgrObj.GetPlayerByPID(pID.(int64))

	//fmt.Println("接收到 ChatMsg  1")
	msg := &pb.ChatMsg{}
	err = proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		fmt.Println(" 消息解压失败 Unmarshal error ", err, " data = ", request.GetData())
		return
	}
	fmt.Println(" msg.MsgType====> ", msg.MsgType)
	fmt.Println(" msg =====> ", msg)

	// 关闭拉流api
	if msg.MsgType == 10 {

	}

	//msg.ObjId = 1 // 暂时没有用
	msg.SendTime = time.Now().Unix()
	// 一对一
	resp := &pb.ChatMsgResp{}
	if msg.ChatType == 3 {
		resp.MsgList = append(resp.MsgList, msg)
		if msg.MsgType == 11 {
			// 拨号 ，只 发给对方

		} else {
			// 一对一 两个人都发
			core.WorldMgrObj.SendChatUserId(msg.ObjId, resp)
			player.SendMsg(pb.S_ChatMsgResp, resp)
		}
	} else {
		resp.MsgList = append(resp.MsgList, msg)
		core.WorldMgrObj.SendAll(pb.S_ChatMsgResp, resp)
	}

}
