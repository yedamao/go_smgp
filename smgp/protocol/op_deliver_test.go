package protocol

import (
	"testing"
)

func TestDeliver(t *testing.T) {
	var (
		msgId      = [10]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		srcTermID  = "1069000000"
		destTermID = "17500000000"
		msgContent = "hello test msg"
	)

	op, err := NewDeliver(
		123,
		msgId,
		IS_REPORT,
		GB18030,
		"", srcTermID, destTermID,
		[]byte(msgContent),

		nil, nil,
		nil,
		nil, nil, nil, nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	parsed, err := ParseOperation(op.Serialize())
	if err != nil {
		t.Fatal(err)
	}

	deliver := parsed.(*Deliver)

	if deliver.MsgID.raw != msgId ||
		deliver.SrcTermID.String() != srcTermID ||
		deliver.DestTermID.String() != destTermID {

		t.Error("parsed deliver not match")
	}

}

func TestDeliverResp(t *testing.T) {
	op, err := NewDeliverResp(1234, [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, STAT_OK)
	if err != nil {
		t.Fatal(err)
	}

	parsed, err := ParseOperation(op.Serialize())
	if err != nil {
		t.Fatal(err)
	}

	resp := parsed.(*DeliverResp)

	if resp.MsgID.String() != op.MsgID.String() ||
		resp.Status.Data() != op.Status.Data() {
		t.Error("DeliverResp parsed not match")
	}
}
