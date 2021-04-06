package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

/**
生成红包推广链接
*/
type RpPromUrlGenerateParams struct {
	GenerateShortUrl     *bool     `json:"generate_short_url,omitempty"`        //是否生成短链接。true-是，false-否，默认false
	GenerateWeApp        *bool     `json:"generate_we_app,omitempty"`           //是否生成小程序链接。true-是，false-否，默认false
	PIdList              *[]string `json:"p_id_list"`                           //推广位列表，例如：["60005_612"]
	CustomParameters     *string   `json:"custom_parameters,omitempty"`         //自定义参数，为链接打上自定义标签。自定义参数最长限制64个字节。
	GenerateWeappWebview *bool     `json:"generate_weapp_webview,omitempty"`    //是否唤起微信客户端， 默认false 否，true 是
	WeAppWebViewShortUrl *bool     `json:"we_app_web_view_short_url,omitempty"` //唤起微信app推广短链接
	WeAppWebWiewUrl      *bool     `json:"we_app_web_wiew_url,omitempty"`       //唤起微信app推广链接
	ChannelType          *int      `json:"channel_type,omitempty"`
}

type RpPromUrlGenerateResult struct {
	RpPromotionUrlGenerateResponse struct {
		UrlList []RpPromUrlGenerateInfo `json:"url_list"`
	} `json:"rp_promotion_url_generate_response"`
	common.CommonResult
}

type RpPromUrlGenerateInfo struct {
	Url                       string `json:"url"`
	ShortUrl                  string `json:"short_url"`
	MobileUrl                 string `json:"mobile_url"`
	MobileShortUrl            string `json:"mobile_short_url"`
	WeAppWebViewUrl           string `json:"we_app_web_view_url"`
	WeAppWebViewShortUrl      string `json:"we_app_web_view_short_url"`
	MultiGroupUrl             string `json:"multi_group_url"`
	MultiGroupShortUrl        string `json:"multi_group_short_url"`
	MultiGroupMobileUrl       string `json:"multi_group_mobile_url"`
	MultiGroupMobileShortUrl  string `json:"multi_group_mobile_short_url"`
	MultiWeAppWebViewUrl      string `json:"multi_we_app_web_view_url"`
	MultiWeAppWebViewShortUrl string `json:"multi_we_app_web_view_short_url"`
	WeAppInfo                 struct {
		AppID             string `json:"app_id"`
		WeAppIconUrl      string `json:"we_app_icon_url"`
		BannerUrl         string `json:"banner_url"`
		Desc              string `json:"desc"`
		SourceDisplayName string `json:"source_display_name"`
		PagePath          string `json:"page_path"`
		UserName          string `json:"user_name"`
		Title             string `json:"title"`
	} `json:"we_app_info"`
}

func (this *DuoduoKe) RpPromUrlGenerate(p *RpPromUrlGenerateParams) (*RpPromUrlGenerateResult, error) {
	apiType := `pdd.ddk.rp.prom.url.generate`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)
	var result RpPromUrlGenerateResult
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
