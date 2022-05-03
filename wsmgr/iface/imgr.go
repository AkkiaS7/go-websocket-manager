package iface

type IMgr interface {
	AddRouter(pattern string, handler IRequestHandler)
	GetConnMgr() IConnManager
	GetRouter() IRouter
}
