package ddk_tool

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

type OauthRpPomUrlGenerateParams struct {
	Amount                     *int64                      `json:"amount,omitempty"`
	ChannelType                *int                        `json:"channel_type,omitempty"`
	DiyOneYuanParam            *DiyOneYuanParam            `json:"diy_one_yuan_param,omitempty"`
	DiyPromoActCollectionParam *DiyPromoActCollectionParam `json:"diy_promo_act_collection_param,omitempty"`
	DiyRedPacketParam          *DiyRedPacketParam          `json:"diy_red_packet_param,omitempty"`
	DiySpRedPacketParam        *DiySpRedPacketParam        `json:"diy_sp_red_packet_param,omitempty"`
	ExtParams                  map[string]string           `json:"ext_params,omitempty"`
	GenerateQqApp              *bool                       `json:"generate_qq_app,omitempty"`
	GenerateSchemaUrl          *bool                       `json:"generate_schema_url,omitempty"`
	GenerateShortLink          *bool                       `json:"generate_short_link,omitempty"`
	GenerateShortUrl           *bool                       `json:"generate_short_url,omitempty"`
	GenerateWeApp              *bool                       `json:"generate_we_app,omitempty"`
	PIdList                    []string                    `json:"p_id_list"`
	ScratchCardAmount          *int64                      `json:"scratch_card_amount,omitempty"`
	TmccParam                  TmccParam                   `json:"tmcc_param,omitempty"`
	ZsDuoId                    *int64                      `json:"zs_duo_id,omitempty"`
	DiyCouponRebateParam       *DiyCouponRebateParam       `json:"diy_coupon_rebate_param,omitempty"`
}
type DiyOneYuanParam struct {
	GoodsSign *string `json:"goods_sign,omitempty"`
}

type DiyPromoActCollectionParam struct {
	CollectionId *int64 `json:"collection_id,omitempty"`
}

type DiyRedPacketParam struct {
	AmountProbability []int64     `json:"amount_probability,omitempty"`
	DisText           *bool       `json:"dis_text,omitempty"`
	NotShowBackground *bool       `json:"not_show_background,omitempty"`
	OptId             *int64      `json:"opt_id,omitempty"`
	RangeItems        *RangeItems `json:"range_items,omitempty"`
}

type RangeItems struct {
	RangeFrom *int64 `json:"range_from,omitempty"`
	RangeId   *int64 `json:"range_id,omitempty"`
	RangeTo   *int64 `json:"range_to,omitempty"`
}

type DiySpRedPacketParam struct {
	GoodsSign *string `json:"goods_sign,omitempty"`
	SkuIdCode *string `json:"sku_id_code,omitempty"`
}

type TmccParam struct {
	GoodsSigns  []string `json:"goods_signs,omitempty"`
	TmcConfigId *int64   `json:"tmc_config_id,omitempty"`
}

type DiyCouponRebateParam struct {
	GoodsSign *string `json:"goods_sign,omitempty"`
}

type OauthRpPomUrlGenerateResult struct {
	RpPromotionUrlGenerateResponse RpPromotionUrlGenerateResponse `json:"rp_promotion_url_generate_response"`
	common.CommonResult
}
type RpPromotionUrlGenerateResponse struct {
	ResourceList []struct {
		Desc string `json:"desc"`
		Url  string `json:"url"`
	} `json:"resource_list"`
	UrlList []struct {
		MobileShortUrl           string `json:"mobile_short_url"`
		MobileUrl                string `json:"mobile_url"`
		MultiGroupMobileShortUrl string `json:"multi_group_mobile_short_url"`
		MultiGroupMobileUrl      string `json:"multi_group_mobile_url"`
		MultiGroupShortUrl       string `json:"multi_group_short_url"`
		MultiGroupUrl            string `json:"multi_group_url"`
		QqAppInfo                struct {
			AppId             string `json:"app_id"`
			BannerUrl         string `json:"banner_url"`
			Desc              string `json:"desc"`
			PagePath          string `json:"page_path"`
			QqAppIconUrl      string `json:"qq_app_icon_url"`
			SourceDisplayName string `json:"source_display_name"`
			Title             string `json:"title"`
			UserName          string `json:"user_name"`
		} `json:"qq_app_info"`
		SchemaUrl   string `json:"schema_url"`
		ShortUrl    string `json:"short_url"`
		TzSchemaUrl string `json:"tz_schema_url"`
		Url         string `json:"url"`
		WeAppInfo   struct {
			AppId             string `json:"app_id"`
			BannerUrl         string `json:"banner_url"`
			Desc              string `json:"desc"`
			PagePath          string `json:"page_path"`
			SourceDisplayName string `json:"source_display_name"`
			Title             string `json:"title"`
			UserName          string `json:"user_name"`
			WeAppIconUrl      string `json:"we_app_icon_url"`
		} `json:"we_app_info"`
		WeixinShortLink string `json:"weixin_short_link"`
	} `json:"url_list"`
}

func (this *DuoduoKeTool) OauthRpPomUrlGenerate(p *OauthRpPomUrlGenerateParams, accessToken string) (*OauthRpPomUrlGenerateResult, error) {
	apiType := `pdd.ddk.oauth.rp.prom.url.generate`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, accessToken, params, paramsURL)
	var result OauthRpPomUrlGenerateResult
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
