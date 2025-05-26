package go_ccoop

import (
	"github.com/asaka1234/go-ccoop/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Merchant_ID          string
	Merchant_Name        string
	Merchant_Title       string
	Exchange_To_Currency string

	CallBack_Key string

	Deposit_Url          string
	Deposit_CallBack_Url string
	DepositReturn_Url    string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, Merchant_ID, Merchant_Name, Merchant_Title string, Exchange_To_Currency string, CallBack_Key string, Deposit_Url, Deposit_CallBack_Url, DepositReturn_Url string) *Client {
	return &Client{
		Merchant_ID:          Merchant_ID,
		Merchant_Name:        Merchant_Name,
		Merchant_Title:       Merchant_Title,
		Exchange_To_Currency: Exchange_To_Currency,

		CallBack_Key: CallBack_Key,

		Deposit_Url:          Deposit_Url,
		Deposit_CallBack_Url: Deposit_CallBack_Url,
		DepositReturn_Url:    DepositReturn_Url,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
