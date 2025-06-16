package go_ccoop

type CCoopInitParams struct {
	MerchantId string `json:"merchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"` // merchantId
	SecretKey  string `json:"secretKey" mapstructure:"secretKey" config:"secretKey"  yaml:"secretKey"`

	BaseUrl   string `json:"baseUrl" mapstructure:"baseUrl" config:"baseUrl"  yaml:"baseUrl"`
	QRCodeUrl string `json:"qrCodeUrl" mapstructure:"qrCodeUrl" config:"qrCodeUrl"  yaml:"qrCodeUrl"`

	DepositBackUrl    string `json:"depositBackUrl" mapstructure:"depositBackUrl" config:"depositBackUrl"  yaml:"depositBackUrl"`
	DepositFeBackUrl  string `json:"depositFeBackUrl" mapstructure:"depositFeBackUrl" config:"depositFeBackUrl"  yaml:"depositFeBackUrl"`
	WithdrawBackUrl   string `json:"withdrawBackUrl" mapstructure:"withdrawBackUrl" config:"withdrawBackUrl"  yaml:"withdrawBackUrl"`
	WithdrawFeBackUrl string `json:"withdrawFeBackUrl" mapstructure:"withdrawFeBackUrl" config:"withdrawFeBackUrl"  yaml:"withdrawFeBackUrl"`
}

// ---------------------------------

/*
	{
		  "order_num": "num***",
		  "create_time": "2023-03-10",
		  "trade_type": "deposit",
		  "deposit": "300001",
		  "order_name": "Admin",
		  "order_status": "0",
		  "mer_id": "*******",
		  "name": "Test",
		  "ref1": "12345",
		  "callback_url": "123",
		  "return_url": "www.ssss.com",
		  "signature": "4cf32d834aa9b005d7b180762f7a4a9f5e2c21e4bb2a145c980a71a8c3ce3ed0",
		  "ex_rate": "34.00"
	}
*/
type CCoopDepositRequest struct {
	OrderNum string `json:"order_num" mapstructure:"order_num"` //商户的订单号
	Deposit  string `json:"deposit" mapstructure:"deposit"`     //数量 amount (ccy就是THB,是写死的)
	//以下sdk来搞`
	//OrderName    string `json:"order_name" mapstructure:"order_name"`       //商户订单username
	//ExchangeRate string `json:"exchange_rate" mapstructure:"exchange_rate"` //TODO 这里要看一下具体实现
	//OrderStatus string `json:"order_status" mapstructure:"order_status"` //写死，都是 0  created
	//TradeType  string `json:"trade_type" mapstructure:"trade_type"`   //写死deposit
	//CreateTime string `json:"create_time" mapstructure:"create_time"` //yyyy-MM-dd, 当前时间
	//MerID       string `json:"mer_id" mapstructure:"mer_id"`             //商户号
	//Name        string `json:"name" mapstructure:"name"`                 //Merchant_Name
	//CallbackURL string `json:"callback_url" mapstructure:"callback_url"`
	//ReturnURL   string `json:"return_url" mapstructure:"return_url"`
	//Ref1        string `json:"ref1" mapstructure:"ref1"` //这个是一个md5签名,让sdk来计算
	//Signature        string `json:"signature" mapstructure:"signature"`                 //签名
}

type CCoopDepositResponse struct {
	Status     string `json:"status"`      //成功的话是 "ok"
	AmountUSD  string `json:"amount_usd"`  //充值金额(USD)->传入的金额是美金金额
	AmountTHB  int    `json:"amount_thb"`  //充值金额(THB)->用汇率做转换
	AcName     string `json:"ac_name"`     //商户订单号.   --->  下单人的姓名,对应request里的 OrderName,
	CreateTime string `json:"create_time"` //2023-03-10
	Ref1       string `json:"ref1"`        //psp三方的订单号

	//manual
	QRCodeUrl string `json:"QRCodeUrl"` //这个是自己构造的，并非api返回
}

// ----------deposit callback-------------------------

type CCoopDepositBackReq struct {
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

type CCoopDepositBackRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}

// ----------withdraw-------------------------

/*
	{
	  "order_num": "num***",
	  "create_time": "2023-03-10",
	  "trade_type": "withdraw",
	  "withdraw": "30001",
	  "order_name": "Admin",
	  "order_status": "0",
	  "mer_id": "*******",
	  "name": "Test",
	  "ref1": "12345",
	  "callback_url": "123",
	  "return_url": "www.ssss.com",
	  "ex_rate": "34.00"
	}
*/
type CCoopWithdrawRequest struct {
	OrderNum string `json:"order_num" mapstructure:"order_num"` //商户的订单号
	Withdraw string `json:"withdraw" mapstructure:"withdraw"`   //数量 (ccy就是THB,是写死的)
	//以下sdk来搞
	//ExchangeRate string `json:"ex_rate" mapstructure:"ex_rate"` //写死1
	//OrderName string `json:"order_name" mapstructure:"order_name"` //这里最好也传merchantOrderNo, 因为回调中有该字段,但是没有order_num
	//OrderStatus string `json:"order_status" mapstructure:"order_status"` //写死，都是 0
	//TradeType  string `json:"trade_type" mapstructure:"trade_type"`   //写死 withdraw
	//CreateTime string `json:"create_time" mapstructure:"create_time"` //yyyy-MM-dd, 当前时间
	//MerID       string `json:"mer_id" mapstructure:"mer_id"`             //商户号
	//Name        string `json:"name" mapstructure:"name"`                 //Merchant_Name
	//CallbackURL string `json:"callback_url" mapstructure:"callback_url"`
	//ReturnURL   string `json:"return_url" mapstructure:"return_url"`
	//Ref1        string `json:"ref1" mapstructure:"ref1"` //这个是一个md5签名,让sdk来计算
	//Signature   string `json:"signature"`
}

type CCoopWithdrawResponse struct {
	Status     string `json:"status"`      //成功的话是 "ok"
	AmountUSD  string `json:"amount_usd"`  //充值金额(USD)->传入的金额是美金金额
	AmountTHB  int    `json:"amount_thb"`  //充值金额(THB)->用汇率做转换
	AcName     string `json:"ac_name"`     //商户订单号.   --->  下单人的姓名,对应request里的 OrderName,
	CreateTime string `json:"create_time"` //2023-03-10
	Ref1       string `json:"ref1"`        //psp 三方订单号

	//manual
	//QRCodeUrl string `json:"QRCodeUrl"` //这个是自己构造的，并非api返回
}
