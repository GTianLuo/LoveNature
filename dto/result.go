package dto

import "lovenature/pkg/e"

type Result struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Err  error       `json:"err"`
}

func Fail(code int, err error) *Result {
	return &Result{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: err,
	}
}

func Success(code int, data interface{}) *Result {
	return &Result{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	}
}
