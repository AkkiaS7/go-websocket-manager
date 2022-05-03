package wsmgr

import "github.com/AkkiaS7/go-websocket-mgr/wsmgr/iface"

type Router struct {
	routes map[string]iface.IRequestHandler
}

// Add a route to the router
func (r *Router) Add(pattern string, handler iface.IRequestHandler) {
	r.routes[pattern] = handler
}

// Get a route from the router
func (r *Router) Get(pattern string) (iface.IRequestHandler, bool) {
	handler, ok := r.routes[pattern]
	return handler, ok
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]iface.IRequestHandler),
	}
}
