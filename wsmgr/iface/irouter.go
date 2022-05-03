package iface

/*
	路由接口
*/
type IRouter interface {
	// Add 添加路由
	Add(pattern string, handler IRequestHandler)
	// Get 删除路由
	Get(pattern string) (IRequestHandler, bool)
}
