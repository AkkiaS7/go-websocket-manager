package wsmgr

type RequestHandlerFunc func(*Request)

type RequestHandler struct {
	WorkerList     []*MsgWorker // Worker列表
	WorkerPoolSize uint64       //工作池的worker数量
	mgr            *Mgr         //隶属的websocket管理器
}

// HandleRequest 将request分发给worker进行处理
func (rh *RequestHandler) HandleRequest(req *Request) {
	//TODO 分发任务
}
