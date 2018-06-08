package protocol

import (
	"testing"
)

func TestId(t *testing.T) {
	id := &Id{
		raw: [10]byte{
			0x01, 0x00,
			0x61, 0x01,
			0x16, 0x17,
			0x00, 0x01,
			0x23, 0x45,
		},
	}

	if "01006101161700012345" != id.String() {
		t.Error("id string not match")
	}
}
