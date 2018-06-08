package protocol

import (
	"bytes"
	"errors"
	"fmt"
)

type Report struct {
	Id         Id
	Sub        string
	Dlvrd      string
	SubmitDate string
	DoneDate   string
	Stat       string
	Err        string
	Txt        string
}

var ErrReportLen = errors.New("ParseReport: Length error")

//
// 根据协议状态报告格式
// 严格按照 字节位置解析
//
// warning: 如果运营商网关实现不规范, 可考虑更换为 匹配字段名解析
func ParseReport(data []byte) (*Report, error) {

	if 122 != len(data) {
		return nil, ErrReportLen
	}

	report := &Report{}

	// format
	// id:IIIIIIIIII sub:SSS dlvrd:DDD Submit date:YYMMDDhhmm done date:YYMMDDhhmm stat:DDDDDDD err:E Text:……
	p := 0

	// Id
	p = p + 3
	copy(report.Id.raw[:], data[p:p+10])
	p = p + 10 + 1

	// sub
	p = p + 4
	report.Sub = string(data[p : p+3])
	p = p + 3 + 1

	// dlvrd
	p = p + 6
	report.Dlvrd = string(data[p : p+3])
	p = p + 3 + 1

	// submit date
	p = p + 12
	report.SubmitDate = string(data[p : p+10])
	p = p + 10 + 1

	// done date
	p = p + 10
	report.DoneDate = string(data[p : p+10])
	p = p + 10 + 1

	// stat
	p = p + 5
	report.Stat = string(data[p : p+7])
	p = p + 7 + 1

	// err
	p = p + 4
	report.Err = string(data[p : p+3])
	p = p + 3 + 1

	// text
	p = p + 5
	report.Txt = string(data[p:])

	return report, nil
}

func NewReport(
	id [10]byte,
	sub, dlvrd, submitDate, doneDate, stat, err, txt string,
) []byte {

	pId := &OctetString{Data: id[:], FixedLen: 10}
	pSub := &OctetString{Data: []byte(sub), FixedLen: 3}
	pDlvrd := &OctetString{Data: []byte(dlvrd), FixedLen: 3}
	pSubmitDate := &OctetString{Data: []byte(submitDate), FixedLen: 10}
	pDoneDate := &OctetString{Data: []byte(doneDate), FixedLen: 10}
	pStat := &OctetString{Data: []byte(stat), FixedLen: 7}
	pErr := &OctetString{Data: []byte(err), FixedLen: 3}
	pTxt := &OctetString{Data: []byte(txt), FixedLen: 20}

	report := append([]byte("id:"), pId.Byte()...)

	report = append(report, []byte(" sub:")...)
	report = append(report, pSub.Byte()...)

	report = append(report, []byte(" dlvrd:")...)
	report = append(report, pDlvrd.Byte()...)

	report = append(report, []byte(" Submit date:")...)
	report = append(report, pSubmitDate.Byte()...)

	report = append(report, []byte(" done date:")...)
	report = append(report, pDoneDate.Byte()...)

	report = append(report, []byte(" stat:")...)
	report = append(report, pStat.Byte()...)

	report = append(report, []byte(" err:")...)
	report = append(report, pErr.Byte()...)

	report = append(report, []byte(" Text:")...)
	report = append(report, pTxt.Byte()...)

	return report
}

func (r *Report) String() string {
	var b bytes.Buffer

	fmt.Println(&b, "id: ", r.Id)
	fmt.Println(&b, "sub: ", r.Sub)
	fmt.Println(&b, "dlvrd: ", r.Dlvrd)
	fmt.Println(&b, "submitDate: ", r.SubmitDate)
	fmt.Println(&b, "doneDate: ", r.DoneDate)
	fmt.Println(&b, "stat: ", r.Stat)
	fmt.Println(&b, "err: ", r.Err)
	fmt.Println(&b, "Text: ", r.Txt)

	return b.String()
}
