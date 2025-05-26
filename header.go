package go_ccoop

func getHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
		//复刻
		"User-Agent":                   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET, POST",
		"Access-Control-Allow-Headers": "X-Requested-With",
	}
}
