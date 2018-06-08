package protocol

type ActiveTest struct {
	*Header
}

func NewActiveTest(sequenceID uint32) (*ActiveTest, error) {

	hdr := &Header{
		PacketLength: 4 + 4 + 4,
		RequestID:    SMGP_ACTIVE_TEST,
		SequenceID:   sequenceID,
	}

	return &ActiveTest{Header: hdr}, nil
}

func ParseActiveTest(hdr *Header, data []byte) (*ActiveTest, error) {
	return &ActiveTest{Header: hdr}, nil
}

func (op *ActiveTest) Serialize() []byte {
	return op.Header.Serialize()
}

func (op *ActiveTest) String() string {
	return op.Header.String()
}

func (op *ActiveTest) GetStatus() Status {
	return STAT_OK
}

type ActiveTestResp struct {
	*Header
}

func NewActiveTestResp(sequenceID uint32) (*ActiveTestResp, error) {

	hdr := &Header{
		PacketLength: 4 + 4 + 4,
		RequestID:    SMGP_ACTIVE_TEST_RESP,
		SequenceID:   sequenceID,
	}

	return &ActiveTestResp{Header: hdr}, nil
}

func ParseActiveTestResp(hdr *Header, data []byte) (*ActiveTestResp, error) {
	return &ActiveTestResp{Header: hdr}, nil
}

func (op *ActiveTestResp) Serialize() []byte {
	return op.Header.Serialize()
}

func (op *ActiveTestResp) String() string {
	return op.Header.String()
}

func (op *ActiveTestResp) GetStatus() Status {
	return STAT_OK
}

type Exit struct {
	*Header
}

func NewExit(sequenceID uint32) (*Exit, error) {

	hdr := &Header{
		PacketLength: 4 + 4 + 4,
		RequestID:    SMGP_EXIT,
		SequenceID:   sequenceID,
	}

	return &Exit{Header: hdr}, nil
}

func ParseExit(hdr *Header, data []byte) (*Exit, error) {
	return &Exit{Header: hdr}, nil
}

func (op *Exit) Serialize() []byte {
	return op.Header.Serialize()
}

func (op *Exit) String() string {
	return op.Header.String()
}

func (op *Exit) GetStatus() Status {
	return STAT_OK
}

type ExitResp struct {
	*Header
}

func NewExitResp(sequenceID uint32) (*ExitResp, error) {

	hdr := &Header{
		PacketLength: 4 + 4 + 4,
		RequestID:    SMGP_EXIT_RESP,
		SequenceID:   sequenceID,
	}

	return &ExitResp{Header: hdr}, nil
}

func ParseExitResp(hdr *Header, data []byte) (*ExitResp, error) {
	return &ExitResp{Header: hdr}, nil
}

func (op *ExitResp) Serialize() []byte {
	return op.Header.Serialize()
}

func (op *ExitResp) String() string {
	return op.Header.String()
}

func (op *ExitResp) GetStatus() Status {
	return STAT_OK
}
