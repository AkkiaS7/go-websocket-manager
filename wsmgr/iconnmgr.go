package wsmgr

/*
	连接管理器的抽象层
*/
type IConnManager interface {
	// Add 将连接加入连接管理器
	Add(conn IConnection)
	// Remove 将连接从连接管理器中删除
	Remove(conn IConnection)
	// Get 从连接管理器中获取连接
	Get(connID uint64) (IConnection, error)
}
