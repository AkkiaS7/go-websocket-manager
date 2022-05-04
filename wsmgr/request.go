package wsmgr

/*
	从Message中分离出来的具体请求内容
	用于传递给请求处理器
*/

type Request struct {
	Conn *Connection
	*MsgBody

	rawMsg *[]byte
}
