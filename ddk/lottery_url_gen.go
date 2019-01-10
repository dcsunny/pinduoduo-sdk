package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

type LotteryUrlGenParams struct {
	PidList              []string `json:"pid_list"`
	GenerateWeappWebview *bool    `json:"generate_weapp_webview"`
	GenerateShortUrl     *bool    `json:"generate_short_url"`
	MultiGroup           *bool    `json:"multi_group"`
	CustomParameters     string   `json:"custom_parameters"`
}

type LotteryUrlGenResult struct {
	LotteryUrlResponse struct {
		Total   int `json:"total"`
		UrlList struct {
			SingleUrlList ResourceUrlGenUrlList `json:"single_url_list"`
			MultiUrlList  ResourceUrlGenUrlList `json:"multi_url_list"`
		} `json:"url_list"`
	} `json:"lottery_url_response"`
	common.CommonResult
}

func (this *DuoduoKe) LotteryUrlGen(p *LotteryUrlGenParams) (*LotteryUrlGenResult, error) {
	apiType := `pdd.ddk.lottery.url.gen`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)
	var result LotteryUrlGenResult
	err := util.HttpPOST(url, nil, &result)
	if err != nil {
		return nil, err
	}

	err = common.CheckErrCode(result.CommonResult)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
