package wsmgr

import (
	"github.com/AkkiaS7/go-websocket-mgr/wsmgr/conf"
	"log"
)

type RequestHandlerFunc func(*Request)

type RequestHandler struct {
	WorkerList     map[uint64]*ReqWorker // Worker列表
	WorkerPoolSize uint64                //工作池的worker数量
	mgr            *Mgr                  //隶属的websocket管理器
}

type ReqWorker struct {
	ID       uint64          // worker的ID
	rh       *RequestHandler // 所属的请求处理器
	ReqQueue chan *Request   // 请求队列
}

//NewRequestHandler 创建一个请求处理器
func NewRequestHandler(mgr *Mgr) *RequestHandler {
	return &RequestHandler{
		WorkerList:     make(map[uint64]*ReqWorker),
		WorkerPoolSize: conf.WorkerPoolSize,
		mgr:            mgr,
	}
}

//StartWorker 启动worker工作池
func (rh *RequestHandler) StartWorker() {
	for i := uint64(0); i < rh.WorkerPoolSize; i++ {
		go rh.StartOneWorker(i)
	}
}

//StartOneWorker 启动一个worker
func (rh *RequestHandler) StartOneWorker(id uint64) {
	worker := &ReqWorker{
		ID:       id,
		rh:       rh,
		ReqQueue: make(chan *Request, conf.WorkerReqQueueSize),
	}
	rh.WorkerList[id] = worker
	for {
		select {
		case req := <-worker.ReqQueue:
			// 获取该请求对应的路由
			handleFunc, ok := rh.mgr.Router.Get(req.CMD)
			if !ok {
				log.Println("No route for", req.CMD)
				continue
			}
			// 调用路由
			handleFunc(req)
		}
	}
}

// HandleRequest 将request分发给worker进行处理
func (rh *RequestHandler) HandleRequest(req *Request) {
	//暂时交由第一个worker处理
	rh.WorkerList[0].ReqQueue <- req
}
