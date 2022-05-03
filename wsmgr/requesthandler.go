package wsmgr

import "github.com/AkkiaS7/go-websocket-mgr/wsmgr/iface"

type RequestHandler struct{}

// Handle 处理请求的方法
func (rh *RequestHandler) Handle(request iface.IRequest) {}
