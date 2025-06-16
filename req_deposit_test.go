package go_ccoop

import (
	"fmt"
	"testing"
)

func TestGenQRCode(t *testing.T) {

	//构造client
	cli := NewClient(nil, &CCoopInitParams{Merchant_ID, SECRET_KEY, IP, BASE_URL, QRCode_Url, Deposit_CallBack_Url, DepositFeBackUrl, WithdrawBackUrl, WithdrawFeBackUrl})
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
		Deposit:   "100",
		OrderNum:  "9438965111", //Unique customer id in your system. 业务系统里的唯一客户id
		OrderName: "你好",
	}
}
