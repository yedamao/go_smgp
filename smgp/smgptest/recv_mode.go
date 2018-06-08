package smgptest

import (
	"log"
	"net"
	"time"

	"github.com/yedamao/go_smgp/smgp/protocol"
)

// RECEIVE_MODE session
type RecvSession struct {
	*Session
}

func runRecvSession(s *Session) {
	r := &RecvSession{s}

	go r.recvWorker()
	go r.sendWorker()
}

// 模拟定时发送 Deliver 至 SP
func (s *RecvSession) sendWorker() {

	for {
		tick := time.NewTicker(5 * time.Second)
		select {
		case <-tick.C:
			// Deliver
			if err := s.Deliver(); err != nil {
				log.Println("Deliver: ", err)
				return
			}
		}
	}
}

// 接收DeliverResp
func (s *RecvSession) recvWorker() {
	defer s.Close()

	for {
		op, err := s.Read()
		if err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
				continue
			}
			log.Println("TestServer: read error exit. ", err)
			return
		}

		switch op.GetHeader().RequestID {
		case protocol.SMGP_DELIVER_RESP:
			log.Println("DeliverResp ...")
			log.Println(op)

		case protocol.SMGP_ACTIVE_TEST:
			log.Println("ActiveTest ...")
			s.ActiveTestResp(op.GetHeader().SequenceID)
		}
	}
}
