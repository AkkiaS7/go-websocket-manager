package iface

type IRequestHandler interface {
	// Handle 处理请求对应的业务
	Handle(request IRequest)
}
