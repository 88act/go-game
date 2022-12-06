/**
 * @Author: 10512203@qq.com
 * @Description:
 * @File: ping_router.go
 * @Date: 2022/9/27 10:47
 **/

package routers

import (
	"encoding/json"
	"fmt"

	goServer "github.com/88act/go-server"
	"github.com/88act/go-server/demo/common"
	"github.com/88act/go-server/demo/common/ProtoMsg"
	"github.com/88act/go-server/demo/config"
	"github.com/88act/go-server/demo/server/managers"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/proto"
)

type DevInfoRouter struct {
	goServer.BaseRouter
}

// func (r *DevInfoRouter) Handle(request goServer.IRequest, message proto.Message) {
func (r *DevInfoRouter) Handle(request goServer.IRequest, message proto.Message) {
	playerInfo := new(managers.PlayerInfo)
	if config.ConfigMgr.JsonOrProto == config.Protocol_JSON {
		err := json.Unmarshal([]byte(request.GetDataJson()), playerInfo)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		msg := message.(*ProtoMsg.C2S_DevInfo)
		_ = copier.Copy(&playerInfo, msg)
	}

	s2cDevInfo := &ProtoMsg.S2C_DevInfo{
		Status: 1,
		Msg:    "正常。。。",
	}
	s2cDevInfo.DevList = []*ProtoMsg.C2S_DevInfo{}
	player := managers.PlayerMgr.Get(request.GetSession().GetConnId())
	player.Info = *playerInfo
	//fmt.Println("Handle player.Info ====")
	list := managers.PlayerMgr.GetAllPlayer()
	//fmt.Println(list)
	fmt.Println("获取在线用户列表--1----")
	for _, oplayer := range list {
		devInfo := ProtoMsg.C2S_DevInfo{}
		_ = copier.Copy(&devInfo, oplayer.Info)
		devInfo.ConnId = int32(oplayer.Session.GetConnId())
		s2cDevInfo.DevList = append(s2cDevInfo.DevList, &devInfo)
		// 给其他人发送上线通知
		if oplayer.Session.GetConnId() != request.GetSession().GetConnId() {
			devInfo2 := ProtoMsg.C2S_DevInfo{}
			_ = copier.Copy(&devInfo2, playerInfo)
			devInfo2.ConnId = int32(request.GetSession().GetConnId())
			s2cOnLine := &ProtoMsg.S2C_OnLine{
				Status: 1,
				Msg:    "上线",
			}
			s2cOnLine.DevList = []*ProtoMsg.C2S_DevInfo{}
			s2cOnLine.DevList = append(s2cOnLine.DevList, &devInfo2)
			s2cOnLineStr, err := json.Marshal(s2cOnLine)
			if err != nil {
				fmt.Println("Umarshal failed: s2cOnLine ", err)
				//return
			}
			oplayer.Session.SendJsonMsg(int32(ProtoMsg.CMD_DEV_S_OnLine), string(s2cOnLineStr))
		}
	}

	fmt.Println("发送回复消息 。。。。")
	fmt.Println(s2cDevInfo)
	str, err := json.Marshal(s2cDevInfo)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		//return
	}
	common.ServiceSendMsgJson(request, int32(ProtoMsg.CMD_DEV_S_DevInfo), string(str))
}
