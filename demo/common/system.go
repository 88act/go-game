/**
 * @Author: 10512203@qq.com
 * @Description:
 * @File: system.go
 * @Date: 2022/9/27 10:58
 **/

package common

import (
	"fmt"

	goServer "github.com/88act/go-server"
	"github.com/88act/go-server/demo/common/ProtoMsg"
	"google.golang.org/protobuf/proto"
)

func ServiceSendMsg(request goServer.IRequest, msgId int32, message proto.Message) {
	data, err := proto.Marshal(message)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var connId = uint32(0)

	property, _ := request.GetSession().GetProperty("ConnID")

	if property != nil {
		connId = property.(uint32)
	}

	dataTransfer := &ProtoMsg.S2G_DataTransfer{
		MsgID:  msgId,
		ConnId: connId,
		Data:   data,
	}
	fmt.Println(dataTransfer)
	//_ = request.SendMsg(int32(ProtoMsg.CMD_SERVICE_S2G_DATA_TRANSFER), dataTransfer)
	//_ = request.SendMsg(int32(ProtoMsg.CMD_SERVICE_S2G_DATA_TRANSFER), )
}

func ServiceSendMsgJson(request goServer.IRequest, msgId int32, jsonStr string) {
	// data, err := proto.Marshal(message)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// var connId = uint32(0)

	// property, _ := request.GetSession().GetProperty("ConnID")

	// if property != nil {
	// 	connId = property.(uint32)
	// }

	// dataTransfer := &ProtoMsg.S2G_DataTransfer{
	// 	MsgID:    msgId,
	// 	ConnId:   connId,
	// 	DataJson: jsonStr,
	// }
	// fmt.Println(dataTransfer)
	// _ = request.SendMsgJson(int32(ProtoMsg.CMD_SERVICE_S2G_DATA_TRANSFER), jsonStr)
	_ = request.SendMsgJson(msgId, jsonStr)

}
