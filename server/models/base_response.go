package models

type BaseResp struct {
	Code   int         `json:"code"`
	ErrMsg string      `json:"err_msg"`
	Data   interface{} `json:"data"`
}
