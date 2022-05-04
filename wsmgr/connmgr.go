package wsmgr

import (
	"github.com/AkkiaS7/go-websocket-mgr/wsmgr/utils"
	"github.com/gorilla/websocket"
	"sync"
)

/*
	连接管理模块的具体实现
*/
type ConnManager struct {
	connections map[uint64]*Connection //存放所有链接ID的map
	connLock    sync.RWMutex           //读写连接的读写锁
	mgr         *Mgr                   //连接管理模块所属的管理器
	connChan    chan *Connection       //连接管理模块的连接通道
	closeChan   chan *Connection       //连接管理模块的关闭通道
}

// NewConnManager 创建一个链接管理模块的实例
func NewConnManager(mgr *Mgr) *ConnManager {
	return &ConnManager{
		connections: make(map[uint64]*Connection),
		mgr:         mgr,
		connChan:    make(chan *Connection),
		closeChan:   make(chan *Connection),
	}
}

// GetConnID

// Add 添加链接
func (cm *ConnManager) Add(wsConn *websocket.Conn) *Connection {
	// 构造Connection实例
	conn := NewConnection(wsConn, utils.GetCid(), cm.mgr.MsgHandler, cm)
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	cm.connections[conn.GetConnID()] = conn
	cm.connChan <- conn
	return conn
}

// Remove 删除链接
func (cm *ConnManager) Remove(conn *Connection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	delete(cm.connections, conn.GetConnID())
	return
}

// Get 得到链接
func (cm *ConnManager) Get(connID uint64) (*Connection, error) {
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()

	if conn, ok := cm.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, nil
	}
}

//Start 启动链接管理模块
func (cm *ConnManager) Start() {
	for {
		select {
		case conn := <-cm.connChan:
			go conn.Start()
		case conn := <-cm.closeChan:
			cm.Remove(conn)
		}
	}
}
