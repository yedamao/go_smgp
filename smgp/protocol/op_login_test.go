package protocol

import (
	"testing"
)

func TestLogin(t *testing.T) {
	op, err := NewLogin(1, "testid", "testauth", SEND_MODE)
	if err != nil {
		t.Error(err)
	}

	parsed, err := ParseOperation(op.Serialize())
	if err != nil {
		t.Fatal(err)
	}

	parsedLogin := parsed.(*Login)

	if parsedLogin.ClientID.String() != op.ClientID.String() ||
		parsedLogin.AuthenticatorClient.String() != op.AuthenticatorClient.String() ||
		parsedLogin.LoginMode != op.LoginMode ||
		parsedLogin.TimeStamp != op.TimeStamp ||
		parsedLogin.ClientVersion != op.ClientVersion {
		t.Error("parsedLogin not equal")
	}
}

func TestLoginResp(t *testing.T) {
	op, err := NewLoginResp(1, 0, "AuthServer", VERSION)
	if err != nil {
		t.Error(err)
	}

	parsed, err := ParseOperation(op.Serialize())
	if err != nil {
		t.Fatal(err)
	}

	parsedLoginResp := parsed.(*LoginResp)

	if parsedLoginResp.Status != op.Status ||
		parsedLoginResp.AuthenticatorServer.String() != op.AuthenticatorServer.String() ||
		parsedLoginResp.ServerVersion != op.ServerVersion {
		t.Error("parsedLoginResp not equal")
	}
}
