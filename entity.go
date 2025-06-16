package go_ccoop

type CCoopInitParams struct {
	MerchantId string `json:"merchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"` // merchantId
	SecretKey  string `json:"secretKey" mapstructure:"secretKey" config:"secretKey"  yaml:"secretKey"`
	Ip         string `json:"ip" mapstructure:"ip" config:"ip"  yaml:"ip"` //三方psp做callback时的服务器ip(做白名单验证)

	BaseUrl   string `json:"baseUrl" mapstructure:"baseUrl" config:"baseUrl"  yaml:"baseUrl"`
	QRCodeUrl string `json:"qrCodeUrl" mapstructure:"qrCodeUrl" config:"qrCodeUrl"  yaml:"qrCodeUrl"`

	DepositBackUrl   string `json:"depositBackUrl" mapstructure:"depositBackUrl" config:"depositBackUrl"  yaml:"depositBackUrl"`
	DepositFeBackUrl string `json:"depositFeBackUrl" mapstructure:"depositFeBackUrl" config:"depositFeBackUrl"  yaml:"depositFeBackUrl"`
	WithdrawBackUrl  string `json:"withdrawBackUrl" mapstructure:"withdrawBackUrl" config:"withdrawBackUrl"  yaml:"withdrawBackUrl"`
}

// ---------------------------------

type CCoopDepositRequest struct {
	OrderNum  string `json:"order_num" mapstructure:"order_num"`   //商户的订单号
	Deposit   string `json:"deposit" mapstructure:"deposit"`       //数量 amount (ccy就是THB,是写死的)
	OrderName string `json:"order_name" mapstructure:"order_name"` //商户订单username
	//以下sdk来搞`
	//ExchangeRate string `json:"exchange_rate" mapstructure:"exchange_rate"` //要真实汇率  USDTHB
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
	OrderNum    string `json:"order_num"`   //商户订单号
	CreateTime  string `json:"create_time"` //2025-06-16
	TradeType   string `json:"trade_type"`  //deposit
	CurrencyAmo string `json:"currency_amo"`
	Merchant    string `json:"merchant"`
	OrderName   string `json:"order_name"`
	Deposit     string `json:"deposit"`      //金额
	OrderStatus string `json:"order_status"` //比如: confirm
	MerID       string `json:"mer_id"`       //商户id
	Ref1        string `json:"ref1"`
}

type CCoopDepositBackRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}

// ----------withdraw-------------------------

type CCoopWithdrawRequest struct {
	OrderNum  string `json:"order_num" mapstructure:"order_num"`   //商户的订单号
	Withdraw  string `json:"withdraw" mapstructure:"withdraw"`     //数量 (ccy就是THB,是写死的)
	OrderName string `json:"order_name" mapstructure:"order_name"` //商户的用户name
	//以下sdk来搞
	//ExchangeRate string `json:"ex_rate" mapstructure:"ex_rate"`       //要真实汇率  USDTHB
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
}

// 提现回调
type CCoopWithdrawBackReq struct {
	OrderNum    string `json:"order_num"`   //商户订单号
	CreateTime  string `json:"create_time"` //2025-06-16
	TradeType   string `json:"trade_type"`  //withdraw
	CurrencyAmo string `json:"currency_amo"`
	Merchant    string `json:"merchant"`
	OrderName   string `json:"order_name"` // 商户userName
	AmountUSD   string `json:"amount_usd"` //数量
	AmountTHB   string `json:"amount_thb"` //数量
	BankName    string `json:"bank_name"`
	AcName      string `json:"ac_name"`      //merchant的用户name
	OrderStatus string `json:"order_status"` //confirm
	MerID       string `json:"mer_id"`
	Ref1        string `json:"ref1"`
}

type CCoopWithdrawBackRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}
