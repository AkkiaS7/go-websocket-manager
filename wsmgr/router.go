package wsmgr

/*
	路由接口
*/
type IRouter interface {
	// Add 添加路由
	Add(pattern string, handler IRequestHandler)
	// Get 删除路由
	Get(pattern string) (IRequestHandler, bool)
}

type Router struct {
	routes map[string]IRequestHandler
}

// Add a route to the router
func (r *Router) Add(pattern string, handler IRequestHandler) {
	r.routes[pattern] = handler
}

// Get a route from the router
func (r *Router) Get(pattern string) (IRequestHandler, bool) {
	handler, ok := r.routes[pattern]
	return handler, ok
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]IRequestHandler),
	}
}
