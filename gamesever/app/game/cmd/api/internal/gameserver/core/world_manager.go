package core

import (
	"fmt"
	"sync"

	"go-cms/app/game/cmd/api/internal/gameserver/pb"

	"google.golang.org/protobuf/proto"
)

/*
当前游戏世界的总管理模块
*/
type WorldManager struct {
	AoiMgr  *AOIManager       //当前世界地图的AOI规划管理器
	Players map[int64]*Player //当前在线的玩家集合
	pLock   sync.RWMutex      //保护Players的互斥读写机制
}

// 提供一个对外的世界管理模块句柄
var WorldMgrObj *WorldManager

// 提供WorldManager 初始化方法
func init() {
	WorldMgrObj = &WorldManager{
		Players: make(map[int64]*Player),
		AoiMgr:  NewAOIManager(AOI_MIN_X, AOI_MAX_X, AOI_CNTS_X, AOI_MIN_Y, AOI_MAX_Y, AOI_CNTS_Y),
	}
}

// 提供添加一个玩家的的功能，将玩家添加进玩家信息表Players
func (wm *WorldManager) AddPlayer(player *Player) {
	//将player添加到 世界管理器中
	wm.pLock.Lock()
	wm.Players[player.PID] = player
	wm.pLock.Unlock()

	//将player 添加到AOI网络规划中
	wm.AoiMgr.AddToGrIDByPos(int(player.PID), player.X, player.Z)
}

// 从玩家信息表中移除一个玩家
func (wm *WorldManager) RemovePlayerByPID(pID int64) {
	wm.pLock.Lock()
	delete(wm.Players, pID)
	wm.pLock.Unlock()
}

// 获取对应玩家信息by id
func (wm *WorldManager) GetPlayerByPID(pID int64) *Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()

	return wm.Players[pID]
}

// 获取对应玩家信息by userid
func (wm *WorldManager) GetPlayerByUID(userId int64) *Player {
	var p *Player
	if userId <= 0 {
		return p
	}

	for _, v := range wm.Players {
		if v.Userinfo.UserId == userId {
			p = v
			break
		}
	}
	return p
}

// 剔除相同的用户下线
func (wm *WorldManager) KillSameUser(userId int64, pId int64) {
	//wm.pLock.RLock()
	//defer wm.pLock.RUnlock()
	fmt.Printf(" 剔除相同的用户下线==================== userId =%d pId = %d", userId, pId)
	fmt.Println("")
	for _, v := range wm.Players {
		if v.Userinfo.UserId == userId && v.PID != pId {
			fmt.Println("重名用户 v.Userinfo.UserId =", v.Userinfo.UserId)
			fmt.Println("重名用户 v.PID=", v.PID)
			// resp := &pb.LogoutResp{}
			// resp.UserId = v.Userinfo.UserId
			// resp.Type = 3 // 重复登录下线
			// fmt.Println("重名用户 resp=", resp)
			// v.SendMsg(pb.S_Logout, resp)
			//delete(wm.Players, v.PID)
			//v.LostConnection(v.PID, v.Userinfo.UserId, 3)
			fmt.Println("v.Conn.GetConnection().Close()=====")
			//v.Conn.GetConnection().Close()
			v.Conn.Stop()
			//4 世界管理器将当前玩家从AOI中摘除
			//WorldMgrObj.AoiMgr.RemoveFromGrIDByPos(int(v.PID), v.X, v.Z)
			//WorldMgrObj.RemovePlayerByPID(v.PID)
		} else {
			fmt.Println(" 没有 v.Userinfo.UserId =", v.Userinfo.UserId)
			fmt.Println(" v.PID =", v.PID)
		}
	}
}

// 获取所有玩家 包括自己
func (wm *WorldManager) GetAllPlayers() []*Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()

	//创建返回的player集合切片
	players := make([]*Player, 0)

	//添加切片
	for _, v := range wm.Players {
		players = append(players, v)
	}
	//返回
	return players
}

// 获取Room玩家 包括自己
func (wm *WorldManager) GetRoomPlayers(roomId int64, cuId int64) []*Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()

	//创建返回的player集合切片
	players := make([]*Player, 0)
	//添加切片
	for _, v := range wm.Players {
		if v.Userinfo.CuId == cuId && v.RoomId == roomId {
			players = append(players, v)
		}
	}
	//返回
	return players
}

// 获取Room玩家 包括自己
func (wm *WorldManager) GetRoomUsers(roomId int64, cuId int64) []*pb.UserInfo {
	players := wm.GetRoomPlayers(roomId, cuId)
	userInfoList := make([]*pb.UserInfo, 0)
	//添加切片
	for _, v := range players {
		userInfoList = append(userInfoList, v.Userinfo)
	}
	//返回
	return userInfoList
}

// 获取指定九宫格gID中的所有player信息
func (wm *WorldManager) GetPlayersByGID(gID int) []*Player {
	//通过gID获取 对应 格子中的所有pID
	pIDs := wm.AoiMgr.grIDs[gID].GetPlyerIDs()

	//通过pID找到对应的player对象
	players := make([]*Player, 0, len(pIDs))
	wm.pLock.RLock()
	for _, pID := range pIDs {
		players = append(players, wm.Players[int64(pID)])
	}
	wm.pLock.RUnlock()

	return players
}

// 广播消息 包括自己
func (wm *WorldManager) SendAll(msgId uint32, data proto.Message) {
	//fmt.Printf("before Marshal data = %+v\n", data)
	//将proto Message结构体序列化
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("marshal msg err: ", err)
		return
	}
	//2. 得到当前世界所有的在线玩家
	players := wm.GetAllPlayers()
	//fmt.Println(players)
	//3. 向所有的玩家发送MsgID:200消息
	for _, player := range players {
		//fmt.Println("SendAll msgId ,pid connID = ", msgId, player.PID, player.Conn.GetConnID())
		player.Conn.SendMsg(msgId, msg)
	}
}

// 广播Room包括自己
func (wm *WorldManager) SendRoomAll(msgId uint32, data proto.Message, roomId int64, cuId int64) {
	//fmt.Printf("before Marshal data = %+v\n", data)
	//将proto Message结构体序列化
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("marshal msg err: ", err)
		return
	}
	players := wm.GetRoomPlayers(roomId, cuId)
	for _, player := range players {
		//fmt.Println("SendAll msgId ,pid connID = ", msgId, player.PID, player.Conn.GetConnID())
		player.Conn.SendMsg(msgId, msg)
	}
}

// 广播Room,排除一个人
func (wm *WorldManager) SendRoom(msgId uint32, data proto.Message, roomId int64, noRecId int64, cuId int64) {
	fmt.Println(" 广播Room,排除一个人 ", msgId, roomId, noRecId)
	//将proto Message结构体序列化

	players := wm.GetRoomPlayers(roomId, cuId)
	for _, player := range players {
		if player.PID != noRecId {
			fmt.Println(" 广播Room ,pid connID = ", player.PID, player.Userinfo.UserId)
			player.SendMsg(msgId, data)
		}
	}
}

func (wm *WorldManager) SendChatUserId(userId int64, resp *pb.ChatMsgResp) {
	// wm.pLock.RLock()
	// defer wm.pLock.RUnlock()
	for _, v := range wm.Players {
		if v.Userinfo.UserId == userId {
			// resp := &pb.LogoutResp{}
			// resp.UserId = v.Userinfo.UserId
			// resp.Type = 3 // 重复登录下线
			//fmt.Println("发送一对一 userId= ", userId)
			v.SendMsg(pb.S_ChatMsgResp, resp)
			break
		}
	}
}
