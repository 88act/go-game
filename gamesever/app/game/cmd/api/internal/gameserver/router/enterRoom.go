package router

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

type EnterRoomRouter struct {
	znet.BaseRouter
}

// 进入房间
func (m *EnterRoomRouter) Handle(request ziface.IRequest) {
	// pID, err := request.GetConnection().GetProperty("pID")
	// if err != nil {
	// 	fmt.Println("GetProperty pID error", err)
	// 	request.GetConnection().Stop()
	// 	return
	// }
	// msg := &pb.EnterRoom{}
	// err = proto.Unmarshal(request.GetData(), msg)
	// if err != nil {
	// 	fmt.Println(" 消息解压失败  error ", err, " data = ", request.GetData())
	// 	return
	// }
	// fmt.Printf("接收到 EnterRoom  UserId =%d ，roomid =%d ", msg.UserId, msg.RoomId)
	// //3. 根据pID得到player对象
	// player := core.WorldMgrObj.GetPlayerByPID(pID.(int64))
	// player.Userinfo.RoomId = msg.RoomId
	// resp := &pb.EnterRoomResp{}
	// resp.RoomId = msg.RoomId
	// resp.Type = 1
	// resp.Name = "第一会议室"
	// resp.Password = "123456"
	// resp.MaxUser = 20
	// resp.Image = ""

	// resp.UserList = core.WorldMgrObj.GetRoomUsers(msg.RoomId, player.Userinfo.CuId)
	// // 自己收到所有人
	// player.SendMsg(pb.S_EnterRoomResp, resp)
	// // 其他人收到当前人
	// //TODO: 修改为 每次都推送全部用户  20231024
	// //resp.UserList = []*pb.UserInfo{}
	// //resp.UserList = append(resp.UserList, player.Userinfo)
	// core.WorldMgrObj.SendRoom(pb.S_EnterRoomResp, resp, msg.RoomId, player.Pid)

	// if err != nil {
	// 	zlog.Error(err)
	// }
}
