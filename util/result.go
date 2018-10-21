package util

import (
	"errors"
)

type CommonResult struct {
	ErrorResponse struct {
		ErrorCode int    `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	} `json:"error_response"`
}

var (
	RequestError     = errors.New("request return status error")
	ParamsError      = errors.New("params error")
	SystemError      = errors.New("system error")
	ClientIDError    = errors.New("client id error")
	AccessTokenError = errors.New("access token error")
)

func CheckErrCode(commonResult CommonResult) error {
	errCode := commonResult.ErrorResponse.ErrorCode
	switch errCode {
	case 0:
		return nil
	case 10000:
		return ParamsError
	case 10010:
		return ClientIDError
	case 50000:
		return SystemError
	case 20000:
		return AccessTokenError
	}
	return nil
}
