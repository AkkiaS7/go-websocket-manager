package wsmgr

type Message struct {
	Header
}

type Header struct {
	CMDLength  uint16
	DataLength uint32
}

type Body struct {
	CMD  string
	Data []byte
}
