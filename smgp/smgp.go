package smgp

import (
	"errors"
	"net"

	"github.com/yedamao/go_smgp/smgp/conn"
	"github.com/yedamao/go_smgp/smgp/protocol"
)

// protocol of function generation sequenceID
type SequenceFunc func() uint32

type Smgp struct {
	conn.Conn
	newSeqNum SequenceFunc

	SPID   string // SP的企业代码
	SPCode string // SP的服务代码
}

func NewSmgp(
	addr string, clientID, sharedSecret, spID string,
	mode uint8, newSeqNum SequenceFunc,
) (*Smgp, error) {
	if nil == newSeqNum {
		return nil, errors.New("newSeqNum or hdl must not be nil")
	}

	s := &Smgp{
		newSeqNum: newSeqNum,
		SPID:      spID,
		SPCode:    clientID,
	}

	if err := s.Connect(addr); err != nil {
		return nil, err
	}

	if err := s.Login(clientID, sharedSecret, mode); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Smgp) setup() error {
	// TODO

	return nil
}

func (s *Smgp) Connect(addr string) error {
	connection, err := net.Dial("tcp", addr)
	if err != nil {
		return err

	}
	s.Conn = *conn.NewConn(connection)

	return nil
}

func (s *Smgp) Login(clientID, sharedSecret string, mode uint8) error {

	op, err := protocol.NewLogin(s.newSeqNum(), clientID, sharedSecret, mode)
	if err != nil {
		return err
	}
	if err = s.Write(op); err != nil {
		return err
	}

	// Read block
	var resp protocol.Operation
	if resp, err = s.Read(); err != nil {
		return err
	}

	if resp.GetHeader().RequestID != protocol.SMGP_LOGIN_RESP {
		return errors.New("Login Resp Wrong RequestID")
	}

	status := resp.GetStatus()
	if protocol.STAT_OK != status {
		return status.Error()
	}

	return nil
}

func (s *Smgp) Submit(
	needReport, priority uint8,
	serviceId, feeType, feeCode, fixedFee string,
	msgFormat uint8,
	validTime, atTime, srcTermID, chargeTermID string,
	destTermID []string, msgContent []byte,

	TP_udhi, PkTotal, PkNumber uint8, // 长短信字段
) (uint32, error) {

	pMsgSrc := &protocol.OctetString{Data: []byte(s.SPID), FixedLen: 8}
	pMServiceID := &protocol.OctetString{FixedLen: 21}

	sequence := s.newSeqNum()
	op, err := protocol.NewSubmit(
		sequence,
		needReport, priority,
		serviceId, feeType, feeCode, fixedFee,
		msgFormat,
		validTime, atTime, srcTermID, chargeTermID,
		destTermID, msgContent,

		// 可选参数
		nil, []byte{TP_udhi},
		nil,
		pMsgSrc.Byte(),
		nil, nil, nil,
		nil, nil,
		[]byte{PkTotal}, []byte{PkNumber},
		nil, nil, pMServiceID.Byte(),
	)
	if err != nil {
		return 0, err
	}

	return sequence, s.Write(op)
}

func (s *Smgp) DeliverResp(
	sequenceID uint32, msgID [10]byte, status protocol.Status,
) error {

	op, err := protocol.NewDeliverResp(sequenceID, msgID, status)
	if err != nil {
		return err
	}

	return s.Write(op)
}

// ActiveTest operation
func (s *Smgp) ActiveTest() error {

	op, err := protocol.NewActiveTest(s.newSeqNum())
	if err != nil {
		return err
	}

	return s.Write(op)
}

func (s *Smgp) ActiveTestResp(sequenceID uint32) error {

	op, err := protocol.NewActiveTestResp(sequenceID)
	if err != nil {
		return err
	}

	return s.Write(op)
}

func (s *Smgp) Exit() error {

	op, err := protocol.NewExit(s.newSeqNum())
	if err != nil {
		return err
	}

	return s.Write(op)
}

func (s *Smgp) ExitResp(seq uint32) error {

	op, err := protocol.NewExitResp(seq)
	if err != nil {
		return err
	}

	return s.Write(op)
}
