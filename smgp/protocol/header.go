package protocol

import (
	"bytes"
	"fmt"
)

// 消息头(所有消息公共包头)
type Header struct {
	PacketLength uint32 // 数据包长度
	RequestID    uint32 // 请求标识
	SequenceID   uint32 // 消息流水号
}

func (p *Header) GetHeader() *Header {
	return p
}

func (p *Header) Serialize() []byte {
	b := packUi32(p.PacketLength)
	b = append(b, packUi32(p.RequestID)...)
	b = append(b, packUi32(p.SequenceID)...)

	return b
}

func (p *Header) String() string {
	var b bytes.Buffer
	fmt.Fprintln(&b, "--- Header ---")
	fmt.Fprintln(&b, "Length: ", p.PacketLength)
	fmt.Fprintf(&b, "CmdId: 0x%x\n", p.RequestID)
	fmt.Fprintln(&b, "Sequence: ", p.SequenceID)

	return b.String()

}

func (p *Header) Parse(data []byte) *Header {

	p.PacketLength = unpackUi32(data[:4])
	p.RequestID = unpackUi32(data[4:8])
	p.SequenceID = unpackUi32(data[8:12])

	return p
}

func ParseHeader(data []byte) (*Header, error) {

	h := &Header{}
	h.Parse(data)

	return h, nil
}
