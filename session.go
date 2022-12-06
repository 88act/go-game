/**
 * @Author: 10512203@qq.com
 * @Description:
 * @File: session
 * @Version: 1.0.0
 * @Date: 2022/4/7 11:07
 */

package easysocket

import (
	"context"
	"errors"
	"net"
	"strconv"
	"sync"
	"time"
)

type ISession interface {
	startReader()
	startWriter()
	Start()
	Stop()
	Context() context.Context

	RemoteAddr() net.Addr
	RemoteIP() string
	GetConnId() uint32

	SendMsg(msgId int32, data []byte) error
	SendBuffMsg(msgId int32, data []byte) error
	SendJsonMsg(msgId int32, data string) error

	SetProperty(key string, value interface{}) error
	GetProperty(key string) (interface{}, error)
	RemoveProperty(key string)
}

type Session struct {
	server      IServer
	connId      uint32
	msgHandle   IMessageHandler
	ctx         context.Context
	cancel      context.CancelFunc
	msgBuffChan chan []byte

	sync.RWMutex
	property     map[string]interface{}
	propertyLock sync.RWMutex
	isClosed     bool
}

func (s *Session) GetConnId() uint32 {
	return s.connId
}

// SetProperty 设置链接属性
func (s *Session) SetProperty(key string, value interface{}) error {
	s.propertyLock.Lock()
	defer s.propertyLock.Unlock()

	if s.property == nil {
		s.property = make(map[string]interface{})
	}

	s.property[key] = value
	return nil
}

// GetProperty 获取链接属性
func (s *Session) GetProperty(key string) (interface{}, error) {
	s.propertyLock.RLock()
	defer s.propertyLock.RUnlock()

	if v, ok := s.property[key]; ok {
		return v, nil
	}

	return nil, errors.New("property not found")
}

// RemoveProperty 移除链接属性
func (s *Session) RemoveProperty(key string) {
	s.propertyLock.Lock()
	defer s.propertyLock.Unlock()

	delete(s.property, key)
}

// SendBuffMsg 发送BuffMsg
func (s *Session) SendBuffMsg(msgId int32, data []byte) error {
	s.RLock()
	defer s.RUnlock()
	idleTimeout := time.NewTimer(5 * time.Millisecond)
	defer idleTimeout.Stop()

	if s.isClosed {
		return errors.New("connection is closed when send buff msg")
	}

	msg := DP.Pack(NewMessage(msgId, data))

	select {
	case <-idleTimeout.C:
		return errors.New("send buff msg timeout")
	case s.msgBuffChan <- msg:
		return nil
	}
}

// SendBuffMsg 发送BuffMsg
func (s *Session) SendJsonMsg(msgId int32, data string) error {
	s.RLock()
	defer s.RUnlock()
	idleTimeout := time.NewTimer(5 * time.Millisecond)
	defer idleTimeout.Stop()

	if s.isClosed {
		return errors.New("connection is closed when send buff msg")
	}
	str := strconv.Itoa(int(msgId)) + "|" + data
	//msg := DP.PackJson(NewMessage(msgId, data))
	strByte := []byte(str)
	select {
	case <-idleTimeout.C:
		return errors.New("send buff msg timeout")
	case s.msgBuffChan <- strByte:
		return nil
	}
}

func (s *Session) Stop() {
	s.cancel()
}

func (s *Session) Context() context.Context {
	return s.ctx
}
