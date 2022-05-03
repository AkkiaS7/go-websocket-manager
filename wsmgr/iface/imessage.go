package iface

type IMessage interface {
	GetCmd() string
	GetData() interface{}
	SetCmd(string)
	SetData(interface{})
	//Unmarshal 反序列化消息
	Unmarshal([]byte) IMessage
}
