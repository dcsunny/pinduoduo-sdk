package ddk

import (
	"encoding/json"

	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

/**
多多进宝生成推广链接，生成备案链接generate_authority_url=true
*/

type ReportOrPromotionLinkParams struct {
	CashGiftID                int64             `json:"cash_gift_id,omitempty"`                 //多多礼金ID
	CashGiftName              string            `json:"cash_gift_name,omitempty"`               //自定义礼金标题，用于向用户展示渠道专属福利，不超过12个字
	CustomParameters          json.RawMessage   `json:"custom_parameters,omitempty"`            //自定义参数，为链接打上自定义标签；自定义参数最长限制64个字节；
	ForceDuoID                bool              `json:"force_duo_id,omitempty"`                 //是否使用多多客专属推广计划
	GenerateAuthorityURL      bool              `json:"generate_authority_url,omitempty"`       //是否生成带授权的单品链接。如果未授权，则会走授权流程
	GenerateMallCollectCoupon bool              `json:"generate_mall_collect_coupon,omitempty"` //是否生成店铺收藏券推广链接
	GenerateQQApp             bool              `json:"generate_qq_app,omitempty"`              //是否生成qq小程序
	GenerateSchemaURL         bool              `json:"generate_schema_url,omitempty"`          //是否返回 schema URL
	GenerateShortURL          bool              `json:"generate_short_url,omitempty"`           //是否生成短链接，true-是，false-否
	GenerateWeApp             bool              `json:"generate_we_app,omitempty"`              //是否生成拼多多福利券微信小程序推广信息
	GoodsSignList             []string          `json:"goods_sign_list,omitempty"`              //商品goodsSign列表，例如：["c9r2omogKFFAc7WBwvbZU1ikIb16_J3CTa8HNN"]，支持批量生链。goodsSign是加密后的goodsId, goodsId已下线
	MaterialID                string            `json:"material_id,omitempty"`                  //素材ID，可以通过商品详情接口获取商品素材信息
	MultiGroup                bool              `json:"multi_group,omitempty"`                  //true--生成多人团推广链接 false--生成单人团推广链接（默认false）
	PID                       string            `json:"p_id"`                                   //推广位ID
	SearchID                  string            `json:"search_id,omitempty"`                    //搜索id，建议填写，提高收益。
	SpecialParams             map[string]string `json:"special_params,omitempty"`               //特殊参数
	ZsDuoID                   int64             `json:"zs_duo_id,omitempty"`                    //招商多多客ID
	URLType                   int               `json:"url_type,omitempty"`                     //生成商品链接类型 0-默认 1-百补相似品列表
}

type SpecialParam struct {
	Key   string `json:"$key"`
	Value string `json:"$value"`
}

type ReportOrPromotionLinkResult struct {
	GoodsPromotionUrlGenerateResponse struct {
		GoodsPromotionUrlList []struct {
			MobileShortUrl string   `json:"mobile_short_url"` //对应出参mobile_url的短链接，与mobile_url功能一致。
			MobileUrl      string   `json:"mobile_url"`       //使用此推广链接，用户安装微信的情况下，默认拉起拼多多福利券微信小程序，否则唤起H5页面
			QqAppInfo      struct { //qq小程序信息
				AppId             string `json:"app_id"`              //拼多多小程序id
				BannerUrl         string `json:"banner_url"`          //Banner图
				Desc              string `json:"desc"`                //描述
				PagePath          string `json:"page_path"`           //小程序path值
				QqAppIconUrl      string `json:"qq_app_icon_url"`     //小程序icon
				SourceDisplayName string `json:"source_display_name"` //来源名
				Title             string `json:"title"`               //小程序标题
				UserName          string `json:"user_name"`           //用户名
			} `json:"qq_app_info"`
			SchemaUrl   string   `json:"schema_url"`    //使用此推广链接，用户安装拼多多APP的情况下会唤起APP（需客户端支持schema跳转协议）
			ShortUrl    string   `json:"short_url"`     //对应出参url的短链接，与url功能一致
			TzSchemaUrl string   `json:"tz_schema_url"` //使用此推广链接，用户安装多多团长APP的情况下会唤起APP（需客户端支持schema跳转协议）
			Url         string   `json:"url"`           //普通推广长链接，唤起H5页面
			WeAppInfo   struct { //拼多多福利券微信小程序信息
				AppId             string `json:"app_id"`              //小程序id
				BannerUrl         string `json:"banner_url"`          //Banner图
				Desc              string `json:"desc"`                //描述
				PagePath          string `json:"page_path"`           //小程序path值
				SourceDisplayName string `json:"source_display_name"` //来源名
				Title             string `json:"title"`               //小程序标题
				UserName          string `json:"user_name"`           //用户名
				WeAppIconUrl      string `json:"we_app_icon_url"`     //小程序图片
			} `json:"we_app_info"`
		} `json:"goods_promotion_url_list"`
	} `json:"goods_promotion_url_generate_response"`
	common.CommonResult
}

func (this *DuoduoKe) ReportOrPromotionLink(p *ReportOrPromotionLinkParams) (*ReportOrPromotionLinkResult, error) {
	apiType := `pdd.ddk.oauth.goods.prom.url.generate`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)
	var result ReportOrPromotionLinkResult
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
