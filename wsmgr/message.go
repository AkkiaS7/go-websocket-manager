package wsmgr

import (
	"encoding/json"
)

type IMessage interface {
	GetCmd() string
	GetData() interface{}
	SetCmd(string)
	SetData(interface{})
	//Unmarshal 反序列化消息
	Unmarshal([]byte) IMessage
}

type Msg struct {
	Cmd  string      `json:"cmd"`
	Data interface{} `json:"data"`
}

func (m *Msg) GetCmd() string {
	return m.Cmd
}
func (m *Msg) GetData() interface{} {
	return m.Data
}
func (m *Msg) SetCmd(str string) {
	m.Cmd = str
}
func (m *Msg) SetData(data interface{}) {
	m.Data = data
}

//Unmarshal 反序列化消息
func (m *Msg) Unmarshal(data []byte) IMessage {
	err := json.Unmarshal(data, m)
	if err != nil {
		return nil
	}
	return m
}
func NewMsg(cmd string, data interface{}) *Msg {
	return &Msg{
		Cmd:  cmd,
		Data: data,
	}
}
