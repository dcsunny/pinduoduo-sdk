package ddk_tool

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

/**
多多进宝推广链接--生成备案链接
*/

type OauthGoodsPromotionUrlGenerateParams struct {
	CashGiftId                *int64            `json:"cash_gift_id,omitempty"`                 //多多礼金ID
	CashGiftName              *string           `json:"cash_gift_name,omitempty"`               // 自定义礼金标题
	CustomParameters          *string           `json:"custom_parameters,omitempty"`            //自定义参数
	ForceDuoId                *bool             `json:"force_duo_id,omitempty"`                 //是否使用多多客专属推广计划
	GenerateAuthorityUrl      *bool             `json:"generate_authority_url,omitempty"`       //是否生成带授权的单品链接。如果未授权，则会走授权流程
	GenerateMallCollectCoupon *bool             `json:"generate_mall_collect_coupon,omitempty"` //是否生成店铺收藏券推广链接
	GenerateQqApp             *bool             `json:"generate_qq_app,omitempty"`              //是否生成qq小程序
	GenerateSchemaUrl         *bool             `json:"generate_schema_url,omitempty"`          //是否返回 schema URL
	GenerateShortUrl          *bool             `json:"generate_short_url,omitempty"`           //是否生成短链接，true-是，false-否
	GenerateWeApp             *bool             `json:"generate_we_app,omitempty"`              //是否生成拼多多福利券微信小程序推广信息
	GoodsSignList             []string          `json:"goods_sign_list,omitempty"`              //商品goodsSign列表，例如：["c9r2omogKFFAc7WBwvbZU1ikIb16_J3CTa8HNN"]，支持批量生链
	MaterialId                *string           `json:"material_id,omitempty"`                  //素材ID，可以通过商品详情接口获取商品素材信息
	MultiGroup                *bool             `json:"multi_group,omitempty"`                  //true--生成多人团推广链接 false--生成单人团推广链接（默认false）1、单人团推广链接：用户访问单人团推广链接，可直接购买商品无需拼团。
	Pid                       string            `json:"p_id"`                                   //推广位ID
	SearchId                  *string           `json:"search_id,omitempty"`                    //搜索id，建议填写，提高收益。来自pdd.ddk.goods.recommend.get、pdd.ddk.goods.search、pdd.ddk.top.goods.list.query等接口
	SpecialParams             map[string]string `json:"special_params,omitempty"`               //特殊参数
	UrlType                   *int              `json:"url_type,omitempty"`                     //生成商品链接类型 0-默认 1-百补相似品列表
	ZsDuoId                   *int64            `json:"zs_duo_id,omitempty"`                    //招商多多客ID
	GenerateWeAppLongLink     *bool             `json:"generate_we_app_long_link,omitempty"`    //是否生成小程序schema长链
}

type OauthGoodsPromotionUrlGenerateResult struct {
	GoodsPromotionUrlGenerateResponse struct {
		GoodsPromotionUrlList []GoodsPromotionUrlList `json:"goods_promotion_url_list"`
	} `json:"goods_promotion_url_generate_response"`

	common.CommonResult
}

type GoodsPromotionUrlList struct {
	MobileShortUrl string `json:"mobile_short_url"`
	MobileUrl      string `json:"mobile_url"`
	QqAppInfo      struct {
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
	WeixinLongLink  string `json:"weixin_long_link"`
	WeixinShortLink string `json:"weixin_short_link"`
}

func (this *DuoduoKeTool) OauthGoodsPromotionUrlGenerate(p *OauthGoodsPromotionUrlGenerateParams, accessToken string) (*OauthGoodsPromotionUrlGenerateResult, error) {
	apiType := `pdd.ddk.oauth.goods.prom.url.generate`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, accessToken, params, paramsURL)
	var result OauthGoodsPromotionUrlGenerateResult
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
