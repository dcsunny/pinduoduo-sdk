package common

import (
	"fmt"
	"time"

	"github.com/dcsunny/pinduoduo-sdk/util"
)

type Context struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func NewContext(clientID string, clientSecret string) *Context {
	return &Context{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

func (this *Context) GetURL(apiType string, accessToken string, params map[string]interface{}, paramsURL string) string {
	timestamp := time.Now().Unix()
	sign := util.Signature(apiType, this.ClientID, this.ClientSecret, accessToken, timestamp, params)

	url := fmt.Sprintf(CommonURL, apiType, timestamp, this.ClientID, sign)
	if accessToken != "" {
		url = url + "&access_token=" + accessToken
	}
	url = url + paramsURL
	return url
}
