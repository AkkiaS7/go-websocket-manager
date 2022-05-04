package wsmgr

import "github.com/gorilla/websocket"

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
