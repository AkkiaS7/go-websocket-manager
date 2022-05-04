package wsmgr

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

/*
	Websocket连接的抽象层
*/
type IConnection interface {
	// Start 启动连接，让当前连接开始工作
	Start()
	// Stop 停止连接，结束当前连接的工作
	Stop()
	// GetConnection 获取当前连接绑定的websocket连接
	GetConnection() *websocket.Conn
	// GetConnID 获取当前连接的ID
	GetConnID() uint64
	// SendMsg 发送数据给客户端
	SendMsg(data []byte) error
	// SetProperty 设置连接的属性
	SetProperty(key string, value interface{})
	// GetProperty 获取连接的属性
	GetProperty(key string) (interface{}, error)
	// RemoveProperty 移除连接的属性
	RemoveProperty(key string)
}

type Connection struct {
	Worker     IWorker         //当前链接所属的worker
	Conn       *websocket.Conn //当前链接的Websocket连接
	ConnID     uint64          //链接的ID
	CloseChan  chan bool       //当前链接的关闭通知管道
	MsgHandler IMsgHandler     //当前链接的消息管理器

	isClosed     bool                   //当前链接的关闭状态
	msgChan      chan []byte            //当前链接的消息管道
	property     map[string]interface{} //当前链接的额外属性
	propertyLock sync.RWMutex
}

// Start 启动连接，让当前连接开始工作
func (c *Connection) Start() {
	//TODO 读写协程
}

// Stop 停止连接，结束当前连接的工作
func (c *Connection) Stop() {
	//TODO 关闭连接
}

// GetConnection 获取当前连接绑定的websocket连接
func (c *Connection) GetConnection() *websocket.Conn {
	return c.Conn
}

// GetConnID 获取当前连接的ID
func (c *Connection) GetConnID() uint64 {
	return c.ConnID
}

// SendMsg 发送数据给客户端
func (c *Connection) SendMsg(data []byte) error {
	//TODO 发送数据
	return nil
}

// SetProperty 设置连接的属性
func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	c.property[key] = value
}

// GetProperty 获取连接的属性
func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	if value, ok := c.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("no property found")
	}
}

// RemoveProperty 移除连接的属性
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	delete(c.property, key)
}

//StartReader 启动读协程
func (c *Connection) StartReader() {
	log.Println("Reader Goroutine is running for connID:", c.ConnID, ", remote address:", c.Conn.RemoteAddr().String())
	defer log.Println("remote address:", c.Conn.RemoteAddr().String(), "connID:", c.ConnID, "is closed")
	defer c.Stop()
	for {
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("connID:", c.ConnID, "read error:", err)
			return
		}
		msg := &Msg{}
		msg.Unmarshal(data)
		req := &Request{
			Conn: c,
			Msg:  msg,
		}
		c.MsgHandler.SendMsgToTaskQueue(req)
	}
}
