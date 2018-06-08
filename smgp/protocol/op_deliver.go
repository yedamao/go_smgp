package protocol

import (
	"bytes"
	"fmt"
)

type Deliver struct {
	*Header

	MsgID      *Id          // 短消息流水号
	IsReport   uint8        // 短消息流水号
	MsgFormat  uint8        // 短消息格式
	RecvTime   *OctetString // 短消息接收时间
	SrcTermID  *OctetString // 短消息发送号码
	DestTermID *OctetString // 短消息接收号码
	MsgLength  uint8        //  短消息长度
	MsgContent []byte       // 短消息内容
	Reserve    *OctetString // 保留

	// 可选字段
	Options Options
}

func NewDeliver(
	sequenceID uint32,

	msgID [10]byte,
	isReport, msgFormat uint8,
	recvTime, srcTermID, destTermID string,
	msgContent []byte,

	// 可选参数,  使用slice类型： nil可以区分是否设置此字段.
	// Integer 和 OctetString 由上层转换为[]byte
	TP_pid, TP_udhi []byte,
	LinkID []byte,
	SrcTermType, SrcTermPseudo, SubmitMsgType, SPDealResult []byte,
) (*Deliver, error) {
	op := &Deliver{}

	op.Header = &Header{}
	var length uint32 = 4 + 4 + 4 // header length

	op.MsgID = &Id{raw: msgID}
	length = length + 10

	op.IsReport = isReport
	length = length + 1
	op.MsgFormat = msgFormat
	length = length + 1

	op.RecvTime = &OctetString{Data: []byte(recvTime), FixedLen: 14}
	length = length + 14
	op.SrcTermID = &OctetString{Data: []byte(srcTermID), FixedLen: 21}
	length = length + 21
	op.DestTermID = &OctetString{Data: []byte(destTermID), FixedLen: 21}
	length = length + 21

	msgLen := len(msgContent)
	op.MsgLength = uint8(msgLen)
	length = length + 1

	op.MsgContent = msgContent
	length = length + uint32(msgLen)

	op.Reserve = &OctetString{FixedLen: 8}
	length = length + 8

	// 可选参数
	op.Options = make(Options)

	if TP_pid != nil {
		op.Options[TAG_TP_pid] = TP_pid
		length = length + 1
	}
	if TP_udhi != nil {
		op.Options[TAG_TP_udhi] = TP_udhi
	}
	if LinkID != nil {
		op.Options[TAG_LinkID] = LinkID
	}
	if SrcTermType != nil {
		op.Options[TAG_SrcTermType] = SrcTermType
	}
	if SrcTermPseudo != nil {
		op.Options[TAG_SrcTermPseudo] = SrcTermPseudo
	}
	if SubmitMsgType != nil {
		op.Options[TAG_SubmitMsgType] = SubmitMsgType
	}
	if SPDealResult != nil {
		op.Options[TAG_SPDealResult] = SPDealResult
	}

	// 可选字段长度
	length = length + uint32(op.Options.Len())

	op.PacketLength = length
	op.RequestID = SMGP_DELIVER
	op.SequenceID = sequenceID

	return op, nil
}

func ParseDeliver(hdr *Header, data []byte) (*Deliver, error) {
	p := 0
	op := &Deliver{}
	op.Header = hdr

	op.MsgID = newId(data[p : p+10])
	p = p + 10

	op.IsReport = data[p]
	p = p + 1
	op.MsgFormat = data[p]
	p = p + 1

	op.RecvTime = &OctetString{Data: data[p : p+14], FixedLen: 14}
	p = p + 14
	op.SrcTermID = &OctetString{Data: data[p : p+21], FixedLen: 21}
	p = p + 21
	op.DestTermID = &OctetString{Data: data[p : p+21], FixedLen: 21}
	p = p + 21

	op.MsgLength = data[p]
	p = p + 1

	op.MsgContent = data[p : p+int(op.MsgLength)]
	p = p + int(op.MsgLength)

	op.Reserve = &OctetString{Data: data[p : p+8], FixedLen: 8}
	p = p + 8

	// parse options
	var err error
	if op.Options, err = ParseOptions(data[p:]); err != nil {
		return nil, err
	}

	return op, nil
}

func (op *Deliver) Serialize() []byte {
	b := op.Header.Serialize()

	b = append(b, op.MsgID.raw[:]...)

	b = append(b, packUi8(op.IsReport)...)
	b = append(b, packUi8(op.MsgFormat)...)

	b = append(b, op.RecvTime.Byte()...)
	b = append(b, op.SrcTermID.Byte()...)
	b = append(b, op.DestTermID.Byte()...)

	b = append(b, packUi8(op.MsgLength)...)
	b = append(b, op.MsgContent...)
	b = append(b, op.Reserve.Byte()...)

	// 可选字段
	b = append(b, op.Options.Serialize()...)

	return b
}

func (op *Deliver) String() string {
	var b bytes.Buffer
	b.WriteString(op.Header.String())

	fmt.Fprintln(&b, "--- Deliver ---")
	fmt.Fprintln(&b, "MsgID: ", op.MsgID)

	fmt.Fprintln(&b, "IsReport: ", op.IsReport)
	fmt.Fprintln(&b, "MsgFormat: ", op.MsgFormat)

	fmt.Fprintln(&b, "RecvTime: ", op.RecvTime)
	fmt.Fprintln(&b, "SrcTermID: ", op.SrcTermID)
	fmt.Fprintln(&b, "DestTermID: ", op.DestTermID)

	fmt.Fprintln(&b, "MsgLength: ", op.MsgLength)
	fmt.Fprintln(&b, "MsgContent: ", string(op.MsgContent))

	// TODO
	// print options

	return b.String()
}

func (op *Deliver) GetStatus() Status {
	return STAT_OK
}

// --- Deliver_Resp
type DeliverResp struct {
	*Header

	MsgID  *OctetString
	Status Status
}

func NewDeliverResp(
	sequenceID uint32, msgID [10]byte, status Status,
) (*DeliverResp, error) {
	op := &DeliverResp{}

	op.Header = &Header{}
	var length uint32 = 4 + 4 + 4

	op.MsgID = &OctetString{Data: msgID[:], FixedLen: 10}
	length = length + 10

	op.Status = status
	length = length + 4

	op.PacketLength = length
	op.RequestID = SMGP_DELIVER_RESP
	op.SequenceID = sequenceID

	return op, nil
}

func ParseDeliverResp(hdr *Header, data []byte) (*DeliverResp, error) {
	op := &DeliverResp{}
	op.Header = hdr

	p := 0
	op.MsgID = &OctetString{Data: data[p : p+10], FixedLen: 10}
	p = p + 10

	op.Status = Status(unpackUi32(data[p : p+4]))
	p = p + 4

	return op, nil
}

func (op *DeliverResp) Serialize() []byte {
	b := op.Header.Serialize()

	b = append(b, op.MsgID.Byte()...)
	b = append(b, packUi32(op.Status.Data())...)

	return b
}

func (op *DeliverResp) String() string {
	var b bytes.Buffer
	b.WriteString(op.Header.String())

	fmt.Fprintln(&b, "--- DeliverResp ---")
	fmt.Fprintln(&b, "MsgID: ", op.MsgID)
	fmt.Fprintln(&b, "Status: ", op.Status)

	return b.String()
}

func (op *DeliverResp) GetStatus() Status {
	return op.Status
}
