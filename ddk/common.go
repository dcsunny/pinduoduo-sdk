package ddk

import (
	"fmt"

	"time"

	"github.com/dcsunny/pinduoduo-sdk/util"
)

const (
	CommonURL = "http://gw-api.pinduoduo.com/api/router?type=%s&data_type=JSON&timestamp=%d&client_id=%s&sign=%s"
)

type DuoduoKe struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func NewDuoduoKe(clientID string, clientSecret string) *DuoduoKe {
	return &DuoduoKe{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

func (this *DuoduoKe) GetURL(apiType string, accessToken string, params map[string]interface{}, paramsURL string) string {
	timestamp := time.Now().Unix()
	sign := util.Signature(apiType, this.ClientID, this.ClientSecret, accessToken, timestamp, params)

	url := fmt.Sprintf(CommonURL, apiType, timestamp, this.ClientID, sign)
	url = url + paramsURL
	return url
}
