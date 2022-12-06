/**
 * @Author: 10512203@qq.com
 * @Description:
 * @File: ping_router.go
 **/

package routers

import (
	"encoding/json"
	"fmt"

	goServer "github.com/88act/go-server"
	"github.com/88act/go-server/demo/common"
	"github.com/88act/go-server/demo/common/ProtoMsg"
	"google.golang.org/protobuf/proto"
)

type PingRouter struct {
	goServer.BaseRouter
}

func (r *PingRouter) Handle(request goServer.IRequest, message proto.Message) {
	msg := message.(*ProtoMsg.C2S_DevPing)

	fmt.Println("C2S_DevPing  Handle 。。。。")
	fmt.Println(msg)

	s2cDevPing := &ProtoMsg.S2C_DevPing{
		Status: 1,
		Msg:    "正常。。。pone " + msg.GetMsg(),
	}
	fmt.Println("发送回复消息 。。。。")
	fmt.Println(s2cDevPing)
	str, err := json.Marshal(s2cDevPing)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		//return
	}
	common.ServiceSendMsgJson(request, int32(ProtoMsg.CMD_DEV_S_DevPing), string(str))
}
