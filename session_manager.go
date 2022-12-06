/**
 * @Author: 10512203@qq.com
 * @Description:
 * @File: session_manager
 * @Version: 1.0.0
 * @Date: 2022/4/7 11:26
 */

package easysocket

import (
	"errors"
	"fmt"
	"sync"
)

type ISessionManager interface {
	Add(session ISession)                // 添加链接
	Remove(session ISession)             // 删除链接
	Get(connId uint32) (ISession, error) // 获取链接
	Count() int                          // 获取所有链接数量
	Clear()                              // 删除并停止所有链接
}

type SessionManager struct {
	sessions map[uint32]ISession
	sessLock sync.RWMutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[uint32]ISession),
	}
}

// Add 添加链接
func (m *SessionManager) Add(session ISession) {
	m.sessLock.Lock()
	defer m.sessLock.Unlock()

	if _, ok := m.sessions[session.GetConnId()]; ok {
		fmt.Printf("session[%d] is exist\n", session.GetConnId())
	}

	m.sessions[session.GetConnId()] = session
}

// Remove 删除链接
func (m *SessionManager) Remove(session ISession) {
	m.sessLock.Lock()
	defer m.sessLock.Unlock()

	delete(m.sessions, session.GetConnId())
}

// Get 获取链接
func (m *SessionManager) Get(connId uint32) (ISession, error) {
	m.sessLock.RLock()
	defer m.sessLock.RUnlock()

	if sess, ok := m.sessions[connId]; ok {
		return sess, nil
	}

	return nil, errors.New("session not found")
}

// Count 获取链接数量
func (m *SessionManager) Count() int {
	m.sessLock.RLock()
	defer m.sessLock.RUnlock()

	return len(m.sessions)
}

// Clear 清空所有链接
func (m *SessionManager) Clear() {
	m.sessLock.Lock()
	defer m.sessLock.Unlock()

	for connId, sess := range m.sessions {
		sess.Stop()
		delete(m.sessions, connId)
	}
}
