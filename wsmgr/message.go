package wsmgr

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
