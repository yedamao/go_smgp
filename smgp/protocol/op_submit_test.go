package protocol

import (
	"testing"
)

func TestSubmit(t *testing.T) {
	var (
		TP_pid     uint8 = 1
		MServiceID       = &OctetString{Data: []byte("22396"), FixedLen: 8}
	)

	op, err := NewSubmit(
		123,
		NEED_REPORT, HIGHER_PRIORITY,
		"serviceId", "0", "0", "0",
		GB18030,
		"", "", "10690001111", "",
		[]string{"17600000000", "17800000000", "17700000000"},
		[]byte("hello test msg"),

		// 可选字段
		nil, packUi8(TP_pid),
		nil,
		nil,
		nil, nil, nil,
		nil, nil,
		nil, nil,
		nil, nil, MServiceID.Byte(),
	)
	if err != nil {
		t.Fatal(err)
	}

	parsed, err := ParseOperation(op.Serialize())
	if err != nil {
		t.Fatal(err)
	}

	submit := parsed.(*Submit)

	if op.MsgType != submit.MsgType ||
		op.Priority != submit.Priority ||
		op.ServiceID.String() != submit.ServiceID.String() ||
		op.MsgFormat != submit.MsgFormat ||
		op.SrcTermID.String() != submit.SrcTermID.String() ||
		op.DestTermIDCount != submit.DestTermIDCount {

		t.Error("parsed submit not match")
	}

	if submit.options.TP_udhi() != TP_pid ||
		submit.options.MServiceID() != MServiceID.String() {

		t.Error("options field not match")
	}
}

func TestSubmitResp(t *testing.T) {
	op, err := NewSubmitResp(1234, [10]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, STAT_OK)
	if err != nil {
		t.Fatal(err)
	}

	parsed, err := ParseOperation(op.Serialize())
	if err != nil {
		t.Fatal(err)
	}

	resp := parsed.(*SubmitResp)

	if resp.MsgID.String() != op.MsgID.String() ||
		resp.Status.Data() != op.Status.Data() {
		t.Error("SubmitResp parsed not match")
	}
}
