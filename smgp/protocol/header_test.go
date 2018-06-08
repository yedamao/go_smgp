package protocol

import (
	"testing"
)

func TestParseHeader(t *testing.T) {

	h := &Header{
		PacketLength: 99,
		RequestID:    SMGP_LOGIN,
		SequenceID:   1,
	}

	parsed, _ := ParseHeader(h.Serialize())

	if h.PacketLength != parsed.PacketLength ||
		h.RequestID != parsed.RequestID ||
		h.SequenceID != parsed.SequenceID {
		t.Error("header parsed not equal")
	}
}
