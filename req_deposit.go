package go_ccoop

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/asaka1234/go-ccoop/utils"
	"github.com/mitchellh/mapstructure"
	"time"
)

func (cli *Client) Deposit(req CCoopDepositRequest) (*CCoopDepositResponse, error) {

	rawURL := cli.Params.BaseUrl

	cTime := time.Now()

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["order_status"] = "0"     //写死
	params["trade_type"] = "deposit" //写死
	params["create_time"] = cTime.Format("2006-01-02")
	params["mer_id"] = cli.Params.MerchantId           //cli.merchantID
	params["name"] = "john"                            //cli.merchantName
	params["callback_url"] = cli.Params.DepositBackUrl //ajax回调接口
	params["return_url"] = cli.Params.DepositFeBackUrl //前端回跳地址
	params["ref1"] = req.OrderNum                      //也是商户订单号
	params["exchange_rate"] = 1

	signStr, _ := utils.Sign(cli.Params.MerchantId, cli.Params.SecretKey)
	params["signature"] = signStr

	//返回值会放到这里
	var result CCoopDepositResponse

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

	fmt.Printf("result: %s\n", string(resp.Body()))

	if err != nil {
		return nil, err
	}

	if result.Status == "ok" {
		if result.AmountTHB == 0 {
			return nil, errors.New("convert to thb wrong!")
		}

		if result.Ref1 != "" {
			//对参数做签名
			data := cli.Params.MerchantId + result.Ref1
			h := hmac.New(sha256.New, []byte(cli.Params.SecretKey))
			h.Write([]byte(data))
			sign := hex.EncodeToString(h.Sum(nil))

			url := fmt.Sprintf("%s?ref1=%s&merid=%s&sign=%s", cli.Params.QRCodeUrl, result.Ref1, cli.Params.MerchantId, sign)
			result.QRCodeUrl = url //拼凑收银台地址
		}
	}

	return &result, err
}
