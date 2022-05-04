package wsmgr

type IRequestHandler interface {
	// Handle 处理请求对应的业务
	Handle(request IRequest)
}

type RequestHandler struct{}

// Handle 处理请求的方法
func (rh *RequestHandler) Handle(request IRequest) {}
