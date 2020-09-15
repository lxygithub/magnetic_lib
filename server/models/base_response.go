package models

type BaseResp struct {
	Code   int         `json:"code"`
	Token  string      `json:"token"`
	ErrMsg string      `json:"err_msg"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}
