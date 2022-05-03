package wsmgr

import "github.com/AkkiaS7/go-websocket-mgr/wsmgr/iface"

type Request struct {
	Conn iface.IConnection `json:"-"`
	Msg  iface.IMessage    `json:"-"`
}

//GetConnection 获取请求连接信息
func (r *Request) GetConnection() iface.IConnection {
	return r.Conn
}

//GetMsg 获取请求消息数据
func (r *Request) GetMsg() iface.IMessage {
	return r.Msg
}
