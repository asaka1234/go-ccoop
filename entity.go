package go_ccoop

import "time"

// ---------------------------------

type CCoopDepositRequest struct {
	OrderNum     string `json:"order_num" mapstructure:"order_num"`         //商户的订单号
	Deposit      string `json:"deposit" mapstructure:"deposit"`             //数量
	OrderName    string `json:"order_name" mapstructure:"order_name"`       //商户订单username
	ExchangeRate string `json:"exchange_rate" mapstructure:"exchange_rate"` //TODO 这里要看一下具体实现
	//以下sdk来搞
	//OrderStatus string `json:"order_status" mapstructure:"order_status"` //写死，都是 0
	//TradeType  string `json:"trade_type" mapstructure:"trade_type"`   //写死，都是 deposit
	//CreateTime string `json:"create_time" mapstructure:"create_time"` //yyyy-MM-dd, 当前时间
	//MerID       string `json:"mer_id" mapstructure:"mer_id"`             //商户号
	//Merchant    string `json:"merchant" mapstructure:"merchant"`         //Merchant_Title
	//Name        string `json:"name" mapstructure:"name"`                 //Merchant_Name
	//CurrencyAmo string `json:"currency_amo" mapstructure:"currency_amo"` //Exchange_To_Currency
	//CallbackURL string `json:"callback_url" mapstructure:"callback_url"`
	//ReturnURL   string `json:"return_url" mapstructure:"return_url"`
	//Ref1        string `json:"ref1" mapstructure:"ref1"` //这个是一个md5签名,让sdk来计算
}

type CCoopDepositResponse struct {
	Status      string    `json:"status"` //成功的话是 ok
	OrderNum    string    `json:"order_num"`
	AmountUSD   string    `json:"amount_usd"`
	AmountTHB   string    `json:"amount_thb"`
	BankName    string    `json:"bank_name"`
	AcName      string    `json:"ac_name"`
	CreateTime  time.Time `json:"create_time"` //2023-03-10
	Ref1        string    `json:"ref1"`
	RedirectURL string    `json:"redirectUrl"`
}

// ----------deposit callback-------------------------
// Notice: 12pay是没有任何callback的验签逻辑的, 所以需要自己搞. 一般都是借助ref字段实现

type One2PayDepositBackReq struct {
	RespCode   int     `json:"resp_code"` //200是成功
	RespMsg    string  `json:"resp_msg"`
	Command    string  `json:"command"`
	BankRef    string  `json:"bank_ref"`
	TranxId    string  `json:"tranx_id"`
	One2PayRef string  `json:"one2pay_ref"`
	Datetime   string  `json:"datetime"`
	EffDate    string  `json:"effdate"`
	Amount     float64 `json:"amount"`
	CusName    string  `json:"cusname"`
	Ref1       string  `json:"ref1"` //放业务自己的orderNo (只能是数字/字母)
	Ref2       string  `json:"ref2"`
	Ref3       string  `json:"ref3"`
	Ref4       string  `json:"ref4"`     // 这个用来做签名, 是amount/ref1/authkey的一个md5签名的截断值(18位)
	TransId    string  `json:"trans_id"` //是psp三方的订单id
}

/*
{
	"resp_code":200,
	"resp_msg":"Success“,
	"command":"Payment",
	"bank_ref":"ITMX 13078496",
	"tranx_id":"5oaY6TEYxKPBi34yDxM2",
	"one2pay_ref":"ABC220407080726611",
	"datetime":"20231004133121",
	"effdate":"20231004",
	"amount":18658.26,
	"cusname":"นาย จิณณะ แสงฤทธิ์“,
	"ref1":"SPI729131696400892",
	"ref2":“123456",
	"ref3":“abc123-",
	"ref4":"abc123-",
	"trans_id":"2f0cf4e1ceb53c4053837d0860c19620“
}
*/

type One2PayDepositBackRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}

// ----------withdraw-------------------------

type One2PayWithdrawRequest struct {
	BankAcc       string  `json:"bankacc"`        //required
	BankCode      string  `json:"bankcode"`       //required
	BankName      string  `json:"bankname"`       //required
	AccountName   string  `json:"accname"`        //required
	Amount        float64 `json:"amount"`         //required
	MobileNo      string  `json:"mobileno"`       //required
	TransactionBy string  `json:"transaction_by"` //required
	Ref1          string  `json:"ref1"`           //required
	Ref2          string  `json:"ref2"`
	Ref3          string  `json:"ref3"`
	Ref4          string  `json:"ref4"`
}

/*
{
	"bankacc":“0652078409",
	"bankcode":"004",
	"bankname":"KASIKORN BANK",
	"accname":"Manop Tangngam",
	"amount":1000.50,
	"mobileno":"0805933181",
	"transaction_by":"Jack Developer",
	"ref1": "123456789012345678“
}
*/

type One2PayWithdrawResponse struct {
	Error               string    `json:"error"`
	Status              int       `json:"status"` //Success Case Status 1000 only
	Message             string    `json:"message"`
	Ref1                string    `json:"ref1"`
	TransactionId       string    `json:"transaction_id"`
	TransactionDateTime time.Time `json:"transactionDate_time"` //YYY-MM-DD hh:mm:ss
}
