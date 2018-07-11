package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/yedamao/encoding"

	"github.com/yedamao/go_smgp/smgp"
	"github.com/yedamao/go_smgp/smgp/protocol"
)

var (
	addr         = flag.String("addr", ":8890", "smgw addr(运营商地址)")
	clientID     = flag.String("clientID", "", "登陆账号")
	sharedSecret = flag.String("secret", "", "登陆密码")
	spID         = flag.String("spID", "", "企业代码")

	spNumber   = flag.String("sp-number", "", "SP的接入号码")
	destNumber = flag.String("dest-number", "", "接收手机号码, 86..., 多个使用，分割")
	msg        = flag.String("msg", "", "短信内容")
)

func init() {
	flag.Parse()
}

var sequenceID uint32

func newSeqNum() uint32 {
	sequenceID++

	return sequenceID
}

func main() {

	if "" == *clientID || "" == *sharedSecret {
		fmt.Println("Arg error: clientID or sharedSecret must not be empty .")
		os.Exit(-1)
	}

	destNumbers := strings.Split(*destNumber, ",")
	fmt.Println("destNumbers: ", destNumbers)

	ts, err := smgp.NewSmgp(
		*addr, *clientID, *sharedSecret, *spID, protocol.SEND_MODE, newSeqNum,
	)
	if err != nil {
		fmt.Println("Connection Err", err)
		os.Exit(-1)
	}
	fmt.Println("connect succ")

	// encoding msg
	content := encoding.UTF82GBK([]byte(*msg))
	if len(content) > 140 {
		fmt.Println("msg Err: not suport long sms")
	}

	// send
	_, err = ts.Submit(
		protocol.NEED_REPORT, protocol.NORMAL_PRIORITY,
		"", "00", "0", "0",
		protocol.GB18030,
		"", "", *spNumber, "",
		destNumbers, content,

		0, 0, 0,
	)
	if err != nil {
		fmt.Println("Submit Err", err)
		ts.Close()
		os.Exit(-1)
	}

	for {

		// receive submit resp
		op, err := ts.Read() // This is blocking
		if err != nil {
			fmt.Println("Read Err:", err)
			os.Exit(-1)
		}
		fmt.Println(op)

		switch op.GetHeader().RequestID {
		case protocol.SMGP_SUBMIT_RESP:
			// message_id should match this with seq message
			fmt.Println("MSG ID:", op.GetHeader().SequenceID)
			ts.Exit()
		case protocol.SMGP_EXIT_RESP:
			fmt.Println("Exit response")
			break
		default:
			fmt.Printf("Unexpect RequestID: %0x\n", op.GetHeader().RequestID)
		}
	}
}
