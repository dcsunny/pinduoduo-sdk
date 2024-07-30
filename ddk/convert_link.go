package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

/**
多多进宝转链接口
*/

type ConvertLinkParams struct {
	CustomParameters  *string `json:"custom_parameters,omitempty"`   //自定义参数，为链接打上自定义标签。自定义参数最长限制64个字节。
	GenerateSchemaUrl bool    `json:"generate_schema_url,omitempty"` //是否返回 schema URL
	Pid               string  `json:"pid"`                           //渠道id
	SourceUrl         string  `json:"source_url"`                    //	需转链的链接，支持拼多多商品链接、进宝长链/短链（即为pdd.ddk.goods.promotion.url.generate接口生成的长短链）
}

type ConvertLinkResult struct {
	GoodsZsUnitGenerateResponse struct {
		MobileShortUrl           string `json:"mobile_short_url"`             //推广短链接（可唤起拼多多app）
		MobileUrl                string `json:"mobile_url"`                   //推广长链接（唤起拼多多app）
		MultiGroupMobileShortUrl string `json:"multi_group_mobile_short_url"` //推广短链接（唤起拼多多app）
		MultiGroupMobileUrl      string `json:"multi_group_mobile_url"`       //推广长链接（可唤起拼多多app）
		MultiGroupShortUrl       string `json:"multi_group_short_url"`        //双人团推广短链接
		MultiGroupUrl            string `json:"multi_group_url"`              //双人团推广长链接
		SchemaUrl                string `json:"schema_url"`                   //schema的链接
		ShortUrl                 string `json:"short_url"`                    //单人团推广短链接
		Url                      string `json:"url"`                          //单人团推广长链接
	} `json:"goods_zs_unit_generate_response"`
	common.CommonResult
}

func (this *DuoduoKe) ConvertLink(p *ConvertLinkParams) (*ConvertLinkResult, error) {
	apiType := `pdd.ddk.oauth.goods.zs.unit.url.gen`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)
	var result ConvertLinkResult
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
