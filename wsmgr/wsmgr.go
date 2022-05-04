package wsmgr

type IMgr interface {
	AddRouter(pattern string, handler IRequestHandler)
	GetConnMgr() IConnManager
	GetRouter() IRouter
}

type Mgr struct {
	MsgHandler IMsgHandler  // 消息处理器
	ConnMgr    IConnManager // 连接管理器
	Router     IRouter      // 路由器
}

func NewMgr() *Mgr {
	mgr := &Mgr{}
	mgr.Router = NewRouter()
	mgr.ConnMgr = NewConnManager()
	mgr.MsgHandler = NewMsgHandler(mgr)
	return mgr
}

func (m *Mgr) AddRouter(pattern string, handler IRequestHandler) {
	m.Router.Add(pattern, handler)
}

func (m *Mgr) GetConnMgr() IConnManager {
	return m.ConnMgr
}

func (m *Mgr) GetRouter() IRouter {
	return m.Router
}
