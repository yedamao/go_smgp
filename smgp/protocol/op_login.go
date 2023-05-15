package protocol

import (
	"bytes"
	"fmt"
)

// Login操作的目的是客户端向服务器端注册作为一个
// 合法客户端身份，若注册成功后即建立了应用层的连接，
// 此后客户端可以与此服务器端进行消息的接收和发送。
type Login struct {
	*Header

	ClientID            *OctetString
	AuthenticatorClient *OctetString
	LoginMode           uint8
	TimeStamp           uint32
	ClientVersion       uint8
}

func NewLogin(
	sequenceID uint32, clientID, sharedSecret string, mode uint8,
) (*Login, error) {

	// genAuthenticatorClient
	authClient, err := genAuthenticatorClient(clientID, sharedSecret, genTimestamp())
	if err != nil {
		return nil, err
	}

	op := &Login{}

	op.Header = &Header{}
	var length uint32 = 4 + 4 + 4 // header length

	op.ClientID = &OctetString{Data: []byte(clientID), FixedLen: 8}
	length = length + 8

	op.AuthenticatorClient = &OctetString{Data: []byte(authClient), FixedLen: 16}
	length = length + 16

	op.LoginMode = mode
	length = length + 1

	op.TimeStamp = genTimestamp()
	length = length + 4

	op.ClientVersion = VERSION
	length = length + 1

	op.PacketLength = length
	op.RequestID = SMGP_LOGIN
	op.SequenceID = sequenceID

	return op, nil
}

func ParseLogin(hdr *Header, data []byte) (*Login, error) {
	p := 0
	op := &Login{}
	op.Header = hdr

	op.ClientID = &OctetString{Data: data[p : p+8], FixedLen: 8}
	p = p + 8

	op.AuthenticatorClient = &OctetString{Data: data[p : p+16], FixedLen: 16}
	p = p + 16

	op.LoginMode = data[p]
	p = p + 1

	op.TimeStamp = unpackUi32(data[p : p+4])
	p = p + 4

	op.ClientVersion = data[p]
	p = p + 1

	return op, nil
}

func (l *Login) Serialize() []byte {
	b := l.Header.Serialize()

	b = append(b, l.ClientID.Byte()...)
	b = append(b, l.AuthenticatorClient.Byte()...)
	b = append(b, packUi8(l.LoginMode)...)
	b = append(b, packUi32(l.TimeStamp)...)
	b = append(b, packUi8(l.ClientVersion)...)

	return b
}

func (l *Login) String() string {
	var b bytes.Buffer
	b.WriteString(l.Header.String())

	fmt.Fprintln(&b, "--- Login ---")
	fmt.Fprintln(&b, "ClientID: ", l.ClientID)
	fmt.Fprintln(&b, "AuthenticatorClient: ", l.AuthenticatorClient)
	fmt.Fprintln(&b, "LoginMode: ", l.LoginMode)
	fmt.Fprintln(&b, "TimeStamp: ", l.TimeStamp)
	fmt.Fprintf(&b, "ClientVersion: 0x%x\n", l.ClientVersion)

	return b.String()
}

func (l *Login) GetStatus() Status {
	return STAT_OK
}

type LoginResp struct {
	*Header

	Status              Status       // 请求返回结果
	AuthenticatorServer *OctetString // 服务器端返回给客户端的认证码
	ServerVersion       uint8        // 服务器端支持的最高版本号
}

func NewLoginResp(
	sequenceID uint32, status Status,
	authServer string, serverVersion uint8,
) (*LoginResp, error) {
	op := &LoginResp{}

	op.Header = &Header{}
	var length uint32 = 4 + 4 + 4

	op.Status = status
	length = length + 4

	op.AuthenticatorServer = &OctetString{Data: []byte(authServer), FixedLen: 16}
	length = length + 16

	op.ServerVersion = serverVersion
	length = length + 1

	op.PacketLength = length
	op.RequestID = SMGP_LOGIN_RESP
	op.SequenceID = sequenceID

	return op, nil
}

func ParseLoginResp(hdr *Header, data []byte) (*LoginResp, error) {
	op := &LoginResp{}
	op.Header = hdr

	p := 0
	op.Status = Status(unpackUi32(data[p : p+4]))
	p = p + 4

	op.AuthenticatorServer = &OctetString{Data: data[p : p+16], FixedLen: 16}
	p = p + 16

	op.ServerVersion = data[p]
	p = p + 1

	return op, nil
}

func (l *LoginResp) Serialize() []byte {
	b := l.Header.Serialize()

	b = append(b, packUi32(l.Status.Data())...)
	b = append(b, l.AuthenticatorServer.Byte()...)
	b = append(b, packUi8(l.ServerVersion)...)

	return b
}

func (l *LoginResp) String() string {
	var b bytes.Buffer
	b.WriteString(l.Header.String())

	fmt.Fprintln(&b, "--- LoginResp ---")
	fmt.Fprintln(&b, "Status: ", l.Status)
	fmt.Fprintln(&b, "AuthenticatorServer: ", l.AuthenticatorServer)
	fmt.Fprintf(&b, "ServerVersion: 0x%x\n", l.ServerVersion)

	return b.String()
}

func (l *LoginResp) GetStatus() Status {
	return l.Status
}
