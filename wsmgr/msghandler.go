package wsmgr

/*
	消息处理器
	用于从消息中抽出路由参数与数据构造Request结构体
	然后交由RequestHandler处理
*/

type MsgHandler struct {
	TaskQueue      []chan Request //请求任务的消息队列
	WorkerPoolSize uint64         //工作池的worker数量
	mgr            *Mgr           //隶属的websocket管理器
}

//StartWorker 启动worker工作池
func (mh *MsgHandler) StartWorker(workerPoolSize int) {
}

//SendMsgToTaskQueue 将消息交给TaskQueue,由worker进行处理
func (mh *MsgHandler) SendMsgToTaskQueue(rawData *[]byte) {

}
