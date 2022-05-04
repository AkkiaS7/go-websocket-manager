package conf

const (
	WorkerPoolSize     = 10
	WorkerReqQueueSize = 100
	MaxDataSize        = 4 * 1024 * 1024
	MaxMsgSize         = 5 + MaxDataSize
)
