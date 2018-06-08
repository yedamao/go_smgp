package smgptest

import (
	"log"
	"net"

	"github.com/yedamao/go_smgp/smgp/protocol"
)

// SEND_MODE session
type SendSession struct {
	*Session
}

func runSendSession(s *Session) {
	r := &SendSession{s}

	go r.worker()
}

// handle connection of SEND_MODE
func (s *SendSession) worker() {
	// close session
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
		case protocol.SMGP_SUBMIT:
			log.Println(op)
			fakeId := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
			s.SubmitResp(op.GetHeader().SequenceID, fakeId, protocol.STAT_OK)

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
