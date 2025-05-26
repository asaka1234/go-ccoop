package utils

func GenSign(params map[string]interface{}, backKey string) string {
	// 获取当前时间并格式化为 yyyyMMddHHmmss
	//currentTime := time.Now().Format("20060102150405") // Go的特定时间格式

	// TODO 拼接签名字符串
	/*
		signText := params["mer_id"] + customerName + amount + req.CURRENCY +
			currentTime + props.CCoop_Credentials.CallBack_Key +
			internalReq.TrackingInfo.IPAddress

	*/
	signText := ""

	// 计算MD5签名
	return GetMD5([]byte(signText))
}

func VerifySign(params map[string]interface{}, backKey string, sign string) bool {
	signSelf := GenSign(params, backKey)
	return signSelf == signSelf
}
