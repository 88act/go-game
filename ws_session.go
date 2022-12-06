/**
 * @Author: 10512203@qq.com
 * @Description:
 * @File: ws_session
 * @Version: 1.0.0
 */

package easysocket

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/88act/go-server/demo/config"
	"github.com/gorilla/websocket"
)

type WsSession struct {
	Session
	conn *websocket.Conn
	ip   string
}

func NewWsSession(server IServer, conn *websocket.Conn, connId uint32, handler IMessageHandler, ip string) *WsSession {
	sess := &WsSession{
		Session: Session{
			server:      server,
			connId:      connId,
			msgHandle:   handler,
			msgBuffChan: make(chan []byte, 1024),
			property:    nil,
			isClosed:    false,
		},
		conn: conn,
		ip:   ip,
	}

	sess.server.GetSessMgr().Add(sess)

	return sess
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func (s *WsSession) startReader() {
	fmt.Println("Reader goroutine is running...")
	defer fmt.Println(s.RemoteAddr().String(), " conn reader exit!")
	defer s.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			fmt.Println("读。。。。。")
			_, data, err := s.conn.ReadMessage()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			if config.ConfigMgr.JsonOrProto == config.Protocol_JSON {
				data = bytes.TrimSpace(bytes.Replace(data, newline, space, -1))
				//fmt.Println("解压信息---1----")
				//fmt.Println(data)
				msg := DP.UnPack_json(data)
				if msg.GetMsgId() > 0 {
					//msg.SetData(data[DP.GetHeadLen():])

					req := &Request{
						sess: s,
						msg:  msg,
					}
					s.msgHandle.SendMsgToTaskQueue(req)
				}
			} else if config.ConfigMgr.JsonOrProto == config.Protocol_TCP {

				// // 获取消息头数据
				headData := data[:DP.GetHeadLen()]
				//解析消息头
				msg := DP.UnPack(headData)

				if msg.GetDataLen() > 0 {
					msg.SetData(data[DP.GetHeadLen():])

					req := &Request{
						sess: s,
						msg:  msg,
					}

					s.msgHandle.SendMsgToTaskQueue(req)
				}
			}
		}
	}
}

func (s *WsSession) startWriter() {
	fmt.Println("Writer goroutine is running...")
	defer fmt.Println(s.RemoteAddr().String(), " conn writer exit!")
	for {
		select {
		case data, ok := <-s.msgBuffChan:
			if ok {
				//if err := s.conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
				if err := s.conn.WriteMessage(websocket.TextMessage, data); err != nil {
					fmt.Println("send buff data error:", err, " conn writer exit")
					return
				}
			} else {
				fmt.Println("msgBuffChan is closed")
				break
			}
		case <-s.ctx.Done():
			return
		}
	}
}

func (s *WsSession) Start() {
	s.ctx, s.cancel = context.WithCancel(context.Background())

	go s.startReader()
	go s.startWriter()

	s.server.CallOnConnStart(s)

	select {
	case <-s.ctx.Done():
		s.finalizer()
		return
	}
}

func (s *WsSession) RemoteAddr() net.Addr {
	return s.conn.RemoteAddr()
}

func (s *WsSession) RemoteIP() string {
	if s.ip == "" {
		l := strings.Split(s.conn.RemoteAddr().String(), ":")
		if len(l) > 0 {
			s.ip = l[0]
		}
	}
	return s.ip
}

func (s *WsSession) SendMsg(msgId int32, data []byte) error {
	s.RLock()
	defer s.RUnlock()

	if s.isClosed {
		return errors.New("connection closed when send msg")
	}

	msg := DP.Pack(NewMessage(msgId, data))

	//return s.conn.WriteMessage(websocket.BinaryMessage, msg)
	return s.conn.WriteMessage(websocket.TextMessage, msg)
}

func (s *WsSession) finalizer() {
	s.server.CallOnConnStop(s)

	s.Lock()
	defer s.Unlock()

	if s.isClosed {
		return
	}

	fmt.Println("conn stop()...connId = ", s.connId)

	s.isClosed = true

	_ = s.conn.Close()

	s.server.GetSessMgr().Remove(s)

	close(s.msgBuffChan)
}
