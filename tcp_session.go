/**
 * @Author: 10512203@qq.com
 * @Description:
 * @File: tcp_session
 * @Version: 1.0.0
 * @Date: 2022/4/7 15:37
 */

package easysocket

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
)

type TCPSession struct {
	Session
	conn net.Conn
}

func NewTCPSession(server IServer, conn net.Conn, connId uint32, handler IMessageHandler) *TCPSession {
	sess := &TCPSession{
		Session: Session{
			server:      server,
			connId:      connId,
			msgHandle:   handler,
			msgBuffChan: make(chan []byte, 1024),
			RWMutex:     sync.RWMutex{},
			property:    nil,
			isClosed:    false,
		},
		conn: conn,
	}

	sess.server.GetSessMgr().Add(sess)

	return sess
}

func (s *TCPSession) GetConn() net.Conn {
	return s.conn
}

func (s *TCPSession) startReader() {
	fmt.Println("Reader goroutine is running...")
	defer fmt.Println(s.RemoteAddr().String(), " conn reader exit!")
	defer s.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			headData := make([]byte, DP.GetHeadLen())
			if _, err := io.ReadFull(s.conn, headData); err != nil {
				fmt.Println("read msg head error: ", err)
				return
			}

			msg := DP.UnPack(headData)

			if msg.GetDataLen() > 0 {
				data := make([]byte, msg.GetDataLen())
				if _, err := io.ReadFull(s.conn, data); err != nil {
					fmt.Println("read msg data error")
					return
				}

				msg.SetData(data)

				req := &Request{
					sess: s,
					msg:  msg,
				}

				s.msgHandle.SendMsgToTaskQueue(req)
			}
		}
	}
}

func (s *TCPSession) startWriter() {
	fmt.Println("Writer goroutine is running...")
	defer fmt.Println(s.RemoteAddr().String(), " conn writer exit!")

	for {
		select {
		case data, ok := <-s.msgBuffChan:
			if ok {
				if _, err := s.conn.Write(data); err != nil {
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

func (s *TCPSession) Start() {
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

func (s *TCPSession) RemoteAddr() net.Addr {
	return s.conn.RemoteAddr()
}

func (s *TCPSession) RemoteIP() string {
	return s.conn.RemoteAddr().String()
}

// SendMsg 直接将Message数据发送给远程TCP客户端
func (s *TCPSession) SendMsg(msgId int32, data []byte) error {
	s.RLock()
	defer s.RUnlock()

	if s.isClosed {
		return errors.New("connection closed when send msg")
	}

	msg := DP.Pack(NewMessage(msgId, data))

	_, err := s.conn.Write(msg)
	return err
}

func (s *TCPSession) finalizer() {
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
