package go_ccoop

import "errors"

func (cli *Client) WithdrawCallback(req CCoopWithdrawBackReq, processor func(CCoopWithdrawBackReq) error) error {
	if req.TradeType != "withdraw" {
		return errors.New("trade type error")
	}
	if req.MerID != cli.Params.MerchantId {
		return errors.New("wrong merchant id")
	}

	//开始处理
	return processor(req)
}
