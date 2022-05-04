package wsmgr

/*
	从消息中分离路由信息与请求信息
*/
type IRequest interface {
	//GetConnection 获取请求连接信息
	GetConnection() IConnection
	//GetMsg 获取请求消息数据
	GetMsg() IMessage
}
