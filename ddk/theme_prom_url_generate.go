package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

/**
多多进宝主题活动推广链接生成
*/

type ThemePromUrlGenerateParams struct {
	PID                  string  `json:"pid"`                                 //推广位ID
	ThemeIdList          []int   `json:"theme_id_list"`                       //主题ID列表,例如[1,235]
	GenerateShortUrl     *bool   `json:"generate_short_url,omitempty"`        //是否生成短链接,true-是,false-否
	GenerateMobile       *bool   `json:"generate_mobile,omitempty"`           //是否生成手机跳转链接。true-是,false-否,默认false
	CustomParameters     *string `json:"custom_parameters,omitempty"`         //自定义参数,为链接打上自定义标签。自定义参数最长限制64个字节。
	GenerateWeappWebview *bool   `json:"generate_weapp_webview,omitempty"`    //是否唤起微信客户端， 默认false 否，true 是
	WeAppWebViewShortUrl *bool   `json:"we_app_web_view_short_url,omitempty"` //唤起微信app推广短链接
	WeAppWebWiewUrl      *bool   `json:"we_app_web_wiew_url,omitempty"`       //唤起微信app推广链接
}

type ThemePromUrlGenerateResult struct {
	ThemePromotionUrlGenerateResponse struct {
		UrlList []ThemePromUrlGenerateInfo `json:"url_list"` //主题活动推广url列表
	} `json:"theme_promotion_url_generate_response"`
	common.CommonResult
}

type ThemePromUrlGenerateInfo struct {
	Url                      string `json:"url"`                          //主题活动推广链接
	ShortUrl                 string `json:"short_url"`                    //主题活动推广短链
	MobileUrl                string `json:"mobile_url"`                   //主题活动推广移动链接
	MobileShortUrl           string `json:"mobile_short_url"`             //主题活动推广移动短链接
	MultiGroupUrl            string `json:"multi_group_url"`              //主题活动推广开团链接
	MultiGroupShortUrl       string `json:"multi_group_short_url"`        //主题活动推广开团短链接
	MultiGroupMobileUrl      string `json:"multi_group_mobile_url"`       //主题活动推广开团移动端链接
	MultiGroupMobileShortUrl string `json:"multi_group_mobile_short_url"` //主题活动推广开团移动端短链接
}

func (this *DuoduoKe) ThemePromUrlGenerate(p *ThemePromUrlGenerateParams) (*ThemePromUrlGenerateResult, error) {
	apiType := `pdd.ddk.theme.prom.url.generate`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)
	var result ThemePromUrlGenerateResult
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
