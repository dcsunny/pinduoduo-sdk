package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

/**
多多进宝主题商品查询
*/

type ThemeGoodsSearchParams struct {
	ThemeID int `json:"theme_id"` //主题ID
}

type ThemeGoodsSearchResult struct {
	ThemeListGetResponse struct {
		GoodsList []ThemeGoodsSearchInfo `json:"goods_list"`
		Total     int                    `json:"total"`
	} `json:"theme_list_get_response"`
	common.CommonResult
}

type ThemeGoodsSearchInfo struct {
	PromotionRate        int     `json:"promotion_rate"`          //佣金比例,千分比
	GoodsEvalCount       int     `json:"goods_eval_count"`        //商品评价数量
	GoodsID              int64   `json:"goods_id"`                //商品编码
	GoodsName            string  `json:"goods_name"`              //商品名称
	GoodsDesc            string  `json:"goods_desc"`              //商品描述
	GoodsThumbnailUrl    string  `json:"goods_thumbnail_url"`     //商品缩略图
	GoodsImageUrl        string  `json:"goods_image_url"`         //商品主图
	GoodsQalleryUrls     string  `json:"goods_gallery_urls"`      //商品详情图列表
	MinGroupPrice        int     `json:"min_group_price"`         //最小拼团价格,单位为分
	MinNormalPrice       int     `json:"min_normal_price"`        //最小单买价格,单位为分
	MallName             string  `json:"mall_name"`               //店铺名称
	OptID                int64   `json:"opt_id"`                  //商品标签类目ID,使用pdd.goods.opt.get获取
	OptName              string  `json:"opt_name"`                //商品标签名
	CatIds               []int64 `json:"cat_ids"`                 //商品一~四级类目ID列表
	HasCoupon            bool    `json:"has_coupon"`              //商品是否带券,true-带券,false-不带券
	CouponMinOrderAmount int     `json:"coupon_min_order_amount"` //优惠券门槛价格,单位为分
	CouponDiscount       int     `json:"coupon_discount"`         //优惠券面额,单位为分
	CouponTotalQuantity  int     `json:"coupon_total_quantity"`   //优惠券总数量
	CouponRemainQuantity int     `json:"coupon_remain_quantity"`  //优惠券剩余数量
	CouponStartTime      int64   `json:"coupon_start_time"`       //优惠券生效时间,UNIX时间戳
	CouponEndTime        int64   `json:"coupon_end_time"`         //优惠券失效时间,UNIX时间戳
	SalesTip             string  `json:"sales_tip"`
	DescTxt              string  `json:"desc_txt"`
	ServTxt              string  `json:"serv_txt"`
	LgstTxt              string  `json:"lgst_txt"`
}

func (this *DuoduoKe) ThemeGoodsSearch(p *ThemeGoodsSearchParams) (*ThemeGoodsSearchResult, error) {
	apiType := `pdd.ddk.theme.goods.search`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)
	var result ThemeGoodsSearchResult
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
