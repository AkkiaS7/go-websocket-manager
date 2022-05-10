package wsmgr

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type Upgrader websocket.Upgrader

func (u *Upgrader) Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error) {
	upgrader := new(websocket.Upgrader)
	upgrader = (*websocket.Upgrader)(u)
	return upgrader.Upgrade(w, r, responseHeader)
	//TODO upgrade后直接在连接管理器中添加连接
}
