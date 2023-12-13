package router

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

type QuitRoomRouter struct {
	znet.BaseRouter
}

// quit 房间
func (m *QuitRoomRouter) Handle(request ziface.IRequest) {
	// pID, err := request.GetConnection().GetProperty("pID")
	// if err != nil {
	// 	fmt.Println("GetProperty pID error", err)
	// 	request.GetConnection().Stop()
	// 	return
	// }
	// msg := &pb.QuitRoom{}
	// err = proto.Unmarshal(request.GetData(), msg)
	// if err != nil {
	// 	fmt.Println(" 消息解压失败  error ", err, " data = ", request.GetData())
	// 	return
	// }
	// fmt.Printf("接收到 QuitRoom  UserId =%d ，roomid =%d ", msg.UserId, msg.RoomId)
	// //3. 根据pID得到player对象
	// player := core.WorldMgrObj.GetPlayerByPID(pID.(int64))
	// player.Userinfo.RoomId = 0
	// resp := &pb.QuitRoomResp{}
	// resp.UserInfo = player.Userinfo
	// core.WorldMgrObj.SendRoomAll(pb.S_QuitRoomResp, resp, msg.RoomId, player.Userinfo.CuId)
	// if err != nil {
	// 	zlog.Error(err)
	// }
}
