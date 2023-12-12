package config

import (
	"go-cms/common"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	DB common.DbConf

	Redis    redis.RedisConf
	LocalRes struct {
		BaseUrl  string
		BasePath string
		Path     string
		PathUser string
	}
	LogConf logc.LogConf

	GameConf GameConf
}

type GameConf struct {
	/*
		Server
	*/
	Host    string // The IP address of the current server. (当前服务器主机IP)
	TCPPort int    // The port number on which the server listens for TCP connections.(当前服务器主机监听端口号)
	WsPort  int    // The port number on which the server listens for WebSocket connections.(当前服务器主机websocket监听端口)
	Name    string // The name of the current server.(当前服务器名称)
	//MaxPacketSize    uint32 // The maximum size of the packets that can be sent or received.(读写数据包的最大值)
	MaxConn        int    // The maximum number of connections that the server can handle.(当前服务器主机允许的最大链接个数)
	WorkerPoolSize uint32 // The number of worker pools in the business logic.(业务工作Worker池的数量)
	// MaxWorkerTaskLen uint32 // The maximum number of tasks that a worker pool can handle.(业务工作Worker对应负责的任务队列最大任务存储数量)
	// WorkerMode       string // The way to assign workers to connections.(为链接分配worker的方式)
	// MaxMsgChanLen    uint32 // The maximum length of the send buffer message queue.(SendBuffMsg发送消息的缓冲最大长度)
	// IOReadBuffSize   uint32 // The maximum size of the read buffer for each IO operation.(每次IO最大的读取长度)

	//The server mode, which can be "tcp" or "websocket". If it is empty, both modes are enabled.
	//"tcp":tcp监听, "websocket":websocket 监听 为空时同时开启
	Mode string

	// A boolean value that indicates whether the new or old version of the router is used. The default value is false.
	// 路由模式 false为旧版本路由，true为启用新版本的路由 默认使用旧版本
	// RouterSlicesMode bool

	/*
		logger
	*/
	LogDir string // The directory where log files are stored. The default value is "./log".(日志所在文件夹 默认"./log")

	// The name of the log file. If it is empty, the log information will be printed to stderr.
	// (日志文件名称   默认""  --如果没有设置日志文件，打印信息将打印至stderr)
	LogFile string

	// LogSaveDays int   // 日志最大保留天数
	// LogFileSize int64 // 日志单个日志最大容量 默认 64MB,单位：字节，记得一定要换算成MB（1024 * 1024）
	// LogCons     bool  // 日志标准输出  默认 false

	// The level of log isolation. The values can be 0 (all open), 1 (debug off), 2 (debug/info off), 3 (debug/info/warn off), and so on.
	// 日志隔离级别  -- 0：全开 1：关debug 2：关debug/info 3：关debug/info/warn ...
	LogIsolationLevel int

	/*
		Keepalive
	*/
	// The maximum interval for heartbeat detection in seconds.
	// 最长心跳检测间隔时间(单位：秒),超过改时间间隔，则认为超时，从配置文件读取
	HeartbeatMax int

	// /*
	// 	TLS
	// */
	// CertFile       string // The name of the certificate file. If it is empty, TLS encryption is not enabled.(证书文件名称 默认"")
	// PrivateKeyFile string // The name of the private key file. If it is empty, TLS encryption is not enabled.(私钥文件名称 默认"" --如果没有设置证书和私钥文件，则不启用TLS加密)
}
