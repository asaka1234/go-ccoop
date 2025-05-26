package go_ccoop

import (
	"crypto/tls"
	"errors"
)

// 退款
func (cli *Client) Withdraw(req One2PayWithdrawRequest) (*One2PayWithdrawResponse, error) {

	rawURL := cli.WithdrawURL

	//返回值会放到这里
	var result One2PayWithdrawResponse

	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(req).
		SetHeaders(getAuthHeaders(cli.PartnerCode, cli.AuthKey, cli.Channel, cli.Device)).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	if err != nil {
		return nil, err
	}

	//-----------错误处理------------------------
	if resp.StatusCode() != 1000 {
		if result.Error != "" {
			return nil, errors.New(result.Error)
		}
		return nil, errors.New(result.Message)
	}

	return &result, err
}
