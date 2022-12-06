package config

var (
	Protocol_JSON     int = 1
	Protocol_PROTOBUF int = 2
	Protocol_WS       int = 1
	Protocol_TCP      int = 2
)

type Config struct {
	//端口
	Port int
	// 1= Json 2 = Proto
	JsonOrProto int
	// 1= websocket 2 = tcp
	WsOrTCP int
}

func (m *Config) Init() {
	// m.Port = 19001
	// m.JsonOrProto = 1
}

var ConfigMgr = Config{Port: 8000, JsonOrProto: Protocol_JSON}
