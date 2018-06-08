package protocol

import (
	"errors"
	"fmt"
)

type Operation interface {
	//Header returns the Operation header, decoded. Header fields
	// can be updated before reserialzing .
	GetHeader() *Header

	// SerializeTo encodes Operation to it's binary form,
	// include the header and body
	Serialize() []byte

	// String
	String() string

	// return status
	GetStatus() Status
}

func ParseOperation(data []byte) (Operation, error) {
	if len(data) < 12 {
		return nil, errors.New("Invalide data length")
	}

	header, err := ParseHeader(data)
	if err != nil {
		return nil, err
	}

	if int(header.PacketLength) != len(data) {
		return nil, errors.New("Invalide data length")
	}

	var op Operation

	switch header.RequestID {
	case SMGP_LOGIN:
		op, err = ParseLogin(header, data[12:])
	case SMGP_LOGIN_RESP:
		op, err = ParseLoginResp(header, data[12:])

	case SMGP_SUBMIT:
		op, err = ParseSubmit(header, data[12:])
	case SMGP_SUBMIT_RESP:
		op, err = ParseSubmitResp(header, data[12:])

	case SMGP_DELIVER:
		op, err = ParseDeliver(header, data[12:])
	case SMGP_DELIVER_RESP:
		op, err = ParseDeliverResp(header, data[12:])

	case SMGP_ACTIVE_TEST:
		op, err = ParseActiveTest(header, data[12:])
	case SMGP_ACTIVE_TEST_RESP:
		op, err = ParseActiveTestResp(header, data[12:])

	case SMGP_EXIT:
		op, err = ParseExit(header, data[12:])
	case SMGP_EXIT_RESP:
		op, err = ParseExitResp(header, data[12:])

	default:
		err = fmt.Errorf("Unknow Operation CmdId: 0x%x", header.RequestID)
	}

	return op, err
}
