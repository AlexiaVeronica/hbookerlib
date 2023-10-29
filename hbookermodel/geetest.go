package hbookermodel

type Geetest struct {
	Code string      `json:"code"`
	Tip  string      `json:"tip"`
	Data GeetestData `json:"data"`
}

type GeetestData struct {
	NeedUseGeetest string `json:"need_use_geetest"`
	CodeLen        string `json:"code_len"`
}

type GeetestFirstRegisterStruct struct {
	Success    int    `json:"success"`
	Gt         string `json:"gt"`
	Challenge  string `json:"challenge"`
	NewCaptcha bool   `json:"new_captcha"`
}
