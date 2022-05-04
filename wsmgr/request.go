package wsmgr

/*
	从原始消息中分离出来的具体请求内容
	用于传递给请求处理器
*/

type Request struct {
	Conn *Connection `json:"-"`
	Msg  *Msg        `json:"-"`
}

//GetConnection 获取请求连接信息
func (r *Request) GetConnection() IConnection {
	return r.Conn
}

//GetMsg 获取请求消息数据
func (r *Request) GetMsg() IMessage {
	return r.Msg
}
