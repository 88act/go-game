/**
 * @Author: 10512203@qq.com
 * @Description:
 * @File: main
 * @Version: 1.0.0
 */

package main

import (
	"encoding/json"
	"fmt"

	goServer "github.com/88act/go-server"
	"github.com/88act/go-server/demo/common/ProtoMsg"
	"github.com/88act/go-server/demo/config"

	"github.com/88act/go-server/demo/server/managers"
	"github.com/88act/go-server/demo/server/routers"
	"github.com/jinzhu/copier"
)

var server goServer.IServer

func main() {
	server = goServer.NewServer("ChatServer", goServer.WsServer, "0.0.0.0", config.ConfigMgr.Port)
	server.SetOnConnStart(playerStart)
	server.SetOnConnStop(playerStop)
	server.AddRouter(int32(ProtoMsg.CMD_DEV_C_DevInfo), &routers.DevInfoRouter{}, ProtoMsg.C2S_DevInfo{})
	server.AddRouter(int32(ProtoMsg.CMD_DEV_C_DevPing), &routers.PingRouter{}, ProtoMsg.C2S_DevPing{})

	server.Serve()

	fmt.Println("websocket 启动 Port=", config.ConfigMgr.Port)

	select {}
}

func playerStart(session goServer.ISession) {
	fmt.Println("新玩家连接: ", session.RemoteIP())
	managers.PlayerMgr.Add(&managers.Player{Session: session})
}

func playerStop(session goServer.ISession) {
	fmt.Println("玩家断开连接")
	connId := session.GetConnId()
	ownPlayer := managers.PlayerMgr.Get(connId)
	playerInfo := ownPlayer.Info

	managers.PlayerMgr.Remove(session)
	go sendToOther(connId, playerInfo)
}

// 通知其他玩家**掉线
func sendToOther(connId uint32, info managers.PlayerInfo) {
	fmt.Println("玩家断开 群发通知。。。")
	devInfo2 := ProtoMsg.C2S_DevInfo{}
	_ = copier.Copy(&devInfo2, info)
	devInfo2.ConnId = int32(connId)
	s2cOnLine := &ProtoMsg.S2C_OnLine{
		Status: 2,
		Msg:    "下线",
	}
	s2cOnLine.DevList = []*ProtoMsg.C2S_DevInfo{}
	s2cOnLine.DevList = append(s2cOnLine.DevList, &devInfo2)
	s2cOnLineStr, err := json.Marshal(s2cOnLine)
	if err != nil {
		fmt.Println("Umarshal failed: s2cOnLine ", err)
		//return
	}
	list := managers.PlayerMgr.GetAllPlayer()
	for _, player := range list {
		if player.Session.GetConnId() != connId {
			player.Session.SendJsonMsg(int32(ProtoMsg.CMD_DEV_S_OnLine), string(s2cOnLineStr))
		}
	}
}
