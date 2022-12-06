/**
 * @Author: 10512203@qq.com
 * @Description:
 * @File: player_manager.go
 * @Date: 2022/9/27 9:43
 **/

package managers

import (
	"fmt"
	"sync"

	goServer "github.com/88act/go-server"
)

type Player struct {
	Session goServer.ISession
	Info    PlayerInfo
}
type PlayerInfo struct {
	Name    string `json:"name"`
	Group   string `json:"group"`
	AppKey  string `json:"appKey"`
	Ip      string `json:"ip"`
	IpLocal string `json:"ipLocal"`
	Port    int32  `json:"port"`
}

type PlayerManager struct {
	sync.RWMutex

	players map[uint32]*Player
}

func (pm *PlayerManager) Add(player *Player) {
	pm.Lock()
	defer pm.Unlock()

	if pm.players == nil {
		pm.players = make(map[uint32]*Player)
	}
	pm.players[player.Session.GetConnId()] = player

	// fmt.Println("player 新连接id GetConnId")
	// fmt.Println(player.Session.GetConnId())

	// fmt.Println("存在的player ")
	// for key, value := range pm.players {
	// 	fmt.Println(key, ":", value)
	// }
}

func (pm *PlayerManager) Remove(session goServer.ISession) {
	pm.Lock()
	defer pm.Unlock()

	delete(pm.players, session.GetConnId())
}

func (pm *PlayerManager) Get(connId uint32) *Player {
	pm.RLock()
	defer pm.RUnlock()

	return pm.players[connId]
}

func (pm *PlayerManager) GetAllPlayer() []*Player {
	pm.RLock()
	defer pm.RUnlock()
	list := []*Player{}
	for key, value := range pm.players {
		fmt.Println(key, ":", value)
		list = append(list, value)
	}
	return list
}

var PlayerMgr = PlayerManager{}
