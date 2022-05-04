package wsmgr

import (
	"errors"
	"github.com/AkkiaS7/go-websocket-mgr/wsmgr/conf"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Connection struct {
	Conn         *websocket.Conn        //当前链接的Websocket连接
	ConnID       uint64                 //链接的ID
	CloseChan    chan bool              //当前链接的关闭通知管道
	MsgHandler   *MsgHandler            //当前链接的消息管理器
	ConnMgr      *ConnManager           //当前链接的管理器
	isClosed     bool                   //当前链接的关闭状态
	msgChan      chan *[]byte           //当前链接的消息管道
	property     map[string]interface{} //当前链接的额外属性
	propertyLock sync.RWMutex
}

func NewConnection(conn *websocket.Conn, connID uint64, msgHandler *MsgHandler, connMgr *ConnManager) *Connection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connID,
		CloseChan:  make(chan bool, 1),
		MsgHandler: msgHandler,
		ConnMgr:    connMgr,
		isClosed:   false,
		msgChan:    make(chan *[]byte, conf.MaxMsgSize),
		property:   make(map[string]interface{}),
	}
	return c
}

// Start 启动连接，让当前连接开始工作
func (c *Connection) Start() {
	log.Println("ConnID:", c.ConnID, " start")
	go c.StartReader()
	go c.StartWriter()
}

// Stop 停止连接，结束当前连接的工作
func (c *Connection) Stop() {
	log.Println("ConnID:", c.ConnID, " stop")
	if c.isClosed {
		return
	}
	c.isClosed = true
	err := c.Conn.Close()
	if err != nil {
		return
	}
	c.CloseChan <- true
	// 从连接管理器中删除当前连接
	c.ConnMgr.closeChan <- c
	// 关闭当前连接的管道
	close(c.CloseChan)
	close(c.msgChan)
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
func (c *Connection) SendMsg(msg *Message) error {
	if c.isClosed {
		return errors.New("connection is closed")
	}
	data := msg.Pack()
	c.msgChan <- &data
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
			log.Println("read message error:", err)
			continue
		}
		req := &Request{
			Conn:   c,
			rawMsg: &data,
		}
		c.MsgHandler.SendToTaskQueue(req)
	}
}

// StartWriter 启动写协程
func (c *Connection) StartWriter() {
	log.Println("Writer Goroutine is running for connID:", c.ConnID, ", remote address:", c.Conn.RemoteAddr().String())
	defer log.Println("remote address:", c.Conn.RemoteAddr().String(), "connID:", c.ConnID, "is closed")
	defer c.Stop()
	for {
		select {
		case data := <-c.msgChan:
			if err := c.Conn.WriteMessage(websocket.BinaryMessage, *data); err != nil {
				log.Println("write message error:", err)
				continue
			}
		case <-c.CloseChan:
			return
		}
	}
}
