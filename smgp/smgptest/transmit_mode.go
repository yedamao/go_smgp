package smgptest

import (
	"log"
	"net"
	"time"

	"github.com/yedamao/go_smgp/smgp/protocol"
)

// TransmitSession TRANSMIT_MODE session
type TransmitSession struct {
	*Session
}

func runTransmitSession(s *Session) {
	r := &TransmitSession{s}

	go r.recvWorker()
	go r.sendWorker()
}

// 模拟定时发送 Deliver 至 SP
func (s *TransmitSession) sendWorker() {
	log.Println("TransmitSession: sendWorker")

	for {
		tick := time.NewTicker(10 * time.Second)
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

// 接收Submit, DeliverResp
func (s *TransmitSession) recvWorker() {
	defer s.Close()

	log.Println("TransmitSession: recvWorker")
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

		case protocol.SMGP_SUBMIT:
			log.Println(op)
			fakeID := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
			s.SubmitResp(op.GetHeader().SequenceID, fakeID, protocol.STAT_OK)

		case protocol.SMGP_ACTIVE_TEST:
			log.Println("ActiveTest ...")
			s.ActiveTestResp(op.GetHeader().SequenceID)

		case protocol.SMGP_EXIT:
			log.Println("Exit ...")
			s.ExitResp(op.GetHeader().SequenceID)
			return
		}
	}
}
