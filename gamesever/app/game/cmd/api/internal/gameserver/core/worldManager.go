package core

import (
	"context"
	"errors"
	"go-cms/app/game/cmd/api/internal/svc"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logc"
)

// 当前游戏世界的总管理模块
type WorldManager struct {
	Players  map[int64]*Player     //  当前在线的玩家集合)
	pLock    sync.RWMutex          //  保护Players的互斥读写机制
	RoomList map[int64]*AOIManager //  场景的AOI管理器
	SvcCtx   *svc.ServiceContext
	//ConnList map[uint64]ziface.IConnection //  当前在线的玩家集合
	//AoiMgr   *AOIManager        // 当前世界地图的AOI规划管理器
}

var once_WorldMgrObj sync.Once = sync.Once{}
var WorldMgrObj *WorldManager

// 世界管理器
func GetWM() *WorldManager {
	once_WorldMgrObj.Do(func() {
		WorldMgrObj = &WorldManager{}
		WorldMgrObj.Players = make(map[int64]*Player)
		WorldMgrObj.RoomList = make(map[int64]*AOIManager)
		var RoomList = []int64{1000, 1001, 1002}
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

// 移除一个玩家
func (m *WorldManager) RemovePlayerByPID(pID int64) {
	m.pLock.Lock()
	delete(m.Players, pID)
	m.pLock.Unlock()
}

// 通过ID获取玩家
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

// 给某个玩家发消息
func (m *WorldManager) SendOne(ctx context.Context, msgId uint32, data proto.Message, pid int64) error {
	p := m.GetPlayerByPID(pid)
	if p != nil {
		err := p.SendMsgObj(msgId, data)
		return err
	} else {
		return errors.New("用户不存在")
	}

}

// 世界广播
func (m *WorldManager) SendWorld(ctx context.Context, msgId uint32, data proto.Message) error {
	msgByte, err := proto.Marshal(data)
	if err != nil {
		return errors.New("结构体序列化错误" + err.Error())
	}
	for _, p := range m.Players {
		err = p.SendMsg(msgId, msgByte)
		if err != nil {
			logc.Error(ctx, "SendWorld发送消息出错，pid=%d,msgId=%d ", p.Pid, msgId)
		}
	}
	return err
}

// 房间广播
func (m *WorldManager) SendRoom(ctx context.Context, msgId uint32, data proto.Message, roomId int64) error {
	msgByte, err := proto.Marshal(data)
	if err != nil {
		return errors.New("结构体序列化错误" + err.Error())
	}
	for _, p := range m.Players {
		if p.RoomId == roomId {
			err = p.SendMsg(msgId, msgByte)
			if err != nil {
				logc.Error(ctx, "SendRoom发送消息出错，roomid =%d, pid=%d,msgId=%d ", roomId, p.Pid, msgId)
			}
		}
	}
	return err
}

// 可视图内广播
func (m *WorldManager) SendView(ctx context.Context, msgId uint32, data proto.Message, roomId int64, gID int) error {
	viewPlayers := m.GetPlayersByGID(roomId, gID)
	msgByte, err := proto.Marshal(data)
	if err != nil {
		return errors.New("结构体序列化错误" + err.Error())
	}
	for _, p := range viewPlayers {
		err = p.SendMsg(msgId, msgByte)
		if err != nil {
			logc.Error(ctx, "SendView发送消息出错 roomid =%d, pid=%d,msgId=%d,gID=%d ", roomId, p.Pid, msgId, gID)
		}
	}
	return err
}
