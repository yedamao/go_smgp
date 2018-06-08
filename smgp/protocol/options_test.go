package protocol

import (
	"testing"
)

func TestOptions(t *testing.T) {
	var (
		udhi    uint16 = 1
		pkTotal uint16 = 10
	)
	ops := make(Options)
	ops[TAG_TP_udhi] = packUi16(udhi)
	ops[TAG_PkTotal] = packUi16(pkTotal)

	serialized := ops.Serialize()

	if int(len(serialized)) != ops.Len() {
		t.Error("serialized len not match")
	}

	parsed, err := ParseOptions(serialized)
	if err != nil {
		t.Error()
	}

	if udhi != unpackUi16(parsed[TAG_TP_udhi]) ||
		pkTotal != unpackUi16(parsed[TAG_PkTotal]) {
		t.Error("value not match")
	}
}
