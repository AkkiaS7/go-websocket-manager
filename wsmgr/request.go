package wsmgr

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
