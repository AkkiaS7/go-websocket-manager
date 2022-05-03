package wsmgr

import "github.com/AkkiaS7/go-websocket-mgr/wsmgr/iface"

type Mgr struct {
	MsgHandler iface.IMsgHandler  // 消息处理器
	ConnMgr    iface.IConnManager // 连接管理器
	Router     iface.IRouter      // 路由器
}

func NewMgr() *Mgr {
	mgr := &Mgr{}
	mgr.Router = NewRouter()
	mgr.ConnMgr = NewConnManager()
	mgr.MsgHandler = NewMsgHandler(mgr)
	return mgr
}

func (m *Mgr) AddRouter(pattern string, handler iface.IRequestHandler) {
	m.Router.Add(pattern, handler)
}

func (m *Mgr) GetConnMgr() iface.IConnManager {
	return m.ConnMgr
}

func (m *Mgr) GetRouter() iface.IRouter {
	return m.Router
}
