package go_ccoop

import (
	"fmt"
	"testing"
)

func TestWithdraw(t *testing.T) {

	cc := One2payDepositStatusSuccess
	fmt.Printf("==wsx====>%d\n", int(cc))

	//构造client
	cli := NewClient(nil, PARTNER_CODE, AUTH_KEY, DEVICE, CHANNEL, DEPOSIT_URL, PAYOUT_URL)

	//发请求
	resp, err := cli.Withdraw(GenWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() One2PayWithdrawRequest {
	return One2PayWithdrawRequest{
		BankAcc:       "0652078409",
		BankCode:      "004",
		BankName:      "KASIKORN BANK",
		AccountName:   "Manop Tangngam",
		Amount:        1000.50,
		MobileNo:      "0805933181",
		TransactionBy: "Jack Developer",
		Ref1:          "123456789012345678",
	}
}
