package protocol

import (
	"bytes"
	"errors"
	"fmt"
)

type Submit struct {
	*Header

	MsgType         uint8          // 短消息类型
	NeedReport      uint8          // SP是否要求返回状态报告
	Priority        uint8          // 短消息发送优先级
	ServiceID       *OctetString   // 业务代码
	FeeType         *OctetString   // 收费类型
	FeeCode         *OctetString   // 资费代码
	FixedFee        *OctetString   // 包月费/封顶费
	MsgFormat       uint8          // 短消息格式
	ValidTime       *OctetString   // 短消息有效时间
	AtTime          *OctetString   // 短消息定时发送时间
	SrcTermID       *OctetString   // 短信息发送方号码
	ChargeTermID    *OctetString   // 计费用户号码
	DestTermIDCount uint8          // 短消息接收号码总数
	DestTermID      []*OctetString // 短消息接收号码
	MsgLength       uint8          // 短消息长度
	MsgContent      []byte         // 短消息内容
	Reserve         *OctetString   // 保留

	// 可选参数
	Options Options
}

func NewSubmit(
	sequenceID uint32,
	needReport, priority uint8,
	serviceID, feeType, feeCode, fixedFee string,
	msgFormat uint8,
	validTime, atTime, srcTermID, chargeTermID string,
	destTermID []string, msgContent []byte,

	// 可选参数,  使用slice类型： nil可以区分是否设置此字段.
	// Integer 和 OctetString 由上层转换为[]byte
	TP_pid, TP_udhi []byte,
	LinkID []byte,
	MsgSrc []byte,
	ChargeUserType, ChargeTermType, ChargeTermPseudo []byte,
	DestTermType, DestTermPseudo []byte,
	PkTotal, PKNumber []byte,
	SubmitMsgType, SPDealResult, MServiceID []byte,

) (*Submit, error) {
	op := &Submit{}

	op.Header = &Header{}
	var length uint32 = 4 + 4 + 4 // header length

	op.MsgType = MT
	length = length + 1

	op.NeedReport = needReport
	length = length + 1

	op.Priority = priority
	length = length + 1

	op.ServiceID = &OctetString{Data: []byte(serviceID), FixedLen: 10}
	length = length + 10

	op.FeeType = &OctetString{Data: []byte(feeType), FixedLen: 2}
	length = length + 2

	op.FeeCode = &OctetString{Data: []byte(feeCode), FixedLen: 6}
	length = length + 6

	op.FixedFee = &OctetString{Data: []byte(fixedFee), FixedLen: 6}
	length = length + 6

	op.MsgFormat = msgFormat
	length = length + 1

	op.ValidTime = &OctetString{Data: []byte(validTime), FixedLen: 17}
	length = length + 17

	op.AtTime = &OctetString{Data: []byte(atTime), FixedLen: 17}
	length = length + 17

	op.SrcTermID = &OctetString{Data: []byte(srcTermID), FixedLen: 21}
	length = length + 21

	op.ChargeTermID = &OctetString{Data: []byte(chargeTermID), FixedLen: 21}
	length = length + 21

	// 短消息接收号码总数
	destTermIDCount := len(destTermID)
	if destTermIDCount > 100 {
		return nil, errors.New("too many destTermID")
	}
	op.DestTermIDCount = uint8(destTermIDCount)
	length = length + 1

	for _, v := range destTermID {
		op.DestTermID = append(
			op.DestTermID, &OctetString{Data: []byte(v), FixedLen: 21})
		length = length + 21
	}

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
	if MsgSrc != nil {
		op.Options[TAG_MsgSrc] = MsgSrc
	}
	if ChargeUserType != nil {
		op.Options[TAG_ChargeUserType] = ChargeUserType
	}
	if ChargeTermType != nil {
		op.Options[TAG_ChargeTermType] = ChargeTermType
	}
	if ChargeTermPseudo != nil {
		op.Options[TAG_ChargeTermPseudo] = ChargeTermPseudo
	}
	if DestTermType != nil {
		op.Options[TAG_DestTermType] = DestTermType
	}
	if DestTermPseudo != nil {
		op.Options[TAG_DestTermPseudo] = DestTermPseudo
	}
	if PkTotal != nil {
		op.Options[TAG_PkTotal] = PkTotal
	}
	if PKNumber != nil {
		op.Options[TAG_PkNumber] = PKNumber
	}
	if SubmitMsgType != nil {
		op.Options[TAG_SubmitMsgType] = SubmitMsgType
	}
	if SPDealResult != nil {
		op.Options[TAG_SPDealResult] = SPDealResult
	}
	if MServiceID != nil {
		op.Options[TAG_MServiceID] = MServiceID
	}
	// 可选字段长度
	length = length + uint32(op.Options.Len())

	op.PacketLength = length
	op.RequestID = SMGP_SUBMIT
	op.SequenceID = sequenceID

	return op, nil
}

func ParseSubmit(hdr *Header, data []byte) (*Submit, error) {
	p := 0
	op := &Submit{}
	op.Header = hdr

	op.MsgType = data[p]
	p = p + 1
	op.NeedReport = data[p]
	p = p + 1
	op.Priority = data[p]
	p = p + 1

	op.ServiceID = &OctetString{Data: data[p : p+10], FixedLen: 10}
	p = p + 10
	op.FeeType = &OctetString{Data: data[p : p+2], FixedLen: 2}
	p = p + 2
	op.FeeCode = &OctetString{Data: data[p : p+6], FixedLen: 6}
	p = p + 6
	op.FixedFee = &OctetString{Data: data[p : p+6], FixedLen: 6}
	p = p + 6

	op.MsgFormat = data[p]
	p = p + 1

	op.ValidTime = &OctetString{Data: data[p : p+17], FixedLen: 17}
	p = p + 17
	op.AtTime = &OctetString{Data: data[p : p+17], FixedLen: 17}
	p = p + 17

	op.SrcTermID = &OctetString{Data: data[p : p+21], FixedLen: 21}
	p = p + 21
	op.ChargeTermID = &OctetString{Data: data[p : p+21], FixedLen: 21}
	p = p + 21

	op.DestTermIDCount = data[p]
	p = p + 1

	for i := 0; i < int(op.DestTermIDCount); i++ {
		op.DestTermID = append(
			op.DestTermID, &OctetString{Data: data[p : p+21], FixedLen: 21})
		p = p + 21
	}

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

func (op *Submit) Serialize() []byte {
	b := op.Header.Serialize()

	b = append(b, packUi8(op.MsgType)...)
	b = append(b, packUi8(op.NeedReport)...)
	b = append(b, packUi8(op.Priority)...)

	b = append(b, op.ServiceID.Byte()...)
	b = append(b, op.FeeType.Byte()...)
	b = append(b, op.FeeCode.Byte()...)
	b = append(b, op.FixedFee.Byte()...)

	b = append(b, packUi8(op.MsgFormat)...)

	b = append(b, op.ValidTime.Byte()...)
	b = append(b, op.AtTime.Byte()...)
	b = append(b, op.SrcTermID.Byte()...)
	b = append(b, op.ChargeTermID.Byte()...)

	b = append(b, packUi8(op.DestTermIDCount)...)

	for i := 0; i < int(op.DestTermIDCount); i++ {
		b = append(b, op.DestTermID[i].Byte()...)
	}

	b = append(b, packUi8(op.MsgLength)...)

	b = append(b, op.MsgContent...)
	b = append(b, op.Reserve.Byte()...)

	// 可选字段
	b = append(b, op.Options.Serialize()...)

	return b
}

func (op *Submit) String() string {
	var b bytes.Buffer
	b.WriteString(op.Header.String())

	fmt.Fprintln(&b, "--- Submit ---")
	fmt.Fprintln(&b, "MsgType: ", op.MsgType)
	fmt.Fprintln(&b, "NeedReport: ", op.NeedReport)
	fmt.Fprintln(&b, "Priority: ", op.Priority)

	fmt.Fprintln(&b, "ServiceID: ", op.ServiceID)
	fmt.Fprintln(&b, "FeeType: ", op.FeeType)
	fmt.Fprintln(&b, "FeeCode: ", op.FeeCode)
	fmt.Fprintln(&b, "FixedFee: ", op.FixedFee)

	fmt.Fprintln(&b, "MsgFormat: ", op.MsgFormat)
	fmt.Fprintln(&b, "ValidTime: ", op.ValidTime)
	fmt.Fprintln(&b, "AtTime: ", op.AtTime)
	fmt.Fprintln(&b, "SrcTermID: ", op.SrcTermID)
	fmt.Fprintln(&b, "ChargeTermID: ", op.ChargeTermID)

	fmt.Fprintln(&b, "DestTermIDCount: ", op.DestTermIDCount)
	for i := 0; i < int(op.DestTermIDCount); i++ {
		fmt.Fprintln(&b, "DestTermID: ", op.DestTermID[i])
	}

	fmt.Fprintln(&b, "MsgLength: ", op.MsgLength)
	fmt.Fprintln(&b, "MsgContent: ", string(op.MsgContent))

	return b.String()
}

func (op *Submit) GetStatus() Status {
	return STAT_OK
}

type SubmitResp struct {
	*Header

	MsgID  *Id
	Status Status
}

func NewSubmitResp(
	sequenceID uint32, msgID [10]byte, status Status,
) (*SubmitResp, error) {
	op := &SubmitResp{}

	op.Header = &Header{}
	var length uint32 = 4 + 4 + 4

	op.MsgID = &Id{raw: msgID}
	length = length + 10

	op.Status = status
	length = length + 4

	op.PacketLength = length
	op.RequestID = SMGP_SUBMIT_RESP
	op.SequenceID = sequenceID

	return op, nil
}

func ParseSubmitResp(hdr *Header, data []byte) (*SubmitResp, error) {
	op := &SubmitResp{}
	op.Header = hdr

	p := 0
	op.MsgID = newId(data[p : p+10])
	p = p + 10

	op.Status = Status(unpackUi32(data[p : p+4]))
	p = p + 4

	return op, nil
}

func (op *SubmitResp) Serialize() []byte {
	b := op.Header.Serialize()

	b = append(b, op.MsgID.raw[:]...)
	b = append(b, packUi32(op.Status.Data())...)

	return b
}

func (op *SubmitResp) String() string {
	var b bytes.Buffer
	b.WriteString(op.Header.String())

	fmt.Fprintln(&b, "--- SubmitResp ---")
	fmt.Fprintln(&b, "MsgID: ", op.MsgID)
	fmt.Fprintln(&b, "Status: ", op.Status)

	return b.String()
}

func (op *SubmitResp) GetStatus() Status {
	return op.Status
}
