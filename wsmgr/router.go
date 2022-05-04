package wsmgr

/*
	路由接口
*/
type IRouter interface {
	// Add 添加路由
	Add(pattern string, handler RequestHandler)
	// Get 获取路由
	Get(pattern string) (RequestHandler, bool)
}

type Router struct {
	routes map[string]RequestHandler
}

// Add a route to the router
func (r *Router) Add(pattern string, handler RequestHandler) {
	r.routes[pattern] = handler
}

// Get a route from the router
func (r *Router) Get(pattern string) (RequestHandler, bool) {
	handler, ok := r.routes[pattern]
	return handler, ok
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]RequestHandler),
	}
}
