package wsmgr

import (
	"bytes"
	"encoding/binary"
	"github.com/AkkiaS7/go-websocket-mgr/wsmgr/conf"
	"log"
)

/*
	消息处理器
	用于从消息中抽出路由参数与数据构造Request结构体
	然后交由RequestHandler处理
*/

type MsgHandler struct {
	WorkerList     []*MsgWorker // Worker列表
	WorkerPoolSize uint64       //工作池的worker数量
	mgr            *Mgr         //隶属的websocket管理器
}

type MsgWorker struct {
	ID       uint64        // worker的ID
	mh       *MsgHandler   // 所属的消息处理器
	ReqQueue chan *Request // 请求队列
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
	for {
		select {
		case req := <-worker.ReqQueue:
			//构造Message
			dataBuff := bytes.NewReader(*req.rawMsg)
			msg := &Message{}
			// 先读取Header
			msg.Header = &MsgHeader{}
			if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Header); err != nil {
				log.Println("解析消息头失败：", err)
			}
			if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Body); err != nil {
				log.Println("解析消息头失败：", err)
			}
			// 根据Header中的长度，读取Body
			cmd := make([]byte, msg.Header.CMDLength)
			if err := binary.Read(dataBuff, binary.LittleEndian, &cmd); err != nil {
				log.Println("解析CMD失败：", err)
			}
			msg.Body = &MsgBody{
				CMD:  string(cmd),
				Data: make([]byte, msg.Header.DataLength),
			}
			if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Body.Data); err != nil {
				log.Println("解析Data失败：", err)
			}
			//构造request
			req.Msg = msg
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
