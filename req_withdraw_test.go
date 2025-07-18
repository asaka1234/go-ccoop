package go_ccoop

import (
	"fmt"
	"testing"
)

func TestWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &CCoopInitParams{Merchant_ID, SECRET_KEY, IP, BASE_URL, QRCode_Url, Deposit_CallBack_Url, DepositFeBackUrl, WithdrawBackUrl, WithdrawFeBackUrl})
	cli.SetDebugModel(true)

	//发请求
	resp, err := cli.Withdraw(GenWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() CCoopWithdrawRequest {
	return CCoopWithdrawRequest{
		Withdraw:  "700",
		OrderNum:  "9438965", //Unique customer id in your system. 业务系统里的唯一客户id
		OrderName: "你好",
		Name:      "TBT,12345",
	}
}
