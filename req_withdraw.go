package go_ccoop

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-ccoop/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
	"time"
)

func (cli *Client) Withdraw(req CCoopWithdrawRequest) (*CCoopWithdrawResponse, error) {

	rawURL := cli.Params.BaseUrl

	cTime := time.Now()

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["order_status"] = "0"      //写死
	params["trade_type"] = "withdraw" //写死
	params["create_time"] = cTime.Format("2006-01-02")
	params["mer_id"] = cli.Params.MerchantId            //cli.merchantID
	params["callback_url"] = cli.Params.WithdrawBackUrl //ajax回调接口
	params["return_url"] = cli.Params.WithdrawFeBackUrl //前端回跳地址
	params["ref1"] = req.OrderNum                       //也是商户订单号
	params["exchange_rate"] = 1

	signStr, _ := utils.Sign(cli.Params.MerchantId, cli.Params.SecretKey)
	params["signature"] = signStr

	//返回值会放到这里
	var result CCoopWithdrawResponse

	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(params).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetLogger(cli.logger).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp))
	cli.logger.Infof("PSPResty#ccoop#withdraw->%+v", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("status code: %d", resp.StatusCode())
	}

	if resp.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp.Error(), resp.Body())
	}

	return &result, err
}
