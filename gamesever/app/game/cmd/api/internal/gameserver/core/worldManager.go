package core

import (
	"sync"
)

// 当前游戏世界的总管理模块
type WorldManager struct {
	//AoiMgr   *AOIManager        // 当前世界地图的AOI规划管理器
	Players  map[int64]*Player     //  当前在线的玩家集合)
	pLock    sync.RWMutex          //  保护Players的互斥读写机制
	RoomList map[int64]*AOIManager //  场景的AOI管理器
}

var RoomList = []int64{1000, 1001, 1002}

var once_WorldMgrObj sync.Once = sync.Once{}
var WorldMgrObj *WorldManager

// 世界管理器
func GetWM() *WorldManager {
	once_WorldMgrObj.Do(func() {
		WorldMgrObj = &WorldManager{}
		WorldMgrObj.Players = make(map[int64]*Player)
		for _, roomId := range RoomList {
			room := NewAOIManager(roomId, AOI_MIN_X, AOI_MAX_X, AOI_CNTS_X, AOI_MIN_Y, AOI_MAX_Y, AOI_CNTS_Y)
			WorldMgrObj.RoomList[roomId] = room
		}
	})
	return WorldMgrObj
}

// 获取场景管理器
func (m *WorldManager) GetRoom(roomId int64) *AOIManager {
	return m.RoomList[roomId]
}

// 将玩家添加进玩家信息表Players
func (m *WorldManager) AddPlayer(player *Player) {
	m.pLock.Lock()
	m.Players[player.Pid] = player
	m.pLock.Unlock()
	m.GetRoom(player.RoomId).AddToGrIDByPos(player.Pid, player.X, player.Z)
}

// 从玩家信息表中移除一个玩家
func (m *WorldManager) RemovePlayerByPID(pID int64) {
	m.pLock.Lock()
	delete(m.Players, pID)
	m.pLock.Unlock()
}

// 通过玩家ID 获取对应玩家信息
func (m *WorldManager) GetPlayerByPID(pID int64) *Player {
	m.pLock.RLock()
	defer m.pLock.RUnlock()
	return m.Players[pID]
}

// 获取所有玩家的信息
func (m *WorldManager) GetAllPlayers() []*Player {
	m.pLock.RLock()
	defer m.pLock.RUnlock()

	players := make([]*Player, 0)
	for _, v := range m.Players {
		players = append(players, v)
	}
	return players
}

// 获取指定gID中的所有player信息
func (m *WorldManager) GetPlayersByGID(roomId int64, gID int) []*Player {

	pIDs := m.GetRoom(roomId).grIDs[gID].GetPlyerIDs()
	players := make([]*Player, 0, len(pIDs))
	m.pLock.RLock()
	for _, pID := range pIDs {
		players = append(players, m.Players[pID])
	}
	m.pLock.RUnlock()
	return players
}
