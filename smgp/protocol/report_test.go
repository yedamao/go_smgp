package protocol

import (
	"testing"
)

func TestParseReport(t *testing.T) {
	oneReport := NewReport([10]byte{}, "001", "001", "1805301557", "1805301557", "DELIVRD", "000", "")

	tests := [][]byte{
		oneReport,

		[]byte("id:\020\020\005\005\060\025W\222c9 sub:001 dlvrd:001 submit date:1805301557 done date:1805301557 stat:DELIVRD err:000 Text:fillfillfillfillfill"),
		[]byte("id:\005p\023\005\060\025PVu\001sSub:001sDlvrd:000sSubmit_Date:1805301550sDone_Date:1805301550sStat:DELIVRDsErr:000sText:fillfillfillfillfill"),
	}

	for _, test := range tests {
		report, err := ParseReport(test)
		if err != nil {
			t.Error(err)
		}

		if "DELIVRD" != report.Stat {
			t.Error("parsed report stat not match")
		}
	}
}
