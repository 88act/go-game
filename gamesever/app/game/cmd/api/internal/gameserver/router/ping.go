package router

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
)

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// Ping Handle
func (m *PingRouter) Handle(req ziface.IRequest) {
	//fmt.Println("接收到 pong ", req.GetMsgID())
	zlog.Infof("Recv Heartbeat from %s, MsgID = %+v, Data = %s", req.GetConnection().RemoteAddr(), req.GetMsgID(), string(req.GetData()))

	// msg := &pb.Ping{}
	// err := proto.Unmarshal(req.GetData(), msg)
	// if err != nil {
	// 	//fmt.Println("ping  Unmarshal error ", err, " data = ", request.GetData())
	// 	return
	// }
	// //fmt.Println("接收到ping  2 sn=", msg.Sn)
	// pingResp := &pb.PingResp{}
	// pingResp.NowTime = time.Now().Format("2016-01-02 15:04:05")
	// pingResp.Sn = msg.Sn

	// data, err := proto.Marshal(pingResp)
	// if err != nil {
	// 	fmt.Println("对象转字节错误 ,err = ", err, " msg = ", msg)
	// 	return
	// }

	//err = request.GetConnection().SendMsg(pb.S_PingResp, data)

	// if err != nil {
	// 	zlog.Error(err)
	// }

}
