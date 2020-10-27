package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

/**
生成多多进宝商品推广链接
*/
type GoodsPromotionUrlGenerateParams struct {
	PID                       string    `json:"p_id"`                                   //推广位ID
	GoodsIdList               []int64   `json:"goods_id_list"`                          //商品ID，仅支持单个查询
	GenerateShortUrl          bool      `json:"generate_short_url"`                     //是否生成短链接，true-是，false-否
	MultiGroup                *bool     `json:"multi_group,omitempty"`                  //true--生成多人团推广链接 false--生成单人团推广链接（默认false）1、单人团推广链接：用户访问单人团推广链接，可直接购买商品无需拼团。2、多人团推广链接：用户访问双人团推广链接开团，若用户分享给他人参团，则开团者和参团者的佣金均结算给推手
	CustomParameters          *string   `json:"custom_parameters,omitempty"`            //自定义参数，为链接打上自定义标签。自定义参数最长限制64个字节。
	PullNew                   *bool     `json:"pull_new,omitempty"`                     //是否开启订单拉新，true表示开启（订单拉新奖励特权仅支持白名单，请联系工作人员开通）
	GenerateWeappWebview      *bool     `json:"generate_weapp_webview,omitempty"`       //是否生成唤起微信客户端链接，true-是，false-否，默认false
	ZsDuoID                   *int64    `json:"zs_duo_id,omitempty"`                    //招商多多客ID
	GenerateWeApp             *bool     `json:"generate_we_app,omitempty"`              //是否生成小程序推广
	GenerateMallCollectCoupon *bool     `json:"generate_mall_collect_coupon,omitempty"` //是否生成店铺收藏券推广链接
	GenerateQqApp             *bool     `json:"generate_qq_app,omitempty"`              //是否生成qq小程序
	GenerateSchemaUrl         *bool     `json:"generate_schema_url,omitempty"`          //是否返回 schema URL
	GenerateWeiboappWebview   *bool     `json:"generate_weiboapp_webview,omitempty"`    //是否生成微博推广链接
	SearchId                  *string   `json:"search_id,omitempty"`                    //搜索id，建议填写，提高收益。来自pdd.ddk.goods.recommend.get、pdd.ddk.goods.search、pdd.ddk.top.goods.list.query等接口
	RoomIdList                *[]string `json:"room_id_list,omitempty"`                 //直播间id列表，如果生成直播间推广链接该参数必填，goods_id_list填[1]
	TargetIdList              *[]string `json:"target_id_list,omitempty"`               //直播预约id列表，如果生成直播间预约推广链接该参数必填，goods_id_list填[1]，room_id_list不填
	GenerateAuthorityUrl      *bool     `json:"generateAuthorityUrl,omitempty"`         //是否生成带授权的单品链接。如果未授权，则会走授权流程
}

type GoodsPromotionUrlGenerateResult struct {
	GoodsPromotionUrlGenerateResponse struct {
		GoodsPromotionUrlList []GoodsPromotionUrlGenerateInfo `json:"goods_promotion_url_list"`
	} `json:"goods_promotion_url_generate_response"`
	common.CommonResult
}

type GoodsPromotionUrlGenerateInfo struct {
	WeAppWebViewShortUrl    string `json:"we_app_web_view_short_url"`
	WeAppWebViewUrl         string `json:"we_app_web_view_url"`
	MobileShortUrl          string `json:"mobile_short_url"`
	MobileUrl               string `json:"mobile_url"`
	ShortUrl                string `json:"short_url"`
	Url                     string `json:"url"`
	SchemaUrl               string `json:"schema_url"`
	WeiboAppWebViewShortUrl string `json:"weibo_app_web_view_short_url"`
	WeiboAppWebViewUrl      string `json:"weibo_app_web_view_url"`
	WeAppInfo               struct {
		WeAppIconUrl      string `json:"we_app_icon_url"`
		BannerUrl         string `json:"banner_url"`
		Desc              string `json:"desc"`
		SourceDisplayName string `json:"source_display_name"`
		PagePath          string `json:"page_path"`
		UserName          string `json:"user_name"`
		Title             string `json:"title"`
		AppId             string `json:"app_id"`
	} `json:"we_app_info"`

	QqAppInfo struct {
		AppId             string `json:"app_id"`
		BannerUrl         string `json:"banner_url"`
		Desc              string `json:"desc"`
		PagePath          string `json:"page_path"`
		QqAppIconUrl      string `json:"qq_app_icon_url"`
		SourceDisplayName string `json:"source_display_name"`
		UserName          string `json:"user_name"`
		Title             string `json:"title"`
	} `json:"qq_app_info"`
}

func (this *DuoduoKe) GoodsPromotionUrlGenerate(p *GoodsPromotionUrlGenerateParams) (*GoodsPromotionUrlGenerateResult, error) {
	apiType := `pdd.ddk.goods.promotion.url.generate`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)

	var result GoodsPromotionUrlGenerateResult
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
