package ddk_tool

import "github.com/dcsunny/pinduoduo-sdk/common"

type DuoduoKeTool struct {
	*common.Context
}

func NewDuoduoKeTool(context *common.Context) *DuoduoKeTool {
	service := new(DuoduoKeTool)
	service.Context = context
	return service
}
