package smgptest

import (
	"testing"
)

func TestServer(t *testing.T) {
	_, err := NewServer(":8890")
	if err != nil {
		t.Fatal(err)
	}
}
