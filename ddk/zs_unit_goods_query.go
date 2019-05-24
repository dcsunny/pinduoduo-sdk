package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

type ZsUnitGoodsQueryParams struct {
	ZSDuoID  int64 `json:"zs_duo_id"`
	Page     *int  `json:"page,omitempty"`
	PageSize int   `json:"page_size"`
}

type ZsUnitGoodsQueryResult struct {
	ZsUnitGoodsQueryResponse struct {
		TotalCount int                         `json:"total_count"`
		List       []ZsUnitGoodsQueryGoodsInfo `json:"list"`
	} `json:"zs_unit_goods_query_response"`
	common.CommonResult
}

type ZsUnitGoodsQueryGoodsInfo struct {
	CategoryName                string      `json:"category_name"`
	CouponRemainQuantity        int         `json:"coupon_remain_quantity"`
	PromotionRate               int         `json:"promotion_rate"`
	IsOnsale                    bool        `json:"is_onsale"`
	CouponID                    *int64      `json:"coupon_id"`
	LockEdit                    interface{} `json:"lock_edit"`
	MallID                      int64       `json:"mall_id"`
	ShareDesc                   string      `json:"share_desc"`
	MallName                    string      `json:"mall_name"`
	CouponPrice                 int         `json:"coupon_price"`
	Rank                        int         `json:"rank"`
	MallCouponEndTime           int64       `json:"mall_coupon_end_time"`
	MarketFee                   int         `json:"market_fee"`
	LgstTxt                     string      `json:"lgst_txt"`
	GoodsName                   string      `json:"goods_name"`
	GoodsEvalCount              int         `json:"goods_eval_count"`
	GoodsID                     int64       `json:"goods_id"`
	GoodsGalleryUrls            []string    `json:"goods_gallery_urls"`
	GoodsDesc                   string      `json:"goods_desc"`
	OptName                     string      `json:"opt_name"`
	OptIDs                      []int       `json:"opt_ids"`
	GoodsImageUrl               string      `json:"goods_image_url"`
	HasMallCoupon               bool        `json:"has_mall_coupon"`
	IsValid                     bool        `json:"is_valid"`
	MinGroupPrice               int         `json:"min_group_price"`
	CouponStartTime             int64       `json:"coupon_start_time"`
	CouponEndTime               int64       `json:"coupon_end_time"`
	CouponDiscount              int         `json:"coupon_discount"`
	AvgDesc                     int         `json:"avg_desc"`
	AvgLgst                     int         `json:"avg_lgst"`
	AvgServ                     int         `json:"avg_serv"`
	DescPct                     float64     `json:"desc_pct"`
	LgstPct                     float64     `json:"lgst_pct"`
	ServPct                     float64     `json:"serv_pct"`
	MallCouponRemainQuantity    int64       `json:"mall_coupon_remain_quantity"` //店铺券余量
	PlanType                    int         `json:"plan_type"`
	SoldQuantity                int         `json:"sold_quantity"`
	CatIds                      []int       `json:"cat_ids"`
	CouponMinOrderAmount        int         `json:"coupon_min_order_amount"`
	GoodsFactPrice              int         `json:"goods_fact_price"`         //商品实际价格
	MallCouponDiscountPct       int         `json:"mall_coupon_discount_pct"` //店铺券折扣
	GoodsEvalScore              float64     `json:"goods_eval_score"`
	ActivityType                int         `json:"activity_type"`
	CatID                       *int64      `json:"cat_id,omitempty"`
	CouponTotalQuantity         int         `json:"coupon_total_quantity"`
	MallCouponMinOrderAmount    int         `json:"mall_coupon_min_order_amount"` //最小使用金额
	UnitID                      int64       `json:"unit_id"`
	GoodsRate                   int         `json:"goods_rate"`
	MerchantType                int         `json:"merchant_type"`
	Quantity                    int         `json:"quantity"`
	SalesTip                    string      `json:"sales_tip"`
	GoodsMarkPrice              int         `json:"goods_mark_price"` //商品标准价格
	PlanTypeAll                 int         `json:"plan_type_all"`
	QrCodeImageUrl              string      `json:"qr_code_image_url"` //二维码主图
	MallCouponID                int64       `json:"mall_coupon_id"`    //店铺券id
	DescTxt                     string      `json:"desc_txt"`
	GoodsThumbnailUrl           string      `json:"goods_thumbnail_url"`
	OptId                       int         `json:"opt_id"`
	SaleNumToday                int         `json:"sale_num_today"` //今日成交量
	SaleNum24                   int         `json:"sale_num24"`     //成交量
	MinNormalPrice              int         `json:"min_normal_price"`
	HasCoupon                   bool        `json:"has_coupon"`
	MallCouponStartTime         int64       `json:"mall_coupon_start_time"` //店铺券开始使用时间
	ServTxt                     string      `json:"serv_txt"`
	GoodsMallCouponPrice        interface{} `json:"goods_mall_coupon_price"`
	MallRate                    int         `json:"mall_rate"`
	GoodsType                   int         `json:"goods_type"`                 //商品类型
	MallCouponTotalQuantity     int64       `json:"mall_coupon_total_quantity"` //店铺券总量
	CreateAt                    int64       `json:"create_at"`
	MallCouponMaxDiscountAmount int         `json:"mall_coupon_max_discount_amount"` //最大使用金额
	MallCps                     int         `json:"mall_cps"`
}

func (this *DuoduoKe) ZsUnitGoodsQuery(p *ZsUnitGoodsQueryParams) (*ZsUnitGoodsQueryResult, error) {
	apiType := `pdd.ddk.zs.unit.goods.query`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)
	var result ZsUnitGoodsQueryResult
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
