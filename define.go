/**
 * @Author: 10512203@qq.com
 * @Description:
 * @File: define
 * @Version: 1.0.0
 * @Date: 2022/4/7 10:16
 */

package easysocket

import "google.golang.org/protobuf/proto"

type HookFunc func(session ISession)

type ServerType uint32
type GateHandler func(request IRequest, message proto.Message)

const (
	NullServer ServerType = iota
	TcpServer
	WsServer
	TcpClient
)

const (
	MsgTypeSize = 4
	MsgLenSize  = 4
)
