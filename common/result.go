package common

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CommonResult struct {
	ErrorResponse struct {
		ErrorCode int    `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
		SubMsg    string `json:"sub_msg"`
		SubCode   string `json:"sub_code"`
	} `json:"error_response"`
}

var (
	ParamsError      = errors.New("params error")
	SystemError      = errors.New("system error")
	ClientIDError    = errors.New("client id error")
	AccessTokenError = errors.New("access token error")
	NotDkkGoodsError = errors.New("非多多客商品")
)

func CheckErrCode(commonResult CommonResult) error {
	if commonResult.ErrorResponse.ErrorCode != 0 {
		j, _ := json.Marshal(commonResult)
		fmt.Println(fmt.Sprintf("err code err:%s", string(j)))
	}
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
	case 50001:
		return NotDkkGoodsError
	case 20000:
		return AccessTokenError
	}
	return nil
}
