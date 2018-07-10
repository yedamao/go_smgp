package smgptest

import (
	"log"
	"net"

	connp "github.com/yedamao/go_smgp/smgp/conn"
	"github.com/yedamao/go_smgp/smgp/protocol"
)

func newSession(rawConn net.Conn) {
	s := &Session{*connp.NewConn(rawConn)}

	// check login
	op, err := s.Read()
	if err != nil {
		log.Println("TestServer: read error exit. ", err)
		return
	}

	// type assert
	login, ok := op.(*protocol.Login)
	if !ok {
		log.Println("TestServer: type assert error .", err)
		return
	}
	s.LoginResp(op.GetHeader().SequenceID, protocol.STAT_OK, "authserver", protocol.VERSION)

	switch login.LoginMode {
	case protocol.SEND_MODE:
		log.Println("login: SEND_MODE")
		runSendSession(s)

	case protocol.RECEIVE_MODE:
		log.Println("login: RECEIVE_MODE")
		runRecvSession(s)

	case protocol.TRANSMIT_MODE:
		log.Println("login: TRANSMIT_MODE")
		runTransmitSession(s)
	default:
		log.Println("Unknown login mode. ")
		return
	}
}

// Session 代表sp->运营商的一条连接
type Session struct {
	connp.Conn
}

// LoginResp operation
func (s *Session) LoginResp(
	sequenceID uint32, status protocol.Status,
	authenticatorServer string, serverVersion uint8,
) error {

	op, err := protocol.NewLoginResp(sequenceID, status, authenticatorServer, serverVersion)
	if err != nil {
		return err
	}

	return s.Write(op)
}

// SubmitResp operation
func (s *Session) SubmitResp(sequenceID uint32, msgID [10]byte, status protocol.Status) error {

	op, err := protocol.NewSubmitResp(sequenceID, msgID, status)
	if err != nil {
		return err
	}

	return s.Write(op)
}

// Deliver operation
func (s *Session) Deliver() error {

	op, err := protocol.NewDeliver(
		123, [...]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, protocol.NOT_REPORT, protocol.ASCII,
		"", "1069000000", "17611111111", []byte("hello test msg"),
		nil, nil, nil, nil, nil, nil, nil,
	)
	if err != nil {
		return err
	}

	return s.Write(op)
}

// ActiveTestResp operation
func (s *Session) ActiveTestResp(sequenceID uint32) error {

	op, err := protocol.NewActiveTestResp(sequenceID)
	if err != nil {
		return err
	}

	return s.Write(op)
}

// ExitResp operation
func (s *Session) ExitResp(sequenceID uint32) error {

	op, err := protocol.NewExitResp(sequenceID)
	if err != nil {
		return err
	}

	return s.Write(op)
}
