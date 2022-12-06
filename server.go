/**
 * @Author: 10512203@qq.com
 * @Description:
 * @File: server
 * @Version: 1.0.0
 * @Date: 2022/4/7 11:14
 */

package easysocket

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type IServer interface {
	start()
	Stop()
	Serve()
	AddPreRouter(handle PreRouterHandle)
	AddRouter(msgId int32, router IRouter, v any)
	SetGateHandler(handler GateHandler)
	GetSessMgr() ISessionManager
	SetOnConnStart(hookFunc HookFunc)
	SetOnConnStop(hookFunc HookFunc)
	CallOnConnStart(session ISession)
	CallOnConnStop(session ISession)
	SendMsg(msgId int32, message proto.Message)
	SendBufferMsg(request IRequest)
}

var connId uint32 = 0

type Server struct {
	serverName    string
	ServerType    ServerType
	host          string
	port          int
	msgHandle     IMessageHandler
	sessMgr       ISessionManager
	serverSession *TCPSession
	OnConnStart   HookFunc
	OnConnStop    HookFunc
	options       Options
}

func NewServer(name string, serverType ServerType, host string, port int, opts ...Option) *Server {
	return &Server{
		serverName: name,
		ServerType: serverType,
		host:       host,
		port:       port,
		msgHandle:  NewMessageHandler(),
		sessMgr:    NewSessionManager(),
		options:    newOptions(opts...),
	}
}

func (s *Server) startTCPServer() {
	go func() {
		l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
		if err != nil {
			fmt.Println("listen tcp error:", err.Error())
			return
		}

		fmt.Println("start ", s.serverName, " success...")

		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("accept error:", err.Error())
				continue
			}

			fmt.Println("new client connect, remote addr = ", conn.RemoteAddr().String())

			connId++
			sess := NewTCPSession(s, conn, connId, s.msgHandle)

			go sess.Start()
		}
	}()
}

func (s *Server) startWsServer() {
	go func() {
		upgrade := &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			conn, err := upgrade.Upgrade(writer, request, nil)

			if err != nil {
				fmt.Println("websocket error:", err)
				return
			}

			realIp := request.Header.Get("X-Forwarded-For")

			connId++
			sess := NewWsSession(s, conn, connId, s.msgHandle, realIp)

			s.sessMgr.Add(sess)

			go sess.Start()
		})

		if s.options.keyFile == "" || s.options.certFile == "" {
			if err := http.ListenAndServe(fmt.Sprintf("%s:%d", s.host, s.port), nil); err != nil {
				fmt.Println("http listen error:", err)
				return
			}
		} else {
			if err := http.ListenAndServeTLS(fmt.Sprintf("%s:%d", s.host, s.port), s.options.certFile, s.options.keyFile, nil); err != nil {
				fmt.Println("http listen error:", err)
				return
			}
		}

		fmt.Println("websocket server is running...")
	}()
}

func (s *Server) startTCPClient() {
	for {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
		if err != nil {
			fmt.Println("connect error", err)
			// 等待重连
			time.Sleep(10 * time.Second)
			continue
		}
		connId++
		s.serverSession = NewTCPSession(s, conn, connId, s.msgHandle)
		go s.serverSession.Start()
		break
	}
}

func (s *Server) start() {
	go s.msgHandle.StartWorkerPool()

	switch s.ServerType {
	case TcpServer:
		s.startTCPServer()
	case WsServer:
		s.startWsServer()
	case TcpClient:
		s.startTCPClient()
	}

	fmt.Printf("%s at %s:%d 已启动...\n", s.serverName, s.host, s.port)
}

func (s *Server) Stop() {
	s.sessMgr.Clear()
}

func (s *Server) Serve() {
	s.start()
}

func (s *Server) AddPreRouter(handle PreRouterHandle) {
	s.msgHandle.AddPreRouter(handle)
}

func (s *Server) AddRouter(msgId int32, router IRouter, v any) {
	s.msgHandle.AddRouter(msgId, router, v)
}

func (s *Server) SetGateHandler(handler GateHandler) {
	s.msgHandle.SetGateHandler(handler)
}

func (s *Server) GetSessMgr() ISessionManager {
	return s.sessMgr
}

func (s *Server) SetOnConnStart(hookFunc HookFunc) {
	s.OnConnStart = hookFunc
}

func (s *Server) SetOnConnStop(hookFunc HookFunc) {
	s.OnConnStop = hookFunc
}

func (s *Server) CallOnConnStart(session ISession) {
	if s.OnConnStart != nil {
		s.OnConnStart(session)
	}
}

func (s *Server) CallOnConnStop(session ISession) {
	if s.OnConnStop != nil {
		s.OnConnStop(session)
	}
}

func (s *Server) SendMsg(msgId int32, message proto.Message) {
	buffer, _ := proto.Marshal(message)

	_ = s.serverSession.SendMsg(msgId, buffer)
}

func (s *Server) SendBufferMsg(request IRequest) {
	s.msgHandle.SendMsgToTaskQueue(request)
}
