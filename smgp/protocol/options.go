package protocol

import (
	"encoding/binary"
	"errors"
)

var ErrLength = errors.New("Options: error length")

type Tag uint16

// 可选参数标签定义  Option Tag
const (
	TAG_TP_pid Tag = 0x0001 + iota
	TAG_TP_udhi
	TAG_LinkID
	TAG_ChargeUserType
	TAG_ChargeTermType
	TAG_ChargeTermPseudo
	TAG_DestTermType
	TAG_DestTermPseudo
	TAG_PkTotal
	TAG_PkNumber
	TAG_SubmitMsgType
	TAG_SPDealResult
	TAG_SrcTermType
	TAG_SrcTermPseudo
	TAG_NodesCount
	TAG_MsgSrc
	TAG_SrcType
	TAG_MServiceID
)

// 可选参数 map
type Options map[Tag][]byte

// 可选参数采用TLV（Tag、Length、Value）形式定义，每个可选参数的Tag、Length、Value的定义见6.3节。
//
// Tag 2 Integer 字段的标签，用于唯一标识可选参数
// Length 2 Integer 字段的长度
// Value 可变长度 可变类型 字段内容
func ParseOptions(rawData []byte) (Options, error) {
	var (
		p      = 0
		ops    = make(Options)
		length = len(rawData)
	)

	for p < length {
		if length-p < 2+2 { // less than Tag len + Length len
			return nil, ErrLength
		}

		tag := binary.BigEndian.Uint16(rawData[p:])
		p += 2

		vlen := binary.BigEndian.Uint16(rawData[p:])
		p += 2

		if length-p < int(vlen) { // remaining not enough
			return nil, ErrLength
		}

		value := rawData[p : p+int(vlen)]
		p += int(vlen)

		ops[Tag(tag)] = value
	}

	return ops, nil
}

func (o Options) String() string {
	// TODO

	return "TODO: Options string"
}

// 返回可选字段部分的长度
func (o Options) Len() int {
	length := 0

	for _, v := range o {
		length += 2 + 2 + len(v)
	}

	return length
}

func (o Options) Serialize() []byte {
	var b []byte

	for k, v := range o {
		b = append(b, packUi16(uint16(k))...)
		b = append(b, packUi16(uint16(len(v)))...)
		b = append(b, v...)
	}

	return b
}

func (o Options) TP_pid() uint8 {
	return o[TAG_TP_pid][0]
}

func (o Options) TP_udhi() uint8 {
	return o[TAG_TP_udhi][0]
}

func (o Options) LinkID() string {
	p := &OctetString{Data: o[TAG_LinkID], FixedLen: 20}
	return p.String()
}

func (o Options) ChargeUserType() uint8 {
	return o[TAG_ChargeUserType][0]
}

func (o Options) ChargeTermType() uint8 {
	return o[TAG_ChargeTermType][0]
}

func (o Options) ChargeTermPseudo() string {
	value := o[TAG_ChargeTermPseudo]
	p := &OctetString{Data: value, FixedLen: int(len(value))}
	return p.String()
}

func (o Options) MsgSrc() string {
	value := o[TAG_MsgSrc]
	p := &OctetString{Data: value, FixedLen: 8}
	return p.String()
}

func (o Options) MServiceID() string {
	value := o[TAG_MServiceID]
	p := &OctetString{Data: value, FixedLen: 21}
	return p.String()
}
