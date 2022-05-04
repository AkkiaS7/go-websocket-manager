package wsmgr

type Mgr struct {
	MsgHandler MsgHandler     // 消息处理器
	ReqHandler RequestHandler // 请求处理器
	ConnMgr    ConnManager    // 连接管理器
	Router     Router         // 路由器
}
