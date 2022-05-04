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

type Request struct {
	Conn IConnection `json:"-"`
	Msg  IMessage    `json:"-"`
}

//GetConnection 获取请求连接信息
func (r *Request) GetConnection() IConnection {
	return r.Conn
}

//GetMsg 获取请求消息数据
func (r *Request) GetMsg() IMessage {
	return r.Msg
}
