package wsmgr

type Mgr struct {
	MsgHandler *MsgHandler     // 消息处理器
	ReqHandler *RequestHandler // 请求处理器
	ConnMgr    *ConnManager    // 连接管理器
	Router     *Router         // 路由器
}

// New 创建一个新的管理器
func New() *Mgr {
	mgr := &Mgr{}
	mgr.ConnMgr = NewConnManager()
	mgr.Router = NewRouter()
	mgr.MsgHandler = NewMsgHandler(mgr)
	mgr.ReqHandler = NewRequestHandler(mgr)
	return mgr
}

func (mgr *Mgr) Start() {
	mgr.MsgHandler.StartWorker()
	mgr.ReqHandler.StartWorker()
	select {}
}
