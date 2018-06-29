package protocol

import (
	"testing"
)

func TestStatus(t *testing.T) {
	status := STAT_OK
	if status.String() != "成功" {
		t.Error("status msg error")
	}

	status = 16
	if status.String() != "Status Unknown: 16" {
		t.Error("status msg error")
	}
}
