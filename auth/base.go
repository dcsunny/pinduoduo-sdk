package auth

import "github.com/dcsunny/pinduoduo-sdk/common"

type Auth struct {
	*common.Context
}

func NewAuth(context *common.Context) *Auth {
	service := new(Auth)
	service.Context = context
	return service
}
