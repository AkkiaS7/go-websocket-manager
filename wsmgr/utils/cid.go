package utils

import "sync"

//TODO 应该有一个更好的方法生成cid
var (
	cid     uint64 = 0
	cidLock sync.Mutex
)

func GetCid() uint64 {
	cidLock.Lock()
	defer cidLock.Unlock()
	cid++
	return cid
}
