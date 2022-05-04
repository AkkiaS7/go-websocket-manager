package wsmgr

import (
	"bytes"
	"encoding/binary"
	"github.com/AkkiaS7/go-websocket-mgr/wsmgr/conf"
	"log"
	"sync"
)

/*
	消息处理器
	用于从消息中抽出路由参数与数据构造Request结构体
	然后交由RequestHandler处理
*/

type MsgHandler struct {
	WorkerList     map[uint64]*MsgWorker // Worker列表
	WorkerPoolSize uint64                //工作池的worker数量
	Lock           sync.RWMutex

	mgr *Mgr //隶属的websocket管理器
}

type MsgWorker struct {
	ID       uint64        // worker的ID
	mh       *MsgHandler   // 所属的消息处理器
	ReqQueue chan *Request // 请求队列
}

func NewMsgHandler(mgr *Mgr) *MsgHandler {
	return &MsgHandler{
		mgr:            mgr,
		WorkerPoolSize: conf.WorkerPoolSize,
		WorkerList:     make(map[uint64]*MsgWorker),
	}
}

//Start 启动消息处理器
func (mh *MsgHandler) Start() {
	log.Println("MsgHandler Start")
	mh.StartWorker()
}

//StartWorker 启动worker工作池
func (mh *MsgHandler) StartWorker() {
	for i := uint64(0); i < mh.WorkerPoolSize; i++ {
		go mh.StartOneWorker(i)
	}
}

func (mh *MsgHandler) StartOneWorker(id uint64) {
	worker := MsgWorker{
		ID:       id,
		mh:       mh,
		ReqQueue: make(chan *Request, conf.WorkerReqQueueSize),
	}
	mh.Lock.Lock()
	mh.WorkerList[id] = &worker
	mh.Lock.Unlock()
	for {
		select {
		case req := <-worker.ReqQueue:
			log.Println("MsgHandler StartOneWorker ReqQueue", req.Conn.ConnID)
			//构造Message
			dataBuff := bytes.NewReader(*req.rawMsg)
			msg := &Message{}
			// 先读取Header
			msg.Header = &MsgHeader{}
			if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Header.CMDLength); err != nil {
				log.Println("解析消息头失败：", err)
			}
			if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Header.DataLength); err != nil {
				log.Println("解析消息头失败：", err)
			}
			log.Println("解析消息头成功：", msg.Header.CMDLength, msg.Header.DataLength)
			// 根据Header中的长度，读取Body
			cmd := make([]byte, msg.Header.CMDLength)
			if err := binary.Read(dataBuff, binary.LittleEndian, &cmd); err != nil {
				log.Println("解析CMD失败：", err)
			}
			req.MsgBody = &MsgBody{
				CMD:  string(cmd),
				Data: make([]byte, msg.Header.DataLength),
			}
			if err := binary.Read(dataBuff, binary.LittleEndian, &req.MsgBody.Data); err != nil {
				log.Println("解析Data失败：", err)
			}
			log.Println("收到消息：", msg)
			//交由RequestHandler处理
			mh.mgr.ReqHandler.HandleRequest(req)
			//处理完成后，返回结果
			//返回结果
		}
	}
}

//SendToTaskQueue 将消息交给TaskQueue,由worker进行处理
func (mh *MsgHandler) SendToTaskQueue(req *Request) {
	//暂时先发给一个Worker
	mh.WorkerList[0].ReqQueue <- req

}
