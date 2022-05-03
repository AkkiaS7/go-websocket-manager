module demo

replace github.com/AkkiaS7/go-websocket-mgr/wsmgr => ../wsmgr // 绝对路径 或 相对路径 都可以

go 1.18

require github.com/AkkiaS7/go-websocket-mgr/wsmgr v0.0.0-00010101000000-000000000000

require github.com/gorilla/websocket v1.5.0 // indirect
