package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

/**
查询多多进宝商品详情
*/
type GoodsDetailParams struct {
	GoodsIdList []int64 `json:"goods_id_list"` //商品ID，仅支持单个查询。例如：[123456]
}

type GoodsDetailResult struct {
	GoodsDetailResponse struct {
		GoodsDetails []GoodsDetailInfo `json:"goods_details"` //多多进宝商品对象列表
	} `json:"goods_detail_response"`
	common.CommonResult
}

type GoodsDetailInfo struct {
	MallCouponID                int64    `json:"mall_coupon_id"`                  //店铺券id
	MallCouponDiscountPct       int      `json:"mall_coupon_discount_pct"`        //店铺券折扣
	MallCouponMinOrderAmount    int      `json:"mall_coupon_min_order_amount"`    //最小使用金额
	MallCouponMaxDiscountAmount int      `json:"mall_coupon_max_discount_amount"` //最大使用金额
	MallCouponTotalQuantity     int64    `json:"mall_coupon_total_quantity"`      //店铺券总量
	MallCouponRemainQuantity    int64    `json:"mall_coupon_remain_quantity"`     //店铺券余量
	MallCouponStartTime         int64    `json:"mall_coupon_start_time"`          //店铺券开始使用时间
	MallCouponEndTime           int64    `json:"mall_coupon_end_time"`            //店铺券结束使用时间
	MallID                      int64    `json:"mall_id"`                         //店铺id
	ServPct                     float64  `json:"serv_pct"`                        //服务评分击败同类店铺百分比
	LgstPct                     float64  `json:"lgst_pct"`                        //物流评分击败同类店铺百分比
	DescPct                     float64  `json:"desc_pct"`                        //描述评分击败同类店铺百分比
	OptIds                      []int64  `json:"opt_ids"`                         //商品标签ID
	AvgServ                     int      `json:"avg_serv"`                        //服务评分
	AvgLgst                     int      `json:"avg_lgst"`                        //物流评分
	AvgDesc                     int      `json:"avg_desc"`                        //描述评分
	CatIds                      []int64  `json:"cat_ids"`                         //商品一~四级类目ID列表
	CatID                       int64    `json:"cat_id"`                          //商品类目ID，使用pdd.goods.cats.get接口获取
	CouponRemainQuantity        int      `json:"coupon_remain_quantity"`          //优惠券剩余数量
	GoodsEvalCount              int      `json:"goods_eval_count"`                //商品评价数
	GoodsEvalScore              float64  `json:"goods_eval_score"`                //商品评价分
	OptName                     string   `json:"opt_name"`                        //商品标签名称
	CouponMinOrderAmount        int      `json:"coupon_min_order_amount"`         //优惠券门槛金额，单位为分
	CouponDiscount              int      `json:"coupon_discount"`                 //优惠券面额，单位为分
	CouponTotalQuantity         int      `json:"coupon_total_quantity"`           //优惠券总数量
	CouponStartTime             int64    `json:"coupon_start_time"`               //优惠券生效时间，UNIX时间戳
	CouponEndTime               int64    `json:"coupon_end_time"`                 //优惠券生效时间，UNIX时间戳
	PromotionRate               int      `json:"promotion_rate"`                  //佣金比例，千分比
	MallName                    string   `json:"mall_name"`                       //店铺名称
	OptID                       int64    `json:"opt_id"`                          //商品标签ID，使用pdd.goods.opt.get接口获取
	GoodsID                     int64    `json:"goods_id"`                        //参与多多进宝的商品ID
	GoodsName                   string   `json:"goods_name"`                      //参与多多进宝的商品标题
	GoodsDesc                   string   `json:"goods_desc"`                      //参与多多进宝的商品描述
	GoodsImageUrl               string   `json:"goods_image_url"`                 //多多进宝商品主图
	GoodsGalleryUrls            []string `json:"goods_gallery_urls"`              //商品轮播图
	SoldQuantity                int      `json:"sold_quantity"`                   //已售卖件数
	MinGroupPrice               int      `json:"min_group_price"`                 //最低价sku的拼团价，单位为分
	MinNormalPrice              int      `json:"min_normal_price"`                //最低价sku的单买价，单位为分
}

func (this *DuoduoKe) GoodsDetail(p *GoodsDetailParams) (*GoodsDetailResult, error) {
	apiType := "pdd.ddk.goods.detail"
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)

	var result GoodsDetailResult
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
