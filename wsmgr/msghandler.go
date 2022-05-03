package wsmgr

import (
	"encoding/json"
	"github.com/AkkiaS7/go-websocket-mgr/wsmgr/conf"
	"github.com/AkkiaS7/go-websocket-mgr/wsmgr/iface"
)

type MsgHandler struct {
	TaskQueue      []chan iface.IRequest //请求任务的消息队列
	WorkerPoolSize uint64                //工作池的worker数量
	WsMgr          iface.IMgr            //隶属的websocket管理器
}

//Marshal 将请求数据转换为字节数组
func (mh *MsgHandler) Marshal(msg iface.IRequest) ([]byte, error) {
	//默认使用json序列化
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//StartWorker 启动worker工作池
func (mh *MsgHandler) StartWorker(workerPoolSize int) {
}

//SendMsgToTaskQueue 将消息交给TaskQueue,由worker进行处理
func (mh *MsgHandler) SendMsgToTaskQueue(request iface.IRequest) {
	msg := request.GetMsg()
	handler, ok := mh.WsMgr.GetRouter().Get(msg.GetCmd())
	if !ok {
		return
	}
	handler.Handle(request)

}

func NewMsgHandler(mgr iface.IMgr) iface.IMsgHandler {
	return &MsgHandler{
		TaskQueue:      make([]chan iface.IRequest, conf.TaskQueueSize),
		WorkerPoolSize: conf.WorkerPoolSize,
		WsMgr:          mgr,
	}
}
