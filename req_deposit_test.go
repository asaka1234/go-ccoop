package go_ccoop

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestGenQRCode(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &CCoopInitParams{Merchant_ID, SECRET_KEY, IP, BASE_URL, QRCode_Url, Deposit_CallBack_Url, DepositFeBackUrl, WithdrawBackUrl})
	cli.SetDebugModel(true)
	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenDepositRequestDemo() CCoopDepositRequest {
	return CCoopDepositRequest{
		Deposit:   "1",
		OrderNum:  "9438965111222", //Unique customer id in your system. 业务系统里的唯一客户id
		OrderName: "你好",
	}
}
