package go_ccoop

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositCallback(req CCoopDepositBackReq, processor func(CCoopDepositBackReq) error) error {
	//TODO 验证签名

	//开始处理
	return processor(req)
}
