package go_ccoop

import "errors"

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositCallback(req CCoopDepositBackReq, processor func(CCoopDepositBackReq) error) error {
	if req.TradeType != "deposit" {
		return errors.New("trade type error")
	}
	if req.MerID != cli.Params.MerchantId {
		return errors.New("wrong merchant id")
	}

	//开始处理
	return processor(req)
}
