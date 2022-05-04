package main

import (
	"github.com/AkkiaS7/go-websocket-mgr/wsmgr"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

func main() {
	u := url.URL{Scheme: "ws", Host: "localhost:8888", Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}
	cmd := "ping"
	data := []byte("hello")
	msg := &wsmgr.Message{}
	msg = msg.BuildMsg(cmd, data)
	err = conn.WriteMessage(websocket.BinaryMessage, msg.Pack())
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("send:", string(msg.Pack()))
	res := &wsmgr.Message{}
	_, m, _ := conn.ReadMessage()
	log.Println("recv:", string(m))
	res = wsmgr.Unpack(m)
	log.Println(string(res.Body.Data))
}
