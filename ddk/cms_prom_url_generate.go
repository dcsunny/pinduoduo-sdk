package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

/**
生成商城推广链接接口
*/
type CmsPromUrlGenerateParams struct {
	GenerateShortUrl     *bool   `json:"generate_short_url,omitempty"`     //是否生成短链接，true-是，false-否
	PIdList              string  `json:"p_id_list"`                        //推广位列表，例如：["60005_612"]
	GenerateMobile       *bool   `json:"generate_mobile,omitempty"`        //是否生成手机跳转链接。true-是，false-否，默认false
	MultiGroup           *bool   `json:"multi_group,omitempty"`            //单人团多人团标志。true-多人团，false-单人团 默认false
	CustomParameters     *string `json:"custom_parameters,omitempty"`      //自定义参数，为链接打上自定义标签。自定义参数最长限制64个字节。
	GenerateWeappWebview *bool   `json:"generate_weapp_webview,omitempty"` //是否唤起微信客户端， 默认false 否，true 是
	WeAppWebViewShortUrl bool    `json:"we_app_web_view_short_url"`        //唤起微信app推广短链接
	WeAppWebWiewUrl      bool    `json:"we_app_web_wiew_url"`              //唤起微信app推广链接
	ChannelType          *int    `json:"channel_type,omitempty"`           //0, "1.9包邮"；1, "今日爆款"； 2, "品牌清仓"； 3, "默认商城"；4,"PC端专属商城"；非必填 ,默认是3,
}

type CmsPromUrlGenerateResult struct {
	CmsPromotionUrlGenerateResponse struct {
		UrlList []CmsPromUrlGenerateInfo `json:"url_list"`
		Total   int                      `json:"total"` //数量
	} `json:"cms_promotion_url_generate_response"`
	common.CommonResult
}

type CmsPromUrlGenerateInfo struct {
	Url                      string `json:"url"`                          //商城推广链接
	ShortUrl                 string `json:"short_url"`                    //商城推广短链接
	MobileUrl                string `json:"mobile_url"`                   //商城推广移动链接
	MobileShortUrl           string `json:"mobile_short_url"`             //商城推广移动链接
	MultiGroupUrl            string `json:"multi_group_url"`              //商城推广多人团链接
	MultiGroupShortUrl       string `json:"multi_group_short_url"`        //商城推广多人团短链接
	MultiGroupMobileUrl      string `json:"multi_group_mobile_url"`       //商城推广多人团移动链接
	MultiGroupMobileShortUrl string `json:"multi_group_mobile_short_url"` //商城推广多人团移动短链接
}

func (this *DuoduoKe) CmsPromUrlGenerate(p *CmsPromUrlGenerateParams) (*CmsPromUrlGenerateResult, error) {
	apiType := `pdd.ddk.cms.prom.url.generate`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)

	var result CmsPromUrlGenerateResult
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
