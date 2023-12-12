package pb

const (
	C_Ping           uint32 = 1000
	S_PingResp       uint32 = 2000
	C_Login          uint32 = 100000
	S_UserInfo       uint32 = 200000
	C_EnterRoom      uint32 = 100001
	S_EnterRoomResp  uint32 = 200001
	C_ChatMsg        uint32 = 100002
	S_ChatMsgResp    uint32 = 200002
	C_CreateRoom     uint32 = 100003
	S_CreateRoomResp uint32 = 200003
	C_RoomList       uint32 = 100004
	S_RoomListResp   uint32 = 200004
	C_QuitRoom       uint32 = 100005
	S_QuitRoomResp   uint32 = 200005

	C_AddFriend     uint32 = 100100
	S_AddFriendResp uint32 = 200100
	C_MyFriend      uint32 = 100101
	S_MyFriendResp  uint32 = 200101

	C_Logout uint32 = 100900
	S_Logout uint32 = 200900
)
