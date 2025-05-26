package go_ccoop

import (
	"crypto/tls"
	"github.com/asaka1234/go-ccoop/utils"
	"github.com/mitchellh/mapstructure"
	"time"
)

func (cli *Client) Deposit(req CCoopDepositRequest) (*CCoopDepositResponse, error) {

	rawURL := cli.Deposit_Url

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["order_status"] = "0"     //写死
	params["trade_type"] = "deposit" //写死
	params["create_time"] = time.Now().Format("2006-01-02")
	params["mer_id"] = cli.Merchant_ID
	params["merchant"] = cli.Merchant_Title
	params["name"] = cli.Merchant_Name
	params["currency_amo"] = cli.Exchange_To_Currency
	params["callback_url"] = cli.Deposit_CallBack_Url
	params["return_url"] = cli.DepositReturn_Url

	// Generate signature
	signStr := utils.GenSign(params, cli.CallBack_Key)
	params["ref1"] = signStr //TODO  要重点关注下算签名算法, 感觉之前c#实现的不对

	//返回值会放到这里
	var result CCoopDepositResponse

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(req).
		SetHeaders(getHeaders()).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	//fmt.Printf("result: %s\n", string(resp.Body()))

	if err != nil {
		return nil, err
	}

	return &result, err
}
