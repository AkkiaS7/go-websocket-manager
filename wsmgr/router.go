package wsmgr

/*
	路由接口
*/
type IRouter interface {
	// Add 添加路由
	Add(pattern string, handler RequestHandlerFunc)
	// Get 获取路由
	Get(pattern string) (RequestHandlerFunc, bool)
}

type RouterGroup struct {
	routes map[string]RequestHandlerFunc
}

// Add a route to the router
func (r *RouterGroup) Add(pattern string, handler RequestHandlerFunc) {
	r.routes[pattern] = handler
}

// Get a route from the router
func (r *RouterGroup) Get(pattern string) (RequestHandlerFunc, bool) {
	handler, ok := r.routes[pattern]
	return handler, ok
}

func NewRouter() *RouterGroup {
	return &RouterGroup{
		routes: make(map[string]RequestHandlerFunc),
	}
}
