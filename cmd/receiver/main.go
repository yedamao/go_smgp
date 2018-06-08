package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/yedamao/go_smgp/smgp"
	"github.com/yedamao/go_smgp/smgp/protocol"
)

var (
	addr         = flag.String("addr", ":8890", "smgw addr(运营商地址)")
	clientID     = flag.String("clientID", "", "登陆账号")
	sharedSecret = flag.String("secret", "", "登陆密码")
)

func init() {
	flag.Parse()
}

var sequenceID uint32 = 0

func newSeqNum() uint32 {
	sequenceID++

	return sequenceID
}

func main() {
	if "" == *clientID || "" == *sharedSecret {
		fmt.Println("Arg error: clientID or sharedSecret must not be empty .")
		flag.Usage()
		os.Exit(-1)
	}

	rcv, err := smgp.NewSmgp(
		*addr, *clientID, *sharedSecret, "",
		protocol.RECEIVE_MODE, newSeqNum,
	)
	if err != nil {
		fmt.Println("Connection Err", err)
		os.Exit(-1)
	}
	fmt.Println("connect succ")

	for {

		rcv.SetDeadline(time.Now().Add(1e9))
		op, err := rcv.Read()
		if err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
				continue
			}
			log.Println("Receiver: read error exit. ", err)
			return
		}

		switch op.GetHeader().RequestID {
		case protocol.SMGP_DELIVER:
			dlv, ok := op.(*protocol.Deliver)
			if !ok {
				log.Println("Receiver: Deliver type assert error")
				return
			}
			log.Println(dlv)
			rcv.DeliverResp(op.GetHeader().SequenceID, dlv.MsgID.Raw(), protocol.STAT_OK)

		case protocol.SMGP_ACTIVE_TEST:
			rcv.ActiveTestResp(op.GetHeader().SequenceID)
			log.Println("ActiveTest ...")

		case protocol.SMGP_EXIT:
			rcv.ExitResp(op.GetHeader().SequenceID)
			log.Println("Receiver: smgw request Exit.")
			return

		case protocol.SMGP_EXIT_RESP:
			log.Println("Receiver: Exit.")
			return

		default:
			log.Printf("Unknow Operation CmdId: 0x%x",
				op.GetHeader().RequestID)
		}
	}

	fmt.Println("Done")
}
