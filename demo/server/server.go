package main

import (
	"github.com/AkkiaS7/go-websocket-mgr/wsmgr"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	upgrader = wsmgr.Upgrader{}
)

func main() {
	mgr := wsmgr.New()
	mgr.Router.Add("ping", func(c *wsmgr.Request) {
		log.Println("ping")
		conn := c.Conn
		msg := &wsmgr.Message{}
		msg = msg.BuildMsg("log", c.MsgBody.Data)
		conn.SendMsg(msg)
	})
	go mgr.Start()
	r := gin.Default()
	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		mgr.ConnMgr.Add(conn)
	})
	r.Run(":8888")
	select {}
}
