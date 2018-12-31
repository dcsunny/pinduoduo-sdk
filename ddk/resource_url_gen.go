package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

type ResourceUrlGenParams struct {
	Pid          string  `json:"pid"`
	ResourceType *int    `json:"resource_type"`
	Url          *string `json:"url"`
}

type ResourceUrlGenResult struct {
	ResourceUrlResponse struct {
		SingleUrlList ResourceUrlGenUrlList `json:"single_url_list"`
		MultiUrlList  ResourceUrlGenUrlList `json:"multi_url_list"`
	} `json:"resource_url_response"`
	common.CommonResult
}

type ResourceUrlGenUrlList struct {
	URL                  string `json:"url"`
	ShortUrl             string `json:"short_url"`
	MobileUrl            string `json:"mobile_url"`
	MobileShortUrl       string `json:"mobile_short_url"`
	WeAppWebViewUrl      string `json:"we_app_web_view_url"`
	WeAppWebViewShortUrl string `json:"we_app_web_view_short_url"`
	WeAppPagePath        string `json:"we_app_page_path"`
}

func (this *DuoduoKe) ResourceUrlGen(p *ResourceUrlGenParams) (*ResourceUrlGenResult, error) {
	apiType := `pdd.ddk.resource.url.gen`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)
	var result ResourceUrlGenResult
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
