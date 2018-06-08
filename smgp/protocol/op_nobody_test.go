package protocol

import (
	"testing"
)

func TestActiveTest(t *testing.T) {
	op, err := NewActiveTest(1234)
	if err != nil {
		t.Fatal(err)
	}

	parsed, err := ParseOperation(op.Serialize())
	if err != nil {
		t.Fatal(err)
	}

	activeTest := parsed.(*ActiveTest)
	if op.SequenceID != activeTest.SequenceID {
		t.Error("ActiveTest parsed not match")
	}
}

func TestActiveTestResp(t *testing.T) {
	op, err := NewActiveTestResp(1234)
	if err != nil {
		t.Fatal(err)
	}

	parsed, err := ParseOperation(op.Serialize())
	if err != nil {
		t.Fatal(err)
	}

	activeTest := parsed.(*ActiveTestResp)
	if op.SequenceID != activeTest.SequenceID {
		t.Error("ActiveTestResp parsed not match")
	}
}

func TestExit(t *testing.T) {
	op, err := NewExit(1234)
	if err != nil {
		t.Fatal(err)
	}

	parsed, err := ParseOperation(op.Serialize())
	if err != nil {
		t.Fatal(err)
	}

	activeTest := parsed.(*Exit)
	if op.SequenceID != activeTest.SequenceID {
		t.Error("Exit parsed not match")
	}
}

func TestExitResp(t *testing.T) {
	op, err := NewExitResp(1234)
	if err != nil {
		t.Fatal(err)
	}

	parsed, err := ParseOperation(op.Serialize())
	if err != nil {
		t.Fatal(err)
	}

	activeTest := parsed.(*ExitResp)
	if op.SequenceID != activeTest.SequenceID {
		t.Error("ExitResp parsed not match")
	}
}
