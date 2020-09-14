package models

type BaseResp struct {
	Code   int         `json:"code"`
	ErrMsg string      `json:"err_msg"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}
