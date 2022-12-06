/**
 * @Author: 10512203@qq.com
 * @Description:
 * @File: server_manager.go
 * @Date: 2022/9/26 16:24
 **/

package managers

import (
	"sort"
	"sync"

	goServer "github.com/88act/go-server"
)

type ServerItem struct {
	Session goServer.ISession
	Count   int
}

type ServerManager struct {
	sync.RWMutex

	servers map[uint32]*ServerItem
}

func (m *ServerManager) Add(item *ServerItem) {
	m.Lock()
	defer m.Unlock()

	if m.servers == nil {
		m.servers = make(map[uint32]*ServerItem)
	}

	m.servers[item.Session.GetConnId()] = item
}

func (m *ServerManager) Get(id uint32) *ServerItem {
	m.RLock()
	defer m.RUnlock()

	if v, ok := m.servers[id]; ok {
		return v
	}

	// 新的连接

	if len(m.servers) == 0 {
		return nil
	}

	list := make([]*ServerItem, 0)

	for _, v := range m.servers {
		list = append(list, v)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Count < list[j].Count
	})

	serverItem := list[0]

	serverItem.Count++

	return serverItem
}

func (m *ServerManager) Remove(session goServer.ISession) {
	delete(m.servers, session.GetConnId())
}

var ServerMgr = &ServerManager{}
