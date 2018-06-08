package protocol

import (
	"testing"
)

func TestStatus(t *testing.T) {
	status := STAT_OK
	if status.String() != "成功" {
		t.Error("status msg error")
	}
}
