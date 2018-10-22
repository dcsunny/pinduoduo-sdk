package common

import (
	"fmt"
	"time"

	"github.com/dcsunny/pinduoduo-sdk/util"
)

type Service struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func NewService(clientID string, clientSecret string) *Service {
	return &Service{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

func (this *Service) GetURL(apiType string, accessToken string, params map[string]interface{}, paramsURL string) string {
	timestamp := time.Now().Unix()
	sign := util.Signature(apiType, this.ClientID, this.ClientSecret, accessToken, timestamp, params)

	url := fmt.Sprintf(CommonURL, apiType, timestamp, this.ClientID, sign)
	url = url + paramsURL
	return url
}
