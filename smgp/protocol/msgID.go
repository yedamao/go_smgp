package protocol

import (
	"encoding/hex"
)

// 短消息流水号类型
type Id struct {
	raw [10]byte
}

func (i Id) String() string {
	return hex.EncodeToString(i.raw[:])
}

func (i *Id) Raw() [10]byte {
	return i.raw
}

func newId(data []byte) *Id {
	id := &Id{}

	copy(id.raw[:], data[:10])

	return id
}
