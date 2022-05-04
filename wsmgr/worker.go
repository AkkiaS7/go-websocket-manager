package wsmgr

/*
	worker的抽象层
*/
type IWorker interface {
	//Start 启动worker
	Start()
	//Stop 停止worker
	Stop()
	//Serve 处理业务
	Serve()
}
