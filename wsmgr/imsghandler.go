package wsmgr

/*
	消息处理器抽象层
	用于从消息中抽出路由参数与数据交由RequestHandler处理
*/
type IMsgHandler interface {
	//StartWorker 启动worker工作池
	StartWorker(workerPoolSize int)
	//SendMsgToTaskQueue 将消息交给TaskQueue,由worker进行处理
	SendMsgToTaskQueue(request IRequest)
}
