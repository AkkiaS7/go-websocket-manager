package wsmgr

import (
	"github.com/AkkiaS7/go-websocket-mgr/wsmgr/iface"
	"sync"
)

/*
	连接管理模块的具体实现
*/
type ConnManager struct {
	connections map[uint64]iface.IConnection //存放所有链接ID的map
	connLock    sync.RWMutex                 //读写连接的读写锁
}

// NewConnManager 创建一个链接管理模块的实例
func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint64]iface.IConnection),
	}
}

// GetConnID

// Add 添加链接
func (cm *ConnManager) Add(conn iface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	cm.connections[conn.GetConnID()] = conn
	return
}

// Remove 删除链接
func (cm *ConnManager) Remove(conn iface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	delete(cm.connections, conn.GetConnID())
	return
}

// Get 得到链接
func (cm *ConnManager) Get(connID uint64) (iface.IConnection, error) {
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()

	if conn, ok := cm.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, nil
	}
}
