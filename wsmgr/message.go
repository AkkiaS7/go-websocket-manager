package wsmgr

import (
	"bytes"
	"encoding/binary"
	"log"
)

type Message struct {
	Header *MsgHeader
	Body   *MsgBody
}

type MsgHeader struct {
	CMDLength  uint16
	DataLength uint32
}

type MsgBody struct {
	CMD  string
	Data []byte
}

func (m *Message) BuildMsg(cmd string, data []byte) *Message {
	m.Body = &MsgBody{
		CMD:  cmd,
		Data: data,
	}
	m.Header = &MsgHeader{
		CMDLength:  uint16(len(cmd)),
		DataLength: uint32(len(data)),
	}
	return m
}

func (m *Message) Pack() []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, m.Header.CMDLength)
	binary.Write(buf, binary.LittleEndian, m.Header.DataLength)
	binary.Write(buf, binary.LittleEndian, []byte(m.Body.CMD))
	binary.Write(buf, binary.LittleEndian, m.Body.Data)
	return buf.Bytes()
}

func Unpack(data []byte) *Message {
	m := &Message{
		Header: &MsgHeader{},
		Body:   &MsgBody{},
	}
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.LittleEndian, &m.Header.CMDLength)
	if err != nil {
		log.Println("Unpack CMDLength error:", err)
		return nil
	}
	binary.Read(buf, binary.LittleEndian, &m.Header.DataLength)
	cmd := make([]byte, m.Header.CMDLength)
	binary.Read(buf, binary.LittleEndian, &cmd)
	m.Body.CMD = string(cmd)
	m.Body.Data = make([]byte, m.Header.DataLength)
	binary.Read(buf, binary.LittleEndian, &m.Body.Data)
	return m
}
